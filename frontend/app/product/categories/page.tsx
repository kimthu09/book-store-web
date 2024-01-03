import { Metadata } from "next";
import TableLayout from "./table-layout";
import { withAuth } from "@/lib/role/withAuth";
export const metadata: Metadata = {
  title: "Thể loại",
};
const CategoryPage = ({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) => {
  return <TableLayout searchParams={searchParams} />;
};

export default withAuth(CategoryPage, ["CATEGORY_VIEW"]);
