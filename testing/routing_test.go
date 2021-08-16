package testing

import (
	"bytes"
	"encoding/json"
	"go-article/controller"
	"go-article/entity"
	"go-article/responsegraph"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetArticle(t *testing.T) {
	req, err := http.NewRequest("GET", "/articles", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controller.CommonController{}.GetArticle(w, r)
	})

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[]entity.Article`
	articles := []entity.Article{}
	_ = json.Unmarshal([]byte(rr.Body.String()), &articles)

	if reflect.TypeOf(articles).String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			reflect.TypeOf(articles).String(), expected)
	}
}

func TestAddArticle(t *testing.T) {
	var jsonStr = []byte(`{"author": "Fadhil","title": "Artikel","body": "Artikel baru terbit"}`)

	req, err := http.NewRequest("POST", "/articles", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controller.CommonController{}.AddArticle(w, r)
	})
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `responsegraph.ResponseGenericIn`
	articles := responsegraph.ResponseGenericIn{}
	_ = json.Unmarshal([]byte(rr.Body.String()), &articles)

	if reflect.TypeOf(articles).String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			reflect.TypeOf(articles).String(), expected)
	}

}
