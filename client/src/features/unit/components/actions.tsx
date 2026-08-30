import { Button } from "primereact/button";
import { Dialog } from "primereact/dialog";
import { useState } from "react";
import unitSrv from "../services/services";
import type { UnitFormModel } from "../types/types";
import UnitForm from "./unit-form";

type AddUnitProps = {
  onSuccess: () => void;
};

type UnitActionsProps = {
  id: string;
  onSuccess: () => void;
};

type EditUnitProps = {
  id: string;
  onSuccess: () => void;
};

type DeleteUnitProps = {
  id: string;
  onSuccess: () => void;
};

export function AddUnit({ onSuccess }: AddUnitProps) {
  const [visible, setVisible] = useState(false);

  async function onCreateUnit(req: UnitFormModel) {
    try {
      const res = await unitSrv.create(req);
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
        label="Add Unit"
        icon="icon-[mdi--add]"
        onClick={() => setVisible(true)}
      />
      <Dialog
        header="Add New Unit"
        visible={visible}
        onHide={() => setVisible(false)}
      >
        <UnitForm onSubmit={onCreateUnit} />
      </Dialog>
    </>
  );
}

export function UnitActions({ id, onSuccess }: UnitActionsProps) {
  return (
    <div className="flex items-center justify-end gap-3">
      <EditUnit id={id} onSuccess={onSuccess} />
      <DeleteUnit id={id} onSuccess={onSuccess} />
    </div>
  );
}

export function EditUnit({ id, onSuccess }: EditUnitProps) {
  const [visible, setVisible] = useState(false);

  const [unit, setUnit] = useState<null | UnitFormModel>(null);

  async function onUpdateUnit(req: UnitFormModel) {
    try {
      const res = await unitSrv.update(id, req);
      console.log(res);

      setVisible(false);
      onSuccess();
    } catch (e) {
      console.error(e);
    }
  }

  async function openDialog() {
    try {
      const res = await unitSrv.getByID(id);

      setUnit({
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
        header="Edit Unit"
        visible={visible}
        onHide={() => setVisible(false)}
      >
        {unit && <UnitForm onSubmit={onUpdateUnit} initialValue={unit} />}
      </Dialog>
    </>
  );
}

export function DeleteUnit({ id, onSuccess }: DeleteUnitProps) {
  const [visible, setVisible] = useState(false);

  async function onDeleteUnit() {
    try {
      await unitSrv.delete(id);

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
        key="delete-unit"
        visible={visible}
        onHide={() => setVisible(false)}
      >
        <Button label="Delete" onClick={() => onDeleteUnit()} />
      </Dialog>
    </>
  );
}
