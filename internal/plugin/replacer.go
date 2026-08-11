package plugin

import (
	"go/ast"
	"go/token"
	"sort"
	"strings"

	"github.com/fatih/structtag"
)

type structVisitor struct {
	visitor func(n ast.Node) ast.Visitor
}

type retag struct {
	err  error
	tags map[string]*structtag.Tags
}

func replaceTags(node ast.Node, tags StructTags) error {
	r := retag{}
	f := func(n ast.Node) ast.Visitor {
		if r.err != nil {
			return nil
		}
		if tp, ok := n.(*ast.TypeSpec); ok {
			r.tags = tags[tp.Name.String()]
			return r
		}
		return nil
	}

	ast.Walk(structVisitor{f}, node)

	return r.err
}

func (v structVisitor) Visit(n ast.Node) ast.Visitor {
	if tp, ok := n.(*ast.TypeSpec); ok {
		if _, ok := tp.Type.(*ast.StructType); ok {
			ast.Walk(v.visitor(n), n)
			return nil
		}
	}
	return v
}

func (r retag) Visit(node ast.Node) ast.Visitor {
	if r.err != nil {
		return nil
	}

	if field, ok := node.(*ast.Field); ok {
		if len(field.Names) == 0 {
			return nil
		}

		newTags := r.tags[field.Names[0].String()]
		if newTags == nil {
			return nil
		}

		if field.Tag == nil {
			field.Tag = &ast.BasicLit{
				Kind: token.STRING,
			}
		}

		oldTags, err := structtag.Parse(strings.Trim(field.Tag.Value, "`"))
		if err != nil {
			r.err = err
			return nil
		}

		sort.Stable(newTags)
		for _, tag := range newTags.Tags() {
			oldTags.Set(tag)
		}

		field.Tag.Value = "`" + oldTags.String() + "`"
		return nil
	}

	return r
}
