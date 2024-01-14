"use client";
import { Card } from "@/components/ui/card";
import { StockReport, StockReportDetail } from "@/types";
import getReport from "@/lib/report/getReport";
import { useState } from "react";
import { StockReportTable } from "@/components/report/StockReportTable";
import MonthReportHeader from "@/components/report/MonthReportHeader";
import { toast } from "../ui/use-toast";
import Loading from "../loading";
import { ExportStockReport } from "./excel-export-stock-report";
import TableSkeleton from "../skeleton/table-skeleton";

const StockReport = () => {
  const [data, setData] = useState<StockReport>();
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const onGetStock = async ({
    timeFrom,
    timeTo,
  }: {
    timeFrom: number;
    timeTo: number;
  }) => {
    setIsLoading(true);
    const responseData = await getReport({
      timeFrom: timeFrom,
      timeTo: timeTo,
      type: "stock",
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

  const onExport = () => {
    if (data == undefined || data.details.length < 1) {
      toast({
        variant: "destructive",
        title: "Có lỗi",
        description: "Không có báo cáo tồn kho nào",
      });
    } else {
      ExportStockReport(data, "StockReport.xlsx");
    }
  };

  return (
    <div>
      <div>
        <MonthReportHeader
          title="Báo cáo tồn"
          firstAction="Xem báo cáo"
          secondAction="Tải excel"
          onClick={onGetStock}
          onExport={onExport}
        />
      </div>

      <div>
        <Card className="p-[10px] my-[22px]">
          {isLoading ? (
            <TableSkeleton
              isHasExtensionAction={false}
              isHasFilter={false}
              isHasSearch={false}
              isHasChooseVisibleRow={false}
              isHasCheckBox={false}
              isHasPaging={false}
              numberRow={10}
              cells={[
                {
                  percent: 2,
                },
                {
                  percent: 2,
                },
                {
                  percent: 5,
                },
                {
                  percent: 2,
                },
                {
                  percent: 2,
                },
                {
                  percent: 2,
                },
                {
                  percent: 2,
                },
                {
                  percent: 2,
                },
              ]}
            ></TableSkeleton>
          ) : (
            <StockReportTable
              data={
                data == undefined || data == null
                  ? []
                  : (data!.details as StockReportDetail[])
              }
              report={data}
            />
          )}
        </Card>
      </div>
    </div>
  );
};

export default StockReport;
