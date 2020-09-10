// Copyright (c) 2016 Blue Jay - MIT License
// Additional changes copyright Richard Phillips - MIT License

package view

import (
	"html/template"
	"net/http"
)

// extend safely reads the extend list.
func (v *ViewInfo) extend() template.FuncMap {
	v.extendMutex.RLock()
	list := v.extendList
	v.extendMutex.RUnlock()

	return list
}

// modify safely reads the modify list.
func (v *ViewInfo) modify() []ModifyFunc {
	// Get the setter collection
	v.modifyMutex.RLock()
	list := v.modifyList
	v.modifyMutex.RUnlock()

	return list
}

// SetTemplates will set the root and child templates.
func (v *ViewInfo) SetTemplates(rootTemp string, childTemps []string) {
	v.mutex.Lock()
	v.templateCollection = make(map[string]*template.Template)
	v.mutex.Unlock()

	v.rootTemplate = rootTemp
	v.childTemplates = childTemps
}

// ModifyFunc can modify the view before rendering.
type ModifyFunc func(http.ResponseWriter, *http.Request, *ViewInfo)

// SetModifiers will set the modifiers for the View that run
// before rendering.
func (v *ViewInfo) SetModifiers(fn ...ModifyFunc) {
	v.modifyMutex.Lock()
	v.modifyList = fn
	v.modifyMutex.Unlock()
}

// SetFuncMaps will combine all template.FuncMaps into one map and then set the
// them for each template.
// If a func already exists, it is rewritten without a warning.
func (v *ViewInfo) SetFuncMaps(fms ...template.FuncMap) {
	// Final FuncMap
	fm := make(template.FuncMap)

	// Loop through the maps
	for _, m := range fms {
		// Loop through each key and value
		for k, v := range m {
			fm[k] = v
		}
	}

	// Load the plugins
	v.extendMutex.Lock()
	v.extendList = fm
	v.extendMutex.Unlock()
}
