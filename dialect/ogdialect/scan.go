package ogdialect

import (
	"reflect"

	"github.com/niconical/ogx/schema"
)

func scanner(typ reflect.Type) schema.ScannerFunc {
	return schema.Scanner(typ)
}
