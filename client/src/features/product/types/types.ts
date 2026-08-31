export type ProductDTO = {
  id: string;
  sku: string;
  name: string;
  category_id: string;
  category_name: string;
  brand_id: string;
  brand_name: string;
  unit_id: string;
  unit_name: string;
  min_stock: number;
  weight: number;
  length: number;
  width: number;
  description: string;
};

export type Product = {
  id: string;
  sku: string;
  name: string;
  categoryId: string;
  categoryName: string;
  brandId: string;
  brandName: string;
  unitId: string;
  unitName: string;
  minStock: number;
  weight: number;
  length: number;
  width: number;
  description: string;
};

export type ProductReq = {
  sku: string;
  name: string;
  category_id: string;
  brand_id: string;
  unit_id: string;
  min_stock: number;
  weight: number;
  length: number;
  width: number;
  description: string;
};

export type ProductFormModel = {
  sku: string;
  name: string;
  categoryId: string;
  brandId: string;
  unitId: string;
  minStock: number;
  weight: number;
  length: number;
  width: number;
  description: string;
};
