"use client";
import { CategoryTable } from "@/components/book-manage/category-table";
import CreateCategory from "@/components/book-manage/create-category";
import Loading from "@/components/loading";
import TableSkeleton from "@/components/skeleton/table-skeleton";
import { Button } from "@/components/ui/button";
import { toast } from "@/components/ui/use-toast";
import { endPoint } from "@/constants";
import { useCurrentUser } from "@/hooks/use-user";
import createListCategory from "@/lib/book/createListCategory";
import { includesRoles } from "@/lib/utils";
import { useRouter } from "next/navigation";
import { Suspense, useEffect, useState } from "react";
import { useSWRConfig } from "swr";

const TableLayout = ({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) => {
  const { mutate } = useSWRConfig();

  const router = useRouter();
  const page = searchParams["page"] ?? "1";

  const handleCategoryAdded = (name: string) => {
    mutate(`${endPoint}/v1/categories?page=${page ?? 1}&limit=10`);
  };
  const handleFile = async (reader: FileReader) => {
    let categories: string[] = [];
    const ExcelJS = require("exceljs");
    const wb = new ExcelJS.Workbook();
    reader.onload = () => {
      const buffer = reader.result;
      wb.xlsx.load(buffer).then((workbook: any) => {
        workbook.eachSheet((sheet: any, id: any) => {
          sheet.eachRow((row: any, rowIndex: number) => {
            if (rowIndex > 1) {
              const name = row.getCell(1).value.toString();
              if (name && name != "") {
                categories.push(name);
              }
            }
          });
        });
      });
    };

    const response: Promise<any> = createListCategory({ names: categories });
    const responseData = await response;
    if (responseData.hasOwnProperty("errorKey")) {
      toast({
        variant: "destructive",
        title: "Có lỗi",
        description: responseData.message,
      });
    } else {
      toast({
        variant: "success",
        title: "Thành công",
        description: "Thêm thể loại thành công",
      });
      handleCategoryAdded("");
    }
  };
  const { currentUser } = useCurrentUser();
  return (
    <div className="col">
      <div className="flex flex-row justify-between items-center">
        <h1>Thể loại</h1>
        {currentUser &&
        includesRoles({
          currentUser: currentUser,
          allowedFeatures: ["CATEGORY_CREATE"],
        }) ? (
          <div className="flex gap-4">
            <CreateCategory handleCategoryAdded={handleCategoryAdded}>
              <Button>Thêm thể loại</Button>
            </CreateCategory>
          </div>
        ) : null}
      </div>
      <div className="flex flex-row flex-wrap gap-2"></div>
      <div className="mb-4 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <Suspense
          fallback={
            <TableSkeleton
              isHasExtensionAction={false}
              isHasFilter={false}
              isHasSearch={true}
              isHasChooseVisibleRow={false}
              isHasCheckBox={false}
              isHasPaging={true}
              numberRow={5}
              cells={[
                {
                  percent: 5,
                },
                {
                  percent: 1,
                },
              ]}
            />
          }
        >
          <CategoryTable
            searchParams={searchParams}
            currentUser={currentUser}
          />
        </Suspense>
      </div>
    </div>
  );
};

export default TableLayout;
