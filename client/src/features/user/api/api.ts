import httpClient from "@/core/utils/http-client";
import type { User } from "../types/types";

const userApi = {
  list: () => httpClient.get<Record<"data", User[]>>("/users"),
};

export default userApi;
