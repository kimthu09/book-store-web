import getAllBooks from "@/lib/getAllBook";
import { BookTable } from "./table";
import { Book } from "@/types";

const TableLayout = async ({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) => {
  const page = searchParams["page"] ?? "1";

  const booksData: Promise<{ paging: any; data: Book[] }> = getAllBooks(
    Number(page)
  );
  const books = await booksData;
  const totalPage = Math.floor(books.paging.total / books.paging.limit) + 1;

  return <BookTable data={books.data} totalPage={totalPage} />;
};

export default TableLayout;
