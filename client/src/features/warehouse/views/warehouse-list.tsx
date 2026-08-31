import { Column } from "primereact/column";
import { DataTable } from "primereact/datatable";
import { useEffect, useState } from "react";
import { AddWarehouse, WarehouseActions } from "../components/actions";
import unitSrv from "../services/services";
import type { Warehouse } from "../types/types";

export default function WarehouseList() {
  const [units, setWarehouses] = useState<Warehouse[]>([]);

  async function loadWarehouses() {
    const res = await unitSrv.list();
    setWarehouses(res);
  }

  useEffect(() => {
    unitSrv.list().then((res) => {
      setWarehouses(res);
    });
  }, []);

  return (
    <div className="flex flex-col gap-4 p-4 w-full">
      <div className="flex items-center justify-between">
        <h6 className="font-bold text-xl">Warehouses</h6>
        <AddWarehouse onSuccess={loadWarehouses} />
      </div>
      <div className="">
        <DataTable value={units}>
          <Column field="name" header="Name" />
          <Column field="description" header="Description" />
          <Column
            body={(unit: Warehouse) => (
              <WarehouseActions id={unit.id} onSuccess={loadWarehouses} />
            )}
          />
        </DataTable>
      </div>
    </div>
  );
}
