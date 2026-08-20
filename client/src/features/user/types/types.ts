export type UserDTO = {
  id: string;
  email: string;
  name: string;
  created_at: string;
  updated_at: string;
  deleted_at: string | null;
};

export type User = {
  id: string;
  email: string;
  name: string;
};
