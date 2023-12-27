import React from "react";
import { StaffTable } from "./staff-table";
import getAllStaff from "@/lib/staff/getAllStaff";
import { PagingProps, Staff } from "@/types";

const TableLayout = async ({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) => {
  const page = searchParams["page"] ?? "1";
  const active = searchParams["active"] ?? undefined;
  const search = searchParams["search"] ?? undefined;
  const staffsData: Promise<{ paging: PagingProps; data: Staff[] }> =
    getAllStaff({
      page: +page,
      isActive: active?.toString(),
      search: search?.toString(),
    });
  const staffs = await staffsData;
  const totalPage = Math.ceil(staffs.paging.total / staffs.paging.limit);
  return <StaffTable data={staffs.data} totalPage={totalPage} />;
};

export default TableLayout;
