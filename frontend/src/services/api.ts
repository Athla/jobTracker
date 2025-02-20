import axios from "axios";
import { Job, JobStatus } from "@/types/job";

const API_URL = "http://localhost:8080";

const api = axios.create({
  baseURL: API_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

// Add request interceptor to add auth token
api.interceptors.request.use((config) => {
  const token = localStorage.getItem("auth_token");
  if (token) {
    // Don't add "Bearer" prefix if it's already there
    const tokenValue = token.startsWith("Bearer ") ? token : `Bearer ${token}`;
    config.headers.Authorization = tokenValue;
  }
  return config;
});

// Add response interceptor to handle auth errors
api.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    console.error("API Error:", {
      url: error.config?.url,
      status: error.response?.status,
      data: error.response?.data,
    });

    if (error.response?.status === 401) {
      // Instead of redirecting, return a rejected promise with the error
      return Promise.reject(new Error("Unauthorized"));
    }
    return Promise.reject(error);
  }
);

export const AuthAPI = {
  login: async (username: string, password: string) => {
    try {
      const response = await api.post("/login", { username, password });
      // Store token with Bearer prefix
      const token = response.data.token;
      const tokenWithBearer = token.startsWith("Bearer ")
        ? token
        : `Bearer ${token}`;
      localStorage.setItem("auth_token", tokenWithBearer);
      return response.data;
    } catch (error) {
      console.error("Login error:", error);
      throw error;
    }
  },

  logout: async () => {
    try {
      await api.post("/logout");
      localStorage.removeItem("auth_token");
    } catch (error) {
      console.error("Logout error:", error);
      throw error;
    }
  },
};

export const JobAPI = {
  getAll: async (): Promise<Job[]> => {
    try {
      const response = await api.get("/api/jobs/");
      return response.data;
    } catch (error) {
      console.error("Error fetching jobs:", error);
      throw new Error("Failed to fetch jobs");
    }
  },

  create: async (job: Partial<Job>): Promise<Job> => {
    const response = await api.post("/api/jobs/", job);
    return response.data;
  },

  update: async (id: string, job: Partial<Job>): Promise<Job> => {
    const response = await api.put(`/api/jobs/${id}`, job);
    return response.data;
  },

  updateStatus: async (
    id: string,
    update: { status: JobStatus; version: number }
  ): Promise<Job> => {
    try {
      const response = await api.patch(`/api/jobs/${id}/status`, update);
      return response.data;
    } catch (error) {
      throw new Error(`Failed to update job status due: ${error}`);
    }
  },

  delete: async (id: string, version: number): Promise<void> => {
    await api.delete(`/api/jobs/${id}?version=${version}`);
  },

  deleteAll: async (): Promise<void> => {
    await api.delete("/api/jobs/deleteAll");
  },
};
