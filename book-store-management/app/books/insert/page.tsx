"use client";

import AuthorList from "@/components/author-list";
import CategoryList from "@/components/category-list";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogContent,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { useToast } from "@/components/ui/use-toast";
import createBook from "@/lib/createBook";
import getAllAuthor from "@/lib/getAllAuthor";
import getAllCategory from "@/lib/getAllCategory";
import Link from "next/link";
import { useState } from "react";
import { SubmitHandler, useFieldArray, useForm } from "react-hook-form";
import { AiOutlineClose } from "react-icons/ai";
import { LuCheck } from "react-icons/lu";

export type FormValues = {
  name: string;
  desc: string;
  idBook: string;
  authorIds: {
    idAuthor: string;
  }[];

  categoryIds: {
    idCate: string;
  }[];
};

const InsertNewBook = () => {
  const { toast } = useToast();

  const form = useForm<FormValues>({
    defaultValues: {
      idBook: "",
      name: "",
      desc: "",
      categoryIds: [],
    },
  });
  const [creating, setCreating] = useState(true);
  const { register, handleSubmit, control, watch } = form;

  const {
    fields: fieldsCate,
    append: appendCate,
    remove: removeCate,
    update: updateCate,
  } = useFieldArray({
    control: control,
    name: "categoryIds",
  });

  const {
    fields: fieldsAuthor,
    append: appendAuthor,
    remove: removeAuthor,
    update: updateAuthor,
  } = useFieldArray({
    control: control,
    name: "authorIds",
  });
  const onSubmit: SubmitHandler<FormValues> = async (data) => {
    console.log(data);
    if (data.categoryIds.length < 1) {
      toast({
        title: "Chưa chọn thể loại",
        description: "Vui lòng chọn ít nhất một thể loại",
      });
      return;
    }
    if (data.authorIds.length < 1) {
      toast({
        title: "Chưa chọn tác giả",
        description: "Vui lòng chọn ít nhất một tác giả",
      });
      return;
    }
    const response: Promise<any> = createBook({
      id: data.idBook,
      name: data.name,
      desc: data.desc,
      categoryIds: data.categoryIds.map((item) => item.idCate),
      authorIds: data.authorIds.map((item) => item.idAuthor),
    });
    const responseData = await response;
    console.log("hi there... " + JSON.stringify(responseData));
    if (responseData.hasOwnProperty("data")) {
      if (responseData.data.id != "") {
        setCreating(false);
      }
    } else {
      toast({
        title: "Mã sách đã tồn tại",
        description: "Vui lòng nhập một mã sách khác",
      });
    }
  };
  const { categories, isLoading, isError } = getAllCategory();
  const {
    authors,
    isLoading: isAuthorLoading,
    isError: isAuthorError,
  } = getAllAuthor();

  if (isError || isAuthorError) return <div>Failed to load</div>;
  if (!categories || !authors) {
    console.log(categories);
    return <div>Loading...</div>;
  } else
    return (
      <div className="col items-center">
        <form
          onSubmit={handleSubmit(onSubmit)}
          className="col xl:w-4/5 w-full xl:px-0 md:px-6 px-0"
        >
          <div className="flex flex-row justify-between">
            <h1 className="font-medium text-xxl self-start">Thêm sách mới</h1>
            <Button type="submit">
              <div className="flex flex-wrap gap-1 items-center">
                <LuCheck />
                Thêm
              </div>
            </Button>
          </div>
          <div>
            <div className="flex flex-col flex-1 gap-4 lg:flex-row">
              <Card className="flex-1">
                <CardContent className="flex-col flex gap-5 mt-5">
                  <div className="basis-1/3">
                    <Label htmlFor="masach">Mã sách</Label>
                    <Input
                      {...register("idBook")}
                      id="masach"
                      placeholder="Hệ thống sẽ tự sinh mã nếu để trống"
                      // value={!isNew ? book.id : ""}
                      // readOnly={!isNew}
                    />
                  </div>
                  <div className="flex-1">
                    <Label>Tên sách</Label>
                    <Input required {...register("name")}></Input>
                  </div>

                  <div>
                    {/* Category select */}
                    <Label>Thể loại</Label>
                    <CategoryList
                      checkedCategory={fieldsCate.map((cate) => cate.idCate)}
                      onCheckChanged={(idCate) => {
                        const selectedIndex = fieldsCate.findIndex(
                          (cate) => cate.idCate === idCate
                        );
                        if (selectedIndex > -1) {
                          removeCate(selectedIndex);
                        } else {
                          appendCate({ idCate: idCate });
                        }
                      }}
                    />
                    <div className="flex flex-wrap gap-2 mt-3">
                      {fieldsCate.map((cate, index) => (
                        <div
                          key={cate.id}
                          className="rounded-xl flex  px-3 py-1 h-fit outline-none text-sm text-primary  bg-blue-100 items-center gap-1 group"
                        >
                          {
                            categories.find((item) => item.id === cate.idCate)
                              ?.name
                          }
                          <div className="cursor-pointer w-4">
                            <AiOutlineClose className="group-hover:hidden" />
                            <AiOutlineClose
                              color="red"
                              fill="red"
                              className="text-primary group-hover:flex hidden h-4 w-4"
                              onClick={() => {
                                removeCate(index);
                              }}
                            />
                          </div>
                        </div>
                      ))}
                    </div>
                  </div>
                  <div>
                    {/* Category select */}
                    <Label>Tác giả</Label>
                    <AuthorList
                      checkedAuthor={fieldsAuthor.map(
                        (author) => author.idAuthor
                      )}
                      onCheckChanged={(idAuthor) => {
                        const selectedIndex = fieldsAuthor.findIndex(
                          (cate) => cate.idAuthor === idAuthor
                        );
                        if (selectedIndex > -1) {
                          removeAuthor(selectedIndex);
                        } else {
                          appendAuthor({ idAuthor: idAuthor });
                        }
                      }}
                    />
                    <div className="flex flex-wrap gap-2 mt-3">
                      {fieldsAuthor.map((author, index) => (
                        <div
                          key={author.id}
                          className="rounded-xl flex  px-3 py-1 h-fit outline-none text-sm text-primary  bg-blue-100 items-center gap-1 group"
                        >
                          {
                            authors.find((item) => item.id === author.idAuthor)
                              ?.name
                          }
                          <div className="cursor-pointer w-4">
                            <AiOutlineClose className="group-hover:hidden" />
                            <AiOutlineClose
                              color="red"
                              fill="red"
                              className="text-primary group-hover:flex hidden h-4 w-4"
                              onClick={() => {
                                removeAuthor(index);
                              }}
                            />
                          </div>
                        </div>
                      ))}
                    </div>
                  </div>

                  <div className="flex-1">
                    <Label>Mô tả</Label>
                    <Textarea {...register("desc")} />
                  </div>
                </CardContent>
              </Card>
              {/* <Card className="basis-2/5">
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
            </Card> */}
            </div>
          </div>
        </form>

        <AlertDialog open={!creating}>
          <AlertDialogContent>
            <AlertDialogHeader>
              <AlertDialogTitle>Đã thêm thành công</AlertDialogTitle>
            </AlertDialogHeader>
            <AlertDialogFooter>
              <Link href={"/books"}>
                <AlertDialogAction>OK</AlertDialogAction>
              </Link>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialog>
      </div>
    );
};

export default InsertNewBook;
