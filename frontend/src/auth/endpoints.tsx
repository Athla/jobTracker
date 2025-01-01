import apiClient from "./axios"


export const BASE_URL = 'http://localhost:8080'


export const registerUser = async (username: string, password: string): Promise<void> => {
  await apiClient.post("/register", { username, password })
}
export const loginUser = async (username: string, password: string): Promise<void> => {
  await apiClient.post("/login", { username, password })
}
export const logoutUser = async (token: string): Promise<void> => {
  await apiClient.post("/logout", { token })
}
