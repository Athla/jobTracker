import {
  createContext,
  useContext,
  useState,
  ReactNode,
} from "react";
import { useNavigate } from "react-router-dom";

interface AuthContextType {
  isAuthenticated: boolean;
  token: string | null;
  login: (token: string) => void;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [token, setToken] = useState<string | null>(() => {
    const savedToken = localStorage.getItem("auth_token");
    return savedToken;
  });
  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(!!token);
  const navigate = useNavigate();

  const login = (newToken: string) => {
    const tokenWithBearer = newToken.startsWith("Bearer ")
      ? newToken
      : `Bearer ${newToken}`;
    setToken(tokenWithBearer);
    setIsAuthenticated(true);
    localStorage.setItem("auth_token", tokenWithBearer);
    navigate("/dashboard");
  };

  const logout = () => {
    setToken(null);
    setIsAuthenticated(false);
    localStorage.removeItem("auth_token");
    navigate("/login");
  };

  return (
    <AuthContext.Provider value={{ isAuthenticated, token, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
}
