import { Column } from "primereact/column";
import { DataTable } from "primereact/datatable";
import { useEffect, useState } from "react";
import { AddCategory, CategoryActions } from "../components/actions";
import categorySrv from "../services/services";
import type { Category } from "../types/types";

export default function CategoryList() {
  const [categories, setCategories] = useState<Category[]>([]);

  async function loadCategories() {
    const res = await categorySrv.list();
    setCategories(res);
  }

  useEffect(() => {
    categorySrv.list().then((res) => {
      setCategories(res);
    });
  }, []);

  return (
    <div className="flex flex-col gap-4 p-4 w-full">
      <div className="flex items-center justify-between">
        <h6 className="font-bold text-xl">Categories</h6>
        <AddCategory onSuccess={loadCategories} />
      </div>
      <div className="">
        <DataTable value={categories}>
          <Column field="name" header="Name" />
          <Column
            body={(category: Category) => (
              <CategoryActions id={category.id} onSuccess={loadCategories} />
            )}
          />
        </DataTable>
      </div>
    </div>
  );
}
