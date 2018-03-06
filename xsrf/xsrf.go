// Copyright (c) 2016 Blue Jay - MIT License
// Additional changes copyright Richard Phillips - MIT License

// Package xsrf is a container for the gorilla csrf package
package xsrf

import (
	"net/http"

	"github.com/richp10/golib/view"

	"github.com/gorilla/csrf"
)

// Info holds the config.
type Info struct {
	AuthKey string
	Secure  bool
}

// Token sets token in the template to the CSRF token.
func Token(w http.ResponseWriter, r *http.Request, v *view.ViewInfo) {
	v.Vars["token"] = csrf.Token(r)
}
