import TableLayout from "@/components/book-manage/table-layout";
import Loading from "@/components/loading";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import React, { Suspense } from "react";
import { Metadata } from "next";
import { withAuth } from "@/lib/role/withAuth";
import { FaPlus } from "react-icons/fa";
import TableSkeleton from "@/components/skeleton/table-skeleton";
export const metadata: Metadata = {
  title: "Sách",
};
function BookManagement({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  return (
    <div className="col">
      <div className="flex flex-row justify-between items-center">
        <h1>Tất cả sách</h1>
        <div className="flex gap-4">
          <Link href="/product/books/add">
            <Button className="p-2">
              <div className="flex flex-wrap gap-1 items-center">
                <FaPlus />
                Thêm sách mới
              </div>
            </Button>
          </Link>
        </div>
      </div>
      <div className="flex flex-row flex-wrap gap-2"></div>
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
                  percent: 5,
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

export default withAuth(BookManagement, ["BOOK_VIEW"]);
