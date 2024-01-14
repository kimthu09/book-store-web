import Loading from "@/components/loading";

import { Suspense } from "react";
import { Metadata } from "next";
import { withAuth } from "@/lib/role/withAuth";
import TableLayout from "@/components/customer/table-layout";
import CreateDialog from "@/components/customer/create";
import { Button } from "@/components/ui/button";
export const metadata: Metadata = {
  title: "Quản lý nhà cung cấp",
};
function CustomerManage({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  return (
    <div className="col">
      <div className="flex flex-row justify-between ">
        <h1>Danh sách khách hàng</h1>
        <div className="flex gap-4">
          <CreateDialog>
            <Button className="lg:px-4 px-2 whitespace-nowrap">
              Thêm khách hàng
            </Button>
          </CreateDialog>
        </div>
      </div>

      <div className="my-3 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <Suspense fallback={<Loading />}>
          <TableLayout searchParams={searchParams} />
        </Suspense>
      </div>
    </div>
  );
}

export default withAuth(CustomerManage, ["CUSTOMER_VIEW"]);
