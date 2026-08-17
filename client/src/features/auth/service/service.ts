import authApi from "../api/api";
import type { LoginDomainResponse, LoginRequest } from "../types/types";

export const Login = async (
  req: LoginRequest,
): Promise<LoginDomainResponse> => {
  const { data } = await authApi.login(req);
  return {
    accessToken: data.access_token,
    userID: data.user_id,
  };
};
