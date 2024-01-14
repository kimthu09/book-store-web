"use client";
import { Input } from "../ui/input";
import { Button } from "../ui/button";
import { useState } from "react";
import { toast } from "../ui/use-toast";
import sendEmailForgotPassword from "@/lib/forgot-password/sendEmailForgotPassword";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { Label } from "@radix-ui/react-label";
import LoadingSpinner from "../ui/loading-spinner";
import Link from "next/link";

const SendEmailForgotPasswordSchema = z.object({
  email: z.string().email("Email không hợp lệ"),
});

const SendEmailForgotPasswordForm = () => {

  const [isLoading, setIsLoading] = useState<boolean>(false);
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm<z.infer<typeof SendEmailForgotPasswordSchema>>({
    resolver: zodResolver(SendEmailForgotPasswordSchema),
  });

  const onSubmit = async ({ email }: { email: string }) => {
    setIsLoading(true);
    const responseData = await sendEmailForgotPassword({
      email: email,
    });
    if (responseData.hasOwnProperty("data")) {
      if (responseData.data === true)
      {
        toast({
          variant: "success",
          title: "Thành công",
          description: "Đã gửi hướng dẫn đổi mật khẩu đến email bạn vừa nhập.",
        });
        reset({
          email: "",
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
        <h1>Quên mật khẩu</h1>
        <h5 className="text-grey">
          Chúng tôi sẽ gửi email đến cho bạn để đổi mật khẩu.
        </h5>
      </div>
      <div className="flex flex-col gap-4 p-6 bg-white border">
        <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-4">
          <div>
            <Label htmlFor="email">Email</Label>
            <Input id="email" {...register("email")}></Input>
            {errors.email && (
              <span className="error___message">{errors.email.message}</span>
            )}
          </div>
          <Button type="submit" disabled={isLoading}>
            {isLoading ? (
              <LoadingSpinner className={"h-4 w-4 text-white"} />
            ) : (
              "Xác nhận"
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

export default SendEmailForgotPasswordForm;
