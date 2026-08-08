import React, {
  useState,
  useEffect,
  createContext,
  useContext,
  Children,
} from "react";
import { login as apiLogin } from "../api/auth";

type AuthContextType = {
  token: string | null;
  login: (email: string, password: string) => Promise<void>;
  logout: () => void;
};

const AuthContext = createContext<AuthContextType | null>(null);

export default function AuthProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const [token, setToken] = useState<string | null>(null); //tokenの状態管理

  //ページを開いた瞬間にトークンをstateにわたす
  useEffect(() => {
    const savedToken = localStorage.getItem("token");
    if (savedToken) {
      setToken(savedToken);
    }
  }, []);

  async function login(email: string, password: string) {
    const data = await apiLogin(email, password);

    localStorage.setItem("token", data.token);
    setToken(data.token);
  }

  function logout() {
    localStorage.removeItem("token");
    setToken(null);
  }

  return (
    //children: タグの中に書いたのが渡される
    <AuthContext.Provider value={{ token, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  return useContext(AuthContext);
}
