import TableLayout from "@/components/invoice/table-layout";
import Loading from "@/components/loading";
import TableSkeleton from "@/components/skeleton/table-skeleton";
import { withAuth } from "@/lib/role/withAuth";
import { Metadata } from "next";
import { Suspense } from "react";
export const metadata: Metadata = {
  title: "Hóa đơn",
};
const Invoice = ({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) => {
  return (
    <div className="col">
      <div className="flex flex-row justify-between ">
        <h1>Danh sách hóa đơn</h1>
        {/* <CreateDialog /> */}
      </div>

      <div className="my-3 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <Suspense
          fallback={
            <TableSkeleton
              isHasExtensionAction={false}
              isHasFilter={true}
              isHasSearch={true}
              isHasChooseVisibleRow={true}
              isHasCheckBox={false}
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
                  percent: 1,
                },
                {
                  percent: 1,
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

export default withAuth(Invoice, ["INVOICE_VIEW"]);
