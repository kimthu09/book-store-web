import { withAuth } from "@/lib/role/withAuth";
import EditStaff from "./detail-layout";

const StaffDetailScreen = ({ params }: { params: { staffId: string } }) => {
  return <EditStaff params={params} />;
};

export default withAuth(StaffDetailScreen, ["USER_VIEW"]);
