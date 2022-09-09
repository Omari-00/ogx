module github.com/niconical/ogx/example/tx-composition

go 1.18

replace github.com/niconical/ogx => ../..

replace github.com/niconical/ogx/dbfixture => ../../dbfixture

replace github.com/niconical/ogx/extra/ogxdebug => ../../extra/ogxdebug

replace github.com/niconical/ogx/dialect/ogdialect => ../../dialect/ogdialect

require (
	gitee.com/opengauss/openGauss-connector-go-pq v1.0.3
	github.com/niconical/ogx v1.1.7
	github.com/niconical/ogx/dialect/ogdialect v0.0.0-20220903032934-13a3cbc7d42c
	github.com/niconical/ogx/extra/ogxdebug v1.1.7
)

require (
	github.com/fatih/color v1.13.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/tjfoc/gmsm v1.4.1 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/crypto v0.0.0-20220829220503-c86fa9a7ed90 // indirect
	golang.org/x/sys v0.0.0-20220907062415-87db552b00fd // indirect
	golang.org/x/xerrors v0.0.0-20220609144429-65e65417b02f // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
