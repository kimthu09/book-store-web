import { Metadata } from "next";
import TableLayout from "./table-layout";
import { withAuth } from "@/lib/role/withAuth";
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

export default withAuth(AuthorPage, ["AUTHOR_VIEW"]);
