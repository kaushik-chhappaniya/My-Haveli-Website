package handlers

import (
	"net/http"
	"time"

	logger "github.com/kaushik-chhappnaiya/myHaweli/internal/middleware/logger"
)

func (a *App) IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	logger.Debug("MainPageHandler invoked")

	// Clear and set fresh data for this request
	a.ClearAndSetData(nil)
	data := map[string]any{
		"contentTemplate": "indexContent",
		"Title":           "Main Page - Shree Nath Ji's Haweli",
		"subTitle":        "Welcome to the Main Entrance",
		"message":         "This is the main page of the Haweli - the entrance door!",
		"timestamp":       time.Now().Format("2006-01-02 15:04:05"),
		"ImgSrc":          "/static/images/Shreenathji.jpeg",
	}
	// a.Data["PageTemplate"] = "index.html"
	// a.Data["Title"] = "Main Page - Shree Nath Ji's Haweli"
	// a.Data["subTitle"] = "Welcome to the Main Entrance"
	// a.Data["message"] = "This is the main page of the Haweli - the entrance door!"
	// a.Data["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	// a.Data["ImgSrc"] = "/static/images/Shreenathji.jpeg"

	// Add cache-busting headers
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	logger.Info("Rendering index template with image")
	// Use the parsed template from TemplatesMap

	name, tmpl := a.Render("index", data)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := tmpl.ExecuteTemplate(w, name, data); err != nil {
		logger.Errorf("render %q error: %v | defined: %s", name, err, tmpl.DefinedTemplates())
		if err := tmpl.ExecuteTemplate(w, "internalServerError", nil); err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	}
}

// if err := a.Templates.ExecuteTemplate(w, "index", data); err != nil {
// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// 	return
// }
// if tmpl, exists := a.Templates["index"]; exists {
// 	err := tmpl.ExecuteTemplate(w, "index", a.Data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// } else {
// 	http.Error(w, "Template not found", http.StatusInternalServerError)
// }

// log.Printf("Defined templates: %s", a.Templates.DefinedTemplates())

// if a.Templates.Lookup("base") == nil {
//     log.Println("missing template 'base'")
// }
// if a.Templates.Lookup("index") == nil {
//     log.Println("missing template 'index'")
// }
// if a.Templates.Lookup("content") == nil {
//     log.Println("missing template 'content'")
// }

// }
