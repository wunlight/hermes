import { create } from "zustand";

type AuthState = {
  userID: string | null;
  accessToken: string | null;
  isLoading: boolean;

  setAuth: (userID: string, accessToken: string) => void;
  resetAuth: () => void;
  setLoading: (value: boolean) => void;
};

const useAuthStore = create<AuthState>()((set) => ({
  userID: null,
  accessToken: null,
  isLoading: true,

  setAuth: (userID, accessToken) =>
    set({
      userID,
      accessToken,
    }),

  resetAuth: () =>
    set({
      userID: null,
      accessToken: null,
    }),

  setLoading: (value) =>
    set({
      isLoading: value,
    }),
}));

export default useAuthStore;
