import Loading from "@/components/loading";
import CreateStaffDialog from "@/components/staff/create-staff-dialog";
import TableLayout from "@/components/staff/table-layout";

import { Metadata } from "next";
import { Suspense } from "react";
export const metadata: Metadata = {
  title: "Quản lý nhân viên",
};
const StaffManage = async ({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) => {
  return (
    <div className="col">
      <div className="flex flex-row justify-between items-center">
        <h1>Danh sách nhân viên</h1>
        <CreateStaffDialog />
      </div>
      <div className="mb-4 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <Suspense fallback={<Loading />}>
          <TableLayout searchParams={searchParams} />
        </Suspense>
      </div>
    </div>
  );
};

export default StaffManage;
