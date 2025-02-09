import { useState, useCallback } from "react";

interface AuthState {
  isAuthenticated: boolean;
  token: string | null;
  login: (token: string) => void;
  logout: () => void;
}

export const useAuth = (): AuthState => {
  const [token, setToken] = useState<string | null>(
    localStorage.getItem("auth_token"),
  );

  const login = useCallback((token: string) => {
    localStorage.setItem("auth_token", token);
    setToken(token);
  }, []);

  const logout = useCallback(() => {
    localStorage.removeItem("auth_token");
    setToken(null);
  }, []);

  return {
    isAuthenticated: !!token,
    token,
    login,
    logout,
  };
};
