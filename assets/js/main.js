/**
 * Main JavaScript file for Shree Nathji Haweli Website
 * Handles theme switching, navigation, and interactive features
 */

// ============================================================================
// THEME MANAGEMENT
// ============================================================================

/**
 * Toggle between dark and light themes
 */
function toggleTheme() {
  console.log("🔄 Toggle theme called");
  const body = document.body;
  const themeIcon = document.getElementById("theme-icon");

  const currentTheme = body.getAttribute("data-theme");
  console.log("📊 Current theme:", currentTheme);

  if (currentTheme === "dark") {
    setTheme("light");
  } else {
    setTheme("dark");
  }
}

/**
 * Set specific theme and update UI
 * @param {string} theme - 'light' or 'dark'
 */
function setTheme(theme) {
  console.log("🎯 Setting theme to:", theme);
  const body = document.body;
  const themeIcon = document.getElementById("theme-icon");

  body.setAttribute("data-theme", theme);
  localStorage.setItem("theme", theme);
  console.log("💾 Theme saved to localStorage:", theme);

  if (themeIcon) {
    if (theme === "dark") {
      themeIcon.className = "bi bi-moon-fill";
    } else {
      themeIcon.className = "bi bi-sun-fill";
    }
    console.log("🎨 Theme icon class updated:", themeIcon.className);
  } else {
    console.error("❌ Theme icon element not found during setTheme!");
  }

  // Dispatch custom event for theme change
  window.dispatchEvent(new CustomEvent("themeChanged", { detail: { theme } }));
  console.log("📢 Theme change event dispatched");
}

/**
 * Initialize theme on page load
 */
function initializeTheme() {
  const savedTheme = localStorage.getItem("theme") || "light";
  const themeIcon = document.getElementById("theme-icon");

  console.log("🎨 Initializing theme:", savedTheme);
  console.log("🔍 Theme icon element:", themeIcon);

  document.body.setAttribute("data-theme", savedTheme);

  if (themeIcon) {
    if (savedTheme === "dark") {
      themeIcon.className = "bi bi-moon-fill";
    } else {
      themeIcon.className = "bi bi-sun-fill";
    }
    console.log("✅ Theme icon updated:", themeIcon.className);
  } else {
    console.warn("⚠️ Theme icon element not found!");
  }
}

// ============================================================================
// NAVIGATION MANAGEMENT
// ============================================================================

/**
 * Highlight active navigation link based on current URL
 */
function highlightActiveNavLink() {
  const currentPath = window.location.pathname;
  const navLinks = document.querySelectorAll(".navbar-nav .nav-link");

  navLinks.forEach((link) => {
    const href = link.getAttribute("href");
    if (href === currentPath) {
      link.classList.add("active");
    } else {
      link.classList.remove("active");
    }
  });
}

/**
 * Handle mobile navbar collapse on link click
 */
function handleMobileNavigation() {
  const navbarToggler = document.querySelector(".navbar-toggler");
  const navbarCollapse = document.querySelector(".navbar-collapse");
  const navLinks = document.querySelectorAll(".navbar-nav .nav-link");

  // Close mobile menu when clicking on nav links
  navLinks.forEach((link) => {
    link.addEventListener("click", () => {
      if (navbarCollapse && navbarCollapse.classList.contains("show")) {
        const bsCollapse = new bootstrap.Collapse(navbarCollapse, {
          toggle: false,
        });
        bsCollapse.hide();
      }
    });
  });
}

// ============================================================================
// SCROLL MANAGEMENT
// ============================================================================

/**
 * Show/hide back to top button based on scroll position
 */
function handleScrollButton() {
  const backToTopBtn = document.getElementById("backToTopBtn");

  if (backToTopBtn) {
    if (
      document.body.scrollTop > 20 ||
      document.documentElement.scrollTop > 20
    ) {
      backToTopBtn.style.display = "block";
    } else {
      backToTopBtn.style.display = "none";
    }
  }
}

/**
 * Smooth scroll to top of page
 */
function scrollToTop() {
  document.body.scrollTop = 0;
  document.documentElement.scrollTop = 0;
}

/**
 * Initialize smooth scrolling for anchor links
 */
function initializeSmoothScrolling() {
  const links = document.querySelectorAll('a[href^="#"]');

  links.forEach((link) => {
    link.addEventListener("click", function (e) {
      e.preventDefault();
      const targetId = this.getAttribute("href");
      const target = document.querySelector(targetId);

      if (target) {
        target.scrollIntoView({
          behavior: "smooth",
          block: "start",
        });
      }
    });
  });
}

// ============================================================================
// MODAL MANAGEMENT
// ============================================================================

/**
 * Handle sign-in form submission
 */
function handleSignInForm() {
  const signInForm = document.querySelector("#signinModal form");

  if (signInForm) {
    signInForm.addEventListener("submit", function (e) {
      e.preventDefault();

      const email = document.getElementById("email").value;
      const password = document.getElementById("password").value;
      const rememberMe = document.getElementById("rememberMe").checked;

      // TODO: Implement actual sign-in logic
      console.log("Sign-in attempt:", { email, rememberMe });

      // Close modal after successful sign-in
      const modal = bootstrap.Modal.getInstance(
        document.getElementById("signinModal"),
      );
      if (modal) {
        modal.hide();
      }

      // Show success message (replace with actual implementation)
      showNotification("Sign-in functionality coming soon!", "info");
    });
  }
}

// ============================================================================
// UTILITY FUNCTIONS
// ============================================================================

/**
 * Show notification/toast message
 * @param {string} message - Message to display
 * @param {string} type - Type of notification ('success', 'error', 'info', 'warning')
 */
function showNotification(message, type = "info") {
  // Create toast element
  const toastHtml = `
        <div class="toast align-items-center text-bg-${type === "error" ? "danger" : type} border-0" role="alert" aria-live="assertive" aria-atomic="true">
            <div class="d-flex">
                <div class="toast-body">
                    <i class="bi bi-${type === "success" ? "check-circle" : type === "error" ? "exclamation-triangle" : type === "warning" ? "exclamation-triangle" : "info-circle"} me-2"></i>
                    ${message}
                </div>
                <button type="button" class="btn-close btn-close-white me-2 m-auto" data-bs-dismiss="toast" aria-label="Close"></button>
            </div>
        </div>
    `;

  // Create or get toast container
  let toastContainer = document.querySelector(".toast-container");
  if (!toastContainer) {
    toastContainer = document.createElement("div");
    toastContainer.className = "toast-container position-fixed top-0 end-0 p-3";
    toastContainer.style.zIndex = "1080";
    document.body.appendChild(toastContainer);
  }

  // Add toast to container
  toastContainer.insertAdjacentHTML("beforeend", toastHtml);

  // Show toast
  const toastElement = toastContainer.lastElementChild;
  const toast = new bootstrap.Toast(toastElement, {
    delay: 5000,
  });
  toast.show();

  // Remove toast element after it's hidden
  toastElement.addEventListener("hidden.bs.toast", () => {
    toastElement.remove();
  });
}

/**
 * Debounce function to limit function calls
 * @param {Function} func - Function to debounce
 * @param {number} wait - Wait time in milliseconds
 * @param {boolean} immediate - Trigger on leading edge
 */
function debounce(func, wait, immediate) {
  let timeout;
  return function executedFunction(...args) {
    const later = () => {
      timeout = null;
      if (!immediate) func(...args);
    };
    const callNow = immediate && !timeout;
    clearTimeout(timeout);
    timeout = setTimeout(later, wait);
    if (callNow) func(...args);
  };
}

// ============================================================================
// PAGE ANIMATIONS
// ============================================================================

/**
 * Add entrance animations to elements
 */
function initializeAnimations() {
  // Add fade-in animation to main content
  document.body.classList.add("fade-in");

  // Animate cards on scroll (intersection observer)
  if ("IntersectionObserver" in window) {
    const cardObserver = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          if (entry.isIntersecting) {
            entry.target.style.animationDelay = Math.random() * 0.5 + "s";
            entry.target.classList.add("fade-in");
            cardObserver.unobserve(entry.target);
          }
        });
      },
      {
        threshold: 0.1,
        rootMargin: "50px",
      },
    );

    // Observe all cards
    document.querySelectorAll(".card, .content-section").forEach((card) => {
      cardObserver.observe(card);
    });
  }
}

// ============================================================================
// EVENT LISTENERS & INITIALIZATION
// ============================================================================

/**
 * Initialize all functionality when DOM is loaded
 */
document.addEventListener("DOMContentLoaded", function () {
  console.log(
    "🚀 DOM Content Loaded - Initializing Shree Nathji Haweli website...",
  );

  // Check if main elements exist
  const themeButton = document.querySelector(".theme-toggle");
  const navbarBrand = document.querySelector(".navbar-brand");
  const signinButton = document.querySelector(".btn-signin");

  console.log("🔍 Element check:");
  console.log("  - Theme button:", themeButton ? "✅ Found" : "❌ Missing");
  console.log("  - Navbar brand:", navbarBrand ? "✅ Found" : "❌ Missing");
  console.log("  - Sign-in button:", signinButton ? "✅ Found" : "❌ Missing");

  // Theme management
  initializeTheme();

  // Navigation
  highlightActiveNavLink();
  handleMobileNavigation();

  // Scroll functionality
  initializeSmoothScrolling();

  // Modal functionality
  handleSignInForm();

  // Animations
  initializeAnimations();

  console.log("🕉️ Shree Nathji Haweli website initialized successfully!");
});

// Scroll event listener (debounced for performance)
window.addEventListener("scroll", debounce(handleScrollButton, 100));

// Global functions for HTML onclick handlers
window.toggleTheme = toggleTheme;
window.scrollToTop = scrollToTop;
window.Function = scrollToTop; // Alias for backward compatibility

// Expose utility functions globally
window.SreeNathjiUtils = {
  showNotification,
  setTheme,
  debounce,
};

// ============================================================================
// DEVELOPMENT HELPERS
// ============================================================================

// Add helpful development tools in development mode
if (
  window.location.hostname === "localhost" ||
  window.location.hostname === "127.0.0.1"
) {
  console.log("🔧 Development mode detected");

  // Add keyboard shortcuts for development
  document.addEventListener("keydown", function (e) {
    // Ctrl/Cmd + Shift + T to toggle theme
    if ((e.ctrlKey || e.metaKey) && e.shiftKey && e.key === "T") {
      e.preventDefault();
      toggleTheme();
      showNotification("Theme toggled via keyboard shortcut", "info");
    }
  });
}
