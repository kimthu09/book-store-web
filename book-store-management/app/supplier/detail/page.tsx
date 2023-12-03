"use client";
import { Card, CardContent } from "@/components/ui/card";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { importNotes, supplierDebts, suppliers } from "@/constants";
import React from "react";
import { ImportTable } from "@/components/supplier-manage/import-table";
import { DebtTable } from "@/components/supplier-manage/debt-table";
import Link from "next/link";
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
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { cn } from "@/lib/utils";
import { CalendarIcon } from "@radix-ui/react-icons";
import { format } from "date-fns";
import { Calendar } from "@/components/ui/calendar";

const SupplierDetail = ({ searchParams }: { searchParams: { id: string } }) => {
  const supplier = suppliers.find((item) => item.id === searchParams.id);
  const [date, setDate] = React.useState<Date>();

  return (
    <div className="col items-center">
      <div className="col xl:w-4/5 w-full xl:px-0 md:px-8 px-0">
        <div className="flex justify-between">
          <h1 className="xl:text-3xl text-2xl">{supplier?.name}</h1>
          <Dialog>
            <DialogTrigger asChild>
              <Button>Tạo phiếu chi</Button>
            </DialogTrigger>
            <DialogContent className="xl:max-w-[720px] max-w-[472px] p-0 bg-white">
              <DialogHeader>
                <DialogTitle className="p-6 pb-0">Tạo phiếu chi </DialogTitle>
              </DialogHeader>
              <div className="border-y-[1px]">
                <div className="p-6 flex flex-col gap-4">
                  <div>
                    <Label htmlFor="idPhieu">Mã phiếu</Label>
                    <Input
                      id="idPhieu"
                      placeholder="Hệ thống sẽ tự sinh mã nếu để trống"
                    ></Input>
                  </div>
                  <div>
                    <Label htmlFor="idNcc">Mã nhà cung cấp</Label>
                    <Input
                      readOnly
                      value={supplier?.id}
                      id="idNcc"
                      placeholder="Hệ thống sẽ tự sinh mã nếu để trống"
                    ></Input>
                  </div>
                  <div>
                    <Label htmlFor="nameNcc">Tên nhà cung cấp</Label>
                    <Input readOnly value={supplier?.name} id="nameNcc"></Input>
                  </div>
                  <div className="flex gap-4">
                    <div className="flex-1">
                      <Label htmlFor="price">Giá trị</Label>
                      <Input id="price" type="number"></Input>
                    </div>
                    <div className="flex-1">
                      <Label>Ngày ghi nhận</Label>
                      <Popover>
                        <PopoverTrigger asChild>
                          <Button
                            variant={"outline"}
                            className={cn(
                              "w-full justify-start text-left font-normal",
                              !date && "text-muted-foreground"
                            )}
                          >
                            <CalendarIcon className="mr-2 h-4 w-4" />
                            {date ? (
                              format(date, "dd/MM/yy")
                            ) : (
                              <span>Chọn ngày</span>
                            )}
                          </Button>
                        </PopoverTrigger>
                        <PopoverContent className="w-auto p-0" align="start">
                          <Calendar
                            mode="single"
                            selected={date}
                            onSelect={setDate}
                            initialFocus
                          />
                        </PopoverContent>
                      </Popover>
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

        <Card>
          <CardContent className="p-6 flex flex-col gap-4">
            <div className="flex gap-4 lg:flex-row flex-col">
              <div className="flex-1">
                <Label htmlFor="id">Mã nhà cung cấp</Label>
                <Input id="id" value={supplier?.id} readOnly></Input>
              </div>
              <div className="flex-1">
                <Label htmlFor="email">Email</Label>
                <Input id="email" value={supplier?.email} readOnly></Input>
              </div>
            </div>
            <div className="flex gap-4 lg:flex-row flex-col">
              <div className="flex-1">
                <Label>Số điện thoại</Label>
                <Input value={supplier?.phone} readOnly></Input>
              </div>
              <div className="flex-1">
                <Label>Nợ hiện tại</Label>
                <Input type="number" value={supplier?.debt} readOnly></Input>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardContent className="p-6 flex flex-col gap-4">
            <Tabs defaultValue="import" className="flex flex-col gap-4">
              <TabsList className="inline-flex  w-full    h-9 items-center text-muted-foreground justify-start rounded-none bg-transparent p-0">
                <TabsTrigger className="tab___trigger" value="import">
                  Lịch sử nhận hàng
                </TabsTrigger>
                <TabsTrigger className="tab___trigger" value="debt">
                  Công nợ
                </TabsTrigger>
              </TabsList>
              <TabsContent value="import">
                <ImportTable
                  data={importNotes.filter(
                    (item) => item.supplierId === supplier?.id
                  )}
                ></ImportTable>
              </TabsContent>
              <TabsContent value="debt">
                <DebtTable
                  data={supplierDebts.filter(
                    (item) => item.idSupplier === supplier?.id
                  )}
                />
              </TabsContent>
            </Tabs>
          </CardContent>
        </Card>
      </div>
    </div>
  );
};

export default SupplierDetail;
