package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllTimestamps(t *testing.T) {
	// GET /ptlist?period=1h&tz=Europe/Athens&t1=20210714T204603Z&t2=20210715T123456Z
	testCases := map[string]struct {
		params     map[string]string
		statusCode int
	}{
		"good params": {
			map[string]string{
				"period": "1h",
				"tz":     "Europe/Athens",
				"t1":     "20210714T204603Z",
				"t2":     "20210715T123456Z",
			},
			http.StatusOK,
		},
	}

	for tc, tp := range testCases {
		req, _ := http.NewRequest("GET", "/", nil)
		q := req.URL.Query()
		for k, v := range tp.params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
		rec := httptest.NewRecorder()
		GetAllTimestamps(rec, req)
		res := rec.Result()
		if res.StatusCode != tp.statusCode {
			t.Errorf("`%v` failed, got %v, expected %v", tc, res.StatusCode, tp.statusCode)
		}
	}
}
