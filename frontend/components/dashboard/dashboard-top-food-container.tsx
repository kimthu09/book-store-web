import { Card } from "../ui/card";
import { DashboardTopFoodTable } from "./dashboard-top-food-table";

const DashboardTopFoodContainer = (props: any) => {
  const { foods } = props;
  return (
    <Card className="flex flex-col flex-1 w-full h-full gap-2 p-4">
      <p className="font-semibold text-xl">Sản phẩm bán chạy</p>
      <DashboardTopFoodTable data={foods} />
    </Card>
  );
};

export default DashboardTopFoodContainer;
