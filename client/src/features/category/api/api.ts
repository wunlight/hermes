import httpClient from "@/core/utils/http-client";
import type { CategoryDTO, CategoryReq } from "../types/types";

const categoryApi = {
  list: () => httpClient.get<CategoryDTO[]>("category"),
  getByID: (id: string) => httpClient.get<CategoryDTO>(`category/${id}`),
  create: (req: CategoryReq) => httpClient.post<CategoryDTO>("category", req),
  update: (id: string, req: CategoryReq) =>
    httpClient.put<CategoryDTO>(`category/${id}`, req),
  delete: (id: string) => httpClient.delete(`category/${id}`),
};

export default categoryApi;
