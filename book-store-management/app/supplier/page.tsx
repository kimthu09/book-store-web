import { SupplierTable } from "@/components/supplier-manage/table";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import React from "react";

const SupplierManage = () => {
  return (
    <div className="col">
      <div className="flex flex-row justify-between ">
        <h1>Danh sách nhà cung cấp</h1>
        <Dialog>
          <DialogTrigger asChild>
            <Button>Thêm nhà cung cấp</Button>
          </DialogTrigger>
          <DialogContent className="xl:max-w-[720px] max-w-[472px] p-0 bg-white">
            <DialogHeader>
              <DialogTitle className="p-6 pb-0">Thêm nhà cung cấp</DialogTitle>
            </DialogHeader>
            <div className="border-y-[1px]">
              <div className="p-6 flex flex-col gap-4">
                <div>
                  <Label htmlFor="idNcc">Mã nhà cung cấp</Label>
                  <Input
                    id="idNcc"
                    placeholder="Hệ thống sẽ tự sinh mã nếu để trống"
                  ></Input>
                </div>
                <div>
                  <Label htmlFor="nameNcc">Tên nhà cung cấp</Label>
                  <Input id="nameNcc"></Input>
                </div>
                <div>
                  <Label htmlFor="email">Email</Label>
                  <Input id="email"></Input>
                </div>
                <div className="flex gap-4">
                  <div className="flex-1">
                    <Label htmlFor="phone">Số điện thoại</Label>
                    <Input id="phone"></Input>
                  </div>
                  <div className="flex-1">
                    <Label htmlFor="noBanDau">Công nợ</Label>
                    <Input id="noBanDau" type="number"></Input>
                  </div>
                </div>
              </div>
            </div>
            <DialogFooter className="p-6 pt-0">
              <DialogClose asChild>
                <div className="flex gap-4">
                  <Button type="button" variant={"outline"}>
                    Huỷ
                  </Button>

                  <Button type="submit">Thêm</Button>
                </div>
              </DialogClose>
            </DialogFooter>
          </DialogContent>
        </Dialog>
      </div>

      <div className="my-4 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <SupplierTable />
      </div>
    </div>
  );
};

export default SupplierManage;
