import type { Brand, BrandDTO } from "../types/types";

export function dtoToBrand(dto: BrandDTO): Brand {
  return {
    code: dto.code,
    createdAt: dto.created_at,
    id: dto.id,
    name: dto.name,
    updatedAt: dto.updated_at,
  };
}
