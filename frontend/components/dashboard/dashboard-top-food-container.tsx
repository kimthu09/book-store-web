import TableSkeleton from "../skeleton/table-skeleton";
import { Card } from "../ui/card";
import { DashboardTopFoodTable } from "./dashboard-top-food-table";

const DashboardTopFoodContainer = (props: any) => {
  const { foods, isLoading } = props;
  return (
    <Card className="flex flex-col flex-1 w-full h-full gap-2 p-4">
      <p className="font-semibold text-base">Sản phẩm bán chạy</p>
      {isLoading ? (
        <TableSkeleton
          isHasExtensionAction={false}
          isHasFilter={false}
          isHasSearch={false}
          isHasChooseVisibleRow={false}
          isHasCheckBox={false}
          isHasPaging={false}
          numberRow={5}
          cells={[
            {
              percent: 1
            },
            {
              percent: 5
            },
            {
              percent: 2
            },
            {
              percent: 2
            }
          ]}
        />
      ) : (
        <DashboardTopFoodTable data={foods} />
      )}
    </Card>
  );
};

export default DashboardTopFoodContainer;
