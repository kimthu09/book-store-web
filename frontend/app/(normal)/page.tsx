import DashboardComponent from "@/components/dashboard/dashboard";
import { getUser } from "@/lib/auth/action";
import { withAuth } from "@/lib/role/withAuth";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Trang chủ",
};
const DashboardPage = () => {
  return <DashboardComponent />;
};

export default withAuth(DashboardPage, ["REPORT_VIEW_SALE"]);