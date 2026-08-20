import { Button } from "primereact/button";
import { Link } from "react-router";

const DashboardView = () => {
  return (
    <Link to="/users">
      <Button label="go to user list" />
    </Link>
  );
};

export default DashboardView;
