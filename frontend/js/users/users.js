// =========================================================
// IDINEX — Users feature
// Profile read/update UI logic.
// =========================================================

import { apiClient } from "../api/client.js";

export async function getProfile() {
  return apiClient.get("/users/me");
}

export async function updateProfile(payload) {
  return apiClient.put("/users/me", payload);
}
