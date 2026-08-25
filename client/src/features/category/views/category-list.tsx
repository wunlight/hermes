import { Column } from "primereact/column";
import { DataTable } from "primereact/datatable";
import { useEffect, useState } from "react";
import { AddCategory, CategoryActions } from "../components/actions";
import categorySrv from "../services/services";

export default function CategoryList() {
  const [categories, setCategories] = useState([]);

  useEffect(() => {
    const loadCategories = async () => {
      const res = await categorySrv.list();
      setCategories(res);
    };

    loadCategories();
  }, []);

  return (
    <div className="flex flex-col gap-4 p-4 w-full">
      <div className="flex items-center justify-between">
        <h6 className="font-bold text-xl">Categories</h6>
        <AddCategory />
      </div>
      <div className="">
        <DataTable value={categories}>
          <Column field="code" header="Code" />
          <Column field="name" header="Name" />
          <Column body={<CategoryActions />} />
        </DataTable>
      </div>
    </div>
  );
}
