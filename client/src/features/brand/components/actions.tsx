import { Button } from "primereact/button";
import { Dialog } from "primereact/dialog";
import { useState } from "react";
import Form from "../components/form";

export function BrandActions() {
  return (
    <div className="flex items-center justify-end gap-3">
      <EditBrand />
      <DeleteBrand />
    </div>
  );
}

export function AddBrand() {
  const [visible, setVisible] = useState(false);

  return (
    <>
      <Button
        label="Add Brand"
        icon="icon-[mdi--add]"
        onClick={() => setVisible(true)}
      />
      <Dialog
        header="Add New Brand"
        visible={visible}
        onHide={() => setVisible(false)}
      >
        <Form />
      </Dialog>
    </>
  );
}

export function EditBrand() {
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
        header="Edit Brand"
        visible={visible}
        onHide={() => setVisible(false)}
      >
        <Form />
      </Dialog>
    </>
  );
}

export function DeleteBrand() {
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
