// =========================================================
// IDINEX — Router
// Minimal client-side path matcher. Idinex ships pages as
// static HTML today, so this only handles active-link state;
// swap in a real router if pages move to a single-page app.
// =========================================================

export function highlightActiveLink() {
  const path = window.location.pathname;
  document.querySelectorAll("a[href]").forEach((a) => {
    if (a.getAttribute("href") && path.endsWith(a.getAttribute("href"))) {
      a.setAttribute("aria-current", "page");
    }
  });
}
