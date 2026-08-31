import type { Product, ProductDTO } from "../types/types";

export function dtoToProduct(dto: ProductDTO): Product {
  return {
    brandId: dto.brand_id,
    brandName: dto.brand_name,
    categoryId: dto.category_id,
    categoryName: dto.category_name,
    unitId: dto.unit_id,
    unitName: dto.unit_name,
    description: dto.description,
    id: dto.id,
    length: dto.length,
    minStock: dto.min_stock,
    name: dto.name,
    sku: dto.sku,
    weight: dto.weight,
    width: dto.width,
  };
}
