import type { Brand, BrandDTO, BrandOption } from "../types/types";

export function dtoToBrand(dto: BrandDTO): Brand {
  return {
    code: dto.code,
    createdAt: dto.created_at,
    id: dto.id,
    name: dto.name,
    updatedAt: dto.updated_at,
  };
}

export function dtoToBrandOption(dto: BrandDTO): BrandOption {
  return {
    id: dto.id,
    name: dto.name,
  };
}
