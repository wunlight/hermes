import categoryApi from "../api/api";
import type { Category, CategoryFormModel, CategoryReq } from "../types/types";
import { dtoToCategory } from "./mapper";

const categorySrv = {
  list: async (): Promise<Category[]> => {
    const { data } = await categoryApi.list();
    return data.map((c) => dtoToCategory(c));
  },
  getByID: async (id: string): Promise<Category> => {
    const { data } = await categoryApi.getByID(id);
    return dtoToCategory(data);
  },
  create: async (form: CategoryFormModel): Promise<Category> => {
    const req: CategoryReq = {
      code: form.code,
      name: form.name,
      parent_id: form.parentId,
    };

    const { data } = await categoryApi.create(req);
    return dtoToCategory(data);
  },
  update: async (id: string, form: CategoryFormModel): Promise<Category> => {
    const req: CategoryReq = {
      code: form.code,
      name: form.name,
      parent_id: form.parentId,
    };

    const { data } = await categoryApi.update(id, req);
    return dtoToCategory(data);
  },
  delete: async (id: string) => {
    await categoryApi.delete(id);
  },
};

export default categorySrv;
