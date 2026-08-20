import axios from "axios";
import useAuthStore from "../../features/auth/stores/authStore";

const httpClient = axios.create({
  baseURL: "http://localhost:8080",
  timeout: 30_000,
  withCredentials: true,
});

httpClient.interceptors.request.use((config) => {
  const accessToken = useAuthStore.getState().accessToken;

  if (accessToken) {
    config.headers.Authorization = `Bearer ${accessToken}`;
  }

  return config;
});

export default httpClient;
