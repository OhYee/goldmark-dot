// Package dot is a extension for the goldmark(http://github.com/yuin/goldmark).
//
// This extension adds svg picture output from dot language using
// godot(https://github.com/OhYee/godot).
package dot

import (
	"bytes"

	"github.com/OhYee/godot"
	ext "github.com/OhYee/goldmark-fenced_codeblock_extension"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
)

// Default dot extension when there is no other fencedCodeBlock goldmark render extensions
var Default = NewDotExtension("dot")

func NewDotExtension(languageName string) goldmark.Extender {
	return ext.NewExt([]ext.RenderMap{
		ext.RenderMap{
			Language:       []string{languageName},
			RenderFunction: NewDot(languageName).Renderer,
		},
	}...)
}

type Dot struct {
	LanguageName string
}

func NewDot(languageName string) *Dot {
	return &Dot{languageName}
}

func (d *Dot) Renderer(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.FencedCodeBlock)
	language := n.Language(source)
	if string(language) == d.LanguageName {
		if !entering {
			svg, _ := godot.Dot(d.getLines(source, node))
			w.Write(svg)
		}
	}
	return ast.WalkContinue, nil
}

func (d *Dot) getLines(source []byte, n ast.Node) []byte {
	buf := bytes.NewBuffer([]byte{})
	l := n.Lines().Len()
	for i := 0; i < l; i++ {
		line := n.Lines().At(i)
		buf.Write(line.Value(source))
	}
	return buf.Bytes()
}
