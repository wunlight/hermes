import authApi from "../api/api";
import type {
  LoginRequest,
  LoginResponse,
  RefreshResponse,
} from "../types/types";

export const Login = async (req: LoginRequest): Promise<LoginResponse> => {
  const { data } = await authApi.login(req);
  return {
    accessToken: data.access_token,
    userID: data.user_id,
  };
};

export const Refresh = async (): Promise<RefreshResponse> => {
  const { data } = await authApi.refresh();
  return {
    accessToken: data.access_token,
    userID: data.user_id,
  };
};
