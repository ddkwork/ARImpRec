package ipmrec

import (
	"strings"
	"testing"

	"github.com/ddkwork/bindgen/c2go"
)

func TestGenerate(t *testing.T) {
	c2go.Generate(t, []c2go.BindgenConfig{{
		HeadersDir:  ".",
		OutputDir:   ".",
		PackageName: "ipmrec",
		HeaderOrder: []string{"ARImpRec.h"},
		BindDll:     true,
		DllName:     "ARImpRec.dll",
		DllFuncFilter: func(name string) bool {
			return !strings.HasPrefix(name, "__builtin") && !strings.HasPrefix(name, "_")
		},
	}})
}
