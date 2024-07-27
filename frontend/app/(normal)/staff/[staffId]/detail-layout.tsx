"use client";

import Loading from "@/components/loading";
import RoleList from "@/components/staff/role-list";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import getStaff from "@/lib/staff/getStaff";
import { useEffect, useState } from "react";
import { AiOutlineClose } from "react-icons/ai";
import { GoArrowSwitch } from "react-icons/go";
import { LuCheck } from "react-icons/lu";
import { z } from "zod";
import Image from "next/image";
import { FaPen } from "react-icons/fa";
import { MdLockReset } from "react-icons/md";
import { endPoint, phoneRegex, required } from "@/constants";
import { SubmitHandler, useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { toast } from "@/components/ui/use-toast";
import axios from "axios";
import ChangeImage from "@/components/staff/change-image";
import { imageUpload } from "@/lib/staff/uploadImage";
import updateStaff from "@/lib/staff/updateStaff";
import ConfirmDialog from "@/components/confirm-dialog";
import { GoPerson } from "react-icons/go";
import { getApiKey } from "@/lib/auth/action";
import StatusList from "@/components/status-list";
import changeStaffStatus from "@/lib/staff/changeStaffStatus";
import { useCurrentUser } from "@/hooks/use-user";
import { includesRoles, isAdmin } from "@/lib/utils";
import { useLoading } from "@/hooks/loading-context";
import StaffDetailSkeleton from "@/components/skeleton/staff-detail";
import { useRouter } from "next/navigation";

const FormSchema = z.object({
  name: required,
  phone: z.string().regex(phoneRegex, "Số điện thoại không hợp lệ"),
  address: z.string(),
});
const PasswordSchema = z.object({
  pass: required,
});

const EditStaff = ({ params }: { params: { staffId: string } }) => {
  const router = useRouter();
  const [role, setRole] = useState("");
  const { showLoading, hideLoading } = useLoading();

  const [status, setStatus] = useState(true);
  const displayStatus = {
    trueText: "Đang làm việc",
    falseText: "Đã nghỉ việc",
  };
  const [open, setOpen] = useState(false);
  const [readOnly, setReadOnly] = useState(true);
  const { data, isLoading, isError, mutate } = getStaff(params.staffId);
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
  useEffect(() => {
    if (data) {
      reset({
        name: data.name,
        phone: data.phone,
        address: data.address,
      });
      setRole(data.role.id);
      setStatus(data.isActive);
    }
  }, [data]);
  const passForm = useForm<z.infer<typeof PasswordSchema>>({
    resolver: zodResolver(PasswordSchema),
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
    showLoading();
    const token = await getApiKey();
    const res = axios
      .patch(
        `${endPoint}/v1/users/${params.staffId}/reset`,
        {
          userSenderPass: data.pass,
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
        description: "Đặt lại mật khẩu nhân viên thành công",
      });
      setOpen(false);
      router.refresh();
    }
  };
  const onSubmit: SubmitHandler<z.infer<typeof FormSchema>> = async (data) => {
    setReadOnly(true);
    showLoading();
    const response: Promise<any> = updateStaff({
      id: params.staffId,
      address: data.address,
      phone: data.phone,
      name: data.name,
    });
    const responseData = await response;
    hideLoading();
    if (data.hasOwnProperty("errorKey")) {
      toast({
        variant: "destructive",
        title: "Có lỗi",
        description: responseData.message,
      });
    } else {
      toast({
        variant: "success",
        title: "Thành công",
        description: "Chỉnh sửa thông tin nhân viên thành công",
      });
      mutate();
      router.refresh();
    }
  };
  const handleOpen = (value: boolean) => {
    if (value) {
      resetPass({ pass: "" });
    }
    setOpen(value);
  };
  const [image, setImage] = useState<any>();
  const handleImageSelected = async () => {
    if (!image) {
      return;
    }
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
      const response: Promise<any> = updateStaff({
        id: params.staffId,
        img: imgRes.data,
      });
      const data = await response;
      hideLoading();

      if (data.hasOwnProperty("errorKey")) {
        toast({
          variant: "destructive",
          title: "Có lỗi",
          description: data.message,
        });
      } else {
        toast({
          variant: "success",
          title: "Thành công",
          description: "Thay đổi ảnh nhân viên thành công",
        });
        mutate();
        router.refresh();
      }
    }
  };
  const changeRole = async () => {
    const token = await getApiKey();
    const res = axios
      .patch(
        `${endPoint}/v1/users/${params.staffId}/role`,
        {
          roleId: role,
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
        description: "Thay đổi vai trò nhân viên thành công",
      });
      setOpen(false);
      mutate();
      router.refresh();
    }
  };

  const changeStatus = async () => {
    showLoading();
    const responseData = await changeStaffStatus({
      userIds: [params.staffId],
      isActive: status,
    });
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
        description: "Thay đổi trạng thái nhân viên thành công",
      });
      setOpen(false);
      mutate();
      router.refresh();
    }
  };
  const { currentUser } = useCurrentUser();
  const canEdit =
    currentUser &&
    includesRoles({
      currentUser: currentUser,
      allowedFeatures: ["USER_UPDATE_INFO"],
    });
  const isAdminRole = currentUser && isAdmin({ currentUser: currentUser });
  if (isLoading) {
    return <StaffDetailSkeleton />;
  } else {
    return (
      <div className="col items-center">
        <div className="col xl:w-4/5 w-full px-0">
          <div className="flex flex-row justify-between">
            <h1 className="font-medium text-2xl self-start">
              Nhân viên: {data?.id}
            </h1>
          </div>
          <form className="flex flex-col gap-4">
            <Card>
              <CardContent className="p-6 flex gap-6">
                <div className="relative h-min">
                  <div className="relative rounded-full border overflow-clip">
                    <Image
                      src={data.img}
                      alt="ảnh"
                      className="h-[100px] w-[100px] object-cover"
                      width={100}
                      height={100}
                    ></Image>
                  </div>
                  {canEdit ? (
                    <ChangeImage
                      image={image}
                      setImage={setImage}
                      handleImageSelected={handleImageSelected}
                      currentImage={data.img}
                    />
                  ) : null}
                </div>
                <div className="flex-1 flex-col flex gap-4">
                  <div>
                    <Label htmlFor="hoten">Họ và tên</Label>
                    <Input
                      id="hoten"
                      readOnly={readOnly}
                      {...register("name")}
                    ></Input>
                  </div>
                  <div className="flex flex-col gap-4 lg:flex-row">
                    <div className="basis-2/3">
                      <Label htmlFor="email">Email</Label>
                      <Input
                        id="email"
                        type="email"
                        defaultValue={data?.email}
                        readOnly
                      ></Input>
                    </div>
                    <div className="basis-1/3">
                      <Label htmlFor="dienthoai">Điện thoại</Label>
                      <Input
                        id="dienthoai"
                        readOnly={readOnly}
                        {...register("phone")}
                      ></Input>
                    </div>
                  </div>
                  <div>
                    <Label htmlFor="diachi">Địa chỉ</Label>
                    <Input
                      id="diachi"
                      readOnly={readOnly}
                      {...register("address")}
                    ></Input>
                  </div>
                  <div className="flex gap-2 justify-end sm:flex-row flex-col">
                    <div className="flex gap-2 justify-center">
                      {!readOnly ? (
                        <div className="flex gap-2 sm:flex-initial flex-1">
                          <Button
                            variant={"outline"}
                            className="bg-white border-rose-700 text-rose-700 hover:text-rose-700 hover:bg-rose-50/30 px-2 flex gap-1 flex-nowrap whitespace-nowrap flex-1"
                            onClick={() => {
                              setReadOnly(true);
                              reset({
                                name: data.name,
                                phone: data.phone,
                                address: data.address,
                              });
                            }}
                          >
                            <AiOutlineClose className="h-5 w-5" />
                            Hủy
                          </Button>
                          <ConfirmDialog
                            title={"Xác nhận"}
                            description="Bạn xác nhận chỉnh sửa nhân viên này ?"
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
                      ) : canEdit ? (
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
                    {isAdminRole ? (
                      <Dialog open={open} onOpenChange={handleOpen}>
                        <DialogTrigger asChild>
                          <Button
                            title="Đặt lại mật khẩu"
                            className="bg-teal-600 hover:bg-teal-600/90 px-2 flex gap-1 flex-nowrap sm:flex-initial whitespace-nowrap"
                            type="button"
                          >
                            <MdLockReset className="h-8 w-8" />
                            Đặt lại mật khẩu
                          </Button>
                        </DialogTrigger>
                        <DialogContent className="max-w-[472px] p-0 bg-white">
                          <DialogHeader>
                            <DialogTitle className="p-6 pb-0">
                              Xác nhận đặt lại mật khẩu ?
                            </DialogTitle>
                          </DialogHeader>
                          <form onSubmit={handleSubmitPass(onSubmitPass)}>
                            <div className="p-6 flex flex-col gap-4 border-y-[1px]">
                              <div>
                                <Label htmlFor="pass">
                                  Mật khẩu người dùng
                                </Label>
                                <Input
                                  type="password"
                                  id="pass"
                                  {...registerPass("pass")}
                                ></Input>
                                {errorPass.pass && (
                                  <span className="error___message">
                                    {errorPass.pass.message}
                                  </span>
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

                                <Button>Xác nhận</Button>
                              </div>
                            </div>
                          </form>
                        </DialogContent>
                      </Dialog>
                    ) : null}
                  </div>
                </div>
              </CardContent>
            </Card>
          </form>
          <div className="flex lg:flex-row flex-col gap-4">
            <Card className="flex-1">
              <CardContent className="p-6 flex items-end justify-between gap-2">
                {isAdminRole ? (
                  <>
                    <div className="flex-1">
                      <Label>Vai trò</Label>
                      <RoleList role={role} setRole={setRole} />
                    </div>
                    <ConfirmDialog
                      title={"Xác nhận"}
                      description="Bạn xác nhận thay đổi vai trò nhân viên ?"
                      handleYes={() => changeRole()}
                      handleNo={() => setRole(data.role.id)}
                    >
                      <Button
                        disabled={role === data.role.id}
                        variant={"outline"}
                        className="px-2 flex gap-1 flex-nowrap w-[8.5rem]"
                      >
                        <GoPerson className="w-4 h-4" />
                        Đổi vai trò
                      </Button>
                    </ConfirmDialog>
                  </>
                ) : (
                  <div className="flex-1">
                    <Label>Vai trò</Label>
                    <Input readOnly value={data.role.name} />
                  </div>
                )}
              </CardContent>
            </Card>
            <Card className="flex-1">
              <CardContent className="p-6 flex items-end justify-between gap-2">
                {isAdminRole ? (
                  <>
                    <div className="flex-1">
                      <Label>Trạng thái</Label>
                      <StatusList
                        status={status}
                        setStatus={setStatus}
                        display={displayStatus}
                      />
                    </div>
                    <ConfirmDialog
                      title={"Xác nhận"}
                      description="Bạn xác nhận thay đổi trạng thái nhân viên ?"
                      handleYes={() => changeStatus()}
                      handleNo={() => setStatus(data.isActive)}
                    >
                      <Button
                        disabled={status === data.isActive}
                        variant={"outline"}
                        className="px-2 flex gap-1 flex-nowrap w-[8.5rem]"
                      >
                        <GoArrowSwitch className="w-4 h-4" />
                        Đổi trạng thái
                      </Button>
                    </ConfirmDialog>
                  </>
                ) : (
                  <div className="flex-1">
                    <Label>Trạng thái</Label>

                    <Input
                      readOnly
                      value={
                        data.isActive
                          ? displayStatus.trueText
                          : displayStatus.falseText
                      }
                    ></Input>
                  </div>
                )}
              </CardContent>
            </Card>
          </div>
        </div>
      </div>
    );
  }
};

export default EditStaff;
