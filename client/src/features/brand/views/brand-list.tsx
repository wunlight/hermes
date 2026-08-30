import { Column } from "primereact/column";
import { DataTable } from "primereact/datatable";
import { useEffect, useState } from "react";
import { AddBrand, BrandActions } from "../components/actions";
import brandSrv from "../services/services";
import type { Brand } from "../types/types";

export default function BrandList() {
  const [brands, setBrands] = useState([]);

  async function loadBrands() {
    const res = await brandSrv.list();
    setBrands(res);
  }

  useEffect(() => {
    brandSrv.list().then((res) => {
      setBrands(res);
    });
  }, []);

  return (
    <div className="flex flex-col gap-4 p-4 w-full">
      <div className="flex items-center justify-between">
        <h6 className="font-bold text-xl">Brands</h6>
        <AddBrand onSuccess={loadBrands} />
      </div>
      <div className="">
        <DataTable value={brands}>
          <Column field="name" header="Name" />
          <Column
            body={(brand: Brand) => (
              <BrandActions id={brand.id} onSuccess={loadBrands} />
            )}
          />
        </DataTable>
      </div>
    </div>
  );
}
