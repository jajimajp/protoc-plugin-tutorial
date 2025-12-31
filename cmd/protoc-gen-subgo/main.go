package main

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func main() {
	opts := protogen.Options{}
	opts.Run(func(plugin *protogen.Plugin) error {
		for _, f := range plugin.Files {
			fn := f.GeneratedFilenamePrefix + ".sub.go"
			g := plugin.NewGeneratedFile(fn, f.GoImportPath)
			g.P("package ", f.GoPackageName)

			for _, m := range f.Messages {
				err := processMessage(g, m)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
}

func processMessage(g *protogen.GeneratedFile, m *protogen.Message) error {
	g.P("type ", m.GoIdent, " struct {")
	for _, f := range m.Fields {
		ty, err := goTypeName(f)
		if err != nil {
			return err
		}
		g.P(f.GoName, " ", ty)
	}
	g.P("}")

	return nil
}

func goTypeName(f *protogen.Field) (string, error) {
	switch k := f.Desc.Kind(); k {
	case protoreflect.Int32Kind:
		return "int32", nil
	case protoreflect.StringKind:
		return "string", nil
	case protoreflect.MessageKind:
		return f.Message.GoIdent.GoName, nil
	default:
		return "", fmt.Errorf("未対応の Field: %s %v", f.GoName, k)
	}
}
