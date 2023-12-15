"use client";

import CategoryList from "@/components/category-list";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { books } from "@/constants";
import { useState } from "react";
import { AiOutlineClose } from "react-icons/ai";
import { LuCheck } from "react-icons/lu";

const EditBook = ({ searchParams }: { searchParams: { id: string } }) => {
  const book = books.find((item) => item.id === searchParams.id);
  const [category, setCategory] = useState(book?.category!);
  const [readOnly, setReadOnly] = useState(true);
  return (
    <div className="col items-center">
      <div className="col xl:w-4/5 w-full xl:px-0 md:px-6 px-0">
        <div className="flex flex-row justify-between">
          <h1 className="font-medium text-xxl self-start">
            Mã sách: {book?.id}
          </h1>
          {!readOnly ? (
            <div className="flex gap-2">
              <Button
                variant={"outline"}
                className="bg-white border-primary text-primary hover:text-primary"
                onClick={() => setReadOnly(true)}
              >
                <div className="flex flex-wrap gap-1 items-center">
                  <AiOutlineClose />
                  Huỷ
                </div>
              </Button>
              <Button onClick={() => setReadOnly(true)}>
                <div className="flex flex-wrap gap-1 items-center">
                  <LuCheck />
                  Lưu
                </div>
              </Button>
            </div>
          ) : (
            <Button
              onClick={() => {
                setReadOnly(false);
              }}
            >
              Chỉnh sửa
            </Button>
          )}
        </div>
        <form>
          <div className="flex flex-col flex-1 gap-4">
            <Card>
              <CardContent>
                <div className="flex flex-col gap-5 mt-5 lg:flex-row">
                  <div className="lg:basis-2/3">
                    <Label htmlFor="name">Tên sách</Label>
                    <Input
                      id="name"
                      defaultValue={book?.name}
                      readOnly={readOnly}
                    />
                  </div>
                  <div className="lg:basis-1/3">
                    {/* Category select */}
                    <Label>Thể loại</Label>
                    <CategoryList
                      category={category}
                      setCategory={setCategory}
                      canAdd
                      readonly={readOnly}
                    />
                  </div>
                </div>
              </CardContent>
            </Card>
            <Card>
              <CardContent>
                <div className="flex flex-col gap-5 mt-5 ">
                  <div className="flex-1">
                    <Label>Nhà xuất bản</Label>
                    <Input readOnly={readOnly} defaultValue={book?.nxb} />
                  </div>
                  <div className="flex flex-row gap-5">
                    <div className="flex-1">
                      <Label>Lần tái bản</Label>
                      <Input
                        type="number"
                        readOnly={readOnly}
                        defaultValue={book?.price}
                        min={0}
                      />
                    </div>
                    <div className="flex-1">
                      <Label>Đơn giá</Label>
                      <Input
                        readOnly={readOnly}
                        defaultValue={book?.price}
                        type="number"
                        min={0}
                      />
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

export default EditBook;
