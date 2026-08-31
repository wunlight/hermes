export type CategoryDTO = {
  id: string;
  parent_id: string | null;
  code: string;
  name: string;
  created_at: string;
  updated_at: string;
};

export type Category = {
  id: string;
  parentId: string | null;
  code: string;
  name: string;
  createdAt: string;
  updatedAt: string;
};

export type CategoryOption = {
  id: string;
  name: string;
};

export type CategoryReq = {
  parent_id: string | null;
  code: string;
  name: string;
};

export type CategoryFormModel = {
  parentId: string | null;
  code: string;
  name: string;
};
