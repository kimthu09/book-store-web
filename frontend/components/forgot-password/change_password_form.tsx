"use client";
import { Input } from "../ui/input";
import { Button } from "../ui/button";
import { useState } from "react";
import { toast } from "../ui/use-toast";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { Label } from "@radix-ui/react-label";
import LoadingSpinner from "../ui/loading-spinner";
import { required } from "@/constants";
import verifyForgotPassword from "@/lib/forgot-password/verifyForgotPassword";
import Link from "next/link";

const ChangePasswordFormSchema = z
  .object({
    newPassword: z.string().min(6, "Ít nhất 6 ký tự"),
    confirmNewPass: required,
  })
  .refine((data) => data.newPassword === data.confirmNewPass, {
    message: "Mật khẩu mới không khớp",
    path: ["confirmNewPass"],
  });

const ChangePasswordForm = (props: any) => {
  const { token } = props;
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const {
    register,
    reset,
    handleSubmit,
    formState: { errors },
  } = useForm<z.infer<typeof ChangePasswordFormSchema>>({
    resolver: zodResolver(ChangePasswordFormSchema),
  });

  const onSubmit = async ({ newPassword }: { newPassword: string }) => {
    setIsLoading(true);
    const responseData = await verifyForgotPassword({
      forgetPasswordToken: token,
      newPassword: newPassword,
    });
    if (responseData.hasOwnProperty("data")) {
      if (responseData.data === true) {
        toast({
          variant: "success",
          title: "Thành công",
          description: "Đã thay đổi mật khẩu thành công.",
        });
        setIsLoading(false);
        reset({
          newPassword: "",
          confirmNewPass:"",
        });
      }
    } else if (responseData.hasOwnProperty("errorKey")) {
      toast({
        variant: "destructive",
        title: "Có lỗi",
        description: responseData.message,
      });
    } else {
      toast({
        variant: "destructive",
        title: "Có lỗi",
        description: "Vui lòng thử lại sau",
      });
    }
    setIsLoading(false);
  };

  return (
    <div className="flex flex-col gap-8 flex-1 w-[500px]">
      <div className="flex flex-col gap-2">
        <h1>Đổi mật khẩu</h1>
        <h5 className="text-grey">Hãy nhập mật khẩu bạn muốn đổi.</h5>
      </div>
      <div className="flex flex-col gap-4 p-6 bg-white border">
        <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-4">
          <div>
            <Label htmlFor="new">Mật khẩu mới</Label>
            <Input
              type="password"
              id="new"
              {...register("newPassword")}
            ></Input>
            {errors.newPassword && (
              <span className="error___message">
                {errors.newPassword.message}
              </span>
            )}
          </div>
          <div>
            <Label htmlFor="new">Xác nhận mật khẩu mới</Label>
            <Input
              type="password"
              id="confirm"
              {...register("confirmNewPass")}
            ></Input>
            {errors.confirmNewPass && (
              <span className="error___message">
                {errors.confirmNewPass.message}
              </span>
            )}
          </div>
          <Button type="submit" disabled={isLoading}>
            {isLoading ? (
              <LoadingSpinner className={"h-4 w-4 text-white"} />
            ) : (
              "Đổi mật khẩu"
            )}
          </Button>
        </form>
        <div className="flex flex-row self-center gap-1">
          <p>Quay lại</p>
          <Link href={"/login"}>
            <p className="text-primary">đăng nhập</p>
          </Link>
        </div>
      </div>
    </div>
  );
};

export default ChangePasswordForm;
