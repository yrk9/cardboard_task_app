import { useState, useEffect, createContext, useContext } from "react";

export default function AuthContext() {
  const [token, setToken] = useState<string | null>(null); //tokenの状態管理

  //ページを開いた瞬間にトークンをstateにわたす
  useEffect(() => {
    const savedToken = localStorage.getItem("token");
    if (savedToken) {
      setToken(savedToken);
    }
  }, []);

  const AuthContext = createContext();
}
