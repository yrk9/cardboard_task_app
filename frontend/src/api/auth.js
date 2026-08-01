// 認証が必要な動作を定義する。(登録とログイン)
import { apiFetch } from "./client";

export async function register(email, password) {
  return apiFetch("/api/v1/auth/register", {
    method: "POST",
    body: JSON.stringify({ email, password }), //jsonを文字列に変換
  });
}

export async function login(email, password) {
  return apiFetch("/api/v1/auth/login", {
    method: "POST",
    body: JSON.stringify({ email, password }),
  });
}
