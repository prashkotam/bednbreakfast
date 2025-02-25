package render

import (
	"net/http"
	"testing"

	"github.com/prashkotam/bednbreakfast/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var tdd models.TemplateData
	r, err := getSession()
	if err != nil {
		t.Error(err)
		return
		}
	if r == nil {
			t.Error("getSession returned nil request")
			return
		}
	session.Put(r.Context(), "flash", "123")
	result := AddDefaultData(&tdd, r)

	if result.Flash != "123" {
		t.Error("Flash value of 123 not found in session")
	}

}

func getSession() (*http.Request, error) {

	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}
	ctx := r.Context()
	ctx, err = session.Load(ctx, r.Header.Get("X-Session"))
	if err != nil {
		return nil, err
	}
	r = r.WithContext(ctx)
	return r, nil
}