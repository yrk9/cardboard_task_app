import React, {
  useState,
  useEffect,
  createContext,
  useContext,
  Children,
} from "react";
import { login as apiLogin } from "../api/auth";
import { register as apiRegister } from "../api/auth";

type AuthContextType = {
  token: string | null;
  isLoading: boolean;
  register: (email: string, password: string) => Promise<void>;
  login: (email: string, password: string) => Promise<void>;
  logout: () => void;
};

const AuthContext = createContext<AuthContextType | null>(null);

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [token, setToken] = useState<string | null>(null); //tokenの状態管理
  const [isLoading, setIsLoading] = useState(true);

  //ページを開いた瞬間にトークンをstateにわたす
  useEffect(() => {
    const savedToken = localStorage.getItem("token");
    if (savedToken) {
      setToken(savedToken);
    }
    setIsLoading(false);
  }, []);

  async function register(email: string, password: string) {
    const data = await apiRegister(email, password);

    localStorage.setItem("token", data.token);
    setToken(data.token);
  }

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
    <AuthContext.Provider value={{ token, isLoading, register, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);

  if (!context) {
    throw new Error("useAuthはAuthProviderの中で使ってください");
  }
  return context;
}
