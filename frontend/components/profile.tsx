import { useCurrentUser } from "@/hooks/use-user";
import { Button } from "./ui/button";
import Image from "next/image";
import { Popover, PopoverContent, PopoverTrigger } from "./ui/popover";
import { useEffect, useState } from "react";
import { getApiKey, logOut } from "@/lib/auth/action";
import { LuLogOut } from "react-icons/lu";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "./ui/dialog";
import { Label } from "./ui/label";
import { Input } from "./ui/input";
import { endPoint, required } from "@/constants";
import { z } from "zod";
import { SubmitErrorHandler, SubmitHandler, useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { MdOutlineLock } from "react-icons/md";
import axios from "axios";
import { toast } from "./ui/use-toast";
const PasswordSchema = z
  .object({
    oldPassword: z.string().min(6, "Ít nhất 6 ký tự"),
    newPassword: z.string().min(6, "Ít nhất 6 ký tự"),
    confirmNewPass: required,
  })
  .refine((data) => data.newPassword === data.confirmNewPass, {
    message: "Mật khẩu mới không khớp",
    path: ["confirmNewPass"],
  });
const Profile = () => {
  const { currentUser } = useCurrentUser();
  const [open, setOpen] = useState(false);
  const [openPass, setOpenPass] = useState(false);

  const [avatar, setAvatar] = useState(null);
  const [name, setName] = useState({ name: "", id: "" });
  useEffect(() => {
    if (currentUser) {
      const json = JSON.stringify(currentUser);
      const user = JSON.parse(json);
      const image = user.data.img;
      if (image) {
        setAvatar(image);
      }

      const userName = user.data.name;
      const userId = user.data.id;
      if (name && userId) {
        setName({ name: userName, id: userId });
      }
    }
  }, [currentUser]);
  const passForm = useForm<z.infer<typeof PasswordSchema>>({
    resolver: zodResolver(PasswordSchema),
    defaultValues: {
      oldPassword: "",
      newPassword: "",
      confirmNewPass: "",
    },
  });
  const {
    register: registerPass,
    handleSubmit: handleSubmitPass,
    control: controlPass,
    reset: resetPass,
    formState: { errors: errorPass },
  } = passForm;
  const onSubmitPass: SubmitHandler<z.infer<typeof PasswordSchema>> = async (
    data
  ) => {
    const token = await getApiKey();
    const res = axios
      .patch(
        `${endPoint}/v1/password`,
        {
          newPassword: data.newPassword,
          oldPassword: data.oldPassword,
        },
        {
          headers: {
            accept: "application/json",
            Authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
          },
        }
      )
      .then((response) => {
        if (response) return response.data;
      })
      .catch((error) => {
        console.error("Error:", error);
        return error.response.data;
      });
    const responseData = await res;
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
        description: "Đổi mật khẩu thành công",
      });
      setOpen(false);
    }
    setOpenPass(false);
  };

  const onOpenPass = (value: boolean) => {
    if (value) {
      resetPass();
    }
    setOpenPass(value);
  };
  return (
    <div>
      <Popover open={open} onOpenChange={setOpen}>
        <PopoverTrigger asChild>
          <Button className="p-0 h-fit rounded-full bg-background overflow-clip">
            <div>
              <Image
                src={`${avatar && avatar !== "" ? avatar : "/avatar.png"}`}
                className="h-9 w-9"
                alt="avatar"
                width={36}
                height={36}
              ></Image>
            </div>
          </Button>
        </PopoverTrigger>
        <PopoverContent className="w-44 mx-4 flex flex-col gap-2 p-0 py-2">
          <div className="flex flex-col px-4 pt-2">
            <span className="text-xs">{name.id}</span>
            <h1 className="text-lg font-semibold">{name.name}</h1>
          </div>
          <Dialog open={openPass} onOpenChange={onOpenPass}>
            <DialogTrigger asChild>
              <Button variant={"ghost"} className="rounded-none justify-start">
                <div className="flex gap-2 items-center">
                  <MdOutlineLock className="h-5 w-5" />
                  Đổi mật khẩu
                </div>
              </Button>
            </DialogTrigger>
            <DialogContent className="max-w-[472px] p-0 bg-white">
              <DialogHeader>
                <DialogTitle className="p-6 pb-0">Đổi mật khẩu</DialogTitle>
              </DialogHeader>
              <form onSubmit={handleSubmitPass(onSubmitPass)}>
                <div className="p-6 flex flex-col gap-4 border-y-[1px]">
                  <div>
                    <Label htmlFor="pass">Mật khẩu hiện tại</Label>
                    <Input
                      type="password"
                      id="pass"
                      {...registerPass("oldPassword")}
                    ></Input>
                    {errorPass.oldPassword && (
                      <span className="error___message">
                        {errorPass.oldPassword.message}
                      </span>
                    )}
                  </div>
                  <div>
                    <Label htmlFor="new">Mật khẩu mới</Label>
                    <Input
                      type="password"
                      id="new"
                      {...registerPass("newPassword")}
                    ></Input>
                    {errorPass.newPassword && (
                      <span className="error___message">
                        {errorPass.newPassword.message}
                      </span>
                    )}
                  </div>
                  <div>
                    <Label htmlFor="new">Xác nhận mật khẩu mới</Label>
                    <Input
                      type="password"
                      id="confirm"
                      {...registerPass("confirmNewPass")}
                    ></Input>
                    {errorPass.confirmNewPass && (
                      <span className="error___message">
                        {errorPass.confirmNewPass.message}
                      </span>
                    )}
                  </div>
                </div>
                <div className="p-4 flex-1 flex justify-end">
                  <div className="flex gap-4">
                    <Button
                      type="button"
                      variant={"outline"}
                      onClick={() => setOpenPass(false)}
                    >
                      Huỷ
                    </Button>

                    <Button>Xác nhận</Button>
                  </div>
                </div>
              </form>
            </DialogContent>
          </Dialog>
          <form action={logOut} className="w-full">
            <Button
              variant={"ghost"}
              className="rounded-none w-full justify-start"
            >
              <div className="flex gap-2 text-primary">
                <LuLogOut className="h-5 w-5 " />
                Đăng xuất
              </div>
            </Button>
          </form>
        </PopoverContent>
      </Popover>
    </div>
  );
};

export default Profile;
