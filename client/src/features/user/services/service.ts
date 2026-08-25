import userApi from "../api/api";
import type { User } from "../types/types";

export const listUser = async (): Promise<User[]> => {
  const { data } = await userApi.list();

  if (!Array.isArray(data.data)) return [];

  return data.data.map((user) => ({
    email: user.email,
    id: user.id,
    name: user.name,
  }));
};
