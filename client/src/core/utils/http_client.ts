import axios from "axios";

const httpClient = axios.create({
  baseURL: "http://localhost:8080",
  timeout: 30_000,
  withCredentials: true,
});

export default httpClient;
