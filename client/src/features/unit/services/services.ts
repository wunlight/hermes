import unitApi from "../api/api";
import type { Unit, UnitFormModel, UnitOption, UnitReq } from "../types/types";
import { dtoToUnit, dtoToUnitOption } from "./mapper";

const unitSrv = {
  list: async (): Promise<Unit[]> => {
    const { data } = await unitApi.list();
    return data.map((b) => dtoToUnit(b));
  },
  getOptions: async (): Promise<UnitOption[]> => {
    const { data } = await unitApi.list();
    return data.map((b) => dtoToUnitOption(b));
  },
  getByID: async (id: string): Promise<Unit> => {
    const { data } = await unitApi.getByID(id);
    return dtoToUnit(data);
  },
  create: async (form: UnitFormModel): Promise<Unit> => {
    const req: UnitReq = {
      code: form.code,
      name: form.name,
    };

    const { data } = await unitApi.create(req);
    return dtoToUnit(data);
  },
  update: async (id: string, form: UnitFormModel): Promise<Unit> => {
    const req: UnitReq = {
      code: form.code,
      name: form.name,
    };

    const { data } = await unitApi.update(id, req);
    return dtoToUnit(data);
  },
  delete: async (id: string) => {
    await unitApi.delete(id);
  },
};

export default unitSrv;
