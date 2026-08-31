import { Button } from "primereact/button";
import { Dialog } from "primereact/dialog";
import { useState } from "react";
import productSrv from "../services/services";
import type { ProductFormModel } from "../types/types";
import ProductForm from "./product-form";

type AddProductProps = {
  onSuccess: () => void;
};

type ProductActionsProps = {
  id: string;
  onSuccess: () => void;
};

type EditProductProps = {
  id: string;
  onSuccess: () => void;
};

type DeleteProductProps = {
  id: string;
  onSuccess: () => void;
};

export function AddProduct({ onSuccess }: AddProductProps) {
  const [visible, setVisible] = useState(false);

  async function onCreateProduct(req: ProductFormModel) {
    try {
      const res = await productSrv.create(req);
      console.log(res);

      setVisible(false);
      onSuccess();
    } catch (e) {
      console.error(e);
    }
  }

  return (
    <>
      <Button
        label="Add Product"
        icon="icon-[mdi--add]"
        onClick={() => setVisible(true)}
      />
      <Dialog
        header="Add New Product"
        visible={visible}
        onHide={() => setVisible(false)}
      >
        <ProductForm onSubmit={onCreateProduct} />
      </Dialog>
    </>
  );
}

export function ProductActions({ id, onSuccess }: ProductActionsProps) {
  return (
    <div className="flex items-center justify-end gap-3">
      <EditProduct id={id} onSuccess={onSuccess} />
      <DeleteProduct id={id} onSuccess={onSuccess} />
    </div>
  );
}

export function EditProduct({ id, onSuccess }: EditProductProps) {
  const [visible, setVisible] = useState(false);
  const [product, setProduct] = useState<null | ProductFormModel>(null);

  async function onUpdateProduct(req: ProductFormModel) {
    try {
      const res = await productSrv.update(id, req);
      console.log(res);

      setVisible(false);
      onSuccess();
    } catch (e) {
      console.error(e);
    }
  }

  async function openDialog() {
    try {
      const res = await productSrv.getByID(id);

      setProduct({
        brandId: res.brandId,
        categoryId: res.categoryId,
        description: res.description,
        length: res.length,
        minStock: res.minStock,
        name: res.name,
        sku: res.sku,
        unitId: res.unitId,
        weight: res.weight,
        width: res.width,
      });

      setVisible(true);
    } catch (e) {
      console.error(e);

      setVisible(false);
    }
  }

  return (
    <>
      <Button
        icon="icon-[mdi--edit]"
        text
        rounded
        onClick={() => openDialog()}
      />
      <Dialog
        header="Edit Product"
        visible={visible}
        onHide={() => setVisible(false)}
      >
        {product && (
          <ProductForm onSubmit={onUpdateProduct} initialValue={product} />
        )}
      </Dialog>
    </>
  );
}

export function DeleteProduct({ id, onSuccess }: DeleteProductProps) {
  const [visible, setVisible] = useState(false);

  async function onDeleteProduct() {
    try {
      await productSrv.delete(id);

      setVisible(false);
      onSuccess();
    } catch (e) {
      console.error(e);
    }
  }

  return (
    <>
      <Button
        icon="icon-[mdi--delete]"
        severity="danger"
        text
        rounded
        onClick={() => setVisible(true)}
      />
      <Dialog
        key="delete-product"
        visible={visible}
        onHide={() => setVisible(false)}
      >
        <Button label="Delete" onClick={() => onDeleteProduct()} />
      </Dialog>
    </>
  );
}
