import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "./context/AuthContext";
import { register } from "./api/auth";

export function LoginPage() {
  const { login } = useAuth();
  const navigate = useNavigate();
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [errorMessage, setErrorMessage] = useState<string | null>(null); //エラーメッセージ表示用
  const [mode, setMode] = useState<"login" | "register">("login"); //登録ページかログインページか

  async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    setErrorMessage(null);
    try {
      if (mode === "login") {
        await login(email, password);
      } else {
        await register(email, password);
      }
      navigate("/");
    } catch (error) {
      setErrorMessage("ログインに失敗しました");
    }
  }

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <label>
          Eメール
          <input
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          ></input>
        </label>
        <label>
          パスワード
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          ></input>
        </label>

        <button type="submit">{mode === "login" ? "ログイン" : "登録"}</button>
        <button
          type="button"
          onClick={() => setMode(mode === "login" ? "register" : "login")}
        >
          {mode === "login" ? "新規登録はこちら" : "ログインへ"}
        </button>
        {errorMessage && <p>{errorMessage}</p>}
      </form>
    </div>
  );
}
