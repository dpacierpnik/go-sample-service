package router

import (
	"github.com/dpacierpnik/go-sample-service/internal/app/theservice/echo"
	"github.com/dpacierpnik/go-sample-service/internal/app/theservice/github"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServeMux(t *testing.T) {

	Convey("The Service", t, func() {

		Convey("echo endpoints", func() {

			endpoints := &Endpoints{
				Echo: echo.NewEndpoints(),
			}
			mux := NewServeMux(endpoints)

			Convey("should rewrite request body", func() {

				// Given:
				reqPayload := `{
					"field1": "value1",
					"field2": "value2"
				}`
				req, err := http.NewRequest("POST", "/echo/body", strings.NewReader(reqPayload))
				if err != nil {
					t.Fatal(err)
				}
				respRecorder := httptest.NewRecorder()

				// When:
				mux.ServeHTTP(respRecorder, req)

				// Then:
				So(respRecorder.Code, ShouldEqual, http.StatusOK)
				So(respRecorder.Body.String(), ShouldEqual, reqPayload)
			})

			Convey("should rewrite request headers", func() {

				// Given:
				req, err := http.NewRequest("POST", "/echo/headers", nil)
				req.Header.Set("Content-Type", "text/html")
				req.Header.Set("Authorization", "Bearer abc")
				if err != nil {
					t.Fatal(err)
				}
				respRecorder := httptest.NewRecorder()

				// When:
				mux.ServeHTTP(respRecorder, req)

				// Then:
				So(respRecorder.Code, ShouldEqual, http.StatusOK)

				expectedResBody := `{"Authorization":["Bearer abc"],"Content-Type":["text/html"]}`
				So(respRecorder.Body.String(), ShouldEqual, expectedResBody)
			})
		})

		Convey("zen endpoint", func() {

			// Start a local HTTP server
			githubStub := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				// Send response to be tested
				rw.Header().Set("Content-Type", "text/plain;charset=utf-8")
				rw.WriteHeader(http.StatusOK)
				rw.Write([]byte("Responsive is better than fast."))
			}))
			// Close the server when test finishes
			defer githubStub.Close()

			clientConfig := &github.ClientConfig{
				GitHubApiUrl: githubStub.URL,
			}
			client := github.NewClient(clientConfig)

			endpoints := &Endpoints{
				Github: github.NewEndpoints(client),
			}
			mux := NewServeMux(endpoints)

			Convey("should return zen quotation from GitHub", func() {

				// Given:
				req, err := http.NewRequest("GET", "/github/zen", nil)
				if err != nil {
					t.Fatal(err)
				}
				req.Header.Set("Content-Type", "application/json")
				respRecorder := httptest.NewRecorder()

				// When:
				mux.ServeHTTP(respRecorder, req)

				// Then:
				So(respRecorder.Code, ShouldEqual, http.StatusOK)
			})
		})
	})
}
