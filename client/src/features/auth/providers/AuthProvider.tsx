import { useEffect, type ReactNode } from "react";
import { Refresh } from "../services/service";
import useAuthStore from "../stores/authStore";
import type { RefreshResponse } from "../types/types";

let bootstrapRefresh: Promise<RefreshResponse> | null = null;

const getBootstrapRefresh = () => {
  bootstrapRefresh ??= Refresh();
  return bootstrapRefresh;
};

const AuthProvider = ({ children }: { children: ReactNode }) => {
  const setAuth = useAuthStore((state) => state.setAuth);
  const resetAuth = useAuthStore((state) => state.resetAuth);
  const setLoading = useAuthStore((state) => state.setLoading);

  useEffect(() => {
    const refresh = async () => {
      try {
        const res = await getBootstrapRefresh();
        setAuth(res.userID, res.accessToken);
      } catch (e) {
        console.error(e);
        resetAuth();
      } finally {
        setLoading(false);
      }
    };

    refresh();
  }, [setAuth, resetAuth, setLoading]);

  return children;
};

export default AuthProvider;
