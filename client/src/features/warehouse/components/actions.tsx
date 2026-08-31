import { Button } from "primereact/button";
import { Dialog } from "primereact/dialog";
import { useState } from "react";
import warehouseSrv from "../services/services";
import type { WarehouseFormModel } from "../types/types";
import WarehouseForm from "./warehouse-form";

type AddWarehouseProps = {
  onSuccess: () => void;
};

type WarehouseActionsProps = {
  id: string;
  onSuccess: () => void;
};

type EditWarehouseProps = {
  id: string;
  onSuccess: () => void;
};

type DeleteWarehouseProps = {
  id: string;
  onSuccess: () => void;
};

export function AddWarehouse({ onSuccess }: AddWarehouseProps) {
  const [visible, setVisible] = useState(false);

  async function onCreateWarehouse(req: WarehouseFormModel) {
    try {
      const res = await warehouseSrv.create(req);
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
        label="Add Warehouse"
        icon="icon-[mdi--add]"
        onClick={() => setVisible(true)}
      />
      <Dialog
        header="Add New Warehouse"
        visible={visible}
        onHide={() => setVisible(false)}
      >
        <WarehouseForm onSubmit={onCreateWarehouse} />
      </Dialog>
    </>
  );
}

export function WarehouseActions({ id, onSuccess }: WarehouseActionsProps) {
  return (
    <div className="flex items-center justify-end gap-3">
      <EditWarehouse id={id} onSuccess={onSuccess} />
      <DeleteWarehouse id={id} onSuccess={onSuccess} />
    </div>
  );
}

export function EditWarehouse({ id, onSuccess }: EditWarehouseProps) {
  const [visible, setVisible] = useState(false);

  const [warehouse, setWarehouse] = useState<null | WarehouseFormModel>(null);

  async function onUpdateWarehouse(req: WarehouseFormModel) {
    try {
      const res = await warehouseSrv.update(id, req);
      console.log(res);

      setVisible(false);
      onSuccess();
    } catch (e) {
      console.error(e);
    }
  }

  async function openDialog() {
    try {
      const res = await warehouseSrv.getByID(id);

      setWarehouse({
        code: res.code,
        name: res.name,
        description: res.description,
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
        header="Edit Warehouse"
        visible={visible}
        onHide={() => setVisible(false)}
      >
        {warehouse && (
          <WarehouseForm
            onSubmit={onUpdateWarehouse}
            initialValue={warehouse}
          />
        )}
      </Dialog>
    </>
  );
}

export function DeleteWarehouse({ id, onSuccess }: DeleteWarehouseProps) {
  const [visible, setVisible] = useState(false);

  async function onDeleteWarehouse() {
    try {
      await warehouseSrv.delete(id);

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
        key="delete-warehouse"
        visible={visible}
        onHide={() => setVisible(false)}
      >
        <Button label="Delete" onClick={() => onDeleteWarehouse()} />
      </Dialog>
    </>
  );
}
