package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"gq", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"ms", "/majors-suites", "GET", []postData{}, http.StatusOK},
	{"sa", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},

	{"post-search-availability", "/search-availability", "POST", []postData{
		{key: "start", value: "24-02-2024"},
		{key: "end", value: "25-02-2024"},
	}, http.StatusOK},

	{"post-search-availability-json", "/search-availability-json", "POST", []postData{
		{key: "start", value: "24-02-2024"},
		{key: "end", value: "25-02-2024"},
	}, http.StatusOK},

	{"make reservation", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "Prashanth"},
		{key: "last_name", value: "Kotamraju"},
		{key: "email", value: "p@g.com"},
		{key: "phome", value: "555-555"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {

	routes := getRoutes()

	testserver := httptest.NewTLSServer(routes)
	defer testserver.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := testserver.Client().Get(testserver.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s. expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
			
			values := url.Values{} //url.Values is a built in type which holds information as a post request for our variable
			
			for _, x := range e.params {

				values.Add(x.key, x.value)
				
			}
			resp, err := testserver.Client().PostForm(testserver.URL + e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s. expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}

			
		
	}
}