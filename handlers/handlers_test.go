package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func strPtr(s string) *string { return &s }

func testRequest(t *testing.T, ts *httptest.Server, method string,
	path string, contentType *string, body *string) (*http.Response, string) {
	var url_body io.Reader
	if body != nil {
		url_body = strings.NewReader(*body)
	}
	client := ts.Client()
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	req, err := http.NewRequest(method, ts.URL+path, url_body)
	require.NoError(t, err)
	if contentType != nil {
		req.Header.Set("Content-Type", *contentType)
	}

	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	return resp, string(respBody)
}
func TestRouter(t *testing.T) {
	RegisterHandlers()
	ts := httptest.NewServer(Main_router)
	defer ts.Close()
	type want struct {
		code        int
		contentType string
	}
	type dataTest struct {
		url    string
		body   *string
		want   want
		method string
	}
	var testTable = []dataTest{
		{"/", strPtr("https://www.yandex.ru"), want{201, "text/plain"}, "POST"},
	}
	for _, v := range testTable {
		resp, shortUrl := testRequest(t, ts, v.method, v.url, &v.want.contentType, v.body)
		assert.Equal(t, v.want.code, resp.StatusCode)
		parsedURL, err := url.Parse(shortUrl)
		require.NoError(t, err)
		resp_get, _ := testRequest(t, ts, "GET", parsedURL.Path, &v.want.contentType, nil)
		assert.Equal(t, 307, resp_get.StatusCode)
	}
}
