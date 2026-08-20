import useAuthStore from "@/features/auth/stores/authStore";
import { Navigate, Outlet } from "react-router";

const ProtectedRoute = () => {
  const { accessToken, isLoading } = useAuthStore();

  if (isLoading) {
    return <div className="grid place-content-center h-dvh">Loading...</div>;
  }

  if (accessToken === null) {
    return <Navigate to="/login" replace />;
  }

  return <Outlet />;
};

export default ProtectedRoute;
