"use client";
import { useState } from "react";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "../ui/dialog";
import { Button } from "../ui/button";
import { Label } from "../ui/label";
import { Input } from "../ui/input";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { phoneRegex, required } from "@/constants";
import { SubmitHandler, useForm } from "react-hook-form";
import { FaPlus } from "react-icons/fa";
import RoleList from "./role-list";
import createStaff from "@/lib/staff/createStaff";
import { toast } from "../ui/use-toast";
import { useRouter } from "next/navigation";
import { imageUpload } from "@/lib/staff/uploadImage";
import { useCurrentUser } from "@/hooks/use-user";
import { isAdmin } from "@/lib/utils";
import { useLoading } from "@/hooks/loading-context";

const StaffSchema = z.object({
  name: required,
  email: z.string().email("Email không hợp lệ"),
  phone: z.string().regex(phoneRegex, "Số điện thoại không hợp lệ"),
  address: z.string(),
  roleId: z.string().min(1, "Không để trống trường này"),
  img: z.string(),
});

const CreateStaffDialog = () => {
  const {
    register,
    handleSubmit,
    reset,
    setValue,
    trigger,
    formState: { errors },
  } = useForm<z.infer<typeof StaffSchema>>({
    shouldUnregister: false,
    resolver: zodResolver(StaffSchema),
    defaultValues: {
      name: "",
      email: "",
      phone: "",
      address: "",
      roleId: "",
      img: "/avatar.png",
    },
  });
  const { showLoading, hideLoading } = useLoading();
  const router = useRouter();
  const onSubmit: SubmitHandler<z.infer<typeof StaffSchema>> = async (data) => {
    if (image) {
      let formData = new FormData();

      formData.append("file", image);
      formData.append("folderName", "avatars");
      showLoading();
      const imgRes = await imageUpload(formData);
      if (imgRes.hasOwnProperty("errorKey")) {
        toast({
          variant: "destructive",
          title: "Có lỗi",
          description: imgRes.message,
        });
        return;
      } else {
        data.img = imgRes.data;
      }
    }

    const response: Promise<any> = createStaff({ staff: data });
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
        description: "Thêm nhân viên thành công",
      });
      setOpen(false);

      router.refresh();
    }
  };
  const [open, setOpen] = useState(false);
  const [role, setRole] = useState("");

  const handleRole = (role: string) => {
    setRole(role);
    setValue("roleId", role);
    trigger("roleId");
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
  const isAdminRole = currentUser && isAdmin({ currentUser: currentUser });
  if (!isAdminRole) {
    return null;
  } else
    return (
      <Dialog
        open={open}
        onOpenChange={(open) => {
          if (open) {
            reset({
              name: "",
              email: "",
              phone: "",
              address: "",
              roleId: "",
              img: "/avatar.png",
            });
            setRole("");
            setImage(null);
          }
          register("roleId");
          setOpen(open);
        }}
      >
        <DialogTrigger asChild>
          <Button className="lg:px-4 px-2 whitespace-nowrap">
            <div className="flex flex-wrap gap-1 items-center">
              <FaPlus />
              Thêm nhân viên
            </div>
          </Button>
        </DialogTrigger>
        <DialogContent className="xl:max-w-[720px] max-w-[472px] p-0 bg-white">
          <DialogHeader>
            <DialogTitle className="p-6 pb-0"> Thêm nhân viên</DialogTitle>
          </DialogHeader>
          <form onSubmit={handleSubmit(onSubmit)}>
            <div className="p-6 flex flex-col gap-4 border-y-[1px]">
              <div>
                <Label htmlFor="nameNcc">Tên nhân viên</Label>
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
              <div>
                <Label htmlFor="address">Địa chỉ</Label>
                <Input id="address" {...register("address")}></Input>
                {errors.address && (
                  <span className="error___message">
                    {errors.address.message}
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
                <div className="flex-1">
                  <Label>Phân quyền</Label>
                  <RoleList role={role} setRole={handleRole} />
                  {errors.roleId && (
                    <span className="error___message">
                      Không để trống trường này
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

export default CreateStaffDialog;
