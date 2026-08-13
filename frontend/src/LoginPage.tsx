import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "./context/AuthContext";
import { register } from "./api/auth";
import "./LoginPage.css";

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
      if (mode === "login") {
        setErrorMessage("ログインに失敗しました");
      } else {
        setErrorMessage("登録に失敗しました");
      }
    }
  }

  function handleRegester() {
    if (mode === "login") {
      setMode("register");
    } else {
      setMode("login");
    }
    setErrorMessage(null);
  }

  return (
    <div className="login">
      <h2>{mode === "login" ? "ログイン" : "登録"}</h2>
      {errorMessage && <p>{errorMessage}</p>}
      <form onSubmit={handleSubmit} className="form-group">
        <label className="form-label">
          Eメール
          <input
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="form-control"
            placeholder="your@email.com"
          ></input>
        </label>
        <label className="form-label">
          パスワード
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="form-control"
            placeholder="パスワードを入力"
          ></input>
        </label>

        <button type="submit" className="btn-login">
          {mode === "login" ? "ログイン" : "登録"}
        </button>
        <button type="button" onClick={handleRegester} className="btn-login">
          {mode === "login" ? "新規登録はこちら" : "ログインへ"}
        </button>
      </form>
    </div>
  );
}
