import RecordNotFound from "@/components/record-notfound";
import { Card, CardContent } from "@/components/ui/card";
import getInvoiceDetail from "@/lib/invoice/getInvoiceDetail";
import { toVND } from "@/lib/utils";

const InvoiceDetails = async ({
  params,
}: {
  params: { invoiceId: string };
}) => {
  const response: Promise<any> = getInvoiceDetail({
    idInvoice: params.invoiceId,
  });
  const responseData = await response;
  console.log(responseData);
  if (responseData.hasOwnProperty("errorKey")) {
    return <RecordNotFound />;
  } else {
    console.log(responseData.details);
    const details = responseData.details as {
      book: { id: string; name: string };
      qty: number;
      unitPrice: number;
    }[];
    return (
      <div className="col items-center">
        <div className="col xl:w-4/5 w-full xl:px-0 md:px-8 px-0">
          <h1 className="lg:text-3xl text-2xl">
            Mã hóa đơn: {params.invoiceId}
          </h1>
          <Card>
            <CardContent className="p-6 px-2 flex flex-col gap-4">
              <h1 className="text-base px-4">Thông tin chung</h1>

              <div className="flex flex-row justify-between pb-2 px-4 border-b">
                <h2>Ngày tạo: </h2>
                <span>
                  {new Date(responseData.createdAt).toLocaleDateString("vi-VN")}
                </span>
              </div>
              <div className="flex flex-row justify-between pb-2 px-4 border-b">
                <h2>Người tạo: </h2>
                <span>{responseData.createdBy.name}</span>
              </div>
              <div className="flex flex-row justify-between px-4 ">
                <h2>Tổng hóa đơn: </h2>
                <span>{toVND(responseData.totalPrice)}</span>
              </div>
            </CardContent>
          </Card>
          <Card>
            <CardContent className="p-6 px-2 flex flex-col gap-4">
              <h1 className="text-base px-4">Thông tin thanh toán</h1>
              {details.map((item, index) => {
                return (
                  <div
                    key={item.book.id}
                    className={`flex flex-row justify-between pb-2 px-4 ${
                      index < details.length - 1 ? "border-b" : ""
                    }`}
                  >
                    <span>{item.book.name}</span>
                    <div>
                      <span className="text-center"> {item.qty} x </span>
                      <span className="text-right">
                        {toVND(item.unitPrice)}
                      </span>
                    </div>
                  </div>
                );
              })}
            </CardContent>
          </Card>
        </div>
      </div>
    );
  }
};

export default InvoiceDetails;
