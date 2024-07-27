"use client";
import { AuthorTable } from "@/components/book-manage/author-table";
import CreateAuthor from "@/components/book-manage/create-author";
import ImportSheet from "@/components/book-manage/import-sheet";
import { Button } from "@/components/ui/button";
import { toast } from "@/components/ui/use-toast";
import { endPoint } from "@/constants";
import { useCurrentUser } from "@/hooks/use-user";
import { getUser } from "@/lib/auth/action";
import createListAuthor from "@/lib/book/createListAuthor";
import { includesRoles } from "@/lib/utils";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import { useSWRConfig } from "swr";

const TableLayout = ({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) => {
  const { mutate } = useSWRConfig();

  const router = useRouter();
  const page = searchParams["page"] ?? "1";

  const handleAuthorAdded = (idAuthor: string) => {
    mutate(`${endPoint}/v1/authors?page=${page ?? 1}&limit=10`);
    router.refresh();
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

    const response: Promise<any> = createListAuthor({ names: categories });
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
      handleAuthorAdded("");
    }
  };

  const { currentUser } = useCurrentUser();
  return (
    <div className="col">
      <div className="flex flex-row justify-between items-center">
        <h1>Tác giả</h1>
        {currentUser &&
        includesRoles({
          currentUser: currentUser,
          allowedFeatures: ["AUTHOR_CREATE"],
        }) ? (
          <div className="flex gap-4">
            <CreateAuthor handleAuthorAdded={handleAuthorAdded}>
              <Button>Thêm tác giả</Button>
            </CreateAuthor>
          </div>
        ) : null}
      </div>
      <div className="flex flex-row flex-wrap gap-2"></div>
      <div className="mb-4 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <AuthorTable searchParams={searchParams} currentUser={currentUser} />
      </div>
    </div>
  );
};

export default TableLayout;
