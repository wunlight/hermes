import warehouseApi from "../api/api";
import type {
  Warehouse,
  WarehouseFormModel,
  WarehouseOption,
  WarehouseReq,
} from "../types/types";
import { dtoToWarehouse, dtoToWarehouseOption } from "./mapper";

const warehouseSrv = {
  list: async (): Promise<Warehouse[]> => {
    const { data } = await warehouseApi.list();
    return data.map((b) => dtoToWarehouse(b));
  },
  getOptions: async (): Promise<WarehouseOption[]> => {
    const { data } = await warehouseApi.list();
    return data.map((b) => dtoToWarehouseOption(b));
  },
  getByID: async (id: string): Promise<Warehouse> => {
    const { data } = await warehouseApi.getByID(id);
    return dtoToWarehouse(data);
  },
  create: async (form: WarehouseFormModel): Promise<Warehouse> => {
    const req: WarehouseReq = {
      code: form.code,
      name: form.name,
      description: form.description,
    };

    const { data } = await warehouseApi.create(req);
    return dtoToWarehouse(data);
  },
  update: async (id: string, form: WarehouseFormModel): Promise<Warehouse> => {
    const req: WarehouseReq = {
      code: form.code,
      name: form.name,
      description: form.description,
    };

    const { data } = await warehouseApi.update(id, req);
    return dtoToWarehouse(data);
  },
  delete: async (id: string) => {
    await warehouseApi.delete(id);
  },
};

export default warehouseSrv;
