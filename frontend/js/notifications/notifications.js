// =========================================================
// IDINEX — Notifications feature
// UI logic for the notification bell / list.
// =========================================================

import { apiClient } from "../api/client.js";

export async function listNotifications() {
  return apiClient.get("/notifications");
}

export async function markRead(id) {
  return apiClient.put(`/notifications/${id}/read`);
}
