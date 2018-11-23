package appcontext

import (
	"github.com/dpacierpnik/go-sample-service/internal/app/theservice/echo"
	"github.com/dpacierpnik/go-sample-service/internal/app/theservice/github"
	"github.com/dpacierpnik/go-sample-service/internal/app/theservice/html"
	"github.com/dpacierpnik/go-sample-service/internal/app/theservice/router"
	"net/http"
)

type AppConfig interface {
	GetString(key string) string
	GetInt(key string) int
}

func NewServeMux(config AppConfig) *http.ServeMux {

	endpoints := newEndpoints(config)

	return router.NewServeMux(endpoints)
}

func newEndpoints(config AppConfig) *router.Endpoints {

	return &router.Endpoints{
		Html:   newHtmlEndpoints(config),
		Github: newGithubEndpoints(config),
		Echo:   echo.NewEndpoints(),
	}
}

func newHtmlEndpoints(config AppConfig) *html.WebPages {

	webPagesConfig := &html.WebPagesConfig{
		TemplatesDir: config.GetString("webpages.templatesDir"),
	}

	return html.NewWebPages(webPagesConfig)
}

func newGithubEndpoints(config AppConfig) *github.Endpoints {

	githubConfig := &github.ClientConfig{
		GitHubApiUrl: config.GetString("github.api.url"),
	}

	githubClient := github.NewClient(githubConfig)

	return github.NewEndpoints(githubClient)
}
