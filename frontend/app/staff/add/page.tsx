"use client";

import { useState } from "react";
import RoleList from "@/components/staff/role-list";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { LuCheck } from "react-icons/lu";
const AddStaff = () => {
  const [role, setRole] = useState("");

  return (
    <div className="col items-center">
      <div className="col xl:w-4/5 w-full xl:px-0 md:px-8 px-0">
        <div className="flex flex-row justify-between">
          <h1 className="font-medium text-xxl self-start">Thêm nhân viên</h1>
          <Button>
            <div className="flex flex-wrap gap-1 items-center">
              <LuCheck />
              Thêm
            </div>
          </Button>
        </div>
        <form className="flex flex-col gap-4">
          <Card>
            <CardContent className="p-6 flex flex-col gap-4">
              <div className="flex flex-col gap-4 lg:flex-row">
                <div className="basis-2/3">
                  <Label htmlFor="hoten">Họ và tên</Label>
                  <Input id="hoten"></Input>
                </div>
                <div className="basis-1/3">
                  <Label>Phân quyền</Label>
                  <RoleList role={role} setRole={setRole} />
                </div>
              </div>
              <div className="flex flex-col gap-4 lg:flex-row">
                <div className="basis-2/3">
                  <Label htmlFor="email">Email</Label>
                  <Input id="email" type="email"></Input>
                </div>
                <div className="basis-1/3">
                  <Label htmlFor="dienthoai">Điện thoại</Label>
                  <Input id="dienthoai"></Input>
                </div>
              </div>
            </CardContent>
          </Card>
        </form>
      </div>
    </div>
  );
};

export default AddStaff;
