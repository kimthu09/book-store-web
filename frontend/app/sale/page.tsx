import { Metadata } from "next";
import SaleScreen from "./page-layout";
import { withAuth } from "@/lib/role/withAuth";
export const metadata: Metadata = {
  title: "Bán hàng",
};
export const Sale = ({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) => {
  return <SaleScreen />;
};

export default withAuth(Sale, ["INVOICE_CREATE"]);
