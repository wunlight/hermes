import productApi from "../api/api";
import type { Product, ProductFormModel, ProductReq } from "../types/types";
import { dtoToProduct } from "./mapper";

const productSrv = {
  list: async (): Promise<Product[]> => {
    const { data } = await productApi.list();
    return data.map((p) => dtoToProduct(p));
  },
  getByID: async (id: string): Promise<Product> => {
    const { data } = await productApi.getByID(id);
    return dtoToProduct(data);
  },
  create: async (form: ProductFormModel): Promise<Product> => {
    const req: ProductReq = {
      brand_id: form.brandId,
      category_id: form.categoryId,
      description: form.description,
      length: form.length,
      min_stock: form.minStock,
      name: form.name,
      sku: form.sku,
      unit_id: form.unitId,
      weight: form.weight,
      width: form.width,
    };

    const { data } = await productApi.create(req);
    return dtoToProduct(data);
  },
  update: async (id: string, form: ProductFormModel): Promise<Product> => {
    const req: ProductReq = {
      brand_id: form.brandId,
      category_id: form.categoryId,
      description: form.description,
      length: form.length,
      min_stock: form.minStock,
      name: form.name,
      sku: form.sku,
      unit_id: form.unitId,
      weight: form.weight,
      width: form.width,
    };

    const { data } = await productApi.update(id, req);
    return dtoToProduct(data);
  },
  delete: async (id: string) => {
    await productApi.delete(id);
  },
};

export default productSrv;
