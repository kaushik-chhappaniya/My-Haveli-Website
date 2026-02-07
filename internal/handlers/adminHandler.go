package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	logger "github.com/kaushik-chhappaniya/myHaweli/middleware/logger"
	"github.com/kaushik-chhappaniya/myHaweli/utils"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}

type Director struct {
	SrNo  int    `json:"srNo"`
	Name  string `json:"name"`
	Phone string `json:"Phone"`
	Role  string `json:"Role"`
}

type DirectorsData struct {
	Directors []Director `json:"directors"`
}

// AdminHandler handles admin login via AJAX from modal
func (a *App) AdminHandler(w http.ResponseWriter, r *http.Request) {
	logger.Debug("AdminHandler invoked.")

	switch r.Method {
	case "POST":
		a.processLogin(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// AdminDashboardHandler shows the admin dashboard with directors table
func (a *App) AdminDashboardHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("AdminDashboardHandler invoked")

	// TODO: Add session/token validation here

	switch r.Method {
	case "GET":
		a.showAdminDashboard(w, r)
	case "POST":
		a.updateDirectorsData(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (a *App) processLogin(w http.ResponseWriter, r *http.Request) {
	logger.Info("Processing login request")

	// Handle both form data and JSON
	var email, password string

	contentType := r.Header.Get("Content-Type")
	if contentType == "application/json" {
		var loginReq LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
			logger.Error("Failed to decode JSON login request: " + err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(LoginResponse{
				Success: false,
				Message: "Invalid request format",
			})
			return
		}
		email = loginReq.Email
		password = loginReq.Password
	} else {
		email = r.FormValue("email")
		password = r.FormValue("password")
	}

	// Validate credentials
	if a.validateCredentials(email, password) {
		// TODO: Generate actual JWT token
		token := "admin_token_" + strconv.FormatInt(time.Now().Unix(), 10)

		if contentType == "application/json" {
			// AJAX request - return JSON response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(LoginResponse{
				Success: true,
				Message: "Login successful",
				Token:   token,
			})
		} else {
			// Form request - redirect
			http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
		}
		logger.Info("Admin login successful for email: " + email)
	} else {
		if contentType == "application/json" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(LoginResponse{
				Success: false,
				Message: "Invalid email or password",
			})
		} else {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`
				<html><body>
					<h1>❌ Login Failed</h1>
					<p>Invalid email or password</p>
					<a href="/">Go Home</a>
				</body></html>
			`))
		}
		logger.Warn("Failed login attempt for email: " + email)
	}
}

func (a *App) showAdminDashboard(w http.ResponseWriter, r *http.Request) {
	logger.Info("Displaying admin dashboard")

	// Load directors data
	store := &utils.Store{
		FilePath: "./internal/database/directorial.json",
	}

	rawData, err := store.ReadAll()
	if err != nil {
		logger.Error("Error reading directorial file: " + err.Error())
		http.Error(w, "Error loading directors data", http.StatusInternalServerError)
		return
	}

	// Convert to directors structure
	var directorsData DirectorsData
	directorsBytes, _ := json.Marshal(rawData)
	json.Unmarshal(directorsBytes, &directorsData)

	// Prepare template data
	data := map[string]any{
		"directors": directorsData.Directors,
		"pageTitle": "Admin Dashboard - Manage Directors",
	}

	a.ClearAndSetData(data)
	name, tmpl := a.Render("adminDashboard", data)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(w, name, data); err != nil {
		logger.Errorf("Error rendering admin dashboard: %v", err)
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
	}
}

func (a *App) updateDirectorsData(w http.ResponseWriter, r *http.Request) {
	logger.Info("Processing directors data update")

	if err := r.ParseForm(); err != nil {
		logger.Error("Error parsing form: " + err.Error())
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	var directors []Director

	// Parse form data - directors are submitted as arrays
	names := r.Form["name[]"]
	phones := r.Form["phone[]"]
	roles := r.Form["role[]"]
	srNos := r.Form["srNo[]"]

	for i := range names {
		if i < len(phones) && i < len(roles) && i < len(srNos) {
			srNo, _ := strconv.Atoi(srNos[i])
			directors = append(directors, Director{
				SrNo:  srNo,
				Name:  names[i],
				Phone: phones[i],
				Role:  roles[i],
			})
		}
	}

	// Prepare data structure for JSON file
	directorsData := DirectorsData{
		Directors: directors,
	}

	// Convert to map[string]interface{} for utils.Store
	dataBytes, _ := json.Marshal(directorsData)
	var dataMap map[string]any
	json.Unmarshal(dataBytes, &dataMap)

	// Save to JSON file
	store := &utils.Store{
		FilePath: "./internal/database/directorial.json",
	}

	if err := store.WriteAll(dataMap); err != nil {
		logger.Error("Error updating directorial file: " + err.Error())
		http.Error(w, "Error saving directors data", http.StatusInternalServerError)
		return
	}

	logger.Infof("Successfully updated directors data with %d records", len(directors))

	// Redirect back to dashboard with success message
	http.Redirect(w, r, "/admin/dashboard?updated=true", http.StatusSeeOther)
}

// validateCredentials checks the provided email and password
func (a *App) validateCredentials(email, password string) bool {
	// TODO: Implement actual credential validation with hashed passwords
	validCredentials := map[string]string{
		"admin@haweli.com":    "admin123",
		"director@haweli.com": "director123",
		"manager@haweli.com":  "manager123",
	}

	if storedPassword, exists := validCredentials[email]; exists {
		return storedPassword == password
	}
	return false
}
