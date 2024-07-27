import Loading from "@/components/loading";
import CreateDialog from "@/components/supplier-manage/create";
import TableLayout from "@/components/supplier-manage/table-layout";

import { Suspense } from "react";
import { Metadata } from "next";
import { withAuth } from "@/lib/role/withAuth";
import { includesRoles } from "@/lib/utils";
import TableSkeleton from "@/components/skeleton/table-skeleton";
import { Button } from "@/components/ui/button";
export const metadata: Metadata = {
  title: "Quản lý nhà cung cấp",
};
function SupplierManage({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  return (
    <div className="col">
      <div className="flex flex-row justify-between ">
        <h1>Danh sách nhà cung cấp</h1>
        <div className="flex gap-4">
          <CreateDialog>
            <Button className="lg:px-4 px-2 whitespace-nowrap">
              Thêm nhà cung cấp
            </Button>
          </CreateDialog>
        </div>
      </div>

      <div className="my-3 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
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
}

export default withAuth(SupplierManage, ["SUPPLIER_VIEW"]);
