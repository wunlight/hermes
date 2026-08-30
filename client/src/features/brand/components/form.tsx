import { InputText } from "primereact/inputtext";

export default function Form() {
  return (
    <form className="flex flex-col gap-4">
      <div className="flex flex-col gap-1">
        <label htmlFor="code">Code</label>
        <InputText placeholder="Enter brand code" />
      </div>
      <div className="flex flex-col gap-1">
        <label htmlFor="name">Name</label>
        <InputText placeholder="Enter brand name" />
      </div>
    </form>
  );
}
