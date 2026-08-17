import { create } from "zustand";
import { persist } from "zustand/middleware";

type AuthState = {
  userID: string | null;
  accessToken: string | null;
  isAuthenticated: boolean;

  setAuth: (userID: string, accessToken: string) => void;
  logout: () => void;
};

const useAuthStore = create<AuthState>()(
  persist(
    (set) => ({
      userID: null,
      accessToken: null,
      isAuthenticated: false,

      setAuth: (userID, accessToken) =>
        set({
          userID,
          accessToken,
          isAuthenticated: true,
        }),

      logout: () =>
        set({
          userID: null,
          accessToken: null,
          isAuthenticated: false,
        }),
    }),
    {
      name: "auth-state",
    },
  ),
);

export default useAuthStore;
