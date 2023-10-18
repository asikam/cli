package datatype

import (
	"fmt"

	"github.com/emicklei/proto"

	"github.com/ignite/cli/ignite/pkg/multiformatname"
	"github.com/ignite/cli/ignite/pkg/protoanalysis/protoutil"
)

// DataBool bool data type definition.
var DataBool = DataType{
	DataType:          func(string) string { return "bool" },
	DefaultTestValue:  "false",
	ValueLoop:         "false",
	ValueIndex:        "false",
	ValueInvalidIndex: "false",
	ProtoType: func(_, name string, index int) string {
		return fmt.Sprintf("bool %s = %d", name, index)
	},
	GenesisArgs: func(name multiformatname.Name, value int) string {
		return fmt.Sprintf("%s: %t,\n", name.UpperCamel, value%2 == 0)
	},
	ToBytes: func(name string) string {
		return fmt.Sprintf(`%[1]vBytes := []byte{0}
					if %[1]v {
						%[1]vBytes = []byte{1}
					}`, name)
	},
	ToString: func(name string) string {
		return fmt.Sprintf("strconv.FormatBool(%s)", name)
	},
	ToProtoField: func(_, name string, index int) *proto.NormalField {
		return protoutil.NewField(name, "bool", index)
	},
	GoCLIImports: []GoImport{{Name: "github.com/spf13/cast"}},
}
