"use client";
import { Input } from "../ui/input";
import { login } from "@/lib/auth/action";
import { Button } from "../ui/button";
import { toast } from "../ui/use-toast";
import { z } from "zod";
import { Label } from "@radix-ui/react-label";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useState } from "react";
import LoadingSpinner from "../ui/loading-spinner";
import Link from "next/link";

const LoginScheme = z.object({
  email: z.string().email("Email không hợp lệ"),
  password: z.string().min(6, "Ít nhất 6 ký tự"),
});

const LoginForm = () => {
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<z.infer<typeof LoginScheme>>({
    resolver: zodResolver(LoginScheme),
  });

  const onSubmit = async ({
    email,
    password,
  }: {
    email: string;
    password: string;
  }) => {
    setIsLoading(true);
    const responseData = await login({
      email: email,
      password: password,
    });
    if (responseData?.error) {
      toast({
        variant: "destructive",
        title: "Đăng nhập thất bại",
        description: "Vui lòng kiểm tra lại email và mật khẩu của bạn",
      });
    }
    setIsLoading(false);
  };

  return (
    <div className="flex flex-col gap-8 flex-1 w-[500px]">
      <div className="flex flex-col gap-2">
        <h1>Đăng nhập vào tài khoản</h1>
        <h5 className="text-grey">
          Xem những gì đang xảy ra với doanh nghiệp của bạn.
        </h5>
      </div>
      <div className="p-6 bg-white border">
        <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-4">
          <div>
            <Label htmlFor="email">Email</Label>
            <Input id="email" {...register("email")}></Input>
            {errors.email && (
              <span className="error___message">{errors.email.message}</span>
            )}
          </div>
          <div>
            <Label htmlFor="pass">Mật khẩu</Label>
            <Input type="password" id="pass" {...register("password")}></Input>
            {errors.password && (
              <span className="error___message">{errors.password.message}</span>
            )}
          </div>
          <Link className="self-end" href={"/forgot-password"}>
            <p className="text-primary">Quên mật khẩu?</p>
          </Link>
          <Button type="submit" disabled={isLoading}>
            {isLoading ? (
              <LoadingSpinner className={"h-4 w-4 text-white"} />
            ) : (
              "Đăng nhập"
            )}
          </Button>
        </form>
      </div>
    </div>
  );
};

export default LoginForm;
function setIsLoading(arg0: boolean) {
  throw new Error("Function not implemented.");
}
