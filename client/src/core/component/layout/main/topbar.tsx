import { Avatar } from "primereact/avatar";

export default function Topbar() {
  return (
    <div className="flex items-center justify-between px-6 py-4 bg-white shadow">
      <div className="flex justify-start gap-3">
        <h3 className="font-bold text-xl">Amigo Motor</h3>
      </div>
      <div className="flex items-center justify-end gap-3">
        <Avatar label="A" shape="circle" className="size-9!" />
      </div>
    </div>
  );
}
