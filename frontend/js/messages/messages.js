// =========================================================
// IDINEX — Messages feature
// UI logic for the messaging thread tied to an access request.
// =========================================================

import { apiClient } from "../api/client.js";

export async function listThreads() {
  return apiClient.get("/messages/threads");
}

export async function sendMessage(threadId, body) {
  return apiClient.post(`/messages/threads/${threadId}`, { body });
}
