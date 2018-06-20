package custommiddleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vicanso/pike/cache"

	"github.com/vicanso/pike/vars"

	"github.com/labstack/echo"
)

func TestHeaderSetter(t *testing.T) {
	headerSetterConfig := HeaderSetterConfig{}
	t.Run("set header(response not set)", func(t *testing.T) {
		fn := HeaderSetter(headerSetterConfig)(func(c echo.Context) error {
			return nil
		})
		e := echo.New()
		pc := NewContext(e.NewContext(nil, nil))
		err := fn(pc)
		if err != vars.ErrResponseNotSet {
			t.Fatalf("response not set should return error")
		}
	})

	t.Run("set header", func(t *testing.T) {
		fn := HeaderSetter(headerSetterConfig)(func(c echo.Context) error {
			pc := c.(*Context)
			if (pc.Response().Header().Get("Token")) != "ABCD" {
				t.Fatalf("set header fail")
			}
			return nil
		})
		e := echo.New()
		c := e.NewContext(nil, &httptest.ResponseRecorder{})
		pc := NewContext(c)

		resp := &cache.Response{
			Header: http.Header{
				"Token": []string{
					"ABCD",
				},
				"Date": []string{
					"Sat, 28 Apr 2018 02:59:16 GMT",
				},
			},
		}
		pc.resp = resp
		fn(pc)
	})
}
