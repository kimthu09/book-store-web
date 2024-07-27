import { RoleTable } from "@/components/staff/role-table";
import Link from "next/link";
import React from "react";
import { Metadata } from "next";
import { withAuth } from "@/lib/role/withAuth";
import { FaPlus } from "react-icons/fa";
export const metadata: Metadata = {
  title: "Phân quyền nhân viên",
};
const RoleSetting = () => {
  return (
    <div className="col">
      <div className="flex flex-row justify-between items-center">
        <h1>Danh sách phân quyền</h1>
        <div>
          <Link
            href="/staff/role/add"
            className="inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 bg-primary text-primary-foreground shadow hover:bg-primary/90 h-9 px-4 py-2"
          >
            <div className="flex flex-wrap gap-1 items-center">
              <FaPlus />
              Thêm phân quyền
            </div>
          </Link>
        </div>
      </div>

      <div className="mb-4 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <RoleTable />
      </div>
    </div>
  );
};

export default withAuth(RoleSetting, [""], true);
