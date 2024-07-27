"use client";

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
import { useToast } from "@/components/ui/use-toast";
import Link from "next/link";
import { useEffect, useState } from "react";
import {
  Controller,
  SubmitErrorHandler,
  SubmitHandler,
  useForm,
} from "react-hook-form";
import { LuCheck } from "react-icons/lu";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import Loading from "@/components/loading";
import { required } from "@/constants";
import PublisherList from "@/components/book-manage/publisher-list";
import { useRouter } from "next/navigation";
import { imageUpload } from "@/lib/staff/uploadImage";
import { useCurrentUser } from "@/hooks/use-user";
import { includesRoles } from "@/lib/utils";
import NoRole from "@/components/no-role";
import getBook from "@/lib/book/getBook";
import ConfirmDialog from "@/components/confirm-dialog";
import { AiOutlineClose } from "react-icons/ai";
import { FaPen } from "react-icons/fa";
import Image from "next/image";
import ChangeImage from "@/components/staff/change-image";
import updateBook from "@/lib/book/updateBook";
import { BookTitle } from "@/types";
import BookTitleSelectEdit from "@/components/book-manage/book-title-select-edit";
import { useLoading } from "@/hooks/loading-context";
import BookDetailSkeleton from "@/components/skeleton/book-detail-skeleton";
import { NumericFormat } from "react-number-format";

const FormSchema = z.object({
  bookTitleId: z.string().min(1, "Vui lòng chọn một đầu sách"),
  edition: z.coerce
    .number({ invalid_type_error: "Lần tái bản phải là một số" })
    .gte(1, "Lần tái bản phải lớn hơn hoặc bằng 1")
    .refine((value) => Number.isInteger(value), {
      message: "Lần tái bản phải là số nguyên",
    }),
  publisherId: z.string().min(1, "Vui lòng chọn một nhà xuất bản"),
  listedPrice: z.coerce
    .number({ invalid_type_error: "Giá niêm yết phải là một số" })
    .gt(0, "Giá niêm yết phải lớn hơn 0")
    .refine((value) => Number.isInteger(value), {
      message: "Giá niêm yết phải là số nguyên",
    }),
  sellPrice: z.coerce
    .number({ invalid_type_error: "Giá bán phải là một số" })
    .gt(0, "Giá bán phải lớn hơn 0")
    .refine((value) => Number.isInteger(value), {
      message: "Giá bán phải là số nguyên",
    }),
  image: z.string(),
});

const EditBook = ({ params }: { params: { bookId: string } }) => {
  const {
    data,
    isLoading,
    isError,
    mutate: mutateBook,
  } = getBook(params.bookId);
  const [readOnly, setReadOnly] = useState(true);
  const [title, setTitle] = useState("");
  const { toast } = useToast();
  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
  });
  const [creating, setCreating] = useState(true);
  const {
    register,
    handleSubmit,
    control,
    setValue,
    trigger,
    reset,
    formState: { errors, isDirty },
  } = form;

  const [publisherId, setPublisherId] = useState("");
  const router = useRouter();
  const handleTitleSet = (idTitle: string) => {
    setTitle(idTitle);
    setValue("bookTitleId", idTitle, { shouldDirty: true });
    trigger("bookTitleId");
  };
  const handlePublisherIdSet = (idPublisher: string) => {
    setPublisherId(idPublisher);
    setValue("publisherId", idPublisher, { shouldDirty: true });
    trigger("publisherId");
  };
  const { showLoading, hideLoading } = useLoading();
  const onSubmit: SubmitHandler<z.infer<typeof FormSchema>> = async (data) => {
    showLoading();
    if (image) {
      let formData = new FormData();

      formData.append("file", image);
      formData.append("folderName", "images");

      const imgRes = await imageUpload(formData);
      if (imgRes.hasOwnProperty("errorKey")) {
        toast({
          variant: "destructive",
          title: "Có lỗi",
          description: imgRes.message,
        });
        return;
      }

      data.image = imgRes.data;
    }
    const response: Promise<any> = updateBook(data, params.bookId);
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
        description: "Chỉnh sửa sách thành công",
      });
      mutateBook();
      setReadOnly(true);
      router.refresh();
    }
  };
  const onErrors: SubmitErrorHandler<z.infer<typeof FormSchema>> = (data) => {
    toast({
      variant: "destructive",
      title: "Có lỗi",
      description: "Vui lòng thử lại sau",
    });
  };
  const [image, setImage] = useState<any>();
  const [imagePreviews, setImagePreviews] = useState<any>();
  const handleMultipleImage = (event: any) => {
    const file = event.target.files[0];
    if (file) {
      if (file && file.type.includes("image")) {
        setImage(file);
        console.log(file.type);
        const reader = new FileReader();
        reader.onload = () => {
          setImagePreviews(reader.result);
        };
        reader.readAsDataURL(file);
      } else {
        setImage(null);
        toast({
          variant: "destructive",
          title: "Có lỗi",
          description: "File không hợp lệ",
        });
        console.log("file không hợp lệ");
      }
    } else {
      setImage(null);
    }
  };
  useEffect(() => {
    if (data) {
      resetForm();
    }
  }, [data]);
  const resetForm = () => {
    reset({
      bookTitleId: data.bookTitle.id,
      edition: data.edition,
      publisherId: data.publisher.id,
      listedPrice: data.listedPrice,
      sellPrice: data.sellPrice,
      image: data.image,
    });
    handlePublisherIdSet(data.publisher.id);
    setTitle(data.bookTitle.id);
    setImage(null);
  };

  const { currentUser } = useCurrentUser();
  if (!currentUser || isLoading) {
    return <BookDetailSkeleton />;
  } else if (
    currentUser &&
    !includesRoles({
      currentUser: currentUser,
      allowedFeatures: ["BOOK_UPDATE"],
    })
  ) {
    return <NoRole></NoRole>;
  } else
    return (
      <div className="col items-center">
        <form
          onSubmit={handleSubmit(onSubmit, onErrors)}
          className="col 2xl:w-4/5 w-full 2xl:px-0 px-0"
        >
          <div className="flex flex-row justify-between">
            <h1 className="font-medium text-2xl self-start">
              Sách: {params.bookId}
            </h1>
            <div className="flex gap-2 justify-end">
              {!readOnly ? (
                <div className="flex gap-2 sm:flex-initial flex-1">
                  <Button
                    variant={"outline"}
                    className="bg-white border-rose-700 text-rose-700 hover:text-rose-700 hover:bg-rose-50/30 px-2 flex gap-1 flex-nowrap whitespace-nowrap flex-1"
                    onClick={() => {
                      setReadOnly(true);
                      resetForm();
                    }}
                  >
                    <AiOutlineClose className="h-5 w-5" />
                    Hủy
                  </Button>
                  <ConfirmDialog
                    title={"Xác nhận"}
                    description="Bạn xác nhận chỉnh sửa sách này ?"
                    handleYes={() => handleSubmit(onSubmit, onErrors)()}
                  >
                    <Button
                      className="px-2 flex gap-1 flex-nowrap whitespace-nowrap flex-1"
                      disabled={!isDirty && !image}
                    >
                      <LuCheck className="h-5 w-5" />
                      Lưu
                    </Button>
                  </ConfirmDialog>
                </div>
              ) : (
                <Button
                  title="Chỉnh sửa"
                  className="px-2 sm:flex-initial flex-1 flex gap-1 flex-nowrap whitespace-nowrap"
                  type="button"
                  onClick={() => {
                    setReadOnly(false);
                  }}
                >
                  <FaPen />
                  Chỉnh sửa
                </Button>
              )}
            </div>
          </div>
          <div>
            <div className="flex flex-col flex-1 gap-4 xl:flex-row">
              <Card className="flex-1">
                <CardContent className="flex-col flex gap-1 mt-5">
                  <BookTitleSelectEdit
                    handleTitleSet={handleTitleSet}
                    titleId={title}
                    readOnly={readOnly}
                  />
                  {errors.bookTitleId && (
                    <span className="error___message">
                      {errors.bookTitleId.message}
                    </span>
                  )}
                </CardContent>
              </Card>
              <Card className="flex-1">
                <CardContent className=" flex gap-5 mt-5">
                  <div className="flex flex-col items-center gap-4">
                    <div className="rounded-sm border overflow-clip w-fit">
                      {image && imagePreviews ? (
                        <img
                          src={imagePreviews}
                          alt={`Preview`}
                          className="h-[120px] w-auto object-cover"
                        />
                      ) : (
                        <Image
                          src={data.image ?? "/no-image.jpg"}
                          alt="ảnh"
                          className="h-[120px] w-auto  object-cover"
                          height={120}
                          width={120}
                        ></Image>
                      )}
                    </div>
                    <Input
                      disabled={readOnly}
                      className="w-32"
                      id="img"
                      type="file"
                      onChange={handleMultipleImage}
                    ></Input>
                  </div>

                  <div className="flex-col flex-1 flex gap-5">
                    <div className="flex lg:gap-4 gap-3 xl:flex-col sm:flex-row flex-col">
                      <div className="flex-1">
                        <Label>Nhà xuất bản</Label>
                        <PublisherList
                          readOnly={readOnly}
                          canAdd={!readOnly}
                          publisherId={publisherId}
                          setPublisherId={handlePublisherIdSet}
                        />
                        {errors.publisherId && (
                          <span className="error___message">
                            {errors.publisherId.message}
                          </span>
                        )}
                      </div>
                      <div className="flex-1">
                        <Label>Lần tái bản</Label>
                        <Input
                          readOnly={readOnly}
                          type="number"
                          {...register("edition")}
                        />
                        {errors.edition && (
                          <span className="error___message">
                            {errors.edition.message}
                          </span>
                        )}
                      </div>
                    </div>
                    <div className="flex lg:gap-4 gap-3 xl:flex-col sm:flex-row flex-col">
                      <div className="flex-1">
                        <Label>Giá niêm yết (VNĐ)</Label>
                        <Controller
                          name="listedPrice"
                          control={control}
                          render={({ field }) => (
                            <NumericFormat
                              value={field.value}
                              onValueChange={(values) => {
                                const numericValue = parseFloat(
                                  values.value.replace(/,/g, "")
                                );

                                field.onChange(numericValue);
                              }}
                              readOnly={readOnly}
                              thousandSeparator="."
                              decimalSeparator=","
                              valueIsNumericString
                              customInput={Input}
                            />
                          )}
                        />
                        {errors.listedPrice && (
                          <span className="error___message">
                            {errors.listedPrice.message}
                          </span>
                        )}
                      </div>
                      <div className="flex-1">
                        <Label>Đơn giá (VNĐ)</Label>
                        <Controller
                          name="sellPrice"
                          control={control}
                          render={({ field }) => (
                            <NumericFormat
                              value={field.value}
                              onValueChange={(values) => {
                                const numericValue = parseFloat(
                                  values.value.replace(/,/g, "")
                                );
                                field.onChange(numericValue);
                              }}
                              readOnly={readOnly}
                              thousandSeparator="."
                              decimalSeparator=","
                              valueIsNumericString
                              customInput={Input}
                            />
                          )}
                        />
                        {errors.sellPrice && (
                          <span className="error___message">
                            {errors.sellPrice.message}
                          </span>
                        )}
                      </div>
                    </div>
                  </div>
                </CardContent>
              </Card>
            </div>
          </div>
        </form>

        <AlertDialog open={!creating}>
          <AlertDialogContent>
            <AlertDialogHeader>
              <AlertDialogTitle>Đã thêm thành công</AlertDialogTitle>
            </AlertDialogHeader>
            <AlertDialogFooter>
              <Link href={"/product/books"}>
                <AlertDialogAction>OK</AlertDialogAction>
              </Link>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialog>
      </div>
    );
};

export default EditBook;
