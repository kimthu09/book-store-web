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
import createPublisher from "@/lib/book/createPublisher";
import { useLoading } from "@/hooks/loading-context";
const required = z.string().min(1, "Không để trống trường này");

const FormSchema = z.object({
  name: required,
});

const CreatePublisher = ({
  handlePublisherAdded,
  children,
}: {
  handlePublisherAdded: (publiserId: string) => void;
  children: React.ReactNode;
}) => {
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
  });
  const router = useRouter();
  const { showLoading, hideLoading } = useLoading();
  const onSubmit: SubmitHandler<z.infer<typeof FormSchema>> = async (data) => {
    setOpen(false);

    const response: Promise<any> = createPublisher(data);
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
        description: "Thêm nhà xuất bản thành công",
      });
      handlePublisherAdded(responseData.data);
      setOpen(false);
    }
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
      <DialogTrigger asChild>{children}</DialogTrigger>
      <DialogContent className="max-w-[472px] p-0 bg-white">
        <DialogHeader>
          <DialogTitle className="p-6 pb-0"> Thêm nhà xuất bản</DialogTitle>
        </DialogHeader>
        <form>
          <div className="p-6 flex flex-col gap-4 border-y-[1px]">
            <div>
              <Label htmlFor="nameNcc">Tên nhà xuất bản</Label>
              <Input id="nameNcc" {...register("name")}></Input>
              {errors.name && (
                <span className="error___message">{errors.name.message}</span>
              )}
            </div>
          </div>
          <div className="p-4 flex-1 flex justify-end">
            <div className="flex gap-4">
              <Button
                type="button"
                variant={"outline"}
                onClick={() => setOpen(false)}
              >
                Huỷ
              </Button>

              <Button type="button" onClick={handleSubmit(onSubmit)}>
                Thêm
              </Button>
            </div>
          </div>
        </form>
      </DialogContent>
    </Dialog>
  );
};

export default CreatePublisher;
