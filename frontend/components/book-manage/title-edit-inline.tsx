import { BookTitle } from "@/types";
import React, { useEffect, useState } from "react";
import { Label } from "../ui/label";
import { Input } from "../ui/input";
import { FaPen } from "react-icons/fa";
import { Button } from "../ui/button";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { required } from "@/constants";
import { useToast } from "../ui/use-toast";
import {
  SubmitErrorHandler,
  SubmitHandler,
  useFieldArray,
  useForm,
} from "react-hook-form";
import CategoryList from "./category-list";
import { AiOutlineCheck, AiOutlineClose } from "react-icons/ai";
import { Textarea } from "../ui/textarea";
import AuthorList from "./author-list";
import updateBookTitle from "@/lib/book/updateBookTitle";
import ConfirmDialog from "../confirm-dialog";
import { useCurrentUser } from "@/hooks/use-user";
import { includesRoles } from "@/lib/utils";
import { useLoading } from "@/hooks/loading-context";
const FormSchema = z.object({
  idBook: z.string().max(12, "Tối đa 12 ký tự"),
  name: required,
  desc: z.string(),
  authorIds: z
    .array(z.object({ idAuthor: z.string() }))
    .nonempty("Vui lòng chọn ít nhất một tác giả"),
  categoryIds: z
    .array(z.object({ idCate: z.string() }))
    .nonempty("Vui lòng chọn ít nhất một thể loại"),
});
const TitleEditInline = ({
  book,
  handleTitleEdited,
}: {
  book: BookTitle;
  handleTitleEdited: () => void;
}) => {
  const { toast } = useToast();
  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
  });
  const {
    register,
    handleSubmit,
    control,
    watch,
    reset,
    formState: { errors, isDirty },
  } = form;

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
  const onError: SubmitErrorHandler<z.infer<typeof FormSchema>> = async (
    data
  ) => {};
  const { showLoading, hideLoading } = useLoading();
  const onSubmit: SubmitHandler<z.infer<typeof FormSchema>> = async (data) => {
    setIsEdit(false);
    const response: Promise<any> = updateBookTitle({
      id: data.idBook,
      name: data.name,
      desc: data.desc,
      categoryIds: data.categoryIds.map((item) => item.idCate),
      authorIds: data.authorIds.map((item) => item.idAuthor),
    });
    showLoading();
    const responseData = await response;
    hideLoading();
    if (responseData.hasOwnProperty("errorKey")) {
      toast({
        variant: "destructive",
        title: "Có lỗi",
        description: responseData.message,
      });
    } else {
      toast({
        variant: "success",
        title: "Thành công",
        description: "Chỉnh sửa đầu sách mới thành công",
      });
      handleTitleEdited();
    }
  };
  const [isEdit, setIsEdit] = useState(false);
  useEffect(() => {
    // Optionally log the error to an error reporting service
    reset({
      idBook: book.id,
      name: book.name,
      desc: book.desc,
      authorIds: book.authors.map((item) => {
        return { idAuthor: item.id };
      }),
      categoryIds: book.categories.map((item) => {
        return { idCate: item.id };
      }),
    });
  }, [book]);

  const { currentUser } = useCurrentUser();
  return (
    <div className="flex bg-background lg:flex-row flex-col p-4 px-6 gap-6">
      <div className="flex basis-1/2 flex-col gap-4 items-start ">
        <div className="flex lg:flex-row flex-col items-start w-full gap-2">
          <span className="font-medium min-w-[5rem]">Đầu sách: </span>
          <div className="w-full">
            <Input
              className="bg-white w-full"
              readOnly={!isEdit}
              {...register("name")}
            />
            {errors.name && (
              <span className="error___message">{errors.name.message}</span>
            )}
          </div>
        </div>
        <div className="flex lg:flex-row flex-col items-start w-full gap-2">
          <span className="font-medium min-w-[5rem]">Mô tả: </span>
          <Textarea
            className="bg-white flex-1 w-full h-24"
            {...register("desc")}
            readOnly={!isEdit}
          />
        </div>
      </div>
      <div className="flex lg:flex-row flex-col gap-4 flex-1">
        <div className="flex-1">
          <div className="flex flex-col gap-4">
            <div>
              <CategoryList
                isEdit={isEdit}
                canAdd
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
                onRemove={removeCate}
              />
              {errors.categoryIds && (
                <span className="error___message">
                  {errors.categoryIds.message}
                </span>
              )}
            </div>
            <div>
              <AuthorList
                isEdit={isEdit}
                canAdd
                checkedAuthor={fieldsAuthor.map((author) => author.idAuthor)}
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
                onRemove={removeAuthor}
              />
              {errors.authorIds && (
                <span className="error___message">
                  {errors.authorIds.message}
                </span>
              )}
            </div>
          </div>
        </div>
        <div className="flex lg:flex-col flex-row lg:justify-start justify-end gap-2">
          {isEdit ? (
            <>
              <Button
                variant={"ghost"}
                size={"icon"}
                className="rounded-full bg-rose-200/60 hover:bg-rose-200/90 text-rose-600 hover:text-rose-600"
                title="Hủy"
                onClick={() => {
                  setIsEdit(false);
                  reset({
                    idBook: book.id,
                    name: book.name,
                    desc: book.desc,
                    authorIds: book.authors.map((item) => {
                      return { idAuthor: item.id };
                    }),
                    categoryIds: book.categories.map((item) => {
                      return { idCate: item.id };
                    }),
                  });
                }}
              >
                <AiOutlineClose className="h-5 w-5" />
              </Button>
              <ConfirmDialog
                title={"Xác nhận"}
                description="Bạn xác nhận chỉnh sửa đầu sách này ?"
                handleYes={() => {
                  handleSubmit(onSubmit, onError)();
                }}
              >
                <Button
                  variant={"ghost"}
                  size={"icon"}
                  disabled={!isDirty}
                  title="Lưu"
                  className="rounded-full bg-green-200/60 hover:bg-green-200/90 text-green-600 hover:text-green-600"
                >
                  <AiOutlineCheck className="h-5 w-5" />
                </Button>
              </ConfirmDialog>
            </>
          ) : currentUser &&
            includesRoles({
              currentUser: currentUser,
              allowedFeatures: ["BOOK_TITLE_UPDATE"],
            }) ? (
            <Button
              variant={"ghost"}
              size={"icon"}
              className="rounded-full bg-blue-200/60 hover:bg-blue-200/90 text-primary hover:text-primary"
              title="Chỉnh sửa"
              onClick={() => setIsEdit(true)}
            >
              <FaPen />
            </Button>
          ) : null}
        </div>
      </div>
    </div>
  );
};

export default TitleEditInline;
