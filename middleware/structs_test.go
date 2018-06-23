package custommiddleware

import (
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo"
)

func TestContext(t *testing.T) {
	t.Run("new context", func(t *testing.T) {
		e := echo.New()
		pc := NewContext(e.NewContext(nil, nil))
		if pc.status != 0 {
			t.Fatalf("status of context should be 0")
		}
		if pc.identity != nil {
			t.Fatalf("identity of context should be nil")
		}
		if pc.director != nil {
			t.Fatalf("director of context should be nil")
		}
		if pc.resp != nil {
			t.Fatalf("resp of context should be nil")
		}
		if pc.fresh != false {
			t.Fatalf("fresh of context should be false")
		}
		if time.Now().Unix()-pc.createdAt.Unix() > 5 {
			t.Fatalf("created at should be now time")
		}
		ReleaseContext(pc)
	})
}

func TestBodyDump(t *testing.T) {
	t.Run("new body dump response", func(t *testing.T) {
		w := NewBodyDumpResponseWriter()
		if w.body.Len() != 0 {
			t.Fatalf("the body should be empty")
		}
		if len(w.headers) != 0 {
			t.Fatalf("the header should be empty")
		}
		ReleaseBodyDumpResponseWriter(w)
	})
}

func TestServerTiming(t *testing.T) {
	t.Run("server timing", func(t *testing.T) {
		st := ServerTiming{}
		st.Init()
		st.GetRequestStatusStart()
		time.Sleep(time.Millisecond * 5)
		st.GetRequestStatusEnd()
		st.End()
		serverTiming := st.String()
		if len(strings.Split(serverTiming, ",")) != 2 {
			t.Fatalf("get server timing fail")
		}
	})
}