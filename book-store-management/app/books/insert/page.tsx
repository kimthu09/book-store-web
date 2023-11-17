"use client";

import CategoryList from "@/components/category-list";
import BookList from "@/components/stock-manage/book-list";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Book } from "@/types";
import { useState } from "react";
import { LuCheck } from "react-icons/lu";

const InsertNewBook = () => {
  const [category, setCategory] = useState("");
  const [book, setBook] = useState<Partial<Book>>({ id: "" });
  const [isNew, setIsNew] = useState(false);
  const handleBookConfirm = (newValue: boolean) => {
    setIsNew(newValue);
    setCategory(book.category!);
  };

  const setBookHandler = (book: Partial<Book>) => {
    setBook(book);
    setCategory(book.category!);
  };
  return (
    <div className="col items-center">
      <div className="col xl:w-4/5 w-full xl:px-0 md:px-6 px-0">
        <div className="flex flex-row justify-between">
          <h1 className="font-medium text-xxl self-start">Thêm sách mới</h1>
          <Button>
            <div className="flex flex-wrap gap-1 items-center">
              <LuCheck />
              Thêm
            </div>
          </Button>
        </div>
        <form>
          <div className="flex flex-col flex-1 gap-4 lg:flex-row">
            <Card className="basis-3/5">
              <CardContent className="flex-col flex gap-5 mt-5">
                <div className="basis-1/3">
                  <Label htmlFor="masach">Mã sách</Label>
                  <Input
                    id="masach"
                    placeholder="Hệ thống sẽ tự sinh mã nếu để trống"
                    value={!isNew ? book.id : ""}
                    readOnly={!isNew}
                  />
                </div>
                <div className="basis-2/3">
                  <Label>Tên sách</Label>
                  <BookList
                    book={book}
                    setBook={setBookHandler}
                    isNew={isNew}
                    setIsNew={handleBookConfirm}
                  ></BookList>
                </div>
                <div>
                  {/* Category select */}
                  <Label>Thể loại</Label>
                  <CategoryList
                    category={!isNew ? book.category! : category}
                    setCategory={setCategory}
                    canAdd
                    readonly={!isNew}
                  />
                </div>
              </CardContent>
            </Card>
            <Card className="basis-2/5">
              <CardContent>
                <div className="flex flex-col gap-5 mt-5 ">
                  <div className="flex-1">
                    <Label>Nhà xuất bản</Label>
                    <Input />
                  </div>
                  <div className="flex flex-row md:gap-5 gap-3 lg:flex-col">
                    <div className="flex-1">
                      <Label>Lần tái bản</Label>
                      <Input type="number" min={0} />
                    </div>
                    <div className="flex-1">
                      <Label>Đơn giá</Label>
                      <Input type="number" min={0} />
                    </div>
                  </div>
                </div>
              </CardContent>
            </Card>
          </div>
        </form>
      </div>
    </div>
  );
};

export default InsertNewBook;
