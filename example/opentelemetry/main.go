package main

import (
	"context"
	"database/sql"

	"github.com/uptrace/opentelemetry-go-extra/otelplay"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/niconical/ogx"
	"github.com/niconical/ogx/dialect/ogdialect"
	"github.com/niconical/ogx/extra/ogxotel"

	_ "gitee.com/opengauss/openGauss-connector-go-pq"
)

var tracer = otel.Tracer("github.com/niconical/ogx/example/opentelemetry")

func main() {
	ctx := context.Background()

	shutdown := otelplay.ConfigureOpentelemetry(ctx)
	defer shutdown()

	connStr := "host=192.168.20.40 port=26000 user=cuih password=Gauss@123 dbname=test sslmode=disable"
	opengaussdb, err := sql.Open("opengauss", connStr)
	if err != nil {
		panic(err)
	}
	opengaussdb.SetMaxOpenConns(1)

	db := ogx.NewDB(opengaussdb, ogdialect.New())
	db.AddQueryHook(ogxotel.NewQueryHook())

	if err := db.ResetModel(ctx, (*TestModel)(nil)); err != nil {
		panic(err)
	}
	defer func() {
		_ = dropSchema(ctx, db, (*TestModel)(nil))
	}()

	ctx, span := tracer.Start(ctx, "handleRequest")
	defer span.End()

	if err := handleRequest(ctx, db); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	otelplay.PrintTraceID(ctx)
}

type TestModel struct {
	ID   int64 `ogx:",pk,autoincrement"`
	Name string
}

func handleRequest(ctx context.Context, db *ogx.DB) error {
	model := &TestModel{
		Name: gofakeit.Name(),
	}
	if _, err := db.NewInsert().Model(model).Exec(ctx); err != nil {
		return err
	}

	// Check that data can be selected without any errors.
	if err := db.NewSelect().Model(model).WherePK().Scan(ctx); err != nil {
		return err
	}

	return nil
}

func dropSchema(ctx context.Context, db *ogx.DB, models ...interface{}) error {
	for _, model := range models {
		if _, err := db.NewDropTable().Model(model).IfExists().Cascade().Exec(ctx); err != nil {
			return err
		}
	}
	return nil
}
