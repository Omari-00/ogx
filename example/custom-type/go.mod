module github.com/niconical/ogx/example/custom-type

go 1.18

replace github.com/niconical/ogx => ../..

replace github.com/niconical/ogx/dbfixture => ../../dbfixture

replace github.com/niconical/ogx/extra/ogxdebug => ../../extra/ogxdebug

replace github.com/niconical/ogx/dialect/ogdialect => ../../dialect/ogdialect

require (
	gitee.com/opengauss/openGauss-connector-go-pq v1.0.4
	github.com/niconical/ogx v1.1.7
	github.com/niconical/ogx/dialect/ogdialect v0.0.0-20220922053747-c69468f75cc6
	github.com/niconical/ogx/extra/ogxdebug v0.0.0-20220922053747-c69468f75cc6
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
	golang.org/x/crypto v0.0.0-20220919173607-35f4265a4bc0 // indirect
	golang.org/x/sys v0.0.0-20220919091848-fb04ddd9f9c8 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
