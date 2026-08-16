import { Route, Routes } from "react-router";
import LoginView from "./features/auth/views/LoginView";

const Router = () => {
  return (
    <Routes>
      <Route path="/login" element={<LoginView />} />
    </Routes>
  );
};

export default Router;
