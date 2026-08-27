import { Column } from "primereact/column";
import { DataTable } from "primereact/datatable";
import { useEffect, useState } from "react";
import unitSrv from "../services/services";

export default function UnitList() {
  const [units, setUnits] = useState([]);

  useEffect(() => {
    const loadUnits = async () => {
      const res = await unitSrv.list();
      setUnits(res);
    };

    loadUnits();
  }, []);

  return (
    <div className="flex flex-col gap-4 p-4 w-full">
      <div className="flex items-center justify-between">
        <h6 className="font-bold text-xl">Units</h6>
      </div>
      <div className="">
        <DataTable value={units}>
          <Column field="code" header="Code" />
          <Column field="name" header="Name" />
        </DataTable>
      </div>
    </div>
  );
}
