// =========================================================
// IDINEX — Frontend config
// Single source of truth for environment-dependent values.
// Swap API_BASE_URL when deploying (see deployments/production).
// =========================================================

export const config = {
  API_BASE_URL: "http://localhost:8080",
  HEALTH_ENDPOINT: "/health",
  REQUEST_TIMEOUT_MS: 4000,
};
