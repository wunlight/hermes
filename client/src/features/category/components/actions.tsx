import { Button } from "primereact/button";
import { Dialog } from "primereact/dialog";
import { useState } from "react";
import Form from "../components/form";

export function CategoryActions() {
  return (
    <div className="flex items-center justify-end gap-3">
      <EditCategory />
      <DeleteCategory />
    </div>
  );
}

export function AddCategory() {
  const [visible, setVisible] = useState(false);

  return (
    <>
      <Button
        label="Add Category"
        icon="icon-[mdi--add]"
        onClick={() => setVisible(true)}
      />
      <Dialog
        header="Add New Category"
        visible={visible}
        onHide={() => setVisible(false)}
      >
        <Form />
      </Dialog>
    </>
  );
}

export function EditCategory() {
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
        header="Edit Category"
        visible={visible}
        onHide={() => setVisible(false)}
      >
        <Form />
      </Dialog>
    </>
  );
}

export function DeleteCategory() {
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
