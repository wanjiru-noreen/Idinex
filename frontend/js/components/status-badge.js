// =========================================================
// IDINEX — Status badge
// Renders "Backend Connected" / "backend offline" into any
// element with [data-status-badge], driven by the health check.
// =========================================================

import { checkBackendHealth } from "../api/health.js";

async function renderStatus(el) {
  const { online } = await checkBackendHealth();

  el.textContent = online ? "Backend Connected" : "backend offline";
  el.classList.toggle("status-badge--online", online);
  el.classList.toggle("status-badge--offline", !online);
  el.setAttribute("aria-live", "polite");
}

export function initStatusBadges() {
  document.querySelectorAll("[data-status-badge]").forEach(renderStatus);
}
