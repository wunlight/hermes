import httpClient from "@/core/utils/http-client";
import type { UnitDTO, UnitReq } from "../types/types";

const unitApi = {
  list: () => httpClient.get<UnitDTO[]>("unit"),
  getByID: (id: string) => httpClient.get<UnitDTO>(`unit/${id}`),
  create: (req: UnitReq) => httpClient.post<UnitDTO>("unit", req),
  update: (id: string, req: UnitReq) =>
    httpClient.put<UnitDTO>(`unit/${id}`, req),
  delete: (id: string) => httpClient.delete(`unit/${id}`),
};

export default unitApi;
