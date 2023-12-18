import getAllBooks from "@/lib/book/getAllBook";
import { BookTable } from "./table";
import { Book, PagingProps } from "@/types";

const TableLayout = async ({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) => {
  const page = searchParams["page"] ?? "1";
  const minSellPrice = searchParams["minSellPrice"] ?? undefined;
  const maxSellPrice = searchParams["maxSellPrice"] ?? undefined;
  const search = searchParams["search"] ?? undefined;
  const categories = searchParams["categories"] ?? undefined;
  const authors = searchParams["authors"] ?? undefined;
  const publisher = searchParams["publisher"] ?? undefined;
  const booksData: Promise<{ paging: PagingProps; data: Book[] }> = getAllBooks(
    {
      page: +page,
      maxSellPrice: maxSellPrice?.toString(),
      minSellPrice: minSellPrice?.toString(),
      categoryIds: categories?.toString(),
      authorIds: authors?.toString(),
      publisher: publisher?.toString(),
      search: search?.toString(),
    }
  );
  const books = await booksData;
  const totalPage = Math.ceil(books.paging.total / books.paging.limit);
  return <BookTable data={books.data} totalPage={totalPage} />;
};

export default TableLayout;
