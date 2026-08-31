import { Column } from "primereact/column";
import { DataTable } from "primereact/datatable";
import { useEffect, useState } from "react";
import { AddProduct, ProductActions } from "../components/actions";
import productSrv from "../services/services";
import type { Product } from "../types/types";

export default function ProductList() {
  const [products, setProducts] = useState<Product[]>([]);

  async function loadProducts() {
    const res = await productSrv.list();
    setProducts(res);
  }

  useEffect(() => {
    productSrv.list().then((res) => {
      setProducts(res);
    });
  }, []);

  return (
    <div className="flex flex-col gap-4 p-4 w-full">
      <div className="flex items-center justify-between">
        <h6 className="font-bold text-xl">Products</h6>
        <AddProduct onSuccess={loadProducts} />
      </div>
      <div className="">
        <DataTable value={products}>
          <Column field="sku" header="SKU" />
          <Column field="name" header="Name" />
          <Column field="categoryName" header="Category" />
          <Column field="brandName" header="Brand" />
          <Column field="description" header="Description" />
          <Column
            body={(product: Product) => (
              <ProductActions id={product.id} onSuccess={loadProducts} />
            )}
          />
        </DataTable>
      </div>
    </div>
  );
}
