"use client";
import React from "react";
import { InvoiceDetailTable } from "@/components/invoice/invoice-detail-table";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { toVND } from "@/lib/utils";
import Image from "next/image";
import { useRef } from "react";
import ReactToPrint from "react-to-print";
import { InvoiceDetailProps } from "./page";
import { FaPrint } from "react-icons/fa";
const DetailLayout = (responseData: any) => {
  const componentRef = useRef(null);
  const details = responseData.details as InvoiceDetailProps[];
  return (
    <Card>
      <CardContent className="flex flex-col p-0 gap-4 relative">
        <ReactToPrint
          trigger={() => {
            return (
              <Button
                variant={"outline"}
                className="absolute top-6 right-6 flex gap-2"
              >
                <FaPrint className="w-5 h-5 text-primary" />
                In
              </Button>
            );
          }}
          content={() => componentRef.current}
        />
        <div ref={componentRef} className="p-6 ">
          <div className="flex justify-center gap-2 mb-6">
            <Image
              className="object-contain w-auto h-auto"
              src="/android-chrome-192x192.png"
              priority
              alt="logo"
              width={50}
              height={50}
            ></Image>
            <h1 className=" py-6 text-3xl uppercase font-medium">Nhà sách</h1>
          </div>

          <div className="flex flex-row justify-between p-4 mb-6 border rounded-md">
            <div className="flex gap-2">
              <span className="font-light">Mã hóa đơn:</span>
              <span className="font-semibold">{responseData.id}</span>
            </div>
            <div className="flex flex-col items-end gap-2 text-sm">
              <span className="font-semibold">
                {new Date(responseData.createdAt).toLocaleDateString("vi-VN")}
              </span>
              <span>Nhân viên: {responseData.createdBy.name}</span>
            </div>
          </div>
          <div className="flex flex-col gap-4">
            <InvoiceDetailTable {...details} />
            <div className="flex justify-end space-x-2 py-4 font-semibold">
              <span>Tổng tiền: </span>
              <span>{toVND(responseData.totalPrice)}</span>
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  );
};

export default DetailLayout;
