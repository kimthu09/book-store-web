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
import { Supplier } from "@/types";
import updateSupplier from "@/lib/supplier/updateSupplier";
const phoneRegex = new RegExp(/(0[3|5|7|8|9])+([0-9]{8})\b/g);
const required = z.string().min(1, "Không để trống trường này");

const SupplierSchema = z.object({
  name: required,
  email: z.string().email("Email không hợp lệ"),
  phone: z.string().regex(phoneRegex, "Số điện thoại không hợp lệ"),
});

const EditDialog = ({
  supplier,
  refresh,
}: {
  supplier: Supplier;
  refresh: () => void;
}) => {
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm<z.infer<typeof SupplierSchema>>({
    resolver: zodResolver(SupplierSchema),
    defaultValues: {
      name: supplier.name,
      email: supplier.email,
      phone: supplier.phone,
    },
  });
  const router = useRouter();
  const onSubmit: SubmitHandler<z.infer<typeof SupplierSchema>> = async (
    data
  ) => {
    setOpen(false);
    console.log(data);
    const response: Promise<any> = updateSupplier({
      name: data.name,
      phone: data.phone,
      email: data.email,
      idSupplier: supplier.id,
    });
    const responseData = await response;
    console.log(responseData);
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
        description: "Chỉnh sửa thành công",
      });
      router.refresh();
      refresh();
    }
  };

  const [open, setOpen] = useState(false);
  return (
    <Dialog
      open={open}
      onOpenChange={(open) => {
        if (open) {
          reset({
            name: supplier.name,
            email: supplier.email,
            phone: supplier.phone,
          });
        }
        setOpen(open);
      }}
    >
      <DialogTrigger asChild>
        <Button
          className="lg:px-4 px-2 whitespace-nowrap "
          onClick={() => console.log("hi there " + supplier.name)}
        >
          Chỉnh sửa
        </Button>
      </DialogTrigger>
      <DialogContent className="xl:max-w-[720px] max-w-[472px] p-0 bg-white">
        <DialogHeader>
          <DialogTitle className="p-6 pb-0">
            Chỉnh sửa thông tin nhà cung cấp
          </DialogTitle>
        </DialogHeader>
        <form onSubmit={handleSubmit(onSubmit)}>
          <div className="p-6 flex flex-col gap-4 border-y-[1px]">
            <div>
              <Label htmlFor="nameNcc">Tên nhà cung cấp</Label>
              <Input id="nameNcc" {...register("name")}></Input>
              {errors.name && (
                <span className="error___message">{errors.name.message}</span>
              )}
            </div>
            <div>
              <Label htmlFor="email">Email</Label>
              <Input id="email" {...register("email")}></Input>
              {errors.email && (
                <span className="error___message">{errors.email.message}</span>
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
                type="button"
                variant={"outline"}
                onClick={() => {
                  console.log(supplier);
                  reset({
                    name: supplier.name,
                    email: supplier.email,
                    phone: supplier.phone,
                  });
                }}
              >
                Đặt lại
              </Button>

              <Button type="submit">Lưu</Button>
            </div>
          </div>
        </form>
      </DialogContent>
    </Dialog>
  );
};

export default EditDialog;
