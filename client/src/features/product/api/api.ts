import httpClient from "@/core/utils/http-client";
import type { ProductDTO, ProductReq } from "../types/types";

const productApi = {
  list: () => httpClient.get<ProductDTO[]>("product"),
  getByID: (id: string) => httpClient.get<ProductDTO>(`product/${id}`),
  create: (req: ProductReq) => httpClient.post<ProductDTO>("product", req),
  update: (id: string, req: ProductReq) =>
    httpClient.put<ProductDTO>(`product/${id}`, req),
  delete: (id: string) => httpClient.delete(`product/${id}`),
};

export default productApi;
