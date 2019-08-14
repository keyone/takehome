package handlers

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	t.Run("test home handler", func(t *testing.T) {
		req, err := http.NewRequest("GET", "localhost:8080/", nil)
		if err != nil {
			t.Fatalf("could not create request %v", err)
		}
		rec := httptest.NewRecorder()

		homeHandler(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("could not read response: %v", err)
		}

		if msg := string(b); !strings.Contains(msg, "endpoints") {
			t.Error("expected standard header that contains endpoints")
		}

	})
}

func TestEchoHandlerWithoutFile(t *testing.T) {
	t.Run("test echo handler", func(t *testing.T) {
		req, err := http.NewRequest("POST", "localhost:8080/", nil)
		if err != nil {
			t.Fatalf("could not create request %v", err)
		}
		rec := httptest.NewRecorder()
		echoHandler(rec, req)

		res := rec.Result()
		defer res.Body.Close()
	})
}

func createMultipart(t *testing.T, field, filename string) (bytes.Buffer, *multipart.Writer) {
	var b bytes.Buffer
	var err error
	w := multipart.NewWriter(&b)
	var fw io.Writer
	file, errOpen := os.Open(filename)
	if errOpen != nil {
		panic(err)
	}
	if fw, err = w.CreateFormFile(field, file.Name()); err != nil {
		t.Errorf("Error creating writer: %v", err)
	}
	if _, err = io.Copy(fw, file); err != nil {
		t.Errorf("Error with io.Copy: %v", err)
	}
	w.Close()
	return b, w
}

func TestEchoHandlerWithFile(t *testing.T) {
	tt := []struct {
		name     string
		filepath string
		handler  func(http.ResponseWriter, *http.Request)
	}{
		{name: "valid matrix", filepath: "../matrix.csv", handler: echoHandler},
		{name: "invalid matrix", filepath: "../matrix-invalid.csv", handler: invertHandler},
		{name: "invalid non-square matrix", filepath: "../matrix-non-square.csv", handler: flattenHandler},
		{name: "valid matrix", filepath: "../matrix.csv", handler: sumHandler},
		{name: "multiply matrix", filepath: "../matrix.csv", handler: multiplyHandler},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			b, w := createMultipart(t, "file", tc.filepath)

			req, err := http.NewRequest("POST", "localhost:8080/", &b)

			if err != nil {
				t.Fatalf("could not create request %v", err)
			}

			req.Header.Set("Content-Type", w.FormDataContentType())

			rec := httptest.NewRecorder()
			tc.handler(rec, req)

			res := rec.Result()
			defer res.Body.Close()
		})
	}

}
func TestHandler(t *testing.T) {
	srv := httptest.NewServer(Handler())
	defer srv.Close()

	res, err := http.Get(fmt.Sprintf("%s/", srv.URL))
	if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	_, errResponse := ioutil.ReadAll(res.Body)
	if errResponse != nil {
		t.Fatalf("could not read response: %v", err)
	}
}
