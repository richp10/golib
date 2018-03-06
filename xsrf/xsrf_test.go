// Copyright (c) 2016 Blue Jay - MIT License
// Additional changes copyright Richard Phillips - MIT License

package xsrf_test

import (
	"encoding/base64"
	"fmt"
	"html"
	"net/http"
	"net/http/httptest"
	"os/exec"
	"testing"

	"github.com/richp10/golib/view"
	"github.com/richp10/golib/xsrf"

	"github.com/gorilla/csrf"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMegaChecks(t *testing.T) {

	Convey("Should pass megachecks", t, func() {

		cmd := exec.Command("megacheck")
		res, _ := cmd.Output()
		So(string(res[:]), ShouldBeEmpty)

	})
}

// TestModify ensures token is added to the view.
func TestModify(t *testing.T) {
	viewInfo := &view.ViewInfo{
		BaseURI:   "/",
		Extension: "tmpl",
		Folder:    "testdata/view",
		Caching:   false,
	}

	templates := view.Template{
		Root:     "test",
		Children: []string{},
	}

	authKey := "PzCh6FNAB7/jhmlUQ0+25sjJ+WgcJeKR2bAOtnh9UnfVN+WJSBvY/YC80Rs+rbMtwfmSP4FUSxKPtpYKzKFqFA=="

	// Set up the view
	viewInfo.SetTemplates(templates.Root, templates.Children)

	// Apply the flash modifier
	viewInfo.SetModifiers(
		xsrf.Token,
	)

	// Decode the string
	key, err := base64.StdEncoding.DecodeString(authKey)
	if err != nil {
		t.Fatal(err)
	}

	// Create an instance of view so we can read the variables
	v := viewInfo.New()

	// Mock the HTTP handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := v.Render(w, r)
		if err != nil {
			t.Fatalf("Should not get error: %v", err)
		}
	})

	// Configure the middleware
	cs := csrf.Protect([]byte(key),
		csrf.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Fatal("invalidHandler should not be called.")
		})),
		csrf.FieldName("_token"),
		csrf.Secure(false),
	)(handler)

	// Simulate a request
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Mock the request
	cs.ServeHTTP(w, r)

	// Need to unescape since the string could have characters that were escaped
	// May fail occasionally if you don't
	actual := html.UnescapeString(w.Body.String())
	expected := fmt.Sprintf(`<div>%v</div>`, v.Vars["token"])

	if actual != expected {
		t.Fatalf("\nactual: %v\nexpected: %v", actual, expected)
	}
}

// TestModify ensures token fails.
func TestModifyFail(t *testing.T) {
	viewInfo := &view.ViewInfo{
		BaseURI:   "/",
		Extension: "tmpl",
		Folder:    "testdata/view",
		Caching:   false,
	}

	templates := view.Template{
		Root:     "test",
		Children: []string{},
	}

	authKey := "PzCh6FNAB7/jhmlUQ0+25sjJ+WgcJeKR2bAOtnh9UnfVN+WJSBvY/YC80Rs+rbMtwfmSP4FUSxKPtpYKzKFqFA=="

	// Set up the view
	viewInfo.SetTemplates(templates.Root, templates.Children)

	// Apply the flash modifier
	viewInfo.SetModifiers(
		xsrf.Token,
	)

	// Decode the string
	key, err := base64.StdEncoding.DecodeString(authKey)
	if err != nil {
		t.Fatal(err)
	}

	// Create an instance of view so we can read the variables
	v := viewInfo.New()

	// Mock the HTTP handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := v.Render(w, r)
		if err != nil {
			t.Fatalf("Should not get error: %v", err)
		}
	})

	// Configure the middleware
	cs := csrf.Protect([]byte(key),
		csrf.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Fatal("invalidHandler should not be called.")
		})),
		csrf.FieldName("_token"),
		csrf.Secure(false),
	)(handler)

	// Simulate a request
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Mock the request
	cs.ServeHTTP(w, r)

	// Need to unescape since the string could have characters that were escaped
	// May fail occasionally if you don't
	actual := html.UnescapeString(w.Body.String())
	expected := fmt.Sprintf(`<div>%v</div>`, "nil")

	if actual == expected {
		t.Fatalf("\nactual: %v\nexpected: %v", actual, expected)
	}
}
