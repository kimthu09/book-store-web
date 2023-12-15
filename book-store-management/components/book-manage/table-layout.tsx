import getAllBooks from "@/lib/book/getAllBook";
import { BookTable } from "./table";
import { Book, PagingProps } from "@/types";

const TableLayout = async ({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) => {
  const page = searchParams["page"] ?? "1";

  const booksData: Promise<{ paging: PagingProps; data: Book[] }> = getAllBooks(
    Number(page)
  );
  const books = await booksData;
  const totalPage = Math.ceil(books.paging.total / books.paging.limit);
  return <BookTable data={books.data} totalPage={totalPage} />;
};

export default TableLayout;
