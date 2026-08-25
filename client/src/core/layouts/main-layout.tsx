import { Outlet } from "react-router";
import Sidebar from "../component/layout/main/sidebar";
import Topbar from "../component/layout/main/topbar";

export default function MainLayout() {
  return (
    <div className="flex flex-col h-dvh overflow-hidden">
      <Topbar />

      <div className="flex h-full">
        <Sidebar />

        <Outlet />
      </div>
    </div>
  );
}
