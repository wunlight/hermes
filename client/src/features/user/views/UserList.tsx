import { Button } from "primereact/button";
import { Column } from "primereact/column";
import { DataTable } from "primereact/datatable";
import { useEffect, useState } from "react";
import { listUser } from "../service/service";
import type { User } from "../types/types";

const RowActions = () => {
  return (
    <div className="flex justify-end gap-3">
      <Button icon="icon-[mdi--edit]" className="aspect-square" />
      <Button icon="icon-[mdi--delete]" className="aspect-square" />
    </div>
  );
};

const UserList = () => {
  const [users, setUsers] = useState<User[]>([]);

  const getUser = async () => {
    const userList = await listUser();
    setUsers(userList);
  };

  useEffect(() => {
    getUser();
  }, []);

  return (
    <div className="p-4">
      <DataTable value={users}>
        <Column field="name" header="Name" />
        <Column body={<RowActions />} />
      </DataTable>
    </div>
  );
};

export default UserList;
