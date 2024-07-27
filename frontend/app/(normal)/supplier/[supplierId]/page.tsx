import { Metadata } from "next";
import { withAuth } from "@/lib/role/withAuth";
import SupplierDetail from "./detail-layout";

const SupplierScreen = ({ params }: { params: { supplierId: string } }) => {
  return <SupplierDetail params={params} />;
};

export default withAuth(SupplierScreen, ["SUPPLIER_VIEW"]);
