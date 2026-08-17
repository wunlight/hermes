export type LoginRequest = {
  email: string;
  password: string;
};

export type LoginResponse = {
  user_id: string;
  access_token: string;
};

export type LoginDomainResponse = {
  userID: string;
  accessToken: string;
};
