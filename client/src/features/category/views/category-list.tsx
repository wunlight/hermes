import { Button } from "primereact/button";
import categorySrv from "../services/services";

export default function CategoryList() {
  const loadCategories = async () => {
    const res = await categorySrv.list();
    console.log("response: ", res);
  };

  return (
    <div>
      <Button label="list category" onClick={() => loadCategories()} />
    </div>
  );
}
