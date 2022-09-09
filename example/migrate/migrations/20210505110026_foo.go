package migrations

import (
	"context"
	"fmt"

	"github.com/niconical/ogx"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *ogx.DB) error {
		fmt.Print(" [up migration] ")
		return nil
	}, func(ctx context.Context, db *ogx.DB) error {
		fmt.Print(" [down migration] ")
		return nil
	})
}
