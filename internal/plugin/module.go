package plugin

import (
	"go/parser"
	"go/printer"
	"go/token"
	"path/filepath"
	"strings"

	pgs "github.com/timmonfette1/protoc-gen-star/v2"
	pgsgo "github.com/timmonfette1/protoc-gen-star/v2/lang/go"
)

type module struct {
	*pgs.ModuleBase
	pgsgo.Context
}

func New() pgs.Module {
	return &module{ModuleBase: &pgs.ModuleBase{}}
}

func (m *module) InitContext(ctx pgs.BuildContext) {
	m.ModuleBase.InitContext(ctx)
	m.Context = pgsgo.InitContext(ctx.Parameters())
}

func (module) Name() string {
	return "protago"
}

func (m *module) Execute(targets map[string]pgs.File, packages map[string]pgs.Package) []pgs.Artifact {
	extractor := newTagExtactor(m, m.Context)

	for _, file := range targets {
		tags := extractor.Extract(file)
		fileName := m.Context.OutputPath(file).SetExt(".go").String()

		outdir := m.Parameters().Str("outdir")
		if outdir != "" {
			fileName = filepath.Join(outdir, fileName)
		}

		fs := token.NewFileSet()
		f, err := parser.ParseFile(fs, fileName, nil, parser.ParseComments)
		m.CheckErr(err)

		err = replaceTags(f, tags)
		m.CheckErr(err)
		var buf strings.Builder
		m.CheckErr(printer.Fprint(&buf, fs, f))
		m.OverwriteGeneratorFile(fileName, buf.String())
	}

	return m.Artifacts()
}
