import { Metadata } from "next";
import TableLayout from "./table-layout";
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

export default CategoryPage;
