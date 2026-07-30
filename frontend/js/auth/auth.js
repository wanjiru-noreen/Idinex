// =========================================================
// IDINEX — Auth feature
// UI logic for login/register/forgot-password. Talks to the
// backend only through js/api/client.js.
// =========================================================

import { apiClient } from "../api/client.js";

export async function login(email, password) {
  return apiClient.post("/auth/login", { email, password });
}

export async function register(payload) {
  return apiClient.post("/auth/register", payload);
}
