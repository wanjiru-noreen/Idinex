// =========================================================
// IDINEX — Include loader
// Idinex is a static, no-build frontend, so page composition
// (navbar/footer shared across pages) is done at runtime:
// any element with [data-include="components/x.html"] gets
// that partial fetched and injected in place.
//
// This is what makes layouts/*.html "reusable" — every page
// includes the same navbar.html / footer.html rather than
// duplicating markup.
// =========================================================

async function loadInclude(el) {
  const path = el.getAttribute("data-include");
  try {
    const response = await fetch(path);
    if (!response.ok) throw new Error(`Failed to load ${path}`);
    el.outerHTML = await response.text();
  } catch (err) {
    console.error("[include-loader]", err);
  }
}

export function initIncludes() {
  const includes = document.querySelectorAll("[data-include]");
  return Promise.all(Array.from(includes).map(loadInclude));
}
