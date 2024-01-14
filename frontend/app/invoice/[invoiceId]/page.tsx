import RecordNotFound from "@/components/record-notfound";

import getInvoiceDetail from "@/lib/invoice/getInvoiceDetail";

import DetailLayout from "./detail-layout";
import { getUser } from "@/lib/auth/action";
import { includesRoles } from "@/lib/utils";
import NoRole from "@/components/no-role";

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
  const user = await getUser();
  if (
    !includesRoles({ currentUser: user, allowedFeatures: ["INVOICE_VIEW"] })
  ) {
    return <NoRole />;
  } else {
    const response: Promise<any> = getInvoiceDetail({
      idInvoice: params.invoiceId,
    });
    const responseData = await response;
    if (responseData.hasOwnProperty("errorKey")) {
      return <RecordNotFound />;
    } else {
      return (
        <div className="col items-center">
          <div className="col xl:w-4/5 w-full xl:px-0 md:px-8 px-0">
            <DetailLayout {...responseData} />
          </div>
        </div>
      );
    }
  }
};

export default InvoiceDetails;
