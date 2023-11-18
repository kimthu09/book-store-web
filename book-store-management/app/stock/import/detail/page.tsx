"use client";
import { ExportImportDetail } from "@/components/excel-export";
import ImportDetailTable from "@/components/stock-manage/impor-detail-table";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { importDetails, importNotes } from "@/constants";
import React from "react";
import { FiDownload } from "react-icons/fi";
const ImportNoteDetail = ({
  searchParams,
}: {
  searchParams: { id: string };
}) => {
  const details = importDetails.filter(
    (item) => item.idNote === searchParams.id
  );
  const note = importNotes.find((item) => item.id === searchParams.id);
  return (
    <div className="col items-center">
      <div className="col xl:w-4/5 w-full xl:px-0 md:px-8 px-0">
        <div className="flex justify-between">
          <h1 className="xl:text-3xl text-2xl">Chi tiết phiếu nhập</h1>

          <div className="flex gap-2">
            <Button
              className="p-2"
              variant={"ghost"}
              onClick={() =>
                ExportImportDetail(note!, details, "PhieuNhap.xlsx")
              }
            >
              <div className="flex flex-wrap gap-1 items-center">
                <FiDownload />
                Tải phiếu
              </div>
            </Button>
            <h1 className="xl:text-lg text-base text-primary px-4 py-1 rounded-full bg-blue-200 self-start">
              {note?.status}
            </h1>
          </div>
        </div>

        <Card>
          <CardContent className="p-6 flex flex-col gap-4">
            <div className="flex gap-4 lg:flex-row flex-col">
              <div className="basis-1/3">
                <Label htmlFor="id">Mã phiếu</Label>
                <Input id="id" value={note?.id} readOnly></Input>
              </div>
              <div className="basis-2/3">
                <Label htmlFor="name">Nhà cung cấp</Label>
                <Input id="name" value={note?.supplierId} readOnly></Input>
              </div>
            </div>
            <div className="flex gap-4 lg:flex-row flex-col">
              <div className="basis-1/3">
                <Label>Ngày lập phiếu</Label>
                <Input
                  value={note?.createAt.toLocaleDateString("vi-VN")}
                  readOnly
                ></Input>
              </div>
              <div className="basis-2/3">
                <Label>Người lập phiếu</Label>
                <Input value={note?.supplierId} readOnly></Input>
              </div>
            </div>
          </CardContent>
        </Card>

        <div className=" p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
          <ImportDetailTable importDetails={details} />
        </div>

        <div className="flex justify-end gap-2">
          <Button variant={"outline"} className="bg-white">
            Huỷ
          </Button>
          <Button>Đặt hàng</Button>
          <Button>Nhập kho</Button>
        </div>
      </div>
    </div>
  );
};

export default ImportNoteDetail;
