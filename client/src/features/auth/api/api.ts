import httpClient from "@/core/utils/http_client";
import type { LoginRequest } from "../types/types";

const authApi = {
  login: (req: LoginRequest) => httpClient.post("auth/login", req),
};

export default authApi;
