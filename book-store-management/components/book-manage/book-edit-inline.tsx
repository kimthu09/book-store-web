import { BookProps } from "@/types";
import React from "react";
import { Label } from "../ui/label";
import { Input } from "../ui/input";
import { FaPen } from "react-icons/fa";
import { Button } from "../ui/button";

const BookEditInline = (book: BookProps) => {
  return (
    <div className="flex bg-background p-3 gap-[5%] relative">
      <div className="absolute bottom-4 right-4">
        <Button
          variant={"ghost"}
          size={"icon"}
          className="rounded-full bg-blue-200/60 hover:bg-blue-200/90 text-primary hover:text-primary"
        >
          <FaPen />
        </Button>
      </div>
      <div className="flex basis-1/3 flex-col gap-2 items-start ">
        <div className="flex items-center gap-2">
          <span className="w-20">Đầu sách: </span>
          <span className="font-semibold text-primary">
            {book.bookTitle.name}
          </span>
        </div>
        <div className="flex items-center gap-2">
          <span className="w-20">Thể loại:</span>
          {book.bookTitle.categories.map((cate) => {
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
          <span className="w-20">Tác giả:</span>
          {book.bookTitle.authors.map((author) => {
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
      <div className="flex basis-1/3 flex-col gap-2 items-start">
        <div className="flex items-center gap-2">
          <span className="min-w-[6rem]">Nhà xuất bản: </span>
          <Input
            className=" bg-white"
            readOnly
            defaultValue={book.publisher.name}
          ></Input>
        </div>
        <div className="flex items-center gap-2">
          <span className="min-w-[6rem]">Lần tái bản: </span>
          <Input
            className=" bg-white"
            readOnly
            defaultValue={book.edition}
          ></Input>
        </div>
      </div>
    </div>
  );
};

export default BookEditInline;
