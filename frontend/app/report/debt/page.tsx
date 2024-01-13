import DebtReport from "@/components/report/DebtReport";
import { withAuth } from "@/lib/role/withAuth";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Báo cáo nợ",
};
const DebtReportPage = () => {
  return <DebtReport />;
};

export default withAuth(DebtReportPage, ["REPORT_VIEW_SUPPLIER"]);
