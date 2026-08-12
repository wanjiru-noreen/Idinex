# Idinex — Frontend

Static, no-build frontend (plain HTML/CSS/JS + ES modules). No bundler
required.

## Run locally

Because pages load partials (`navbar.html`, `footer.html`) via `fetch`,
open `index.html` through a local server rather than the `file://`
protocol, or the includes won't load.

```bash
cd frontend
python3 -m http.server 5500
# then open http://localhost:5500
```

Any static server works (`npx serve`, VS Code Live Server, etc.).

## Connect to the backend

Edit `js/config/config.js`:

```js
export const config = {
  API_BASE_URL: "http://localhost:8080", // point at your running Go API
  HEALTH_ENDPOINT: "/health",
};
```

The landing page calls `GET /health` on load and shows **Backend
Connected** or **backend offline** in the top banner.

## Structure

See `../Architecture` doc — this folder mirrors it: `assets/`,
`components/` (shared HTML partials), `css/` (one file per
responsibility), `js/` (one folder per feature + api/config/router/
utils infrastructure), `layouts/` (reference page shells), `pages/`.

## Status

- [x] Folder structure
- [x] CSS architecture
- [x] JS architecture
- [x] Landing page
- [x] Reusable layout (navbar/footer partials + include loader)
- [x] API client + `/health` check
- [ ] Backend endpoints wired beyond `/health`
