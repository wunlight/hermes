import httpClient from "@/core/utils/http_client";
import type { LoginRequest, LoginResponse } from "../types/types";

const authApi = {
  login: (req: LoginRequest) =>
    httpClient.post<LoginResponse>("auth/login", req),
};

export default authApi;
