// =========================================================
// IDINEX — DOM utils
// Small, dependency-free helpers shared across feature modules.
// =========================================================

export const qs = (selector, scope = document) => scope.querySelector(selector);
export const qsa = (selector, scope = document) => Array.from(scope.querySelectorAll(selector));

export function on(el, event, handler) {
  el.addEventListener(event, handler);
  return () => el.removeEventListener(event, handler);
}
