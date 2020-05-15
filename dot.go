// Package dot is a extension for the goldmark(http://github.com/yuin/goldmark).
//
// This extension adds svg picture output from dot language using
// godot(https://github.com/OhYee/godot).
package dot

import (
	"bytes"

	"github.com/OhYee/godot"
	ext "github.com/OhYee/goldmark-fenced_codeblock_extension"
	fp "github.com/OhYee/goutils/functional"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
)

// Default dot extension when there is no other fencedCodeBlock goldmark render extensions
var Default = NewDotExtension("dot")

// RenderMap return the goldmark-fenced_codeblock_extension.RenderMap
func RenderMap(languages ...string) ext.RenderMap {
	return ext.RenderMap{
		Languages:      languages,
		RenderFunction: NewDot(languages...).Renderer,
	}
}

// NewDotExtension return the goldmark.Extender
func NewDotExtension(languages ...string) goldmark.Extender {
	return ext.NewExt(RenderMap(languages...))
}

// Dot render struct
type Dot struct {
	Languages []string
}

// NewDot initial a Dot struct
func NewDot(languages ...string) *Dot {
	return &Dot{languages}
}

// Renderer render function
func (d *Dot) Renderer(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.FencedCodeBlock)
	language := string(n.Language(source))

	if fp.AnyString(func(l string) bool {
		return l == language
	}, d.Languages) {
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
