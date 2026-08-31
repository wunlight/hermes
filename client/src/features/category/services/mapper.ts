import type { Category, CategoryDTO, CategoryOption } from "../types/types";

export function dtoToCategory(dto: CategoryDTO): Category {
  return {
    code: dto.code,
    createdAt: dto.created_at,
    id: dto.id,
    name: dto.name,
    parentId: dto.parent_id,
    updatedAt: dto.updated_at,
  };
}

export function dtoToCategoryOption(dto: CategoryDTO): CategoryOption {
  return {
    id: dto.id,
    name: dto.name,
  };
}
