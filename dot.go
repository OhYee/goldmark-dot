// Package dot is a extension for the goldmark(http://github.com/yuin/goldmark).
//
// This extension adds svg picture output from dot language using
// godot(https://github.com/OhYee/godot).
package dot

import (
	"bytes"
	"crypto/sha1"

	"github.com/OhYee/godot"
	ext "github.com/OhYee/goldmark-fenced_codeblock_extension"
	fp "github.com/OhYee/goutils/functional"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
)

// Default dot extension when there is no other fencedCodeBlock goldmark render extensions
var Default = NewDotExtension(20, "dot")

// RenderMap return the goldmark-fenced_codeblock_extension.RenderMap
func RenderMap(length int, languages ...string) ext.RenderMap {
	return ext.RenderMap{
		Languages:      languages,
		RenderFunction: NewDot(length, languages...).Renderer,
	}
}

// NewDotExtension return the goldmark.Extender
func NewDotExtension(length int, languages ...string) goldmark.Extender {
	return ext.NewExt(RenderMap(length, languages...))
}

// Dot render struct
type Dot struct {
	Languages []string
	buf       map[string][]byte
	MaxLength int
}

// NewDot initial a Dot struct
func NewDot(length int, languages ...string) *Dot {
	return &Dot{Languages: languages, buf: make(map[string][]byte), MaxLength: length}
}

// Renderer render function
func (d *Dot) Renderer(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.FencedCodeBlock)
	language := string(n.Language(source))

	if fp.AnyString(func(l string, idx int) bool {
		return l == language
	}, d.Languages) {
		if !entering {
			raw := d.getLines(source, node)
			h := sha1.New()
			h.Write(raw)
			hash := string(h.Sum([]byte{}))
			if result, exist := d.buf[hash]; exist {
				w.Write([]byte(result))
			} else {
				svg, _ := godot.Dot(raw)
				if len(d.buf) >= d.MaxLength {
					d.buf = make(map[string][]byte)
				}
				d.buf[hash] = svg
				w.Write(svg)
			}
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
