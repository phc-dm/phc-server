package main

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"

	chromahtml "github.com/alecthomas/chroma/formatters/html"
	mathjax "github.com/litao91/goldmark-mathjax"
	highlighting "github.com/yuin/goldmark-highlighting"

	"gopkg.in/yaml.v3"
)

var md goldmark.Markdown

// https://github.com/yuin/goldmark-highlighting/blob/9216f9c5aa010c549cc9fc92bb2593ab299f90d4/highlighting_test.go#L27
func customWrapper(w util.BufWriter, c highlighting.CodeBlockContext, entering bool) {
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

	md = goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Typographer,
			// Questo pacchetto ha un nome stupido perché in realtà si occupa solo del parsing lato server del Markdown mentre lato client usiamo KaTeX.
			mathjax.NewMathJax(),
			highlighting.NewHighlighting(
				highlighting.WithStyle("github"),
				highlighting.WithWrapperRenderer(customWrapper),
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

type ArticleFrontMatter struct {
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	Tags        []string `yaml:"tags,flow"`
	PublishDate string   `yaml:"publish_date"`
	Important   bool     `yaml:"important"`
}

type Article struct {
	Id          string
	Title       string
	Description string
	Tags        []string
	PublishDate time.Time
	Important   bool

	SourceMarkdown string
	OutputHTML     string
}

type ArticleRegistry struct {
	Articles map[string]*Article
}

func removeBlanks(v []string) []string {
	r := []string{}

	for _, s := range v {
		if len(strings.TrimSpace(s)) > 0 {
			r = append(r, s)
		}
	}

	return r
}

func (ar *ArticleRegistry) LoadArticle(name string) (*Article, error) {
	fileBytes, err := os.ReadFile(path.Join("./news/", name+".md"))
	if err != nil {
		return nil, err
	}

	source := string(fileBytes)

	// TODO: Ehm bugia esiste "https://github.com/yuin/goldmark-meta" però boh per ora funge tutto
	parts := removeBlanks(strings.Split(source, "---"))

	frontMatterSource := parts[0]
	markdownSource := strings.Join(parts[1:], "---")

	var frontMatter ArticleFrontMatter
	if err := yaml.Unmarshal([]byte(frontMatterSource), &frontMatter); err != nil {
		return nil, err
	}

	publishDate, err := time.Parse("2006/01/02 15:04", frontMatter.PublishDate)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := md.Convert([]byte(markdownSource), &buf); err != nil {
		return nil, err
	}

	article := &Article{
		Id:             name,
		Title:          frontMatter.Title,
		Description:    frontMatter.Description,
		Tags:           frontMatter.Tags,
		PublishDate:    publishDate,
		Important:      frontMatter.Important,
		SourceMarkdown: markdownSource,
		OutputHTML:     buf.String(),
	}

	ar.Articles[name] = article
	return article, nil
}

func (ar *ArticleRegistry) LoadAll() ([]*Article, error) {
	entries, err := os.ReadDir("./news")
	if err != nil {
		return nil, err
	}

	articles := []*Article{}

	for _, entry := range entries {
		if !entry.IsDir() {
			name := strings.TrimRight(entry.Name(), ".md")
			article, err := ar.LoadArticle(name)
			if err != nil {
				return nil, err
			}

			articles = append(articles, article)
		}
	}

	sort.Slice(articles, func(i, j int) bool {
		return articles[i].PublishDate.After(articles[j].PublishDate)
	})

	return articles, nil
}

func NewArticleRegistry() *ArticleRegistry {
	return &ArticleRegistry{
		map[string]*Article{},
	}
}

func (ar *ArticleRegistry) Render(name string) (string, error) {
	article, err := ar.LoadArticle(name)
	if err != nil {
		return "", err
	}

	return article.OutputHTML, nil
}
