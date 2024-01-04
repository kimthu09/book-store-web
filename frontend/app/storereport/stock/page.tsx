import StockReport from "@/components/report/StockReport";
import { withAuth } from "@/lib/role/withAuth";
import { Metadata } from "next";

export const metadata: Metadata = {
    title: "Báo cáo tồn kho",
};
const StockReportPage = () => {
    return (
        <StockReport />
    );
}

export default withAuth(StockReportPage, ["REPORT_VIEW_STOCK"]);
