export type LoginRequest = {
  email: string;
  password: string;
};

export type LoginDTOResponse = {
  user_id: string;
  access_token: string;
};

export type LoginResponse = {
  userID: string;
  accessToken: string;
};

export type RefreshDTOResponse = LoginDTOResponse;

export type RefreshResponse = LoginResponse;
