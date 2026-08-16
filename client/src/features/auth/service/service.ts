import authApi from "../api/api";
import type { LoginRequest } from "../types/types";

export const Login = async (req: LoginRequest) => {
  const { data } = await authApi.login(req);
  console.log(data);
};
