package articles

import (
	"fmt"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"

	chromahtml "github.com/alecthomas/chroma/formatters/html"
	mathjax "github.com/litao91/goldmark-mathjax"
	highlighting "github.com/yuin/goldmark-highlighting"
)

var articleMarkdown goldmark.Markdown

// https://github.com/yuin/goldmark-highlighting/blob/9216f9c5aa010c549cc9fc92bb2593ab299f90d4/highlighting_test.go#L27
func customCodeBlockWrapper(w util.BufWriter, c highlighting.CodeBlockContext, entering bool) {
	lang, ok := c.Language()
	if entering {
		if ok {
			w.WriteString(fmt.Sprintf(`<pre class="%s"><code>`, lang))
			return
		}
		w.WriteString("<pre><code>")
	} else {
		if ok {
			w.WriteString("</pre></code>")
			return
		}
		w.WriteString("</pre></code>")
	}
}

func init() {
	articleMarkdown = goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Typographer,
			// Questo pacchetto ha un nome stupido perché in realtà si occupa solo del parsing lato server del Markdown mentre lato client usiamo KaTeX.
			mathjax.NewMathJax(),
			highlighting.NewHighlighting(
				highlighting.WithStyle("github"),
				highlighting.WithWrapperRenderer(customCodeBlockWrapper),
				highlighting.WithFormatOptions(
					chromahtml.PreventSurroundingPre(true),
				),
			),
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
}
