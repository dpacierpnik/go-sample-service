package github

import (
	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEndpoints(t *testing.T) {

	Convey("Zen endpoint", t, func() {

		Convey("should return GitHub zen response", func() {

			// Given:
			client := &Client{
				GetZen: func() (Zen, error) {
					return Zen("Responsive is better than fast."), nil
				},
			}
			zenEndpoint := newZenEndpoint(client)

			req, err := http.NewRequest("GET", "/zenMapping", nil)
			if err != nil {
				t.Fatal(err)
			}
			respRecorder := httptest.NewRecorder()

			// When:
			zenEndpoint(respRecorder, req)

			// Then:
			So(respRecorder.Code, ShouldEqual, http.StatusOK)
			expectedResponseBody := `{"zen":"Responsive is better than fast."}`
			So(respRecorder.Body.String(), ShouldEqual, expectedResponseBody)
		})

		Convey("should return 500 if communication failed", func() {

			// Given:
			client := &Client{
				GetZen: func() (Zen, error) {
					return "", errors.New("Any error.")
				},
			}
			zenEndpoint := newZenEndpoint(client)

			req, err := http.NewRequest("GET", "/zenMapping", nil)
			if err != nil {
				t.Fatal(err)
			}
			respRecorder := httptest.NewRecorder()

			// When:
			zenEndpoint(respRecorder, req)

			// Then:
			So(respRecorder.Code, ShouldEqual, http.StatusInternalServerError)
		})
	})
}
