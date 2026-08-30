import brandApi from "../api/api";
import type { Brand, BrandFormModel, BrandReq } from "../types/types";
import { dtoToBrand } from "./mapper";

const brandSrv = {
  list: async (): Promise<Brand[]> => {
    const { data } = await brandApi.list();
    return data.map((b) => dtoToBrand(b));
  },
  getByID: async (id: string): Promise<Brand> => {
    const { data } = await brandApi.getByID(id);
    return dtoToBrand(data);
  },
  create: async (form: BrandFormModel): Promise<Brand> => {
    const req: BrandReq = {
      code: form.code,
      name: form.name,
    };

    const { data } = await brandApi.create(req);
    return dtoToBrand(data);
  },
  update: async (id: string, form: BrandFormModel): Promise<Brand> => {
    const req: BrandReq = {
      code: form.code,
      name: form.name,
    };

    const { data } = await brandApi.update(id, req);
    return dtoToBrand(data);
  },
  delete: async (id: string) => {
    await brandApi.delete(id);
  },
};

export default brandSrv;
