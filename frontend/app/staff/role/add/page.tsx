"use client";
import Loading from "@/components/loading";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Checkbox } from "@/components/ui/checkbox";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import getAllRoleFunction from "@/lib/staff/getAllFunction";
import React from "react";
import { LuCheck } from "react-icons/lu";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { required } from "@/constants";
import { SubmitHandler, useFieldArray, useForm } from "react-hook-form";
import { toast } from "@/components/ui/use-toast";
import createRole from "@/lib/staff/createRole";
import { useCurrentUser } from "@/hooks/use-user";
import { isAdmin } from "@/lib/utils";
import NoRole from "@/components/no-role";
import { useLoading } from "@/hooks/loading-context";
import DropdownSkeleton from "@/components/skeleton/dropdown-skeleton";
import FeatureSkeleton from "@/components/skeleton/feature-skeleton";

const FormSchema = z.object({
  name: required,
  features: z.array(z.object({ idFeature: z.string(), groupName: z.string() })),
});

const AddRole = () => {
  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
  });
  const {
    register,
    handleSubmit,
    control,
    watch,
    formState: { errors },
  } = form;
  const { roleFunctions, isLoading, isError } = getAllRoleFunction();
  const { fields, append, remove } = useFieldArray({
    control: control,
    name: "features",
  });
  const onSelect = (featureId: string, groupName: string) => {
    const selectedIndex = fields.findIndex(
      (feature) => feature.idFeature === featureId
    );
    if (selectedIndex > -1) {
      remove(selectedIndex);
    } else {
      append({ idFeature: featureId, groupName: groupName });
    }
  };
  const { showLoading, hideLoading } = useLoading();
  const onSubmit: SubmitHandler<z.infer<typeof FormSchema>> = async (data) => {
    const response: Promise<any> = createRole({
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
          description: "Thêm phân quyền mới thành công",
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
  };
  const { currentUser } = useCurrentUser();
  const isAdminRole = currentUser && isAdmin({ currentUser: currentUser });
  if (isLoading || !currentUser) {
    return (
      <div className="col items-center">
        <div className="col xl:w-4/5 w-full xl:px-0 md:px-8 px-0">
          <div className="flex flex-row justify-between">
            <h1 className="font-medium text-xxl self-start flex-1">
              Thêm phân quyền
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
  else
    return (
      <div className="col items-center">
        <form
          className="col xl:w-4/5 w-full xl:px-0 md:px-8 px-0"
          onSubmit={handleSubmit(onSubmit)}
        >
          <div className="flex flex-row justify-between">
            <h1 className="font-medium text-xxl self-start flex-1">
              Thêm phân quyền
            </h1>
            <Button type="submit">
              <div className="flex gap-1 items-center flex-nowrap">
                <LuCheck />
                Thêm
              </div>
            </Button>
          </div>
          <div className="flex flex-col gap-4">
            <Card>
              <CardContent className="p-6">
                <Label htmlFor="tenPhanQuyen">Tên phân quyền</Label>
                <Input id="tenPhanQuyen" {...register("name")}></Input>
                {errors.name && (
                  <span className="error___message">{errors.name.message}</span>
                )}
              </CardContent>
            </Card>
            <Card>
              <CardContent className="p-6">
                {!roleFunctions ? (
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
                ) : (
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

                            <span className="text-muted-foreground text-sm">
                              {" (đã chọn "}
                              {
                                fields.filter(
                                  (item) => item.groupName === value
                                ).length
                              }
                              {")"}
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
                                    onClick={() =>
                                      onSelect(item.id, item.groupName)
                                    }
                                  ></Checkbox>
                                  <Label
                                    onClick={() =>
                                      onSelect(item.id, item.groupName)
                                    }
                                  >
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
                )}
              </CardContent>
            </Card>
          </div>
        </form>
      </div>
    );
};

export default AddRole;
