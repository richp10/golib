// Copyright (c) 2016 Blue Jay - MIT License
// Additional changes copyright Richard Phillips - MIT License

// Package flash adds the flashes to the view template.
package flash

import (
	"fmt"
	"net/http"

	"sbase-web/lib/flight"

	"github.com/richp10/golib/view"

	flashlib "github.com/richp10/golib/flash"
)

// Modify adds the flashes to the view.
func Modify(w http.ResponseWriter, r *http.Request, v *view.ViewInfo) {
	c := flight.Context(w, r)

	// Get the flashes for the template
	if flashes := c.Sess.Flashes(); len(flashes) > 0 {
		v.Vars["flashes"] = make([]flashlib.Info, len(flashes))
		for i, f := range flashes {
			switch f.(type) {
			case flashlib.Info:
				v.Vars["flashes"].([]flashlib.Info)[i] = f.(flashlib.Info)
			default:
				v.Vars["flashes"].([]flashlib.Info)[i] = flashlib.Info{fmt.Sprint(f), flashlib.Standard}
			}

		}
		c.Sess.Save(r, w)
	}
}
