import { Column } from "primereact/column";
import { DataTable } from "primereact/datatable";
import { useEffect, useState } from "react";
import { AddProduct, ProductActions } from "../components/actions";
import productSrv from "../services/services";

export default function ProductList() {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    const loadProducts = async () => {
      const res = await productSrv.list();
      setProducts(res);
    };

    loadProducts();
  }, []);

  return (
    <div className="flex flex-col gap-4 p-4 w-full">
      <div className="flex items-center justify-between">
        <h6 className="font-bold text-xl">Products</h6>
        <AddProduct />
      </div>
      <div className="">
        <DataTable value={products}>
          <Column field="name" header="Name" />
          <Column body={<ProductActions />} />
        </DataTable>
      </div>
    </div>
  );
}
