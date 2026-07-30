// =========================================================
// IDINEX — Ideas feature
// UI logic for listing, creating, and viewing ideas.
// =========================================================

import { apiClient } from "../api/client.js";

export async function listIdeas() {
  return apiClient.get("/ideas");
}

export async function getIdea(id) {
  return apiClient.get(`/ideas/${id}`);
}

export async function createIdea(payload) {
  return apiClient.post("/ideas", payload);
}
