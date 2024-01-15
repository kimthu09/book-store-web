import { Invoice, PagingProps } from "@/types";
import InvoiceTable from "./table";
import getAllInvoice from "@/lib/invoice/getAllInvoice";

const TableLayout = async ({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) => {
  const page = searchParams["page"] ?? "1";
  const maxPrice = searchParams["maxPrice"] ?? undefined;
  const minPrice = searchParams["minPrice"] ?? undefined;
  const createdBy = searchParams["createdBy"] ?? undefined;

  const search = searchParams["search"] ?? undefined;
  const staffsData: Promise<{ paging: PagingProps; data: Invoice[] }> =
    getAllInvoice({
      page: +page,
      maxPrice: maxPrice?.toString(),
      minPrice: minPrice?.toString(),
      createdBy: createdBy?.toString(),
      search: search?.toString(),
    });
  const staffs = await staffsData;
  const totalPage = Math.ceil(staffs.paging.total / staffs.paging.limit);
  return <InvoiceTable data={staffs.data} totalPage={totalPage} />;
};

export default TableLayout;
