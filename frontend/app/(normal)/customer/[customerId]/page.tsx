import { withAuth } from "@/lib/role/withAuth";
import CustomerDetail from "./detail-layout";

const CustomerScreen = ({ params }: { params: { customerId: string } }) => {
  return <CustomerDetail params={params} />;
};

export default withAuth(CustomerScreen, ["CUSTOMER_VIEW"]);
