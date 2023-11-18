"use client";
import ImportBooks from "@/components/stock-manage/import-books";
import SupplierList from "@/components/supplier-list";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useState } from "react";
import { useForm } from "react-hook-form";

export type FormValues = {
  books: {
    idBook: string;
    quantity: number;
    price: number;
  }[];
};

const NewImport = () => {
  const form = useForm<FormValues>({
    defaultValues: {
      books: [],
    },
  });
  const [supplier, setSupplier] = useState("");
  return (
    <div className="col items-center">
      <div className="col xl:w-4/5 w-full xl:px-0 md:px-6 px-0">
        <h1 className="font-medium text-xxl self-start">Phiếu nhập sách</h1>
        <form>
          <div className="flex flex-col gap-4">
            <Card>
              <CardContent className="p-6 flex lg:flex-row flex-col gap-5">
                <div className="flex-1">
                  <Label htmlFor="idPhieu">Mã phiếu</Label>
                  <Input
                    id="idPhieu"
                    placeholder="Hệ thống tự sinh mã nếu bỏ trống"
                  ></Input>
                </div>
                <div className="flex-1">
                  <Label htmlFor="idNcc">Nhà cung cấp</Label>
                  <SupplierList supplier={supplier} setSupplier={setSupplier} />
                </div>
              </CardContent>
            </Card>
            <Card>
              <CardContent className="p-6">
                <Label>Thông tin sách</Label>
                <ImportBooks form={form} />
              </CardContent>
            </Card>
          </div>
        </form>
      </div>
    </div>
  );
};

export default NewImport;
