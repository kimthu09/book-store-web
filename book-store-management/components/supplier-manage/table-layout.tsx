import getAllSupplier from "@/lib/supplier/getAllSupplier";
import { PagingProps, Supplier } from "@/types";
import React from "react";
import { SupplierTable } from "./table";

const TableLayout = async ({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) => {
  const page = searchParams["page"] ?? "1";
  const maxDebt = searchParams["maxDebt"] ?? undefined;
  const minDebt = searchParams["minDebt"] ?? undefined;
  const search = searchParams["search"] ?? undefined;
  const booksData: Promise<{ paging: PagingProps; data: Supplier[] }> =
    getAllSupplier({
      page: +page,
      maxDebt: maxDebt?.toString(),
      minDebt: minDebt?.toString(),
      search: search?.toString(),
    });
  const books = await booksData;
  const totalPage = Math.ceil(books.paging.total / books.paging.limit);
  return <SupplierTable data={books.data} totalPage={totalPage} />;
};

export default TableLayout;
