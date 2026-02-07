package handlers

import (
	"fmt"
	"html/template"
	"maps"

	"path/filepath"
)

/*
Initialising App struct to hold templates and shared resources. This will be used across various handlers.
This passes the template instance to handlers for rendering.
*/

type App struct {
	BaseTemplate *template.Template            // parsed base + partials
	Pages        map[string]*template.Template // precomputed clone per pages
	Data         map[string]any                // Shared but cleaned per request

	TemplateDir string
}

// MustLoadBase parses base and partials once at startup.
func (a *App) MustLoadBase() {
	baseDir := filepath.Join(a.TemplateDir, "base")

	a.BaseTemplate = template.Must(
		template.New("site").ParseFiles(
			filepath.Join(baseDir, "header.html"),
			filepath.Join(baseDir, "footer.html"),
			filepath.Join(baseDir, "base.html"),
		),
	)
}

// MustPrecomputePages clones the base and parses one page file per clone.
// It avoids cross-file redefinition conflicts for `block "content"`.
func (a *App) MustPrecomputePages() {
	type pageDef struct {
		Name string // wrapper name (e.g., "index", "about")
		File string // file name (e.g., "index.html")
	}

	defs := []pageDef{
		{Name: "index", File: "index.html"},
		{Name: "aboutUs", File: "aboutUs.html"},
		{Name: "adminDashboard", File: "adminDashboard.html"},
		{Name: "internalServerError", File: "internalServerError.html"},
		{Name: "notFound", File: "notFound.html"},
	}

	a.Pages = make(map[string]*template.Template, len(defs))

	for _, d := range defs {
		clone, err := a.BaseTemplate.Clone()
		if err != nil {
			panic(fmt.Errorf("clone base failed for %s: %w", d.Name, err))
		}
		if _, err := clone.ParseFiles(filepath.Join(a.TemplateDir, d.File)); err != nil {
			panic(fmt.Errorf("parse page failed for %s: %w", d.Name, err))
		}
		// Store in map a[PageName[pageDef[Name]]] = clone template. map of templates with key as page name
		a.Pages[d.Name] = clone
	}
}

// Render renders a precomputed page by wrapper name ("index", "about").
func (a *App) Render(name string, data any) (string, *template.Template) {
	t, ok := a.Pages[name]
	if !ok {
		panic(fmt.Errorf("template not found: %s", name))
	}
	// Enrich common fields here if you like:

	return name, t
}

// type App struct {
// 	Templates    *template.Template
// 	PageTemplate string
// 	Data         map[string]any // Shared but cleaned per request
// }

// NewApp creates a new App instance
// func NewApp() *App {
// 	return &App{
// 		Templates: nil,
// 		Data:      make(map[string]any),
// 	}
// }

// ClearAndSetData clears existing data and sets new data
func (a *App) ClearAndSetData(data map[string]any) {
	// Clear existing data
	if a.Data == nil {
		a.Data = make(map[string]any)
	}
	if data == nil {
		for k := range a.Data {
			delete(a.Data, k)
		}
	}
	// Set new data
	if data != nil {
		maps.Copy(a.Data, data)
	}
}

// render executes the template specified in data["PageTemplate"] with the provided data
// func (a *App) render(w http.ResponseWriter, data map[string]any) {
// 	// Get template name from data, default to "base" if not specified
// 	templateName := "base"
// 	if tmpl, ok := data["PageTemplate"].(string); ok && tmpl != "" {
// 		templateName = tmpl
// 	}

// 	// Execute the specified template
// 	err := a.Templates.ExecuteTemplate(w, templateName, data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }
