package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Double(stdin io.Reader, stdout io.Writer) error {
	buf, err := io.ReadAll(stdin)
	if err != nil {
		return err
	}

	stdout.Write(buf)
	stdout.Write(buf)

	return nil
}

func TestDouble(t *testing.T) {
	stdin := bytes.NewBufferString("foo\n")
	stdout := new(bytes.Buffer)

	err := Double(stdin, stdout)
	assert.NoError(t, err)

	expected := []byte("foo\nfoo\n")
	assert.Equal(t, expected, stdout.Bytes())
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-API-VERSION", "1.0")
	b, _ := io.ReadAll(r.Body)
	// helloをbodyと連結させる
	w.Write(append([]byte("hello"), b...))
	w.WriteHeader(http.StatusOK)
}

func TestHandler(t *testing.T) {
	b := strings.NewReader("foo\n")
	req := httptest.NewRequest(http.MethodGet, "http://localhost/ping", b)

	w := httptest.NewRecorder()

	Handler(w, req)

	if got := w.Header().Get("X-API-VERSION"); got != "1.0" {
		t.Errorf("X-API-VERSION = %q, want %q", got, "1.0")
	}

	// body, _ := io.ReadAll(w.Body)
	// if got := string(body); got != "hellofoo\n" {
	// 	t.Errorf("body = %q, want %q", got, "hellofoo\n")
	// }

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
}
