package articles

import (
	"bytes"
	"os"
	"strings"
	"time"

	"github.com/phc-dm/phc-server/config"
	"gopkg.in/yaml.v3"
)

type Article struct {
	Id          string
	Title       string
	Description string
	Tags        []string
	PublishDate time.Time

	ArticlePath    string
	markdownSource string
	renderedHTML   string
}

func (article *Article) HasTag(tag string) bool {
	for _, t := range article.Tags {
		if t == tag {
			return true
		}
	}

	return false
}

func NewArticle(articlePath string) (*Article, error) {
	article := &Article{
		ArticlePath: articlePath,
	}

	err := article.load()
	if err != nil {
		return nil, err
	}

	return article, nil
}

func trimAll(vs []string) []string {
	r := []string{}
	for _, v := range vs {
		r = append(r, strings.TrimSpace(v))
	}
	return r
}

func (article *Article) load() error {
	fileBytes, err := os.ReadFile(article.ArticlePath)
	if err != nil {
		return err
	}

	source := string(fileBytes)

	// TODO: Ehm bugia pare che esista "https://github.com/yuin/goldmark-meta" però non penso valga la pena aggiungerlo
	parts := strings.SplitN(source, "---", 3)[1:]

	frontMatterSource := parts[0]
	markdownSource := parts[1]

	var frontMatter struct {
		Id          string `yaml:"id"`
		Title       string `yaml:"title"`
		Description string `yaml:"description"`
		Tags        string `yaml:"tags"`
		PublishDate string `yaml:"publish_date"`
	}
	if err := yaml.Unmarshal([]byte(frontMatterSource), &frontMatter); err != nil {
		return err
	}

	publishDate, err := time.Parse("2006/01/02 15:04", frontMatter.PublishDate)
	if err != nil {
		return err
	}

	article.Id = frontMatter.Id
	article.Title = frontMatter.Title
	article.Description = frontMatter.Description
	article.Tags = trimAll(strings.Split(frontMatter.Tags, ","))
	article.PublishDate = publishDate

	article.markdownSource = markdownSource
	article.renderedHTML = ""

	return nil
}

func (article *Article) Render() (string, error) {
	if config.Mode == "development" {
		article.load()
	}

	if article.renderedHTML == "" {
		var buf bytes.Buffer
		if err := articleMarkdown.Convert([]byte(article.markdownSource), &buf); err != nil {
			return "", err
		}

		article.renderedHTML = buf.String()
	}

	return article.renderedHTML, nil
}
