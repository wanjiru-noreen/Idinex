// =========================================================
// IDINEX — API client
// Thin fetch wrapper. Feature JS files (js/ideas, js/auth, ...)
// should call through here rather than using fetch() directly,
// so headers, timeouts, and error handling stay in one place.
// =========================================================

import { config } from "../config/config.js";

class ApiError extends Error {
  constructor(message, status) {
    super(message);
    this.name = "ApiError";
    this.status = status;
  }
}

async function request(path, options = {}) {
  const controller = new AbortController();
  const timeout = setTimeout(() => controller.abort(), config.REQUEST_TIMEOUT_MS);

  try {
    const response = await fetch(`${config.API_BASE_URL}${path}`, {
      headers: { "Content-Type": "application/json", ...options.headers },
      signal: controller.signal,
      ...options,
    });

    if (!response.ok) {
      throw new ApiError(`Request to ${path} failed`, response.status);
    }

    const contentType = response.headers.get("content-type") || "";
    return contentType.includes("application/json") ? response.json() : response.text();
  } finally {
    clearTimeout(timeout);
  }
}

export const apiClient = {
  get: (path, options) => request(path, { method: "GET", ...options }),
  post: (path, body, options) =>
    request(path, { method: "POST", body: JSON.stringify(body), ...options }),
  put: (path, body, options) =>
    request(path, { method: "PUT", body: JSON.stringify(body), ...options }),
  delete: (path, options) => request(path, { method: "DELETE", ...options }),
};

export { ApiError };
