package main

import (
	"bytes"
	"io/ioutil"
	"path"
	"runtime"

	dot "github.com/OhYee/goldmark-dot"
	ext "github.com/OhYee/goldmark-fenced_codeblock_extension"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

var raw = "```go\npackage main\n\nimport ()\nfunc main(){}\n```\n\n```dot-svg\ndigraph{a->b}\n```"

func main() {
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			ext.NewExt(
				dot.RenderMap(20, "dot-svg"),
				ext.RenderMap{
					Languages:      []string{"*"},
					RenderFunction: ext.GetFencedCodeBlockRendererFunc(highlighting.NewHTMLRenderer()),
				},
			),
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(),
	)
	buf := bytes.Buffer{}
	if err := md.Convert([]byte(raw), &buf); err != nil {
		panic(err.Error())
	}

	_, file, _, _ := runtime.Caller(0)
	if err := ioutil.WriteFile(path.Join(path.Dir(file), "output.html"), buf.Bytes(), 777); err != nil {
		panic(err.Error())
	}
}
