"use client";
import React from "react";
import { InvoiceDetailTable } from "@/components/invoice/invoice-detail-table";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { toVND } from "@/lib/utils";
import Image from "next/image";
import { useRef } from "react";
import ReactToPrint, { useReactToPrint } from "react-to-print";
import { FaPrint } from "react-icons/fa";
import { InvoiceDetailProps } from "@/app/invoice/[invoiceId]/page";
import { ShopGeneral } from "@/types";
const PrintInvoice = ({
  onPrint,
  responseData,
}: {
  responseData: any;
  onPrint: () => void;
}) => {
  const componentRef = useRef(null);
  const details = responseData.invoice.details as InvoiceDetailProps[];
  const shop = responseData.shop as ShopGeneral;
  const handlePrint = useReactToPrint({
    content: () => componentRef.current,
  });
  return (
    <Card>
      <CardContent className="flex flex-col p-0 gap-4 relative">
        <div
          className="whitespace-nowrap text-primary-foreground shadow py-2 inline-flex h-8 shrink-0 items-center justify-center rounded-md border bg-green-500 px-3 text-sm font-medium   disabled:pointer-events-none disabled:opacity-50 hover:bg-green-500/90 boder-none"
          onClick={() => {
            handlePrint();
            onPrint();
          }}
        >
          In hóa đơn
        </div>
        <div ref={componentRef} className="printScreen">
          <div className="p-6">
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
                <h1 className="text-3xl uppercase font-medium">{shop.name}</h1>
              </div>
              {shop.address && shop.address !== "" ? (
                <span className="text-lg uppercase font-medium">
                  Địa chỉ: {shop.address}
                </span>
              ) : null}
              {shop.phone && shop.phone !== "" ? (
                <span className="text-base font-medium">
                  Số điện thoại: {shop.phone}
                </span>
              ) : null}
              {shop.wifiPass && shop.wifiPass !== "" ? (
                <span className="text-base font-light">
                  Wifi: {shop.wifiPass}
                </span>
              ) : null}
            </div>

            <div className="flex flex-row justify-between p-4 mb-6 border rounded-md">
              <div className="flex flex-col items-stretch gap-2 text-sm">
                <div className="flex gap-2">
                  <span className="font-light w-[6rem]">Mã hóa đơn:</span>
                  <span className="font-semibold">
                    {responseData.invoice.id}
                  </span>
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
                <div className="grid grid-cols-2  space-x-2 font-semibold mt-10">
                  <span className="col-span-2 text-center italic">
                    Cảm ơn và hẹn gặp lại!
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  );
};

export default PrintInvoice;
