import { Button } from "primereact/button";
import { Divider } from "primereact/divider";
import { InputText } from "primereact/inputtext";
import { useState } from "react";
import type { WarehouseFormModel } from "../types/types";

type WarehouseFormProps = {
  onSubmit: (value: WarehouseFormModel) => void;
  initialValue?: WarehouseFormModel;
};

export default function WarehouseForm({
  onSubmit,
  initialValue,
}: WarehouseFormProps) {
  const [form, setForm] = useState<WarehouseFormModel>(
    initialValue ?? {
      code: "",
      name: "",
      description: "",
    },
  );

  function updateFormField(field: keyof WarehouseFormModel, value: string) {
    setForm((prev) => ({
      ...prev,
      [field]: value,
    }));
  }

  async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();

    onSubmit(form);
  }

  return (
    <form onSubmit={handleSubmit} className="flex flex-col gap-4 w-72">
      <div className="flex flex-col gap-1">
        <label htmlFor="code">Code</label>
        <InputText
          id="code"
          value={form.code}
          onChange={(e) => updateFormField("code", e.target.value)}
          placeholder="Enter warehouse code"
        />
      </div>
      <div className="flex flex-col gap-1">
        <label htmlFor="name">Name</label>
        <InputText
          id="name"
          value={form.name}
          onChange={(e) => updateFormField("name", e.target.value)}
          placeholder="Enter warehouse name"
        />
      </div>
      <div className="flex flex-col gap-1">
        <label htmlFor="description">Description</label>
        <InputText
          id="description"
          value={form.description}
          onChange={(e) => updateFormField("description", e.target.value)}
          placeholder="Enter warehouse description"
        />
      </div>
      <Divider className="my-0!" />
      <div className="flex justify-end gap-3">
        <Button type="submit" label="Submit" />
      </div>
    </form>
  );
}
