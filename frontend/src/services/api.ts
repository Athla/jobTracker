import { Job } from "@/types/job";

const API_URL = "http://localhost:8080";

export interface APIError {
  error: string;
  code?: string;
}

export interface LoginRequest {
  username: string;
  password: string;
}

export interface LoginResponse {
  token: string;
  type: string;
}

async function handleResponse<T>(response: Response): Promise<T> {
  if (!response.ok) {
    if (response.status === 401) {
      localStorage.removeItem("auth_token");
      window.location.reload();
    }
    const error = await response.json();
    throw new Error(error.error || "An error occurred");
  }

  return response.json();
}

async function fetchWithAuth(
  endpoint: string,
  options: RequestInit = {}
): Promise<Response> {
  const token = localStorage.getItem("auth_token");
  const headers = {
    "Content-Type": "application/json",
    ...(token && { Authorization: `Bearer ${token}` }),
    ...options.headers,
  };

  return fetch(`${API_URL}${endpoint}`, {
    ...options,
    headers,
  });
}

export const AuthService = {
  async login(credentials: LoginRequest): Promise<LoginResponse> {
    const response = await fetch(`${API_URL}/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(credentials),
    });

    if (!response.ok) {
      throw new Error("Login failed");
    }

    return response.json();
  },

  async logout(token: string): Promise<void> {
    const response = await fetch(`${API_URL}/logout`, {
      method: "POST",
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      throw new Error("Logout failed");
    }
  },
};

export const JobAPI = {
  getAll: async (): Promise<Job[]> => {
    const response = await fetchWithAuth("/api/jobs");
    return handleResponse<Job[]>(response);
  },

  get: async (id: string): Promise<Job[]> => {
    const response = await fetchWithAuth(`/api/jobs/${id}`);
    return handleResponse<Job[]>(response);
  },

  create: async (job: Partial<Job>): Promise<Job> => {
    const response = await fetchWithAuth("/api/jobs", {
      method: "POST",
      body: JSON.stringify(job),
    });

    return handleResponse<Job>(response);
  },

  update: async (id: string, job: Partial<Job>): Promise<Job> => {
    const response = await fetchWithAuth(`/api/jobs/${id}`, {
      method: "POST",
      body: JSON.stringify(job),
    });

    return handleResponse<Job>(response);
  },

  updateStatus: async (id: string, job: Partial<Job>): Promise<Job> => {
    const response = await fetchWithAuth(`/api/jobs/${id}/status`, {
      method: "POST",
      body: JSON.stringify(job),
    });

    return handleResponse<Job>(response);
  },

  delete: async (id: string, version: number): Promise<Job> => {
    const response = await fetchWithAuth(`/api/jobs/${id}?version=${version}`, {
      method: "DELETE",
    });

    return handleResponse<Job>(response);
  },

  deleteAll: async (): Promise<Job> => {
    const response = await fetchWithAuth(`/api/jobs/deleteAll`, {
      method: "DELETE",
    });

    return handleResponse<Job>(response);
  },
};
