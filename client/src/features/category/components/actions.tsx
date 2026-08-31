import { Button } from "primereact/button";
import { Dialog } from "primereact/dialog";
import { useState } from "react";
import categorySrv from "../services/services";
import type { CategoryFormModel } from "../types/types";
import CategoryForm from "./category-form";

type AddCategoryProps = {
  onSuccess: () => void;
};

type CategoryActionsProps = {
  id: string;
  onSuccess: () => void;
};

type EditCategoryProps = {
  id: string;
  onSuccess: () => void;
};

type DeleteCategoryProps = {
  id: string;
  onSuccess: () => void;
};

export function AddCategory({ onSuccess }: AddCategoryProps) {
  const [visible, setVisible] = useState(false);

  async function onCreateCategory(req: CategoryFormModel) {
    try {
      const res = await categorySrv.create(req);
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
        label="Add Category"
        icon="icon-[mdi--add]"
        onClick={() => setVisible(true)}
      />

      <Dialog
        header="Add New Category"
        visible={visible}
        onHide={() => setVisible(false)}
      >
        <CategoryForm onSubmit={onCreateCategory} />
      </Dialog>
    </>
  );
}

export function CategoryActions({ id, onSuccess }: CategoryActionsProps) {
  return (
    <div className="flex items-center justify-end gap-3">
      <EditCategory id={id} onSuccess={onSuccess} />
      <DeleteCategory id={id} onSuccess={onSuccess} />
    </div>
  );
}

export function EditCategory({ id, onSuccess }: EditCategoryProps) {
  const [visible, setVisible] = useState(false);
  const [category, setCategory] = useState<null | CategoryFormModel>(null);

  async function onUpdateCategory(req: CategoryFormModel) {
    try {
      const res = await categorySrv.update(id, req);
      console.log(res);

      setVisible(false);
      onSuccess();
    } catch (e) {
      console.error(e);
    }
  }

  async function openDialog() {
    try {
      const res = await categorySrv.getByID(id);

      setCategory({
        code: res.code,
        name: res.name,
        parentId: res.parentId,
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
        header="Edit Category"
        visible={visible}
        onHide={() => setVisible(false)}
      >
        {category && (
          <CategoryForm onSubmit={onUpdateCategory} initialValue={category} />
        )}
      </Dialog>
    </>
  );
}

export function DeleteCategory({ id, onSuccess }: DeleteCategoryProps) {
  const [visible, setVisible] = useState(false);

  async function onDeleteCategory() {
    try {
      await categorySrv.delete(id);

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
        key="delete-category"
        visible={visible}
        onHide={() => setVisible(false)}
      >
        <Button label="Delete" onClick={() => onDeleteCategory()} />
      </Dialog>
    </>
  );
}
