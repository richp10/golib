// Copyright (c) 2016 Blue Jay - MIT License
// Additional changes copyright Richard Phillips - MIT License

// Package authlevel adds an AuthLevel variable to the view template.
package authlevel

import (
	"net/http"

	"github.com/richp10/golib/view"
)

// Modify sets AuthLevel in the template to auth if the user is authenticated.
// Sets AuthLevel to anon if not authenticated.
func Modify(w http.ResponseWriter, r *http.Request, v *view.ViewInfo) {
	//c := flight.Context(r)

	v.Vars["AuthLevel"] = "anon"
	// Set the AuthLevel to auth if the user is logged in
	/*
	if c.Sess.Values["id"] != nil {
		v.Vars["AuthLevel"] = "auth"
	} else {
		v.Vars["AuthLevel"] = "anon"
	}
	*/

}
