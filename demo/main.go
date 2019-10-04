package main

import (
	"bytes"
	"fmt"
	"github.com/OhYee/goldmark-dot"
	"github.com/yuin/goldmark"
	// "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

func main() {
	var buf bytes.Buffer
	source := []byte("```go\npackage main\n\nimport ()\n\nfunc main(){}\n```\n\n```dot\ndigraph{a->b}\n```\n\n")

	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			dot.Dot, // or dot.NewDot("dot-svg", highlighting.NewHTMLRenderer()),
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(),
	)

	if err := md.Convert(source, &buf); err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", buf.Bytes())
}
