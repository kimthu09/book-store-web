import { Customer, PagingProps } from "@/types";
import React from "react";
import getListCustomer from "@/lib/customer/getListCustomer";
import { CustomerTable } from "./table";

const TableLayout = async ({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) => {
  const page = searchParams["page"] ?? "1";
  const maxPoint = searchParams["maxPoint"] ?? undefined;
  const minPoint = searchParams["minPoint"] ?? undefined;
  const search = searchParams["search"] ?? undefined;
  const booksData: Promise<{ paging: PagingProps; data: Customer[] }> =
    getListCustomer({
      page: +page,
      maxPoint: maxPoint?.toString(),
      minPoint: minPoint?.toString(),
      search: search?.toString(),
    });
  const books = await booksData;
  const totalPage = Math.ceil(books.paging.total / books.paging.limit);
  return <CustomerTable data={books.data} totalPage={totalPage} />;
};

export default TableLayout;
