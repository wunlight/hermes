import { Column } from "primereact/column";
import { DataTable } from "primereact/datatable";
import { useEffect, useState } from "react";
import { AddUnit, UnitActions } from "../components/actions";
import unitSrv from "../services/services";
import type { Unit } from "../types/types";

export default function UnitList() {
  const [units, setUnits] = useState<Unit[]>([]);

  async function loadUnits() {
    const res = await unitSrv.list();
    setUnits(res);
  }

  useEffect(() => {
    unitSrv.list().then((res) => {
      setUnits(res);
    });
  }, []);

  return (
    <div className="flex flex-col gap-4 p-4 w-full">
      <div className="flex items-center justify-between">
        <h6 className="font-bold text-xl">Units</h6>
        <AddUnit onSuccess={loadUnits} />
      </div>
      <div className="">
        <DataTable value={units}>
          <Column field="name" header="Name" />
          <Column
            body={(unit: Unit) => (
              <UnitActions id={unit.id} onSuccess={loadUnits} />
            )}
          />
        </DataTable>
      </div>
    </div>
  );
}
