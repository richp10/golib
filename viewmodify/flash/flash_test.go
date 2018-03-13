// Copyright (c) 2016 Blue Jay - MIT License
// Additional changes copyright Richard Phillips - MIT License

package flash_test

import (
	/*
	"github.com/richp10/golib/view"
	"github.com/gorilla/sessions"
	"net/http/httptest"
	"fmt"
	"github.com/spf13/viper"
	"os"
	*/
)

/*
import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"testing"

	"github.com/richp10/golib/env"
	"github.com/richp10/golib/flash"
	"github.com/richp10/golib/session"
	"github.com/richp10/golib/view"
	flashmod "github.com/richp10/golib/viewmodify/flash"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"

	"github.com/gorilla/sessions"
)

func TestMegaChecks(t *testing.T) {

	Convey("Should pass megachecks", t, func() {

		cmd := exec.Command("megacheck")
		res, _ := cmd.Output()
		So(string(res[:]), ShouldBeEmpty)

	})
}

// TestModify ensures flashes are added to the view.
func TestModify(t *testing.T) {

	Convey("Should Allow flashes to be added to a view", t, func() {

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

		options := sessions.Options{
			Path:     "/",
			Domain:   "",
			MaxAge:   28800,
			Secure:   false,
			HttpOnly: true,
		}

		s := session.Info{
			AuthKey:    "PzCh6FNAB7/jhmlUQ0+25sjJ+WgcJeKR2bAOtnh9UnfVN+WJSBvY/YC80Rs+rbMtwfmSP4FUSxKPtpYKzKFqFA==",
			EncryptKey: "3oTKCcKjDHMUlV+qur2Ve664SPpSuviyGQ/UqnroUD8=",
			CSRFKey:    "xULAGF5FcWvqHsXaovNFJYfgCt6pedRPROqNvsZjU18=",
			Name:       "sess",
			Options:    options,
		}

		// Set up the view
		viewInfo.SetTemplates(templates.Root, templates.Children)

		// Apply the flash modifier
		viewInfo.SetModifiers(
			flashmod.Modify,
		)

		// Set up the session cookie store
		s.SetupConfig()

		// Simulate a request
		w := httptest.NewRecorder()
		r, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		text := "Success test."

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			v := viewInfo.New()

			// Get the session
			sess, _ := s.Instance(r)

			// Add flashes to the session
			sess.AddFlash(flash.Info{text, flash.Success})
			sess.Save(r, w)

			err := v.Render(w, r)
			if err != nil {
				t.Fatalf("Should not get error: %v", err)
			}
		})

		handler.ServeHTTP(w, r)

		actual := w.Body.String()
		expected := fmt.Sprintf(`<div class="%v">%v</div>`, flash.Success, text)

		if actual != expected {
			t.Fatalf("\nactual: %v\nexpected: %v", actual, expected)
		}

	})
}

// TestModify ensures flashes are not displayed on the page.
func TestModifyFail(t *testing.T) {
	viewInfo := &view.ViewInfo{
		BaseURI:   "/",
		Extension: "tmpl",
		Folder:    "testdata/view",
		Caching:   false,
	}

	templates := view.Template{
		Root:     "test_fail",
		Children: []string{},
	}

	options := sessions.Options{
		Path:     "/",
		Domain:   "",
		MaxAge:   28800,
		Secure:   false,
		HttpOnly: true,
	}

	s := session.Info{
		AuthKey:    "PzCh6FNAB7/jhmlUQ0+25sjJ+WgcJeKR2bAOtnh9UnfVN+WJSBvY/YC80Rs+rbMtwfmSP4FUSxKPtpYKzKFqFA==",
		EncryptKey: "3oTKCcKjDHMUlV+qur2Ve664SPpSuviyGQ/UqnroUD8=",
		CSRFKey:    "xULAGF5FcWvqHsXaovNFJYfgCt6pedRPROqNvsZjU18=",
		Name:       "sess",
		Options:    options,
	}

	// Set up the view
	viewInfo.SetTemplates(templates.Root, templates.Children)

	// Apply the flash modifier
	viewInfo.SetModifiers(
		flashmod.Modify,
	)

	// Set up the session cookie store
	s.SetupConfig()

	// Simulate a request
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	text := "Success test."

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := viewInfo.New()

		// Get the session
		sess, _ := s.Instance(r)

		// Add flashes to the session
		sess.AddFlash(flash.Info{text, flash.Success})
		sess.Save(r, w)

		err := v.Render(w, r)
		if err != nil {
			t.Fatalf("Should not get error: %v", err)
		}
	})

	handler.ServeHTTP(w, r)

	actual := w.Body.String()
	expected := "Failure!"

	if actual != expected {
		t.Fatalf("\nactual: %v\nexpected: %v", actual, expected)
	}

}

// TestFlashDefault ensures flashes are added to the view even if a plain text
// message is added to flashes instead of a flash.Info type
func
TestFlashDefault(t *testing.T) {
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

	options := sessions.Options{
		Path:     "/",
		Domain:   "",
		MaxAge:   28800,
		Secure:   false,
		HttpOnly: true,
	}

	s := session.Info{
		AuthKey:    "PzCh6FNAB7/jhmlUQ0+25sjJ+WgcJeKR2bAOtnh9UnfVN+WJSBvY/YC80Rs+rbMtwfmSP4FUSxKPtpYKzKFqFA==",
		EncryptKey: "3oTKCcKjDHMUlV+qur2Ve664SPpSuviyGQ/UqnroUD8=",
		CSRFKey:    "xULAGF5FcWvqHsXaovNFJYfgCt6pedRPROqNvsZjU18=",
		Name:       "sess",
		Options:    options,
	}

	// Set up the view
	viewInfo.SetTemplates(templates.Root, templates.Children)

	// Apply the flash modifier
	viewInfo.SetModifiers(
		flashmod.Modify,
	)

	// Set up the session cookie store
	s.SetupConfig()

	// Simulate a request
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	text := "Just a string."

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := viewInfo.New()

		// Get the session
		sess, _ := s.Instance(r)

		// Add flashes to the session
		sess.AddFlash(text)
		sess.Save(r, w)

		err := v.Render(w, r)
		if err != nil {
			t.Fatalf("Should not get error: %v", err)
		}
	})

	handler.ServeHTTP(w, r)

	actual := w.Body.String()
	expected := fmt.Sprintf(`<div class="%v">%v</div>`, flash.Standard, text)

	if actual != expected {
		t.Fatalf("\nactual: %v\nexpected: %v", actual, expected)
	}

}

// TestNonStringFlash ensures flashes do not error when added with a non-standard type.
func
TestNonStringFlash(t *testing.T) {
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

	options := sessions.Options{
		Path:     "/",
		Domain:   "",
		MaxAge:   28800,
		Secure:   false,
		HttpOnly: true,
	}

	s := session.Info{
		AuthKey:    "PzCh6FNAB7/jhmlUQ0+25sjJ+WgcJeKR2bAOtnh9UnfVN+WJSBvY/YC80Rs+rbMtwfmSP4FUSxKPtpYKzKFqFA==",
		EncryptKey: "3oTKCcKjDHMUlV+qur2Ve664SPpSuviyGQ/UqnroUD8=",
		CSRFKey:    "xULAGF5FcWvqHsXaovNFJYfgCt6pedRPROqNvsZjU18=",
		Name:       "sess",
		Options:    options,
	}

	// Set up the view
	viewInfo.SetTemplates(templates.Root, templates.Children)

	// Apply the flash modifier
	viewInfo.SetModifiers(
		flashmod.Modify,
	)

	// Set up the session cookie store
	s.SetupConfig()

	// Simulate a request
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	text := 123

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := viewInfo.New()

		// Get the session
		sess, _ := s.Instance(r)

		// Add flashes to the session
		sess.AddFlash(text)
		sess.Save(r, w)

		err := v.Render(w, r)
		if err != nil {
			t.Fatalf("Should not get error: %v", err)
		}
	})

	handler.ServeHTTP(w, r)

	actual := w.Body.String()
	expected := fmt.Sprintf(`<div class="%v">%v</div>`, flash.Standard, text)

	if actual != expected {
		t.Fatalf("\nactual: %v\nexpected: %v", actual, expected)
	}

}

func TestMain(m *testing.M) {
	if viper.GetString("TEST") != "asdfadsfasdfasdf" {
		env.Load()
	}
	m.Run()
	code := m.Run()
	os.Exit(code)
}
*/