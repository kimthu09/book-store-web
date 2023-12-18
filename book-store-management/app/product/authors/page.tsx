import { AuthorTable } from "@/components/book-manage/author-table";
import CreateAuthor from "@/components/book-manage/create-author";
import { Button } from "@/components/ui/button";
import { Metadata } from "next";
import TableLayout from "./table-layout";
export const metadata: Metadata = {
  title: "Tác giả",
};
const AuthorPage = ({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) => {
  return <TableLayout searchParams={searchParams} />;
};

export default AuthorPage;
