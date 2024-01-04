import SaleReport from "@/components/report/SaleReport";
import { withAuth } from "@/lib/role/withAuth";
import { Metadata } from "next";

export const metadata: Metadata = {
    title: "Báo cáo mặt hàng",
};
const SaleReportPage = () => {
    return (
        <SaleReport />
    );
}

export default withAuth(SaleReportPage, ["REPORT_VIEW_SALE"]);
