"use client";
import ConfirmDialog from "@/components/confirm-dialog";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Skeleton } from "@/components/ui/skeleton";
import { toast } from "@/components/ui/use-toast";
import { phoneRegex, required } from "@/constants";
import { useLoading } from "@/hooks/loading-context";
import { useShop } from "@/hooks/use-shop";
import { useCurrentUser } from "@/hooks/use-user";
import updateShop from "@/lib/shop-general/updateShop";
import { isAdmin } from "@/lib/utils";
import { zodResolver } from "@hookform/resolvers/zod";

import { useRouter } from "next/navigation";
import React, { useEffect, useState } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { AiOutlineClose } from "react-icons/ai";
import { FaPen } from "react-icons/fa";
import { LuCheck } from "react-icons/lu";
import { z } from "zod";
const FormSchema = z.object({
  name: required,
  phone: z.string().refine((data) => data === "" || phoneRegex.test(data), {
    message: "Số điện thoại không hợp lệ",
  }),
  email: z
    .string()
    .refine(
      (value) =>
        value === "" || /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i.test(value),
      {
        message: "Email không hợp lệ",
      }
    ),
  address: z.string(),
  wifiPass: z.string(),
  accumulatePointPercent: z.coerce
    .number({ invalid_type_error: "Tỉ lệ tích điểm phải là một số" })
    .gt(0, "Tỉ lệ tích điểm phải lớn hơn 0"),
  usePointPercent: z.coerce
    .number({ invalid_type_error: "Tỉ lệ dùng điểm phải là một số" })
    .gt(0, "Tỉ lệ dùng điểm phải lớn hơn 0"),
});
const DetailLayout = () => {
  const { shop } = useShop();
  const [readOnly, setReadOnly] = useState(true);
  const router = useRouter();
  const { currentUser } = useCurrentUser();
  const isAdminRole = currentUser && isAdmin({ currentUser: currentUser });
  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
  });
  const {
    register,
    handleSubmit,
    control,
    reset,
    formState: { errors, isDirty },
  } = form;
  const handleReset = () => {
    if (shop) {
      reset({
        name: shop.name,
        phone: shop.phone,
        address: shop.address,
        email: shop.email,
        wifiPass: shop.wifiPass,
        accumulatePointPercent: shop.accumulatePointPercent,
        usePointPercent: shop.usePointPercent,
      });
    }
  };
  const { showLoading, hideLoading } = useLoading();
  const onSubmit: SubmitHandler<z.infer<typeof FormSchema>> = async (data) => {
    setReadOnly(true);
    if (!data.email) {
      data.email = "";
    }
    const response: Promise<any> = updateShop(data);
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
        description: "Chỉnh thiết lập cửa hàng thành công",
      });
      router.refresh();
    }
  };
  useEffect(() => {
    handleReset();
  }, [shop]);
  if (!shop) {
    return (
      <div className="col items-center">
        <div className="col xl:w-3/5 w-full xl:px-0 md:px-8 px-0">
          <div className="flex justify-between">
            <h1 className="xl:text-3xl text-2xl">Thiết lập cửa hàng</h1>
          </div>
          <Card>
            <CardContent className="p-6 flex flex-col gap-4">
              <div className="flex gap-4  flex-col">
                <div className="flex gap-4 lg:flex-row flex-col">
                  <div className="basis-2/3">
                    <Skeleton className="h-6 w-full" />
                  </div>
                  <div className="basis-1/3">
                    <Skeleton className="h-6 w-full" />
                  </div>
                </div>
                <div className="flex gap-4 lg:flex-row flex-col">
                  <div className="basis-2/3">
                    <Skeleton className="h-6 w-full" />
                  </div>
                  <div className="basis-1/3">
                    <Skeleton className="h-6 w-full" />
                  </div>
                </div>
                <div className="flex gap-4 lg:flex-row flex-col">
                  <div className="basis-1/2">
                    <Skeleton className="h-6 w-full" />
                  </div>
                  <div className="basis-1/2">
                    <Skeleton className="h-6 w-full" />
                  </div>
                </div>
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
    );
  } else
    return (
      <div className="col items-center">
        <div className="col xl:w-3/5 w-full xl:px-0 md:px-8 px-0">
          <div className="flex justify-between">
            <h1 className="xl:text-3xl text-2xl">Thiết lập cửa hàng</h1>
            <div className="flex gap-2 justify-center">
              {!readOnly ? (
                <div className="flex gap-2 sm:flex-initial flex-1">
                  <Button
                    variant={"outline"}
                    className="bg-white border-rose-700 text-rose-700 hover:text-rose-700 hover:bg-rose-50/30 px-2 flex gap-1 flex-nowrap whitespace-nowrap flex-1"
                    onClick={() => {
                      setReadOnly(true);
                      handleReset();
                    }}
                  >
                    <AiOutlineClose className="h-5 w-5" />
                    Hủy
                  </Button>
                  <ConfirmDialog
                    title={"Xác nhận"}
                    description="Bạn xác nhận chỉnh sửa thiết lập cửa hàng ?"
                    handleYes={() => handleSubmit(onSubmit)()}
                  >
                    <Button
                      className="px-2 flex gap-1 flex-nowrap whitespace-nowrap flex-1"
                      disabled={!isDirty}
                    >
                      <LuCheck className="h-5 w-5" />
                      Lưu
                    </Button>
                  </ConfirmDialog>
                </div>
              ) : isAdminRole ? (
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
              ) : null}
            </div>
          </div>

          <Card>
            <CardContent className="p-6 flex flex-col   gap-4">
              <div className="flex gap-4  flex-col">
                <div className="flex gap-4 lg:flex-row flex-col">
                  <div className="basis-2/3">
                    <Label htmlFor="name">Tên cửa hàng</Label>
                    <Input
                      id="name"
                      readOnly={readOnly}
                      defaultValue={shop?.name}
                      {...register("name")}
                    ></Input>
                    {errors.name && (
                      <span className="error___message">
                        {errors.name.message}
                      </span>
                    )}
                  </div>
                  <div className="basis-1/3">
                    <Label htmlFor="phone">Số điện thoại</Label>
                    <Input
                      id="phone"
                      readOnly={readOnly}
                      defaultValue={shop?.phone}
                      {...register("phone")}
                    ></Input>
                    {errors.phone && (
                      <span className="error___message">
                        {errors.phone.message}
                      </span>
                    )}
                  </div>
                </div>
                <div className="flex gap-4 lg:flex-row flex-col">
                  <div className="basis-2/3">
                    <Label htmlFor="email">Email</Label>
                    <Input
                      id="email"
                      readOnly={readOnly}
                      defaultValue={shop?.email ?? ""}
                      {...register("email")}
                    ></Input>
                    {errors.email && (
                      <span className="error___message">
                        {errors.email.message}
                      </span>
                    )}
                  </div>
                  <div className="basis-1/3">
                    <Label htmlFor="diem">Mật khẩu wifi</Label>
                    <Input
                      id="diem"
                      readOnly={readOnly}
                      defaultValue={shop?.wifiPass}
                      {...register("wifiPass")}
                    ></Input>
                    {errors.wifiPass && (
                      <span className="error___message">
                        {errors.wifiPass.message}
                      </span>
                    )}
                  </div>
                </div>
                <div className="basis-2/3">
                  <Label htmlFor="add">Địa chỉ</Label>
                  <Input
                    id="add"
                    readOnly={readOnly}
                    defaultValue={shop?.address ?? ""}
                    {...register("address")}
                  ></Input>
                  {errors.address && (
                    <span className="error___message">
                      {errors.address.message}
                    </span>
                  )}
                </div>
                <div className="flex gap-4 lg:flex-row flex-col">
                  <div className="basis-1/2">
                    <Label htmlFor="accumulatePointPercent">
                      Tỉ lệ tích điểm
                    </Label>
                    <Input
                      id="accumulatePointPercent"
                      readOnly={readOnly}
                      defaultValue={shop?.accumulatePointPercent}
                      {...register("accumulatePointPercent")}
                    ></Input>
                    {errors.accumulatePointPercent && (
                      <span className="error___message">
                        {errors.accumulatePointPercent.message}
                      </span>
                    )}
                  </div>
                  <div className="basis-1/2">
                    <Label htmlFor="usePointPercent">Tỉ lệ dùng điểm</Label>
                    <Input
                      id="usePointPercent"
                      readOnly={readOnly}
                      defaultValue={shop?.usePointPercent}
                      {...register("usePointPercent")}
                    ></Input>
                    {errors.usePointPercent && (
                      <span className="error___message">
                        {errors.usePointPercent.message}
                      </span>
                    )}
                  </div>
                </div>
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
    );
};

export default DetailLayout;
