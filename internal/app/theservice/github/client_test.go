package github

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient(t *testing.T) {

	Convey("Client", t, func() {

		Convey("should return response from GitHub", func() {

			// given
			// Start a local HTTP server
			githubStub := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				// Send response to be tested
				rw.Header().Set("Content-Type", "text/plain;charset=utf-8")
				rw.WriteHeader(http.StatusOK)
				rw.Write([]byte("Responsive is better than fast."))
			}))
			// Close the server when test finishes
			defer githubStub.Close()

			config := &ClientConfig{
				GitHubApiUrl: githubStub.URL,
			}
			client := NewClient(config)

			// when
			zen, err := client.GetZen()

			// then
			So(err, ShouldBeNil)
			So(zen, ShouldNotBeNil)
			So(zen, ShouldEqual, "Responsive is better than fast.")
		})

		Convey("should return error if backend endpoint not found", func() {

			// given

			// Start a local HTTP server
			githubStub := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.WriteHeader(http.StatusNotFound)
			}))
			// Close the server when test finishes
			defer githubStub.Close()

			config := &ClientConfig{
				GitHubApiUrl: githubStub.URL,
			}
			client := NewClient(config)

			// when
			zen, err := client.GetZen()

			// then
			So(zen, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
		})

		Convey("should return error if backend endpoint is down", func() {

			// given

			// Start a local HTTP server
			githubStub := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				// Send response to be tested
				rw.Header().Set("Content-Type", "text/plain;charset=utf-8")
				rw.WriteHeader(http.StatusInternalServerError)
			}))
			// Close the server when test finishes
			defer githubStub.Close()

			config := &ClientConfig{
				GitHubApiUrl: githubStub.URL,
			}
			client := NewClient(config)

			// when
			zen, err := client.GetZen()

			// then
			So(zen, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
		})

		Convey("should return error if connection to backend endpoint is broken", func() {

			// given
			config := &ClientConfig{
				GitHubApiUrl: "http://localhost:0",
			}
			client := NewClient(config)

			// when
			zen, err := client.GetZen()

			// then
			So(zen, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
		})
	})
}

