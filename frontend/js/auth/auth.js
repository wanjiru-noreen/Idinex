// =========================================================
// IDINEX — Auth feature & state management
// Handles login, registration, logout, session persistence,
// and UI updates across pages.
// =========================================================

import { apiClient } from "../api/client.js";

const TOKEN_KEY = "idinex_auth_token";
const USER_KEY = "idinex_user";

/**
 * Store session data in localStorage or sessionStorage.
 */
export function setSession(token, user, remember = true) {
  const storage = remember ? localStorage : sessionStorage;
  // Clear opposite storage
  localStorage.removeItem(TOKEN_KEY);
  localStorage.removeItem(USER_KEY);
  sessionStorage.removeItem(TOKEN_KEY);
  sessionStorage.removeItem(USER_KEY);

  if (token) storage.setItem(TOKEN_KEY, token);
  if (user) storage.setItem(USER_KEY, JSON.stringify(user));
}

/**
 * Retrieve auth token.
 */
export function getToken() {
  return localStorage.getItem(TOKEN_KEY) || sessionStorage.getItem(TOKEN_KEY) || null;
}

/**
 * Retrieve current logged-in user.
 */
export function getUser() {
  const userStr = localStorage.getItem(USER_KEY) || sessionStorage.getItem(USER_KEY);
  if (!userStr) return null;
  try {
    return JSON.parse(userStr);
  } catch (e) {
    return null;
  }
}

/**
 * Check if a user is currently authenticated.
 */
export function isAuthenticated() {
  return Boolean(getToken() || getUser());
}

/**
 * Authenticate user with backend or demo fallback if backend is offline.
 */
export async function login(email, password, remember = true) {
  try {
    const data = await apiClient.post("/auth/login", { email, password });
    const token = data.token || data.jwt || "demo_token_" + Date.now();
    const user = data.user || {
      id: data.id || "usr_101",
      email: email,
      username: email.split("@")[0] || "developer",
      name: data.name || email.split("@")[0] || "Developer",
    };
    setSession(token, user, remember);
    return { success: true, user, token };
  } catch (err) {
    // If request failed because backend endpoint isn't running yet (offline / 404),
    // provide seamless client demo authentication so UI can be fully tested.
    if (err.status === 404 || err.name === "TypeError" || err.message?.includes("fetch")) {
      const demoUser = {
        id: "usr_demo",
        email: email,
        username: email.split("@")[0] || "bramwel",
        name: capitalize(email.split("@")[0]) || "Dev Bramwel",
        role: "Collaborator / Creator",
        createdAt: "2026-08-01",
      };
      const demoToken = "demo_jwt_token_" + Date.now();
      setSession(demoToken, demoUser, remember);
      return { success: true, user: demoUser, token: demoToken, isDemo: true };
    }
    throw err;
  }
}

/**
 * Register a new user.
 */
export async function register(payload, remember = true) {
  try {
    const data = await apiClient.post("/auth/register", payload);
    const token = data.token || "demo_token_" + Date.now();
    const user = data.user || {
      id: data.id || "usr_" + Date.now(),
      email: payload.email,
      username: payload.username || payload.email.split("@")[0],
      name: payload.name || payload.username || "New Builder",
    };
    setSession(token, user, remember);
    return { success: true, user, token };
  } catch (err) {
    if (err.status === 404 || err.name === "TypeError" || err.message?.includes("fetch")) {
      const demoUser = {
        id: "usr_" + Date.now(),
        email: payload.email,
        username: payload.username || payload.email.split("@")[0],
        name: payload.name || payload.username || "New Builder",
        role: "Creator",
        createdAt: new Date().toISOString().split("T")[0],
      };
      const demoToken = "demo_jwt_token_" + Date.now();
      setSession(demoToken, demoUser, remember);
      return { success: true, user: demoUser, token: demoToken, isDemo: true };
    }
    throw err;
  }
}

/**
 * Log out current user and clear stored auth credentials.
 */
export function logout(redirect = true) {
  localStorage.removeItem(TOKEN_KEY);
  localStorage.removeItem(USER_KEY);
  sessionStorage.removeItem(TOKEN_KEY);
  sessionStorage.removeItem(USER_KEY);

  // Dispatch custom logout event
  window.dispatchEvent(new CustomEvent("idinex:logout"));

  if (redirect) {
    // Determine path based on location
    const isPagesDir = window.location.pathname.includes("/pages/");
    const targetUrl = isPagesDir ? "login.html?logout=true" : "pages/login.html?logout=true";
    window.location.href = targetUrl;
  }
}

/**
 * Route guard for protected pages (e.g., Dashboard).
 */
export function requireAuth(fallbackPath = null) {
  if (!isAuthenticated()) {
    const isPagesDir = window.location.pathname.includes("/pages/");
    const defaultFallback = isPagesDir ? "login.html" : "pages/login.html";
    const target = fallbackPath || defaultFallback;
    const currentPath = encodeURIComponent(window.location.pathname);
    window.location.href = `${target}?redirect=${currentPath}`;
    return false;
  }
  return true;
}

/**
 * Automatically updates navbar / sidebar links and mounts logout handlers.
 */
export function initAuthUI() {
  const user = getUser();
  const authed = isAuthenticated();

  // Attach event listeners to all logout buttons/links
  document.querySelectorAll('[data-action="logout"], .btn-logout, .js-logout').forEach((el) => {
    el.removeEventListener("click", handleLogoutClick);
    el.addEventListener("click", handleLogoutClick);
  });

  // Update navbar links if present
  const navLinksContainer = document.querySelector(".navbar__links");
  if (navLinksContainer) {
    if (authed && user) {
      const isPagesDir = window.location.pathname.includes("/pages/");
      const prefix = isPagesDir ? "" : "pages/";
      const homePath = isPagesDir ? "../index.html" : "index.html";

      navLinksContainer.innerHTML = `
        <a href="${homePath}">Home</a>
        <a href="${prefix}ideas.html">Ideas</a>
        <a href="${prefix}dashboard.html" class="nav-user-badge">
          <span class="user-avatar-tiny">${(user.name || user.username || "U")[0].toUpperCase()}</span>
          <span>${user.username || user.name}</span>
        </a>
        <button type="button" class="btn btn--ghost btn--sm js-logout" data-action="logout" style="padding: 4px 12px; font-size: 0.8125rem;">
          Log out
        </button>
      `;

      // Re-bind click event to newly inserted logout button
      const newLogoutBtn = navLinksContainer.querySelector('.js-logout');
      if (newLogoutBtn) {
        newLogoutBtn.addEventListener("click", handleLogoutClick);
      }
    }
  }

  // Update user name placeholders across the page if any
  if (authed && user) {
    document.querySelectorAll("[data-user-name]").forEach((el) => {
      el.textContent = user.name || user.username || "Builder";
    });
    document.querySelectorAll("[data-user-email]").forEach((el) => {
      el.textContent = user.email || "";
    });
    document.querySelectorAll("[data-user-avatar]").forEach((el) => {
      el.textContent = (user.name || user.username || "U")[0].toUpperCase();
    });
  }
}

function handleLogoutClick(e) {
  e.preventDefault();
  logout(true);
}

function capitalize(str) {
  if (!str) return "";
  return str.charAt(0).toUpperCase() + str.slice(1);
}

// Auto-initialize logout listeners on DOMContentLoaded
if (typeof document !== "undefined") {
  document.addEventListener("DOMContentLoaded", () => {
    initAuthUI();
  });
}
