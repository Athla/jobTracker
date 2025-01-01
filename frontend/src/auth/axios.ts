import axios from "axios";

const apiClient = axios.create({
  baseURL: `http://localhost:8080/users`,
  headers: {
    "Content-Type": "application/json",
  },
});

export default apiClient;
