import type { Warehouse, WarehouseDTO, WarehouseOption } from "../types/types";

export function dtoToWarehouse(dto: WarehouseDTO): Warehouse {
  return {
    code: dto.code,
    description: dto.description,
    id: dto.id,
    name: dto.name,
  };
}

export function dtoToWarehouseOption(dto: WarehouseDTO): WarehouseOption {
  return {
    id: dto.id,
    name: dto.name,
  };
}
