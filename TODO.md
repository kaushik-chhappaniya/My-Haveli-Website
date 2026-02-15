# 🛕 Temple Website – Contribution TODOs

This document defines the upcoming enhancements and contribution areas for the static temple website.
All contributors must follow Bootstrap-first design and keep the site fully static (GitHub Pages compatible).

---

## 🔴 High Priority

### 1. Visiting Hours Table
**Objective:** Clearly display darshan / visiting hours.

- Create a responsive table (Bootstrap)
- Columns:
  - Day
  - Morning Darshan
  - Afternoon Closed
  - Evening Darshan
- Mobile-friendly
- Place on:
  - Home page (summary)
  - Dedicated section or page (full table)
- Data source:
  - Static JSON or inline HTML (decision pending)

---

### 2. Tippani Page (Religious Notes / Announcements)
**Objective:** Add a dedicated spiritual / informational page.

- Create `tippani.html`
- Content examples:
  - Daily Tippani
  - Festival notes
  - Religious instructions
- Use Bootstrap typography (`container`, `lead`, `blockquote`)
- Optional:
  - Load content from static JSON
- Must include:
  - Existing navbar
  - Existing footer

---

## 🟡 Medium Priority

### 3. CSS Enhancements
**Objective:** Improve visual polish without breaking simplicity.

- Improve mobile spacing
- Normalize font sizes across pages
- Improve table readability on small screens
- Standardize:
  - Border radius
  - Shadows
  - Colors via CSS variables
- Avoid:
  - Inline styles
  - Over-customizing Bootstrap core classes

---

### 4. Reference Official Shreenathji Website
**Objective:** Align design and content structure with tradition.

- Study:
  - Layout patterns
  - Section ordering
  - Typography hierarchy
- Do NOT copy assets or content verbatim
- Take inspiration for:
  - Page flow
  - Naming conventions
  - Section grouping

---

## 🟢 Nice to Have (Later Phase)

- Festival calendar page
- Image gallery grid (Bootstrap cards)
- Accessibility improvements (contrast, alt text)
- SEO meta tags
- Page-level last updated timestamps

---

## 📐 Contribution Rules

- Bootstrap only (no Tailwind, no frameworks)
- Vanilla HTML / CSS / JS only
- No backend assumptions
- Keep GitHub Pages compatibility
- Keep JSON schema backward-compatible
- One feature per commit

---

## 📂 Suggested Folder Additions
/pages
└── tippani.html
/data
└── visiting_hours.json (optional)
/assets/css
└── enhancements.css

---

## 🧭 Philosophy

This project prioritizes:
- Stability over complexity
- Clarity over cleverness
- Tradition over trends

Build with respect.
