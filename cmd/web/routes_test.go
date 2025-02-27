package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/prashkotam/bednbreakfast/internal/config"
)

func TestRoutes(t *testing.T) {

	var app config.Appconfig

	mux := Routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		//Do nothing
	default:
		t.Error(fmt.Sprintf("type is not *chi.Mux, type is %T", v))
	}

}