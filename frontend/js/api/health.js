// =========================================================
// IDINEX — Health check
// Calls GET /health and reports whether the backend is reachable.
// =========================================================

import { config } from "../config/config.js";

/**
 * Pings the backend health endpoint.
 * @returns {Promise<{ online: boolean, detail?: string }>}
 */
export async function checkBackendHealth() {
  const controller = new AbortController();
  const timeout = setTimeout(() => controller.abort(), config.REQUEST_TIMEOUT_MS);

  try {
    const response = await fetch(`${config.API_BASE_URL}${config.HEALTH_ENDPOINT}`, {
      method: "GET",
      signal: controller.signal,
    });
    return { online: response.ok };
  } catch (err) {
    return { online: false, detail: err.name === "AbortError" ? "timeout" : err.message };
  } finally {
    clearTimeout(timeout);
  }
}
