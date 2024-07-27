"use client";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { required } from "@/constants";
import {
  SubmitErrorHandler,
  SubmitHandler,
  useFieldArray,
  useForm,
} from "react-hook-form";
import { Card, CardContent } from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import { Checkbox } from "@/components/ui/checkbox";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import getRole from "@/lib/staff/getRole";
import Loading from "@/components/loading";
import { useEffect, useState } from "react";
import { toast } from "@/components/ui/use-toast";
import updateRole from "@/lib/staff/updateRole";
import { useCurrentUser } from "@/hooks/use-user";
import { isAdmin } from "@/lib/utils";
import NoRole from "@/components/no-role";
import { useLoading } from "@/hooks/loading-context";
import DropdownSkeleton from "@/components/skeleton/dropdown-skeleton";
import FeatureSkeleton from "@/components/skeleton/feature-skeleton";

const FormSchema = z.object({
  name: required,
  features: z.array(z.object({ idFeature: z.string() })),
});
const RoleDetail = ({ params }: { params: { roleId: string } }) => {
  const { response, isLoading, isError } = getRole(params.roleId);
  const roleFunctions = response
    ? (response.data! as {
        id: string;
        description: string;
        groupName: string;
        isHas: boolean;
      }[])
    : [];
  const name = response ? response.name : "";
  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
  });
  const {
    register,
    handleSubmit,
    control,
    reset,
    setValue,
    formState: { errors },
  } = form;
  const { fields, append, remove } = useFieldArray({
    control: control,
    name: "features",
  });
  const { showLoading, hideLoading } = useLoading();
  const onSubmit: SubmitHandler<z.infer<typeof FormSchema>> = async (data) => {
    setReadOnly(true);
    const response: Promise<any> = updateRole({
      id: params.roleId,
      name: data.name,
      features: data.features.map((item) => item.idFeature),
    });
    showLoading();
    const responseData = await response;
    hideLoading();
    if (responseData.hasOwnProperty("data")) {
      if (responseData.data) {
        toast({
          variant: "success",
          title: "Thành công",
          description: "Chỉnh sửa phân quyền mới thành công",
        });
      }
    } else if (responseData.hasOwnProperty("errorKey")) {
      reset({
        name: name!,
        features: roleFunctions.map((item) => {
          if (item.isHas) return { idFeature: item.id };
        }),
      });
      toast({
        variant: "destructive",
        title: "Có lỗi",
        description: responseData.message,
      });
    } else {
      reset({
        name: name!,
        features: roleFunctions.map((item) => {
          if (item.isHas) return { idFeature: item.id };
        }),
      });
      toast({
        variant: "destructive",
        title: "Có lỗi",
        description: "Vui lòng thử lại sau",
      });
    }
  };
  const onSelect = (idFeature: string) => {
    if (!readOnly) {
      const index = fields.findIndex(
        (feature) => feature.idFeature === idFeature
      );
      if (index > -1) {
        remove(index);
      } else {
        append({ idFeature: idFeature });
      }
    }
  };
  const onError: SubmitErrorHandler<z.infer<typeof FormSchema>> = async (
    data
  ) => {
    setReadOnly(true);
  };

  const [readOnly, setReadOnly] = useState(true);
  useEffect(() => {
    if (response) {
      const features = roleFunctions
        .filter((item) => item.isHas)
        .map((feature) => {
          return { idFeature: feature.id };
        });
      reset({
        name: name!,
        features: features,
      });
    }
  }, [response]);
  const { currentUser } = useCurrentUser();
  const isAdminRole = currentUser && isAdmin({ currentUser: currentUser });
  if (isLoading || !currentUser) {
    return (
      <div className="col items-center">
        <div className="col xl:w-4/5 w-full xl:px-0 md:px-8 px-0">
          <div className="flex flex-row justify-between">
            <h1 className="font-medium text-xxl self-start flex-1">
              Chi tiết phân quyền
            </h1>
          </div>
          <div className="flex flex-col gap-4">
            <Card>
              <CardContent className="p-6">
                <DropdownSkeleton />
              </CardContent>
            </Card>
            <Card>
              <CardContent className="p-6">
                <FeatureSkeleton />
              </CardContent>
            </Card>
          </div>
        </div>
      </div>
    );
  } else if (!isAdminRole) {
    return <NoRole />;
  } else if (isError) return <div>Failed to load</div>;
  else {
    return (
      <div className="col items-center">
        <form
          className="col xl:w-4/5 w-full xl:px-0 md:px-8 px-0"
          onSubmit={handleSubmit(onSubmit, onError)}
        >
          <div className="flex flex-row justify-between">
            <h1 className="font-medium text-xxl self-start flex-1">
              Chi tiết phân quyền
            </h1>
            <div className="flex gap-2">
              <div className={`flex gap-2 ${readOnly ? "hidden" : "flex"}`}>
                <Button
                  type="button"
                  variant={"outline"}
                  className="bg-white border border-primary text-primary hover:text-primary"
                  onClick={() => {
                    setReadOnly(true);
                    reset({
                      name: name!,
                      features: roleFunctions.map((item) => {
                        if (item.isHas) return { idFeature: item.id };
                      }),
                    });
                  }}
                >
                  Hủy
                </Button>
                <Button type="submit">Lưu</Button>
              </div>

              <Button
                type="button"
                className={`${readOnly ? "flex" : "hidden"}`}
                onClick={() => setReadOnly(false)}
              >
                Chỉnh sửa
              </Button>
            </div>
          </div>
          <div className="flex flex-col gap-4">
            <Card>
              <CardContent className="p-6">
                <Label htmlFor="tenPhanQuyen">Tên phân quyền</Label>
                <Input
                  id="tenPhanQuyen"
                  {...register("name")}
                  readOnly={readOnly}
                ></Input>
                {errors.name && (
                  <span className="error___message">{errors.name.message}</span>
                )}
              </CardContent>
            </Card>
            <Card>
              <CardContent className="p-6">
                {
                  <div className="grid xl:grid-cols-3 lg:grid-cols-2 md:grid-cols-1 sm:grid-cols-2 grid-cols-1 gap-y-6 gap-x-4">
                    {Array.from(
                      new Set(roleFunctions.map((item) => item.groupName))
                    ).map((value) => {
                      return (
                        <div key={value}>
                          <div>
                            <span className="text-base text-primary">
                              {value}
                            </span>
                          </div>
                          <div className="flex flex-col gap-2 mt-2">
                            {roleFunctions.map((item) => {
                              return item.groupName === value ? (
                                <div key={item.id} className="flex gap-2">
                                  <Checkbox
                                    id={item.id}
                                    checked={
                                      fields.findIndex(
                                        (feature) =>
                                          feature.idFeature === item.id
                                      ) > -1
                                    }
                                    onClick={() => onSelect(item.id)}
                                  ></Checkbox>
                                  <Label onClick={() => onSelect(item.id)}>
                                    {item.description}
                                  </Label>
                                </div>
                              ) : null;
                            })}
                          </div>
                        </div>
                      );
                    })}
                  </div>
                }
              </CardContent>
            </Card>
          </div>
        </form>
      </div>
    );
  }
};

export default RoleDetail;
