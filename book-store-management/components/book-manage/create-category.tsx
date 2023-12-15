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
import createSupplier from "@/lib/supplier/createSupplier";
import { useRouter } from "next/navigation";
const phoneRegex = new RegExp(/(0[3|5|7|8|9])+([0-9]{8})\b/g);
const required = z.string().min(1, "Không để trống trường này");

const SupplierSchema = z.object({
  name: required,
});

const CreateCategory = () => {
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm<z.infer<typeof SupplierSchema>>({
    resolver: zodResolver(SupplierSchema),
  });
  const router = useRouter();
  const onSubmit: SubmitHandler<z.infer<typeof SupplierSchema>> = async (
    data
  ) => {
    //TODO:
    // setOpen(false);
    // console.log(data);
    // const response: Promise<any> = createSupplier(data);
    // const responseData = await response;
    // console.log(responseData);
    // if (responseData.hasOwnProperty("errorKey")) {
    //   toast({
    //     variant: "destructive",
    //     title: "Có lỗi",
    //     description: responseData.message,
    //   });
    // } else {
    //   toast({
    //     variant: "success",
    //     title: "Thành công",
    //     description: "Thêm nhà cung cấp thành công",
    //   });
    //   router.refresh();
    // }
  };

  const [open, setOpen] = useState(false);
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
      <DialogTrigger asChild>
        <Button className="lg:px-4 px-2 whitespace-nowrap">
          Thêm thể loại
        </Button>
      </DialogTrigger>
      <DialogContent className="xl:max-w-[720px] max-w-[472px] p-0 bg-white">
        <DialogHeader>
          <DialogTitle className="p-6 pb-0">Thêm thể loại</DialogTitle>
        </DialogHeader>
        <form onSubmit={handleSubmit(onSubmit)}>
          <div className="p-6 flex flex-col gap-4 border-y-[1px]">
            <div>
              <Label htmlFor="nameNcc">Tên thể loại</Label>
              <Input id="nameNcc" {...register("name")}></Input>
              {errors.name && (
                <span className="error___message">{errors.name.message}</span>
              )}
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

export default CreateCategory;
