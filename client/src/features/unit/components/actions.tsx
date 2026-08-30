import { Button } from "primereact/button";
import { Dialog } from "primereact/dialog";
import { useState } from "react";
import Form from "../components/form";

export function UnitActions() {
  return (
    <div className="flex items-center justify-end gap-3">
      <EditUnit />
      <DeleteUnit />
    </div>
  );
}

export function AddUnit() {
  const [visible, setVisible] = useState(false);

  return (
    <>
      <Button
        label="Add Unit"
        icon="icon-[mdi--add]"
        onClick={() => setVisible(true)}
      />
      <Dialog
        header="Add New Unit"
        visible={visible}
        onHide={() => setVisible(false)}
      >
        <Form />
      </Dialog>
    </>
  );
}

export function EditUnit() {
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
        header="Edit Unit"
        visible={visible}
        onHide={() => setVisible(false)}
      >
        <Form />
      </Dialog>
    </>
  );
}

export function DeleteUnit() {
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
