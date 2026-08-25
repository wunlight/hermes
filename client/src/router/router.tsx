import MainLayout from "@/core/layouts/main-layout";
import DashboardView from "@/features/dashboard/views/dashboard-view";
import { Route, Routes } from "react-router";
import LoginView from "../features/auth/views/login-view";
import ProtectedRoute from "./protected-route";
import PublicRoute from "./public-route";

export default function Router() {
  return (
    <Routes>
      <Route element={<PublicRoute />}>
        <Route path="/login" element={<LoginView />} />
      </Route>

      <Route element={<ProtectedRoute />}>
        <Route path="/" element={<MainLayout />}>
          <Route path="/dashboard" element={<DashboardView />} />
          <Route path="/inventory" element={<DashboardView />} />
          <Route path="/order" element={<DashboardView />} />
          <Route path="/report" element={<DashboardView />} />
          <Route path="/category" element={<DashboardView />} />
          <Route path="/brand" element={<DashboardView />} />
          <Route path="/unit" element={<DashboardView />} />
          <Route path="/product" element={<DashboardView />} />
        </Route>
      </Route>
    </Routes>
  );
}
