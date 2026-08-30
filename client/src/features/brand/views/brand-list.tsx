import { Column } from "primereact/column";
import { DataTable } from "primereact/datatable";
import { useEffect, useState } from "react";
import { AddBrand, BrandActions } from "../components/actions";
import brandSrv from "../services/services";

export default function BrandList() {
  const [brands, setBrands] = useState([]);

  useEffect(() => {
    const loadBrands = async () => {
      const res = await brandSrv.list();
      setBrands(res);
    };

    loadBrands();
  }, []);

  return (
    <div className="flex flex-col gap-4 p-4 w-full">
      <div className="flex items-center justify-between">
        <h6 className="font-bold text-xl">Brands</h6>
        <AddBrand />
      </div>
      <div className="">
        <DataTable value={brands}>
          <Column field="name" header="Name" />
          <Column body={<BrandActions />} />
        </DataTable>
      </div>
    </div>
  );
}
