import DashboardCard from "./dashboard-card";

const DashboardCardHolder = (props: any) => {
  const { isLoading, cardInfos } = props;
  return (
    <div className="flex lg:flex-row flex-col gap-4">
      <div className="flex flex-row gap-4 flex-1">
        <DashboardCard
          title={cardInfos[0].title}
          icon={cardInfos[0].icon}
          value={cardInfos[0].value}
          color="bg-[rgba(207,250,218,0.5)] text-[#2f6e3f]"
          isLoading={isLoading}
        />
        <DashboardCard
          title={cardInfos[1].title}
          icon={cardInfos[1].icon}
          value={cardInfos[1].value}
          color="bg-[rgba(177,240,252,0.5)] text-[#0a6273]"
          isLoading={isLoading}
        />
      </div>
      <div className="flex flex-row lg:gap-[6] gap-4 flex-1">
        <DashboardCard
          title={cardInfos[2].title}
          icon={cardInfos[2].icon}
          value={cardInfos[2].value}
          color="bg-[rgba(247,231,161,0.5)] text-[#876a03]"
          isLoading={isLoading}
        />
        <DashboardCard
          title={cardInfos[3].title}
          icon={cardInfos[3].icon}
          value={cardInfos[3].value}
          color="bg-[rgba(209,214,255,0.5)] text-[#081687]"
          isLoading={isLoading}
        />
      </div>
    </div>
  );
};

export default DashboardCardHolder;
