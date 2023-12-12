import { BookTable } from "@/components/book-manage/table";
import { Button } from "@/components/ui/button";
import getAllBooks from "@/lib/getAllBook";
import { Book } from "@/types";
import Link from "next/link";
import React from "react";
import { AiOutlinePlus } from "react-icons/ai";

async function BookManagement({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  const page = searchParams["page"] ?? "1";

  const booksData: Promise<{ paging: any; data: Book[] }> = getAllBooks(
    Number(page)
  );
  const books = await booksData;
  const totalPage = Math.floor(books.paging.total / books.paging.limit) + 1;
  return (
    <div className="col">
      <div className="flex flex-row justify-between items-center">
        <h1>Tất cả sách</h1>
        <div className="flex gap-4">
          <Link href="/books/insert">
            <Button className="p-2">
              <div className="flex flex-wrap gap-1 items-center">
                <AiOutlinePlus />
                Thêm sách mới
              </div>
            </Button>
          </Link>
        </div>
      </div>
      <div className="flex flex-row flex-wrap gap-2"></div>

      <div className="mb-4 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <BookTable data={books.data} totalPage={totalPage} />
      </div>
    </div>
  );
}

export default BookManagement;
