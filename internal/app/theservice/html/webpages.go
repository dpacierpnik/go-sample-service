package html

import (
	"github.com/dpacierpnik/go-sample-service/internal/app/theservice/html/webpages"
	"net/http"
)

type WebPages struct {
	Index http.HandlerFunc
}

type WebPagesConfig struct {
	TemplatesDir string
}

func NewWebPages(config *WebPagesConfig) *WebPages {
	return &WebPages{
		Index: webpages.NewIndex(config.TemplatesDir),
	}
}
