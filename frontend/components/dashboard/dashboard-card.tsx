import { Card } from "../ui/card";

const DashboardCard = (props: any) => {
  const { title, icon, value, color } = props;

  return (
    <Card
      className={`flex flex-col lg:gap-4 p-8 gap-2 flex-1 ${
        color ? color : ""
      }`}
    >
      <div className="flex flex-row gap-2 items-center">
        <p className="text-lg">{title}</p>
        {icon}
      </div>
      <p className="text-2xl font-semibold">{value}</p>
    </Card>
  );
};

export default DashboardCard;
