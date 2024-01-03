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
import { useState } from "react";
import { SubmitHandler, useFieldArray, useForm } from "react-hook-form";
import { LuCheck } from "react-icons/lu";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import Loading from "@/components/loading";
import { endPoint, required } from "@/constants";
import BookTitleSelect from "@/components/book-manage/book-title-select";
import PublisherList from "@/components/book-manage/publisher-list";
import createBook from "@/lib/book/createBook";
import { useSWRConfig } from "swr";
import { useRouter } from "next/navigation";
import { imageUpload } from "@/lib/staff/uploadImage";
import { useCurrentUser } from "@/hooks/use-user";
import { includesRoles } from "@/lib/utils";
import NoRole from "@/components/no-role";

const FormSchema = z.object({
  bookTitleId: z.string().min(1, "Vui lòng chọn một đầu sách"),
  idBook: z.string().max(12, "Tối đa 12 ký tự"),
  name: required,
  edition: z.coerce.number().gte(1, "Lần tái bản phải lớn hơn 0"),
  publisherId: z.string().min(1, "Vui lòng chọn một nhà xuất bản"),
  listedPrice: z.coerce.number().gte(1, "Giá niêm yết phải lớn hơn 0"),
  sellPrice: z.coerce.number().gte(1, "Giá bán phải lớn hơn 0"),
  image: z.string(),
});

const InsertNewBook = () => {
  const { toast } = useToast();
  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      bookTitleId: "",
      idBook: "",
      name: "",
      edition: 1,
      publisherId: "",
      listedPrice: 0,
      sellPrice: 0,
      image: "/no-image.jpg",
    },
  });
  const [creating, setCreating] = useState(true);
  const {
    register,
    handleSubmit,
    control,
    setValue,
    trigger,
    formState: { errors },
  } = form;

  const [publisherId, setPublisherId] = useState("");
  const router = useRouter();
  const handleTitleSet = (idTitle: string) => {
    setValue("bookTitleId", idTitle);
    trigger("bookTitleId");
  };
  const handlePublisherIdSet = (idPublisher: string) => {
    setPublisherId(idPublisher);
    setValue("publisherId", idPublisher);
    trigger("publisherId");
  };
  const { mutate } = useSWRConfig();
  const onSubmit: SubmitHandler<z.infer<typeof FormSchema>> = async (data) => {
    console.log(data);
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
      console.log(imgRes.data);

      data.image = imgRes.data;
    }

    const response: Promise<any> = createBook({
      id: data.idBook,
      name: data.name,
      bookTitleId: data.bookTitleId,
      sellPrice: data.sellPrice,
      listedPrice: data.listedPrice,
      edition: data.edition,
      publisherId: data.publisherId,
      image: data.image,
    });
    const responseData = await response;
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
        description: "Thêm mới sách thành công",
      });
      mutate(`${endPoint}/v1/books/all`);
      router.refresh();
    }
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
  const { currentUser } = useCurrentUser();
  if (!currentUser) {
    return <Loading />;
  } else if (
    currentUser &&
    !includesRoles({
      currentUser: currentUser,
      allowedFeatures: ["BOOK_CREATE"],
    })
  ) {
    return <NoRole></NoRole>;
  } else
    return (
      <div className="col items-center">
        <form
          onSubmit={handleSubmit(onSubmit)}
          className="col 2xl:w-4/5 w-full 2xl:px-0 md:px-6 px-0"
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
            <div className="flex flex-col flex-1 gap-4 xl:flex-row">
              <Card className="flex-1">
                <CardContent className="flex-col flex gap-1 mt-5">
                  <BookTitleSelect handleTitleSet={handleTitleSet} />
                  {errors.bookTitleId && (
                    <span className="error___message">
                      {errors.bookTitleId.message}
                    </span>
                  )}
                </CardContent>
              </Card>
              <Card className="flex-1">
                <CardContent className="flex-col flex gap-5 mt-5">
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
                  <div className="flex lg:gap-4 gap-3 xl:flex-col sm:flex-row flex-col">
                    <div className="flex-1">
                      <Label>Nhà xuất bản</Label>
                      <PublisherList
                        canAdd
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
                      <Input type="number" {...register("edition")} />
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
                      <Input type="number" {...register("listedPrice")} />
                      {errors.listedPrice && (
                        <span className="error___message">
                          {errors.listedPrice.message}
                        </span>
                      )}
                    </div>
                    <div className="flex-1">
                      <Label>Đơn giá (VNĐ)</Label>
                      <Input type="number" {...register("sellPrice")} />
                      {errors.sellPrice && (
                        <span className="error___message">
                          {errors.sellPrice.message}
                        </span>
                      )}
                    </div>
                  </div>
                  <div>
                    <Label htmlFor="img">Hình ảnh</Label>
                    <div className="flex items-center gap-4">
                      <Input
                        className="basis-1/2"
                        id="img"
                        type="file"
                        onChange={handleMultipleImage}
                      ></Input>
                      <div>
                        {image && (
                          <img
                            src={imagePreviews}
                            alt={`Preview`}
                            className="h-24 w-auto"
                          />
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

export default InsertNewBook;
