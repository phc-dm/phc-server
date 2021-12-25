package articles

import (
	"fmt"
	"os"
	"path"
	"sort"
)

type Registry struct {
	RootPath     string
	ArticleCache map[string]*Article
}

func NewRegistry(rootPath string) *Registry {
	return &Registry{
		rootPath,
		map[string]*Article{},
	}
}

func (registry *Registry) loadArticles() error {
	entries, err := os.ReadDir(registry.RootPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			article, err := NewArticle(path.Join(registry.RootPath, entry.Name()))
			if err != nil {
				return err
			}

			registry.ArticleCache[article.Id] = article
		}
	}

	return nil
}

func (registry *Registry) GetArticle(id string) (*Article, error) {
	article, present := registry.ArticleCache[id]
	if !present {
		err := registry.loadArticles()
		if err != nil {
			return nil, err
		}

		article, present := registry.ArticleCache[id]
		if !present {
			return nil, fmt.Errorf(`no article with id "%s"`, id)
		}

		return article, nil
	}

	return article, nil
}

func (registry *Registry) GetArticles() ([]*Article, error) {
	err := registry.loadArticles()
	if err != nil {
		return nil, err
	}

	articles := []*Article{}
	for _, article := range registry.ArticleCache {
		articles = append(articles, article)
	}

	sort.Slice(articles, func(i, j int) bool {
		return articles[i].PublishDate.After(articles[j].PublishDate)
	})

	return articles, nil
}
