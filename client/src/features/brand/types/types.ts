export type BrandDTO = {
  id: string;
  code: string;
  name: string;
  created_at: string;
  updated_at: string;
};

export type Brand = {
  id: string;
  code: string;
  name: string;
  createdAt: string;
  updatedAt: string;
};

export type BrandOption = {
  id: string;
  name: string;
};

export type BrandReq = {
  code: string;
  name: string;
};

export type BrandFormModel = {
  code: string;
  name: string;
};
