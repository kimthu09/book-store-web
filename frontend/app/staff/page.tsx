import Loading from "@/components/loading";
import TableSkeleton from "@/components/skeleton/table-skeleton";
import CreateStaffDialog from "@/components/staff/create-staff-dialog";
import TableLayout from "@/components/staff/table-layout";
import { withAuth } from "@/lib/role/withAuth";

import { Metadata } from "next";
import { Suspense } from "react";
export const metadata: Metadata = {
  title: "Quản lý nhân viên",
};
const StaffManage = ({
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
        <Suspense
          fallback={
            <TableSkeleton
              isHasExtensionAction={true}
              isHasFilter={true}
              isHasSearch={true}
              isHasChooseVisibleRow={true}
              isHasCheckBox={true}
              isHasPaging={true}
              numberRow={5}
              cells={[
                {
                  percent: 1,
                },
                {
                  percent: 2,
                },
                {
                  percent: 2,
                },
                {
                  percent: 1,
                },
                {
                  percent: 2,
                },
                {
                  percent: 1,
                },
                {
                  percent: 1,
                },
              ]}
            />
          }
        >
          <TableLayout searchParams={searchParams} />
        </Suspense>
      </div>
    </div>
  );
};

export default withAuth(StaffManage, ["USER_VIEW"]);
