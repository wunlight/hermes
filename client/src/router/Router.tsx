import DashboardView from "@/features/dashboard/views/DashboardView";
import UserList from "@/features/user/views/UserList";
import { Route, Routes } from "react-router";
import LoginView from "../features/auth/views/LoginView";
import ProtectedRoute from "./ProtectedRoute";
import PublicRoute from "./PublicRoute";

const Router = () => {
  return (
    <Routes>
      <Route element={<PublicRoute />}>
        <Route path="/login" element={<LoginView />} />
      </Route>

      <Route element={<ProtectedRoute />}>
        <Route path="/" element={<DashboardView />} />
        <Route path="/users" element={<UserList />} />
      </Route>
    </Routes>
  );
};

export default Router;
