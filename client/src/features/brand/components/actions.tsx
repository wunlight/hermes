import { Button } from "primereact/button";
import { Dialog } from "primereact/dialog";
import { useState } from "react";
import brandSrv from "../services/services";
import type { BrandFormModel } from "../types/types";
import BrandForm from "./brand-form";

type AddBrandProps = {
  onSuccess: () => void;
};

type BrandActionsProps = {
  id: string;
  onSuccess: () => void;
};

type EditBrandProps = {
  id: string;
  onSuccess: () => void;
};

type DeleteBrandProps = {
  id: string;
  onSuccess: () => void;
};

export function AddBrand({ onSuccess }: AddBrandProps) {
  const [visible, setVisible] = useState(false);

  async function onCreateBrand(req: BrandFormModel) {
    try {
      const res = await brandSrv.create(req);
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
        label="Add Brand"
        icon="icon-[mdi--add]"
        onClick={() => setVisible(true)}
      />
      <Dialog
        header="Add New Brand"
        visible={visible}
        onHide={() => setVisible(false)}
      >
        <BrandForm onSubmit={onCreateBrand} />
      </Dialog>
    </>
  );
}

export function BrandActions({ id, onSuccess }: BrandActionsProps) {
  return (
    <div className="flex items-center justify-end gap-3">
      <EditBrand id={id} onSuccess={onSuccess} />
      <DeleteBrand id={id} onSuccess={onSuccess} />
    </div>
  );
}

export function EditBrand({ id, onSuccess }: EditBrandProps) {
  const [visible, setVisible] = useState(false);

  const [brand, setBrand] = useState<null | BrandFormModel>(null);

  async function onUpdateBrand(req: BrandFormModel) {
    try {
      const res = await brandSrv.update(id, req);
      console.log(res);

      setVisible(false);
      onSuccess();
    } catch (e) {
      console.error(e);
    }
  }

  async function openDialog() {
    try {
      const res = await brandSrv.getByID(id);

      setBrand({
        code: res.code,
        name: res.name,
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
        header="Edit Brand"
        visible={visible}
        onHide={() => setVisible(false)}
      >
        {brand && <BrandForm onSubmit={onUpdateBrand} initialValue={brand} />}
      </Dialog>
    </>
  );
}

export function DeleteBrand({ id, onSuccess }: DeleteBrandProps) {
  const [visible, setVisible] = useState(false);

  async function onDeleteBrand() {
    try {
      await brandSrv.delete(id);

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
        key="delete-brand"
        visible={visible}
        onHide={() => setVisible(false)}
      >
        <Button label="Delete" onClick={() => onDeleteBrand()} />
      </Dialog>
    </>
  );
}
