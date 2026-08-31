import type { Unit, UnitDTO, UnitOption } from "../types/types";

export function dtoToUnit(dto: UnitDTO): Unit {
  return {
    code: dto.code,
    id: dto.id,
    name: dto.name,
  };
}

export function dtoToUnitOption(dto: UnitDTO): UnitOption {
  return {
    id: dto.id,
    name: dto.name,
  };
}
