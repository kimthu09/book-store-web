"use client";
import { SubmitHandler, useForm } from "react-hook-form";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { useState } from "react";
import { toast } from "@/components/ui/use-toast";
import { useRouter } from "next/navigation";
import { phoneRegex, required } from "@/constants";
import { useCurrentUser } from "@/hooks/use-user";
import { includesRoles } from "@/lib/utils";
import createCustomer from "@/lib/customer/createCustomer";
import { useLoading } from "@/hooks/loading-context";

const SupplierSchema = z.object({
  id: z.string().max(12, "Tối đa 12 ký tự"),
  name: required,
  email: z
    .string()
    .refine(
      (value) =>
        value === "" || /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i.test(value),
      {
        message: "Email không hợp lệ",
      }
    ),
  phone: z.string().regex(phoneRegex, "Số điện thoại không hợp lệ"),
});

const CreateDialog = ({
  children,
  handleCustomerAdded,
}: {
  children: React.ReactNode;
  handleCustomerAdded?: (customerId: string) => void;
}) => {
  const { currentUser } = useCurrentUser();
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm<z.infer<typeof SupplierSchema>>({
    resolver: zodResolver(SupplierSchema),
  });
  const router = useRouter();
  const { showLoading, hideLoading } = useLoading();
  const onSubmit: SubmitHandler<z.infer<typeof SupplierSchema>> = async (
    data
  ) => {
    const response: Promise<any> = createCustomer(data);
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
        description: "Thêm khách hàng thành công",
      });
      setOpen(false);

      if (handleCustomerAdded) {
        handleCustomerAdded(responseData.data);
      }
      router.refresh();
    }
  };

  const [open, setOpen] = useState(false);
  if (
    !currentUser ||
    (currentUser &&
      !includesRoles({
        currentUser: currentUser,
        allowedFeatures: ["CUSTOMER_CREATE"],
      }))
  ) {
    return null;
  } else
    return (
      <Dialog
        open={open}
        onOpenChange={(open) => {
          if (open) {
            reset();
          }
          setOpen(open);
        }}
      >
        <DialogTrigger asChild>{children}</DialogTrigger>
        <DialogContent className="xl:max-w-[720px] max-w-[472px] p-0 bg-white">
          <DialogHeader>
            <DialogTitle className="p-6 pb-0">Thêm khách hàng</DialogTitle>
          </DialogHeader>
          <form onSubmit={handleSubmit(onSubmit)}>
            <div className="p-6 flex flex-col gap-4 border-y-[1px]">
              <div>
                <Label htmlFor="idNcc">Mã khách hàng</Label>
                <Input
                  id="idNcc"
                  placeholder="Hệ thống sẽ tự sinh mã nếu để trống"
                  {...register("id")}
                ></Input>
                {errors.id && (
                  <span className="error___message">{errors.id.message}</span>
                )}
              </div>
              <div>
                <Label htmlFor="nameNcc">Tên khách hàng</Label>
                <Input id="nameNcc" {...register("name")}></Input>
                {errors.name && (
                  <span className="error___message">{errors.name.message}</span>
                )}
              </div>
              <div>
                <Label htmlFor="email">Email</Label>
                <Input id="email" {...register("email")}></Input>
                {errors.email && (
                  <span className="error___message">
                    {errors.email.message}
                  </span>
                )}
              </div>
              <div className="flex gap-4">
                <div className="flex-1">
                  <Label htmlFor="phone">Số điện thoại</Label>
                  <Input id="phone" {...register("phone")}></Input>
                  {errors.phone && (
                    <span className="error___message">
                      {errors.phone.message}
                    </span>
                  )}
                </div>
              </div>
            </div>
            <div className="p-4 flex-1 flex justify-end">
              <div className="flex gap-4">
                <Button
                  type="reset"
                  variant={"outline"}
                  onClick={() => setOpen(false)}
                >
                  Huỷ
                </Button>

                <Button type="submit">Thêm</Button>
              </div>
            </div>
          </form>
        </DialogContent>
      </Dialog>
    );
};

export default CreateDialog;
