import httpClient from "@/core/utils/http-client";
import type { WarehouseDTO, WarehouseReq } from "../types/types";

const warehouseApi = {
  list: () => httpClient.get<WarehouseDTO[]>("warehouse"),
  getByID: (id: string) => httpClient.get<WarehouseDTO>(`warehouse/${id}`),
  create: (req: WarehouseReq) =>
    httpClient.post<WarehouseDTO>("warehouse", req),
  update: (id: string, req: WarehouseReq) =>
    httpClient.put<WarehouseDTO>(`warehouse/${id}`, req),
  delete: (id: string) => httpClient.delete(`warehouse/${id}`),
};

export default warehouseApi;
