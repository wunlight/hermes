import httpClient from "@/core/utils/http-client";
import type {
  LoginDTOResponse,
  LoginRequest,
  RefreshDTOResponse,
} from "../types/types";

const authApi = {
  login: (req: LoginRequest) =>
    httpClient.post<LoginDTOResponse>("auth/login", req),
  refresh: () => httpClient.post<RefreshDTOResponse>("auth/refresh"),
};

export default authApi;
