package datatype

import (
	"fmt"

	"github.com/emicklei/proto"

	"github.com/ignite/cli/ignite/pkg/multiformatname"
	"github.com/ignite/cli/ignite/pkg/protoanalysis/protoutil"
)

var (
	// DataInt is an int data type definition.
	DataInt = DataType{
		DataType:          func(string) string { return "int32" },
		DefaultTestValue:  "111",
		ValueLoop:         "int32(i)",
		ValueIndex:        "0",
		ValueInvalidIndex: "100000",
		ProtoType: func(_, name string, index int) string {
			return fmt.Sprintf("int32 %s = %d", name, index)
		},
		GenesisArgs: func(name multiformatname.Name, value int) string {
			return fmt.Sprintf("%s: %d,\n", name.UpperCamel, value)
		},
		ToBytes: func(name string) string {
			return fmt.Sprintf(`%[1]vBytes := make([]byte, 4)
  					binary.BigEndian.PutUint32(%[1]vBytes, uint32(%[1]v))`, name)
		},
		ToString: func(name string) string {
			return fmt.Sprintf("strconv.Itoa(int(%s))", name)
		},
		ToProtoField: func(_, name string, index int) *proto.NormalField {
			return protoutil.NewField(name, "int32", index)
		},
		GoCLIImports: []GoImport{{Name: "github.com/spf13/cast"}},
	}

	// DataIntSlice is an int array data type definition.
	DataIntSlice = DataType{
		DataType:         func(string) string { return "[]int32" },
		DefaultTestValue: "1,2,3,4,5",
		ProtoType: func(_, name string, index int) string {
			return fmt.Sprintf("repeated int32 %s = %d", name, index)
		},
		GenesisArgs: func(name multiformatname.Name, value int) string {
			return fmt.Sprintf("%s: []int32{%d},\n", name.UpperCamel, value)
		},
		ToProtoField: func(_, name string, index int) *proto.NormalField {
			return protoutil.NewField(name, "int32", index, protoutil.Repeated())
		},
		GoCLIImports: []GoImport{{Name: "github.com/spf13/cast"}, {Name: "strings"}},
		NonIndex:     true,
	}
)
