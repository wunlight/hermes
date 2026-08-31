export type WarehouseDTO = {
  id: string;
  code: string;
  name: string;
  description: string;
};

export type Warehouse = {
  id: string;
  code: string;
  name: string;
  description: string;
};

export type WarehouseOption = {
  id: string;
  name: string;
};

export type WarehouseReq = {
  code: string;
  name: string;
  description: string;
};

export type WarehouseFormModel = {
  code: string;
  name: string;
  description: string;
};
