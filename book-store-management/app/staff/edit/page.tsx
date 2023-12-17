"use client";

import RoleList from "@/components/staff/role-list";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useState } from "react";
import { AiOutlineClose } from "react-icons/ai";
import { LuCheck } from "react-icons/lu";

const EditStaff = ({ searchParams }: { searchParams: { id: string } }) => {
  const [role, setRole] = useState("");
  const [readOnly, setReadOnly] = useState(true);
  const staff = { id: "1", name: "hi" };
  return (
    <div className="col items-center">
      <div className="col xl:w-4/5 w-full xl:px-0 md:px-8 px-0">
        <div className="flex flex-row justify-between">
          <h1 className="font-medium text-xxl self-start">
            Nhân viên: {staff?.id}
          </h1>
          <div className="flex gap-2">
            {!readOnly ? (
              <div className="flex gap-2">
                <Button
                  variant={"outline"}
                  className="bg-white border-primary text-primary hover:text-primary px-2"
                  onClick={() => setReadOnly(true)}
                >
                  <div className="flex flex-wrap gap-[2px] items-center">
                    <AiOutlineClose />
                    Huỷ
                  </div>
                </Button>
                <Button className="px-2" onClick={() => setReadOnly(true)}>
                  <div className="flex flex-wrap gap-[2px] items-center">
                    <LuCheck />
                    Lưu
                  </div>
                </Button>
              </div>
            ) : (
              <Button
                className="px-3"
                onClick={() => {
                  setReadOnly(false);
                }}
              >
                Chỉnh sửa
              </Button>
            )}
            <div>
              <Button
                type="button"
                className="bg-green-500 hover:bg-green-500/90 px-3"
              >
                Đổi mật khẩu
              </Button>
            </div>
          </div>
        </div>
        <form className="flex flex-col gap-4">
          <Card>
            <CardContent className="p-6 flex flex-col gap-4">
              <div className="flex flex-col gap-4 lg:flex-row">
                <div className="basis-2/3">
                  <Label htmlFor="hoten">Họ và tên</Label>
                  <Input
                    id="hoten"
                    readOnly={readOnly}
                    defaultValue={staff?.name}
                  ></Input>
                </div>
                <div className="basis-1/3">
                  <Label>Phân quyền</Label>
                  <RoleList role={role} setRole={setRole} />
                </div>
              </div>
              <div className="flex flex-col gap-4 lg:flex-row">
                <div className="basis-2/3">
                  <Label htmlFor="email">Email</Label>
                  <Input id="email" type="email" readOnly={readOnly}></Input>
                </div>
                <div className="basis-1/3">
                  <Label htmlFor="dienthoai">Điện thoại</Label>
                  <Input id="dienthoai" readOnly={readOnly}></Input>
                </div>
              </div>
            </CardContent>
          </Card>
        </form>
      </div>
    </div>
  );
};

export default EditStaff;
