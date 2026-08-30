import httpClient from "@/core/utils/http-client";
import type { BrandDTO, BrandReq } from "../types/types";

const brandApi = {
  list: () => httpClient.get("brand"),
  getByID: (id: string) => httpClient.get<BrandDTO>(`brand/${id}`),
  create: (req: BrandReq) => httpClient.post<BrandDTO>("brand", req),
  update: (id: string, req: BrandReq) =>
    httpClient.put<BrandDTO>(`brand/${id}`, req),
  delete: (id: string) => httpClient.delete(`brand/${id}`),
};

export default brandApi;
