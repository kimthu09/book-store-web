import { Card } from "../ui/card";
import { Skeleton } from "../ui/skeleton";

const DashboardCard = (props: any) => {
  const { title, icon, value, color, isLoading } = props;

  return (
    <Card
      className={`flex flex-col lg:gap-4 p-8 gap-2 flex-1 ${
        color ? color : ""
      }`}
    >
      {isLoading ? (
        <Skeleton className="h-6 w-16" />
      ) : (
        <div className="flex flex-row gap-2 items-center">
          <p className="text-lg">{title}</p>
          {icon}
        </div>
      )}
      {isLoading ? (
        <Skeleton className="h-6 w-20" />
      ) : (
        <p className="text-2xl font-semibold">{value}</p>
      )}
    </Card>
  );
};

export default DashboardCard;
