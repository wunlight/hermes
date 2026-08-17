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
        <Route path="/" />
      </Route>
    </Routes>
  );
};

export default Router;
