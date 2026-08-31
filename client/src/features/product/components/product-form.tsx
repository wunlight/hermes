import brandSrv from "@/features/brand/services/services";
import type { BrandOption } from "@/features/brand/types/types";
import categorySrv from "@/features/category/services/services";
import type { CategoryOption } from "@/features/category/types/types";
import unitSrv from "@/features/unit/services/services";
import type { UnitOption } from "@/features/unit/types/types";
import { Button } from "primereact/button";
import { Divider } from "primereact/divider";
import { Dropdown } from "primereact/dropdown";
import { InputNumber } from "primereact/inputnumber";
import { InputText } from "primereact/inputtext";
import { useEffect, useState } from "react";
import type { ProductFormModel } from "../types/types";

type ProductFormProps = {
  onSubmit: (value: ProductFormModel) => void;
  initialValue?: ProductFormModel;
};

export default function ProductForm({
  onSubmit,
  initialValue,
}: ProductFormProps) {
  const [form, setForm] = useState<ProductFormModel>(
    initialValue ?? {
      brandId: "",
      categoryId: "",
      description: "",
      length: 0,
      minStock: 0,
      name: "",
      sku: "",
      unitId: "",
      weight: 0,
      width: 0,
    },
  );

  const [categoryOptions, setCategoryOptions] = useState<CategoryOption[]>([]);
  const [brandOptions, setBrandOptions] = useState<BrandOption[]>([]);
  const [unitOptions, setUnitOptions] = useState<UnitOption[]>([]);

  function updateFormTextField(field: keyof ProductFormModel, value: string) {
    setForm((prev) => ({
      ...prev,
      [field]: value,
    }));
  }

  function updateFormNumberField(
    field: keyof ProductFormModel,
    value: number | null,
  ) {
    setForm((prev) => ({
      ...prev,
      [field]: value ?? 0,
    }));
  }

  async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();

    onSubmit(form);
  }

  useEffect(() => {
    categorySrv.getOptions().then((res) => {
      setCategoryOptions(res);
    });

    brandSrv.getOptions().then((res) => {
      setBrandOptions(res);
    });

    unitSrv.getOptions().then((res) => {
      setUnitOptions(res);
    });
  }, []);

  return (
    <form onSubmit={handleSubmit} className="flex flex-col gap-4">
      <div className="grid grid-cols-2 gap-4">
        <div className="flex flex-col gap-1">
          <label htmlFor="sku">SKU</label>
          <InputText
            id="sku"
            value={form.sku}
            onChange={(e) => updateFormTextField("sku", e.target.value)}
            placeholder="Enter product sku"
          />
        </div>
        <div className="flex flex-col gap-1">
          <label htmlFor="name">Name</label>
          <InputText
            id="name"
            value={form.name}
            onChange={(e) => updateFormTextField("name", e.target.value)}
            placeholder="Enter product name"
          />
        </div>
        <div className="col-span-2 flex flex-col gap-1">
          <label htmlFor="description">Description</label>
          <InputText
            id="description"
            value={form.description}
            onChange={(e) => updateFormTextField("description", e.target.value)}
            placeholder="Enter product description"
          />
        </div>
      </div>
      <div className="grid grid-cols-2 gap-4">
        <div className="flex flex-col gap-1">
          <label htmlFor="category">Category</label>
          <Dropdown
            id="category"
            value={form.categoryId}
            options={categoryOptions}
            optionValue="id"
            optionLabel="name"
            onChange={(e) => updateFormTextField("categoryId", e.value)}
            placeholder="Select product category"
          />
        </div>
        <div className="flex flex-col gap-1">
          <label htmlFor="brand">Brand</label>
          <Dropdown
            id="brand"
            value={form.brandId}
            options={brandOptions}
            optionValue="id"
            optionLabel="name"
            onChange={(e) => updateFormTextField("brandId", e.value)}
            placeholder="Select product brand"
          />
        </div>
        <div className="flex flex-col gap-1">
          <label htmlFor="unit">Unit</label>
          <Dropdown
            id="unit"
            value={form.unitId}
            options={unitOptions}
            optionValue="id"
            optionLabel="name"
            onChange={(e) => updateFormTextField("unitId", e.value)}
            placeholder="Select product unit"
          />
        </div>
      </div>
      <div className="grid grid-cols-2 gap-4">
        <div className="flex flex-col gap-1">
          <label htmlFor="min_stock">Minimal Stock</label>
          <InputNumber
            id="min_stock"
            value={form.minStock}
            onChange={(e) => updateFormNumberField("minStock", e.value)}
            placeholder="Enter product minimal stock"
          />
        </div>
        <div className="flex flex-col gap-1">
          <label htmlFor="weight">Weight</label>
          <InputNumber
            id="weight"
            value={form.weight}
            onChange={(e) => updateFormNumberField("weight", e.value)}
            placeholder="Enter product weight"
          />
        </div>
        <div className="flex flex-col gap-1">
          <label htmlFor="length">Length</label>
          <InputNumber
            id="length"
            value={form.length}
            onChange={(e) => updateFormNumberField("length", e.value)}
            placeholder="Enter product length"
          />
        </div>
        <div className="flex flex-col gap-1">
          <label htmlFor="width">Width</label>
          <InputNumber
            id="width"
            value={form.width}
            onChange={(e) => updateFormNumberField("width", e.value)}
            placeholder="Enter product width"
          />
        </div>
      </div>
      <Divider className="my-0!" />
      <div className="flex justify-end gap-3">
        <Button type="submit" label="Submit" />
      </div>
    </form>
  );
}
