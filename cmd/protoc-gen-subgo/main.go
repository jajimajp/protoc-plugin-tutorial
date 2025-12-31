package main

import (
	"google.golang.org/protobuf/compiler/protogen"
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
	g.P("}")

	return nil
}
