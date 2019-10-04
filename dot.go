// package dot is a extension for the goldmark(http://github.com/yuin/goldmark).
//
// This extension adds svg picture output from dot language using
// godot(https://github.com/OhYee/godot).
package dot

import (
	"bytes"
	"fmt"

	"github.com/OhYee/godot"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

// Config struct holds options for the extension.
type Config struct {
	languageName      string
	defaultRenderFunc renderer.NodeRendererFunc
}

type dot struct {
	config *Config
}

// Dot default dot extension when there is no other fencedCodeBlock goldmark render extensions
var Dot = NewDot("dot", html.NewRenderer())

// NewDot returns a new extension with given arguments.
func NewDot(languageName string, render renderer.NodeRenderer) goldmark.Extender {
	defaultRenderFunc := getRenderFunction(ast.KindFencedCodeBlock, render)
	if defaultRenderFunc == nil {
		panic(fmt.Sprintf("%T don't render ast.KindFencedCodeBlock(FencedCodeBlock)", render))
	}
	config := &Config{languageName, defaultRenderFunc}
	return &dot{config}
}

// Extend implements goldmark.Extender.
func (d *dot) Extend(m goldmark.Markdown) {
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewHTMLRenderer(d.config), 0),
	))
}

// HTMLRenderer struct is a renderer.NodeRenderer implementation for the extension.
type HTMLRenderer struct {
	config *Config
}

// NewHTMLRenderer builds a new HTMLRenderer with given options and returns it.
func NewHTMLRenderer(config *Config) renderer.NodeRenderer {
	r := &HTMLRenderer{config}
	return r
}

// RegisterFuncs implements NodeRenderer.RegisterFuncs.
func (r *HTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindFencedCodeBlock, r.renderFencedCodeBlock)
}

func (r *HTMLRenderer) getLines(source []byte, n ast.Node) []byte {
	buf := bytes.NewBuffer([]byte{})
	l := n.Lines().Len()
	for i := 0; i < l; i++ {
		line := n.Lines().At(i)
		buf.Write(line.Value(source))
	}
	return buf.Bytes()
}

func (r *HTMLRenderer) renderFencedCodeBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.FencedCodeBlock)
	language := n.Language(source)

	if string(language) == r.config.languageName {
		if !entering {
			svg, _ := godot.Dot(r.getLines(source, node))
			w.Write(svg)
		}
	} else {
		return r.config.defaultRenderFunc(w, source, node, entering)
	}

	return ast.WalkContinue, nil
}
