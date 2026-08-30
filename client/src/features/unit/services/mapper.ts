import type { Unit, UnitDTO } from "../types/types";

export function dtoToUnit(dto: UnitDTO): Unit {
  return {
    code: dto.code,
    id: dto.id,
    name: dto.name,
  };
}
