import { PublisherTable } from "@/components/book-manage/publisher-table";

import { Metadata } from "next";
import TableLayout from "./table-layout";
export const metadata: Metadata = {
  title: "Nhà xuất bản",
};
const PublisherPage = ({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) => {
  return <TableLayout searchParams={searchParams} />;
};

export default PublisherPage;
