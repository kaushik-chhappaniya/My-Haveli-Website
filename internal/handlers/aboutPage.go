package handlers

import (
	"net/http"
	"time"

	logger "github.com/kaushik-chhappaniya/myHaweli/middleware/logger"
	"github.com/kaushik-chhappaniya/myHaweli/utils"
)

var directorialRead utils.Store

func init() {
	logger.Info("AboutPageHandler initialized.")
	directorialRead = utils.Store{
		FilePath: "./internal/database/directorial.json",
	}

}

func (a *App) AboutPageHandler(w http.ResponseWriter, r *http.Request) {
	logger.Debug("About page handler invoked.")
	directorialData, err := directorialRead.ReadAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Debugf("%v", directorialData)
	data := map[string]any{
		"contentTemplate": "aboutContent",
		"Title":           "Mari Haveli - About Us",
		"directorsData":   directorialData["directors"],
		"timestamp":       time.Now().Format("2006-01-02 15:04:05"),
	}
	name, tmpl := a.Render("aboutUs", data)
	if err := tmpl.ExecuteTemplate(w, name, data); err != nil {
		logger.Errorf("render %q error: %v | defined: %s", name, err, tmpl.DefinedTemplates())
		if err := tmpl.ExecuteTemplate(w, "internalServerError", nil); err != nil {
			if err := tmpl.ExecuteTemplate(w, "internalServerError", nil); err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}
	}

}
