package datatype

import (
	"fmt"

	"github.com/emicklei/proto"

	"github.com/ignite/cli/ignite/pkg/multiformatname"
	"github.com/ignite/cli/ignite/pkg/protoanalysis/protoutil"
)

var (
	// DataUint uint data type definition.
	DataUint = DataType{
		DataType:          func(string) string { return "uint64" },
		DefaultTestValue:  "111",
		ValueLoop:         "uint64(i)",
		ValueIndex:        "0",
		ValueInvalidIndex: "100000",
		ProtoType: func(_, name string, index int) string {
			return fmt.Sprintf("uint64 %s = %d", name, index)
		},
		GenesisArgs: func(name multiformatname.Name, value int) string {
			return fmt.Sprintf("%s: %d,\n", name.UpperCamel, value)
		},
		ToBytes: func(name string) string {
			return fmt.Sprintf(`%[1]vBytes := make([]byte, 8)
  					binary.BigEndian.PutUint64(%[1]vBytes, %[1]v)`, name)
		},
		ToString: func(name string) string {
			return fmt.Sprintf("strconv.Itoa(int(%s))", name)
		},
		ToProtoField: func(_, name string, index int) *proto.NormalField {
			return protoutil.NewField(name, "uint64", index)
		},
		GoCLIImports: []GoImport{{Name: "github.com/spf13/cast"}},
	}

	// DataUintSlice uint array data type definition.
	DataUintSlice = DataType{
		DataType:         func(string) string { return "[]uint64" },
		DefaultTestValue: "1,2,3,4,5",
		ProtoType: func(_, name string, index int) string {
			return fmt.Sprintf("repeated uint64 %s = %d", name, index)
		},
		GenesisArgs: func(name multiformatname.Name, value int) string {
			return fmt.Sprintf("%s: []uint64{%d},\n", name.UpperCamel, value)
		},
		ToProtoField: func(_, name string, index int) *proto.NormalField {
			return protoutil.NewField(name, "uint64", index, protoutil.Repeated())
		},
		GoCLIImports: []GoImport{{Name: "github.com/spf13/cast"}, {Name: "strings"}},
		NonIndex:     true,
	}
)
