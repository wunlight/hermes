import { NavLink } from "react-router";

type SidebarItem = {
  name: string;
  icon: string;
  to?: string;
  children?: SidebarItem[];
};

export default function Sidebar() {
  const itemList: SidebarItem[] = [
    {
      to: "/dashboard",
      name: "Dashboard",
      icon: "icon-[mdi--view-dashboard]",
    },
    {
      to: "/inventory",
      name: "Inventory",
      icon: "icon-[mdi--archive]",
    },
    {
      to: "/order",
      name: "Orders",
      icon: "icon-[mdi--cart]",
    },
    {
      to: "/report",
      name: "Reports",
      icon: "icon-[mdi--file-report]",
    },
    {
      name: "Master Data",
      icon: "icon-[mdi--cog]",
      children: [
        {
          to: "/category",
          name: "Categories",
          icon: "icon-[mdi--category]",
        },
        {
          to: "/brand",
          name: "Brands",
          icon: "icon-[mdi--category]",
        },
        {
          to: "/unit",
          name: "Units",
          icon: "icon-[mdi--category]",
        },
        {
          to: "/product",
          name: "Products",
          icon: "icon-[mdi--category]",
        },
        {
          to: "/warehouse",
          name: "Warehouses",
          icon: "icon-[mdi--category]",
        },
      ],
    },
  ];

  const renderItem = (item: SidebarItem) => {
    const content = (
      <>
        <span className={item.icon} />
        <span>{item.name}</span>
      </>
    );

    return (
      <div key={item.name}>
        {item.to ? (
          <NavLink
            to={item.to}
            className={({ isActive }) =>
              `flex h-9 items-center gap-4 rounded-l-full px-4 text-sm ${
                isActive
                  ? "bg-white text-indigo-950"
                  : "text-white hover:bg-indigo-700"
              }`
            }
          >
            {content}
          </NavLink>
        ) : (
          <div className="flex h-9 items-center gap-4 px-4 text-sm text-white hover:bg-indigo-700 rounded-l-full">
            {content}
          </div>
        )}

        {item.children && (
          <div className="ml-5 flex flex-col gap-1">
            {item.children.map(renderItem)}
          </div>
        )}
      </div>
    );
  };

  return (
    <div className="flex shrink-0 h-full w-60 flex-col gap-2.5 bg-indigo-800 py-5 pl-2.5">
      {itemList.map(renderItem)}
    </div>
  );
}
