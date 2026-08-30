import { Button } from "primereact/button";
import { Dialog } from "primereact/dialog";
import { useState } from "react";
import Form from "../components/form";

export function ProductActions() {
  return (
    <div className="flex items-center justify-end gap-3">
      <EditProduct />
      <DeleteProduct />
    </div>
  );
}

export function AddProduct() {
  const [visible, setVisible] = useState(false);

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
        <Form />
      </Dialog>
    </>
  );
}

export function EditProduct() {
  const [visible, setVisible] = useState(false);
  return (
    <>
      <Button
        icon="icon-[mdi--edit]"
        text
        rounded
        onClick={() => setVisible(true)}
      />
      <Dialog
        header="Edit Product"
        visible={visible}
        onHide={() => setVisible(false)}
      >
        <Form />
      </Dialog>
    </>
  );
}

export function DeleteProduct() {
  const [visible, setVisible] = useState(false);
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
        key="delete-category"
        visible={visible}
        onHide={() => setVisible(false)}
      >
        <p>Delete Confirmation</p>
      </Dialog>
    </>
  );
}
