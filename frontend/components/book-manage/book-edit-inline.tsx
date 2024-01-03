import { BookProps } from "@/types";
import React, { useEffect, useState } from "react";
import { Label } from "../ui/label";
import { Input } from "../ui/input";
import { FaPen } from "react-icons/fa";
import { Button } from "../ui/button";
import { toVND } from "@/lib/utils";

const BookEditInline = ({ currentBook }: { currentBook: BookProps }) => {
  const [book, setBook] = useState<BookProps | undefined>(undefined);
  useEffect(() => {
    // Optionally log the error to an error reporting service
    setBook(currentBook);
  }, [currentBook]);
  return (
    <div className="flex bg-background lg:flex-row flex-col p-4 px-6 gap-8">
      <div className="flex flex-1 flex-col gap-2 items-start ">
        <div className="flex items-center gap-2">
          <span className="w-[6rem]">Đầu sách: </span>
          <span className="font-semibold text-primary">
            {book?.bookTitle.name}
          </span>
        </div>
        <div className="flex items-center gap-2">
          <span className="w-[6rem]">Thể loại:</span>
          {book?.bookTitle.categories.map((cate) => {
            return (
              <div
                key={cate.id}
                className="rounded-xl flex px-3 py-1 h-fit outline-none text-sm text-primary whitespace-nowrap bg-blue-100 items-center gap-1 group"
              >
                {cate.name}
              </div>
            );
          })}
        </div>
        <div className="flex items-center gap-2">
          <span className="w-[6rem]">Tác giả:</span>
          {book?.bookTitle.authors.map((author) => {
            return (
              <div
                key={author.id}
                className="rounded-xl flex px-3 py-1 h-fit outline-none text-sm text-primary whitespace-nowrap bg-blue-100 items-center gap-1 group"
              >
                {author.name}
              </div>
            );
          })}
        </div>
      </div>
      <div className="flex flex-1 flex-col gap-2 items-start">
        <div className="flex items-center gap-2">
          <span className="min-w-[6rem]">Nhà xuất bản: </span>
          <span className="font-semibold">{book?.publisher.name}</span>
        </div>
        <div className="flex items-center gap-2">
          <span className="min-w-[6rem]">Lần tái bản: </span>
          <span className="font-semibold">{book?.edition}</span>
        </div>
        <div className="flex items-center gap-2">
          <span className="min-w-[6rem]">Giá nhập </span>
          <span className="font-semibold">{toVND(book?.importPrice!)}</span>
        </div>
      </div>
    </div>
  );
};

export default BookEditInline;
