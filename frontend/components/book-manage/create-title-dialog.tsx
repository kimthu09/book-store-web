"use client";

import AuthorList from "@/components/book-manage/author-list";
import CategoryList from "@/components/book-manage/category-list";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { useToast } from "@/components/ui/use-toast";
import { useState } from "react";
import { SubmitHandler, useFieldArray, useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import Loading from "@/components/loading";
import { required } from "@/constants";
import createBookTitle from "@/lib/book/createBookTitle";
import {
  Dialog,
  DialogContent,
  DialogTitle,
  DialogTrigger,
} from "../ui/dialog";
import { FaPlus } from "react-icons/fa";
import getAllAuthorList from "@/lib/book/getAllAuthorList";
import getAllCategoryList from "@/lib/book/getAllCategoryList";
import { useLoading } from "@/hooks/loading-context";
import DropdownSkeleton from "../skeleton/dropdown-skeleton";
import { useRouter } from "next/navigation";

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

const CreateTitleDialog = ({
  handleTitleAdded,
  children,
}: {
  handleTitleAdded: (titleId: string) => void;
  children: React.ReactNode;
}) => {
  const { toast } = useToast();
  const router = useRouter();
  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
  });
  const {
    register,
    handleSubmit,
    control,
    watch,
    reset,
    formState: { errors },
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

  const [open, setOpen] = useState(false);
  const handleOpen = (value: boolean) => {
    setOpen(value);
    if (value) {
      reset({
        idBook: "",
        name: "",
        desc: "",
        authorIds: [],
        categoryIds: [],
      });
    }
  };
  const { showLoading, hideLoading } = useLoading();
  const onSubmit: SubmitHandler<z.infer<typeof FormSchema>> = async (data) => {
    const response: Promise<any> = createBookTitle({
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
        description: "Thêm đầu sách mới thành công",
      });
      handleTitleAdded(responseData.data);
      router.refresh();

      setOpen(false);
    }
  };
  const { categories, isLoading, isError } = getAllCategoryList();
  const {
    authors,
    isLoading: isAuthorLoading,
    isError: isAuthorError,
  } = getAllAuthorList();

  if (isError || isAuthorError) return <div>Failed to load</div>;
  else
    return (
      <Dialog open={open} onOpenChange={handleOpen}>
        <DialogTrigger asChild>{children}</DialogTrigger>
        {!categories || !authors ? (
          <DialogContent>
            <DropdownSkeleton />
          </DialogContent>
        ) : (
          <DialogContent className="p-0 bg-white">
            <DialogTitle className="p-6 py-4 border-b">
              Thêm đầu sách
            </DialogTitle>
            <div className="col items-center px-6">
              <form
                // onSubmit={handleSubmit(onSubmit)}
                className="w-full py-6 pt-0"
              >
                <div className="flex flex-col gap-4">
                  <div className="flex-col flex gap-5">
                    <div className="basis-1/3">
                      <Label htmlFor="masach">Mã sách</Label>
                      <Input
                        {...register("idBook")}
                        id="masach"
                        placeholder="Hệ thống sẽ tự sinh mã nếu để trống"
                      />
                      {errors.idBook && (
                        <span className="error___message">
                          {errors.idBook.message}
                        </span>
                      )}
                    </div>
                    <div className="flex-1">
                      <Label>Tên sách</Label>
                      <Input {...register("name")}></Input>
                      {errors.name && (
                        <span className="error___message">
                          {errors.name.message}
                        </span>
                      )}
                    </div>

                    <div>
                      {/* Category select */}
                      <Label>Thể loại</Label>
                      <CategoryList
                        isEdit
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
                        onRemove={(index) => {
                          removeCate(index);
                        }}
                      />
                      {errors.categoryIds && (
                        <span className="error___message">
                          {errors.categoryIds.message}
                        </span>
                      )}
                    </div>
                    <div>
                      {/* Category select */}
                      <Label>Tác giả</Label>
                      <AuthorList
                        isEdit
                        canAdd
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
                        onRemove={(index) => {
                          removeAuthor(index);
                        }}
                      />
                      {errors.authorIds && (
                        <span className="error___message">
                          {errors.authorIds.message}
                        </span>
                      )}
                    </div>

                    <div className="flex-1">
                      <Label>Mô tả</Label>
                      <Textarea {...register("desc")} />
                    </div>
                  </div>
                  <div className="flex gap-4 py-4  justify-end">
                    <Button
                      type="button"
                      onClick={() => handleOpen(false)}
                      variant={"outline"}
                    >
                      Huỷ
                    </Button>
                    <Button
                      type="button"
                      onClick={handleSubmit(onSubmit)}
                      className="self-end"
                    >
                      Thêm
                    </Button>
                  </div>
                </div>
              </form>
            </div>
          </DialogContent>
        )}
      </Dialog>
    );
};

export default CreateTitleDialog;
