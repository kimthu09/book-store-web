"use client";
import getDashboard from "@/lib/dashboard/getDashboard";
import { CardDashboardInfo, Dashboard } from "@/types";
import { useState } from "react";
import { toast } from "../ui/use-toast";
import DashboardHeader from "./dashboard-header";
import DashboardCardHolder from "./dashboard-card-container";
import DashboardTopFoodContainer from "./dashboard-top-food-container";
import Loading from "../loading";
import DashboardChartContainer from "./dashboard-chart-container";
import React from "react";
import { FaBook, FaChartLine, FaUserTag } from "react-icons/fa";
import { GiShamrock } from "react-icons/gi";

const DashboardComponent = () => {
  const [data, setData] = useState<Dashboard>({
    timeFrom: new Date(),
    timeTo: new Date(),
    totalSale: 0,
    totalCustomer: 0,
    totalSold: 0,
    totalPoint: 0,
    topSoldBooks: [] as unknown as [
      { id: string; name: string; qty: number; sale: number }
    ],
    chartPriceComponents: [] as unknown as [{ time: Date; value: number }],
    chartProfitComponents: [] as unknown as [{ time: Date; value: number }],
  });
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [chartType, setChartType] = React.useState("day");

  const onGetDashboard = async ({
    timeFrom,
    timeTo,
  }: {
    timeFrom: number;
    timeTo: number;
  }) => {
    setIsLoading(true);
    const responseData = await getDashboard({
      timeFrom: timeFrom,
      timeTo: timeTo,
    });
    if (responseData.hasOwnProperty("data")) {
      if (responseData.data) {
        setData(responseData.data);
      }
    } else if (responseData.hasOwnProperty("errorKey")) {
      toast({
        variant: "destructive",
        title: "Có lỗi",
        description: responseData.message,
      });
    } else {
      toast({
        variant: "destructive",
        title: "Có lỗi",
        description: "Vui lòng thử lại sau",
      });
    }
    setIsLoading(false);
  };
  let cardInfos: CardDashboardInfo[] = [];

  const totalSaleFormat = new Intl.NumberFormat("vi-VN", {
    style: "currency",
    currency: "VND",
  }).format(data.totalSale);

  const pointFormat = data.totalPoint.toLocaleString("vi-VN");

  const customerFormat = data.totalCustomer.toLocaleString("vi-VN");

  const totalSoldFormat = data.totalSold.toLocaleString("vi-VN");

  if (data != undefined) {
    cardInfos.push({
      title: "Doanh thu",
      value: totalSaleFormat,
      icon: <FaChartLine className="h-6 w-6" />,
    });
    cardInfos.push({
      title: "Số điểm tích được",
      value: pointFormat,
      icon: <GiShamrock className="h-6 w-6" />,
    });
    cardInfos.push({
      title: "Số khách đã mua",
      value: customerFormat,
      icon: <FaUserTag className="h-6 w-6" />,
    });
    cardInfos.push({
      title: "Số sản phẩm bán được",
      value: totalSoldFormat,
      icon: <FaBook className="h-6 w-6" />,
    });
  }
  return (
    <div className="flex flex-col lg:gap-[6] gap-4">
      <div>
        <DashboardHeader onClick={onGetDashboard} />
      </div>
      <div className="flex flex-col w-full lg:gap-[6] gap-4">
        <DashboardCardHolder cardInfos={cardInfos} isLoading={isLoading} />
        <div className="flex lg:flex-row flex-col w-full lg:gap-[6] gap-4 h-auto basis-3/5">
          <DashboardChartContainer
            price={data?.chartPriceComponents}
            profit={data?.chartProfitComponents}
            timeFrom={data?.timeFrom}
            timeTo={data?.timeTo}
            chartType={chartType}
            setChartType={setChartType}
            isLoading={isLoading}
          />
          <div className="basis-2/5 flex h-full">
            <DashboardTopFoodContainer foods={data?.topSoldBooks} isLoading={isLoading} />
          </div>
        </div>
      </div>
    </div>
  );
};

export default DashboardComponent;
