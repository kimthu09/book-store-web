"use client";
import { Card, CardContent } from "@/components/ui/card";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useState } from "react";
import { ImportTable } from "@/components/supplier-manage/import-table";
import { DebtTable } from "@/components/supplier-manage/debt-table";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import getSupplier from "@/lib/supplier/getSupplier";
import Loading from "@/components/loading";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { SubmitHandler, useForm } from "react-hook-form";
import { toast } from "@/components/ui/use-toast";
import paySupplier from "@/lib/supplier/paySupplier";
import EditDialog from "@/components/supplier-manage/edit";
import { useSWRConfig } from "swr";
import { endPoint } from "@/constants";
const FormSchema = z.object({
  quantity: z.coerce.number().gte(1, "Giá trị phải lớn hơn 0"), // Force it to be a number
});
const SupplierDetail = ({ params }: { params: { supplierId: string } }) => {
  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
  });
  const {
    register,
    handleSubmit,
    control,
    formState: { errors },
  } = form;
  const { data, isLoading, isError } = getSupplier(params.supplierId);

  const { mutate } = useSWRConfig();

  const onSubmit: SubmitHandler<z.infer<typeof FormSchema>> = async (
    dataCreate
  ) => {
    const response: Promise<any> = paySupplier({
      quantity: dataCreate.quantity,
      idSupplier: params.supplierId,
    });
    const responseData = await response;
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
        description: "Lập phiếu chi thành công",
      });
      mutate(`${endPoint}/v1/suppliers/${data.id}/debts?page=${pageIndex}`);
    }
    setOpenDialog(false);
  };
  const [openDialog, setOpenDialog] = useState(false);
  const [pageIndex, setPageIndex] = useState(1);
  const [tab, setTab] = useState("import");
  const onTabChange = (value: string) => {
    setTab(value);
  };

  if (isError) return <div>Failed to load</div>;
  if (isLoading) {
    return <Loading />;
  } else
    return (
      <div className="col items-center">
        <div className="col xl:w-4/5 w-full xl:px-0 md:px-8 px-0">
          <div className="flex justify-between">
            <h1 className="xl:text-3xl text-2xl">{data.name}</h1>
            <div className="flex gap-2">
              <EditDialog
                supplier={data}
                refresh={() => {
                  mutate(`${endPoint}/v1/suppliers/${data.id}`);
                }}
              />
              <Dialog open={openDialog} onOpenChange={setOpenDialog}>
                <DialogTrigger asChild>
                  <Button>Tạo phiếu chi</Button>
                </DialogTrigger>
                <DialogContent className="xl:max-w-[720px] max-w-[472px] p-0 bg-white">
                  <DialogHeader>
                    <DialogTitle className="p-6 pb-0">
                      Tạo phiếu chi
                    </DialogTitle>
                  </DialogHeader>
                  <form
                    className="border-y-[1px]"
                    onSubmit={handleSubmit(onSubmit)}
                  >
                    <div className="p-6 flex flex-col gap-4">
                      {/* <div>
                      <Label htmlFor="idPhieu">Mã phiếu</Label>
                      <Input
                        id="idPhieu"
                        placeholder="Hệ thống sẽ tự sinh mã nếu để trống"
                      ></Input>
                    </div> */}
                      <div>
                        <Label htmlFor="idNcc">Mã nhà cung cấp</Label>
                        <Input
                          readOnly
                          value={data.id}
                          id="idNcc"
                          placeholder="Hệ thống sẽ tự sinh mã nếu để trống"
                        ></Input>
                      </div>
                      <div>
                        <Label htmlFor="nameNcc">Tên nhà cung cấp</Label>
                        <Input readOnly value={data.name} id="nameNcc"></Input>
                      </div>
                      <div className="flex gap-4">
                        <div className="flex-1">
                          <Label htmlFor="price">Giá trị</Label>
                          <Input
                            id="price"
                            type="number"
                            {...register("quantity")}
                          ></Input>
                          {errors.quantity && (
                            <span className="error___message">
                              {errors.quantity.message}
                            </span>
                          )}
                        </div>
                      </div>
                    </div>
                    <div className="flex gap-4 px-6 py-4 border-t justify-end">
                      <Button
                        type="button"
                        onClick={() => setOpenDialog(false)}
                        variant={"outline"}
                      >
                        Huỷ
                      </Button>
                      <Button type="submit">Thêm</Button>
                    </div>
                  </form>
                </DialogContent>
              </Dialog>
            </div>
          </div>

          <Card>
            <CardContent className="p-6 flex flex-col gap-4">
              <div className="flex gap-4 lg:flex-row flex-col">
                <div className="flex-1">
                  <Label htmlFor="id">Mã nhà cung cấp</Label>
                  <Input id="id" value={data.id} readOnly></Input>
                </div>
                <div className="flex-1">
                  <Label htmlFor="email">Email</Label>
                  <Input id="email" value={data.email} readOnly></Input>
                </div>
              </div>
              <div className="flex gap-4 lg:flex-row flex-col">
                <div className="flex-1">
                  <Label>Số điện thoại</Label>
                  <Input value={data.phone} readOnly></Input>
                </div>
                <div className="flex-1">
                  <Label>Nợ hiện tại</Label>
                  <Input type="number" value={data.debt} readOnly></Input>
                </div>
              </div>
            </CardContent>
          </Card>

          <Card>
            <CardContent className="p-6 flex flex-col gap-4 relative">
              <Tabs
                defaultValue={tab}
                onValueChange={onTabChange}
                className="flex flex-col gap-4"
              >
                <TabsList className="inline-flex w-full  h-9 items-center text-muted-foreground justify-start rounded-none bg-transparent p-0">
                  <TabsTrigger className="tab___trigger" value="import">
                    Lịch sử nhập hàng
                  </TabsTrigger>
                  <TabsTrigger className="tab___trigger" value="debt">
                    Công nợ
                  </TabsTrigger>
                </TabsList>
                <TabsContent value="import">
                  <ImportTable supplierId={params.supplierId}></ImportTable>
                </TabsContent>
                <TabsContent value="debt">
                  <DebtTable
                    supplierId={params.supplierId}
                    pageIndex={pageIndex}
                    setPageIndex={setPageIndex}
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
