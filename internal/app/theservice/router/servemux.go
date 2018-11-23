package router

import (
	"github.com/dpacierpnik/go-sample-service/internal/app/theservice/echo"
	"github.com/dpacierpnik/go-sample-service/internal/app/theservice/github"
	"github.com/dpacierpnik/go-sample-service/internal/app/theservice/html"
	"net/http"
)

type Endpoints struct {
	Html   *html.WebPages
	Github *github.Endpoints
	Echo   *echo.Endpoints
}

func NewServeMux(endpoints *Endpoints) *http.ServeMux {

	mux := http.NewServeMux()

	if endpoints.Github != nil {
		mux.HandleFunc("/github/zen", endpoints.Github.Zen)
	}
	if endpoints.Echo != nil {
		mux.HandleFunc("/echo/body", endpoints.Echo.Body)
		mux.HandleFunc("/echo/headers", endpoints.Echo.Headers)
	}

	if endpoints.Html != nil {
		mux.HandleFunc("/", endpoints.Html.Index)
	}

	return mux
}
