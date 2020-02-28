package dot

import (
	"bytes"
	"testing"

	"github.com/OhYee/godot"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

func Test_default(t *testing.T) {
	var buf bytes.Buffer
	source := []byte("```go\npackage main\n\nimport ()\n\nfunc main(){}\n```\n\n```dot\ndigraph{a->b}\n```\n\n")
	want := `<pre><code class="language-go">package main

import ()

func main(){}
</code></pre>
` + func() string { a, _ := godot.Dot([]byte("digraph{a->b}")); return string(a) }()
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			Default,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(),
	)

	if err := md.Convert(source, &buf); err != nil {
		t.Error(err)
	}
	if bytes.Compare(buf.Bytes(), []byte(want)) != 0 {
		t.Errorf("got %s, excepted %s\n", buf.Bytes(), want)
	}
}
