import { InvoiceDetailTable } from "@/components/invoice/invoice-detail-table";
import RecordNotFound from "@/components/record-notfound";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import getInvoiceDetail from "@/lib/invoice/getInvoiceDetail";
import { toVND } from "@/lib/utils";
import Image from "next/image";
import { useRef } from "react";
import ReactToPrint from "react-to-print";
import DetailLayout from "./detail-layout";

export type InvoiceDetailProps = {
  book: { id: string; name: string };
  qty: number;
  unitPrice: number;
};
const InvoiceDetails = async ({
  params,
}: {
  params: { invoiceId: string };
}) => {
  const response: Promise<any> = getInvoiceDetail({
    idInvoice: params.invoiceId,
  });
  const responseData = await response;
  if (responseData.hasOwnProperty("errorKey")) {
    return <RecordNotFound />;
  } else {
    // console.log(responseData.details);

    return (
      <div className="col items-center">
        <div className="col xl:w-4/5 w-full xl:px-0 md:px-8 px-0">
          <DetailLayout {...responseData} />
        </div>
      </div>
    );
  }
};

export default InvoiceDetails;
