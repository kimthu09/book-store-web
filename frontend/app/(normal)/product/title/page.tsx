import { withAuth } from "@/lib/role/withAuth";
import TableLayout from "./table-layout";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Đầu sách",
};
const BookTitleScreen = () => {
  return <TableLayout />;
};

export default withAuth(BookTitleScreen, ["BOOK_TITLE_VIEW"]);
