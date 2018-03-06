// Copyright (c) 2016 Blue Jay - MIT License
// Additional changes copyright Richard Phillips - MIT License

// Package uri adds URI shortcuts to the view template.
package uri

import (
	"net/http"
	"path"

	"github.com/richp10/golib/view"
)

// Modify sets BaseURI, CurrentURI, ParentURI, and the GrandparentURI
// variables for use in the templates.
func Modify(w http.ResponseWriter, r *http.Request, v *view.ViewInfo) {
	v.Vars["BaseURI"] = v.BaseURI
	v.Vars["CurrentURI"] = r.URL.Path
	v.Vars["ParentURI"] = path.Dir(r.URL.Path)
	v.Vars["GrandparentURI"] = path.Dir(path.Dir(r.URL.Path))
}
