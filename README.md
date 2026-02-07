# 🕉️ Shree Nathji Haweli - Temple Website

A modern, responsive temple website built with Go backend and HTML templates, featuring a comprehensive admin system for content management.

## 📋 Table of Contents
- [Features](#features)
- [Admin System](#admin-system)
- [AJAX Implementation](#ajax-implementation)
- [Architecture Overview](#architecture-overview)
- [Quick Start](#quick-start)
- [API Endpoints](#api-endpoints)
- [Development](#development)

## ✨ Features

### 🎨 **Frontend Features**
- **Responsive Design**: Bootstrap 5 with mobile-first approach
- **Dark/Light Mode**: Theme toggle with localStorage persistence
- **Modern UI**: Beautiful cards, animations, and smooth transitions
- **Temple Branding**: Hindi/Sanskrit text with appropriate styling
- **Interactive Elements**: Smooth scrolling, hover effects, back-to-top button

### 🔐 **Admin System**
- **Modal-based Login**: AJAX authentication through navbar
- **Directors Management**: Complete CRUD operations for temple directors
- **Data Validation**: Client-side and server-side validation
- **Bulk Operations**: Select multiple directors for batch actions
- **Preview System**: Preview changes before saving
- **Statistics Dashboard**: Real-time counts and system status

### 🏗️ **Backend Features**
- **Go Server**: High-performance HTTP server with proper timeouts
- **Template System**: Pre-compiled templates for fast rendering
- **JSON Storage**: File-based data storage with thread-safe operations
- **Structured Logging**: Comprehensive logging with different levels
- **Middleware**: CORS, logging, and authentication middleware

## 🎯 Admin System

### **Login Process**
1. **Access**: Click "Sign In" button in navbar
2. **Modal Login**: Enter credentials in Bootstrap modal
3. **AJAX Authentication**: Credentials sent via JSON API
4. **Redirect**: Successful login redirects to admin dashboard

### **Demo Credentials**
```
Email: admin@haweli.com
Password: admin123

Alternative accounts:
- director@haweli.com / director123
- manager@haweli.com / manager123
```

### **Dashboard Features**
- **Directors Table**: Editable form with all director information
- **Add Directors**: Dynamic row addition with validation
- **Delete Directors**: Individual or bulk deletion with confirmation
- **Form Validation**: Required fields, proper data types
- **Preview Modal**: Review changes before saving
- **Statistics Cards**: Live director count and system status

## 🔄 AJAX Implementation

### **Login AJAX Call**

**Frontend (header.html)**:
```javascript
function handleAdminLogin() {
    const email = document.getElementById('adminEmail').value;
    const password = document.getElementById('adminPassword').value;
    
    fetch('/admin', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            email: email,
            password: password
        })
    })
    .then(response => response.json())
    .then(data => {
        if (data.success) {
            window.location.href = '/admin/dashboard';
        } else {
            alert(data.message);
        }
    })
    .catch(error => {
        console.error('Login error:', error);
        alert('Login error. Please try again.');
    });
}
```

**Backend (adminHandler.go)**:
```go
func (a *App) processLogin(w http.ResponseWriter, r *http.Request) {
    var loginReq LoginRequest
    
    if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(LoginResponse{
            Success: false,
            Message: "Invalid request format",
        })
        return
    }
    
    if a.validateCredentials(loginReq.Email, loginReq.Password) {
        json.NewEncoder(w).Encode(LoginResponse{
            Success: true,
            Message: "Login successful",
        })
    } else {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(LoginResponse{
            Success: false,
            Message: "Invalid credentials",
        })
    }
}
```

### **AJAX Flow Explanation**

1. **Frontend Request**:
   - User clicks "Sign In" button
   - JavaScript collects form data
   - Sends POST request to `/admin` with JSON payload
   - Shows loading state on button

2. **Backend Processing**:
   - Go handler receives JSON request
   - Validates credentials against hardcoded map
   - Returns JSON response with success/failure

3. **Frontend Response**:
   - JavaScript processes JSON response
   - If successful: redirects to `/admin/dashboard`
   - If failed: shows error message
   - Restores button state

### **Why AJAX vs Form Submission?**

**AJAX Benefits**:
- ✅ **No Page Reload**: Smooth user experience
- ✅ **Custom Loading States**: Show progress indicators
- ✅ **Error Handling**: Display errors without navigation
- ✅ **JSON Communication**: Structured data exchange
- ✅ **Modal Integration**: Works perfectly with Bootstrap modals

**Traditional Form Drawbacks**:
- ❌ **Page Reload**: Jarring user experience
- ❌ **Limited Feedback**: Hard to show loading states
- ❌ **Modal Issues**: Form submission closes modal
- ❌ **Error Display**: Requires separate error pages

## 🏗️ Architecture Overview

### **Template-Based Architecture**

**Why Templates over React/SPA?**
- ✅ **SEO Friendly**: Server-side rendering for better search engine indexing
- ✅ **Performance**: Faster initial page load, smaller bundle size
- ✅ **Simplicity**: Less complexity, easier to debug and maintain
- ✅ **Progressive Enhancement**: Works without JavaScript
- ✅ **Temple Content**: Perfect for content-heavy, informational websites

### **Core Components**

#### **1. appStruct.go — Core Architecture**
- **App Struct**: Central application structure
- **Template Management**: Pre-compiled templates for performance
- **Base Templates**: `base.html`, `header.html`, `footer.html`
- **Page Templates**: Individual page templates that extend base

#### **2. main.go — Application Startup**
1. Initialize App struct with configuration
2. Load and parse base templates once at startup
3. Precompute page-specific template clones
4. Set up HTTP router with middleware
5. Start server with production-ready timeouts

#### **3. Template Flow**
**During Startup**:
1. Parse base templates: `base`, `header`, `footer`
2. For each page: Clone base + Parse page templates
3. Store precompiled templates in memory

**During Request**:
1. Handler selects correct template
2. Execute template with data
3. Return rendered HTML

### **Data Storage**
- **JSON Files**: Simple file-based storage for directors data
- **Thread-Safe**: Mutex locks for concurrent access
- **Utils Package**: Centralized file operations

## 🚀 Quick Start

### **Prerequisites**
- Go 1.19 or higher
- Git

### **Installation**
```bash
# Clone repository
git clone <repository-url>


# Install dependencies
go mod tidy

# Run development server
go run cmd/server/main.go

# Or use Air for hot reload
air
```

### **Access Application**
- **Main Site**: http://localhost:8081
- **Admin Login**: Click "Sign In" in navbar
- **Admin Dashboard**: http://localhost:8081/admin/dashboard (after login)

## 🛠️ API Endpoints

### **Public Endpoints**
```
GET  /              - Home page
GET  /home          - Home page (alternative)
GET  /aboutUs       - About us page
GET  /directorial   - Directors listing
GET  /static/*      - Static assets (CSS, JS, images)
```

### **Admin Endpoints**
```
POST /admin                  - AJAX login authentication
GET  /admin/dashboard        - Admin dashboard (requires auth)
POST /admin/dashboard        - Update directors data
```

## 💻 Development

### **Project Structure**
```
shreeNathJi/
├── cmd/server/main.go              # Application entry point
├── internal/
│   ├── handlers/                   # HTTP handlers
│   │   ├── AppStruct.go           # App structure & template management
│   │   ├── adminHandler.go        # Admin login & dashboard
│   │   ├── indexPage.go           # Home page handler
│   │   ├── homepage.go            # Homepage handler
│   │   ├── aboutPage.go           # About page handler
│   │   └── diretorialPage.go      # Directors page handler
│   ├── middleware/                 # HTTP middleware
│   │   └── logger.go              # Logging middleware & functions
│   ├── routes/                     # Route definitions
│   │   └── routes.go              # HTTP routes setup
│   ├── utils/                      # Utility functions
│   │   └── jsonUtils.go           # JSON file operations
│   └── database/                   # Data storage
│       └── directorial.json       # Directors data
├── ui/
│   ├── templates/                  # HTML templates
│   │   ├── base/                  # Base templates
│   │   │   ├── base.html          # Main layout
│   │   │   ├── header.html        # Navigation & login modal
│   │   │   └── footer.html        # Footer & scripts
│   │   ├── index.html             # Home page template
│   │   ├── homepage.html          # Homepage template
│   │   ├── directorial.html       # Directors listing
│   │   └── adminDashboard.html    # Admin interface
│   └── static/                     # Static assets
│       ├── css/styles.css         # Main stylesheet
│       ├── js/main.js             # JavaScript functionality
│       └── images/                # Image assets
├── logs/                           # Log files (auto-generated)
├── go.mod                         # Go module definition
├── go.sum                         # Go module checksums
└── README.md                      # This file
```

---

## 📝 What I Implemented

### **Complete Admin System**

**1. Modal-Based Authentication**:
- Bootstrap modal triggered from navbar
- AJAX login with JSON request/response
- No page reload, smooth user experience
- Error handling with proper feedback

**2. Directors Management Dashboard**:
- Full CRUD operations for temple directors
- Editable table with form validation
- Add/remove directors dynamically
- Bulk operations with select all functionality
- Preview changes before saving

**3. Backend API**:
- RESTful JSON API for authentication
- File-based data storage with thread safety
- Proper HTTP status codes and responses
- Error handling and validation

**4. User Experience Enhancements**:
- Loading states during operations
- Success/error notifications
- Confirmation dialogs for destructive actions
- Statistics dashboard with live counts

The system demonstrates modern web development practices with a traditional server-rendered architecture, perfect for a temple website that needs to be SEO-friendly and accessible to all users.

# ✅ **3. router.go — URL → Handler Mapping**

### What it does:

1.  Creates an `http.ServeMux`.
2.  Registers all routes:
    *   `/` → IndexHandler
    *   `/about` → AboutHandler
3.  Serves static files (`/static/...`)
4.  Returns the mux to main.go

### Purpose:

*   Keep routing clean and separate.
*   Handler functions are methods on `App` so they get access to templates.

***

# ✅ **4. Handlers (indexHandler.go, aboutHandler.go, etc.)**

Each file handles one route.

### What each handler does:

1.  Constructs a simple `ViewData`
    *   sets page title
    *   year auto-filled by Render()

2.  Calls:
        name, tmpl := app.Render("index", data)
        tmpl.ExecuteTemplate(w, name, data)

3.  Uses the **precomputed page template**:
    *   `"index"` template = base + page blocks merged

### Purpose:

*   Handlers do almost nothing except:
    *   gather request-specific data
    *   call the correct precomputed template
*   Very small and easy to maintain.

***

# 🔥 **5. How the Templates Flow (important)**

### During Startup:

1.  Parse base templates:  
    `base`, `header`, `footer`
2.  For each page (like `index.html`):
    *   Clone base
    *   Parse that page’s templates
    *   Store clone in a map (`map[string]*template.Template`)

### During Request:

1.  Handler chooses which template:
    `"index"` or `"about"`
2.  Call ExecuteTemplate with wrapper:
    `"index"` calls `"base"`
3.  `"base"` calls:
    *   header
    *   `block "title"`
    *   `block "content"`
    *   footer

### Result:

**Fast, conflict‑free, production‑grade rendering.**

***

# ⭐ Putting It All Together (conceptual flow)

1.  **Startup**
    *   Build `App`
    *   Load base templates
    *   Precompute per-page clones
    *   Start server

2.  **Request comes in**
    *   Router directs to handler

3.  **Handler**
    *   Creates ViewData
    *   Calls Render() → picks precomputed template
    *   Executes `"index"` or `"about"` wrapper

4.  **Template Execution**
    *   `"index"` → `"base"`
    *   `"base"` → header + content block + footer

No parsing, no redefinition conflicts, super fast.

***

# 🎯 Summary (Very Short)

*   **appStruct.go**: loads & precomputes all templates at startup.
*   **main.go**: initializes App, loads templates, starts server.
*   **router.go**: maps URL paths to handlers.
*   **handlers**: build per-page data and execute the correct template.
*   **templates**: page files override blocks (`title`, `content`) in base.

***

If you want, I can next produce a **diagram** showing the whole flow visually, or a **minimal runnable repo layout**.
