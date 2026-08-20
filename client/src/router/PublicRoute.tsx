import useAuthStore from "@/features/auth/stores/authStore";
import { Navigate, Outlet } from "react-router";

const PublicRoute = () => {
  const { accessToken } = useAuthStore();

  if (accessToken !== null) {
    return <Navigate to="/" replace />;
  }

  return <Outlet />;
};

export default PublicRoute;
