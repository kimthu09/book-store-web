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
import { ShopGeneral } from "@/types";
const DetailLayout = (responseData: any) => {
  const componentRef = useRef(null);
  const details = responseData.invoice.details as InvoiceDetailProps[];
  const shop = responseData.shop as ShopGeneral;

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
          <div className="flex flex-col gap-2 items-center mb-6">
            <div className="flex items-center justify-center gap-2">
              <Image
                className="object-contain w-auto h-auto"
                src="/android-chrome-192x192.png"
                priority
                alt="logo"
                width={36}
                height={36}
              ></Image>
              <h1 className="xl:text-3xl text-2xl uppercase font-medium">
                {shop.name}
              </h1>
            </div>
            {shop.address && shop.address !== "" ? (
              <span className="text-base uppercase font-medium printScreen">
                Địa chỉ: {shop.address}
              </span>
            ) : null}
            {shop.phone && shop.phone !== "" ? (
              <span className="text-base font-medium printScreen">
                Số điện thoại: {shop.phone}
              </span>
            ) : null}
            {shop.wifiPass && shop.wifiPass !== "" ? (
              <span className="text-base font-light printScreen">
                Wifi: {shop.wifiPass}
              </span>
            ) : null}
          </div>

          <div className="flex flex-row justify-between p-4 mb-6 border rounded-md">
            <div className="flex flex-col items-stretch gap-2 text-sm">
              <div className="flex gap-2">
                <span className="font-light w-[6rem]">Mã hóa đơn:</span>
                <span className="font-semibold">{responseData.invoice.id}</span>
              </div>
              {responseData.invoice.customer && (
                <>
                  <div className="flex gap-2">
                    <span className="font-light w-[6rem] whitespace-nowrap">
                      Khách hàng:
                    </span>
                    <div className="font-semibold flex flex-col">
                      {responseData.invoice.customer.name}
                      <span className="font-normal">
                        ({responseData.invoice.customer.phone})
                      </span>
                    </div>
                  </div>
                  <div className="flex gap-2">
                    <span className="font-light w-[6rem] whitespace-nowrap">
                      Điểm:
                    </span>
                    <div className="font-semibold flex gap-2 text-green-700">
                      +{" "}
                      {responseData.invoice.pointReceive.toLocaleString(
                        "vi-VN"
                      )}
                    </div>
                  </div>
                </>
              )}
            </div>

            <div className="flex flex-col items-end gap-2 text-sm">
              <span className="font-semibold">
                {new Date(responseData.invoice.createdAt).toLocaleTimeString(
                  "vi-VN",
                  {
                    hour: "2-digit",
                    minute: "2-digit",
                  }
                )}
                {", "}
                {new Date(responseData.invoice.createdAt).toLocaleDateString(
                  "vi-VN"
                )}
              </span>
              <span>Nhân viên: {responseData.invoice.createdBy.name}</span>
            </div>
          </div>
          <div className="flex flex-col gap-4">
            <InvoiceDetailTable details={details} />
            <div className="flex flex-col gap-2">
              {responseData.invoice.amountPriceUsePoint !== 0 ? (
                <>
                  <div className="grid grid-cols-2  space-x-2 font-semibold">
                    <span className="min-w-[6rem]">Tổng tiền: </span>
                    <span className="text-right">
                      {toVND(responseData.invoice.totalPrice)}
                    </span>
                  </div>
                  <div className="grid grid-cols-2  space-x-2 font-semibold">
                    <span className="min-w-[6rem]">Giảm: </span>
                    <span className="text-right">
                      - {toVND(responseData.invoice.amountPriceUsePoint)}
                    </span>
                  </div>
                </>
              ) : null}
              <div className="grid grid-cols-2  space-x-2 font-semibold">
                <span className="min-w-[6rem]">Thành tiền: </span>
                <span className="text-right">
                  {toVND(responseData.invoice.amountReceived)}
                </span>
              </div>
              <div className="grid grid-cols-2  space-x-2 font-semibold mt-10 printScreen">
                <span className="col-span-2 text-center italic">
                  Cảm ơn và hẹn gặp lại!
                </span>
              </div>
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  );
};

export default DetailLayout;
