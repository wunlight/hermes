import { Button } from "primereact/button";
import { Divider } from "primereact/divider";
import { InputText } from "primereact/inputtext";
import { useState } from "react";
import type { UnitFormModel } from "../types/types";

type UnitFormProps = {
  onSubmit: (value: UnitFormModel) => void;
  initialValue?: UnitFormModel;
};

export default function UnitForm({ onSubmit, initialValue }: UnitFormProps) {
  const [form, setForm] = useState<UnitFormModel>(
    initialValue ?? {
      code: "",
      name: "",
    },
  );

  function updateFormField(field: keyof UnitFormModel, value: string) {
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
          placeholder="Enter unit code"
        />
      </div>
      <div className="flex flex-col gap-1">
        <label htmlFor="name">Name</label>
        <InputText
          id="name"
          value={form.name}
          onChange={(e) => updateFormField("name", e.target.value)}
          placeholder="Enter unit name"
        />
      </div>
      <Divider className="my-0!" />
      <div className="flex justify-end gap-3">
        <Button type="submit" label="Submit" />
      </div>
    </form>
  );
}
