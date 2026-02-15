# 🛕 Contributing Guidelines

Thank you for your interest in contributing to the Temple Website project.
This project is a **static, community-driven, long-term website** intended to be stable, respectful, and easy to maintain.

Please read this document carefully before contributing.

---

## 🧭 Project Principles (Non-Negotiable)

- **Static-first**: No backend, no servers, no databases
- **GitHub Pages compatible**
- **Bootstrap-based UI**
- **Vanilla HTML / CSS / JavaScript only**
- **Cultural and religious respect is mandatory**

If your change violates any of the above, it will be rejected.

---

## 🏗️ Tech Stack

Allowed:
- HTML5
- CSS3
- Bootstrap (official CDN only)
- Vanilla JavaScript
- JSON / YAML for data
- GitHub Actions (already configured)

Not Allowed:
- React, Vue, Angular, Svelte
- Tailwind or other CSS frameworks
- jQuery
- Backend languages (Go, Node, PHP, etc.)
- External UI kits or templates

---

## 🌱 How to Contribute

### 1. Fork & Branch
- Fork the repository
- Create a branch from `main`
- Use clear branch names:
  - `feature/visiting-hours`
  - `fix/navbar-mobile`
  - `content/tippani-page`

---

### 2. One Change = One Purpose
Each PR must:
- Solve **one** problem
- Touch minimal files
- Avoid unrelated refactoring

Bad ❌:
- Feature + CSS cleanup + rename files

Good ✅:
- Add visiting hours table only

---

### 3. Coding Standards

#### HTML
- Semantic tags (`header`, `main`, `section`, `footer`)
- Bootstrap grid system only
- No inline styles unless unavoidable

#### CSS
- Use existing CSS variables
- Mobile-first
- No hardcoded colors unless approved
- Do not override Bootstrap core styles aggressively

#### JavaScript
- Plain JS only
- Fetch data from `/data/*.json`
- No inline `<script>` logic
- Defensive coding (handle empty fields)

---

## 📂 Folder Structure Rules

Do not change structure without approval.

/assets
/css
/images
/data
/pages
/scripts


- New pages → `/pages`
- New JS → `/scripts`
- New data → `/data`

---

## 📝 Content Guidelines

- Use respectful language
- Avoid personal opinions
- No political or controversial content
- Religious text must be accurate
- If unsure → ask before submitting

---

## 🧪 Testing Before PR

Before submitting:
- Open site via GitHub Pages preview
- Test:
  - Mobile view
  - Navbar toggle
  - Links
  - JSON fetch logic
- Ensure no console errors

---

## 🔔 Notifications & Data Updates

- Notifications are managed via **GitHub Issues automation**
- Do NOT manually edit notification JSON unless instructed
- Follow issue template strictly

---

## 📦 Pull Request Checklist

Your PR must include:
- [ ] Clear title
- [ ] Description of change
- [ ] Screenshots (for UI changes)
- [ ] No broken links
- [ ] No console errors

PRs that do not meet this checklist may be closed without review.

---

## 🛑 What Will Get a PR Rejected Immediately

- Adding frameworks
- Changing core design philosophy
- Mixing multiple features
- Poor commit messages
- Disrespectful content
- Breaking mobile layout

---

## 🙏 Final Note

This website represents a place of faith and community.
Contribute with discipline, humility, and care.

Thank you for helping build something meaningful.
