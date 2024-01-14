"use client";
import { Card } from "@/components/ui/card";
import ReportHeader from "@/components/report/ReportHeader";
import getReport from "@/lib/report/getReport";
import { useState } from "react";
import { SaleReportTable } from "@/components/report/SaleReportTable";
import { toast } from "../ui/use-toast";
import Loading from "../loading";
import { SaleReport, SaleReportDetail } from "@/types";
import { ExportSaleReport } from "./excel-export-sale-report";
import TableSkeleton from "../skeleton/table-skeleton";

const SaleReport = () => {
  const [data, setData] = useState<SaleReport>();
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const onGetSale = async ({
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
      type: "sale",
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
        description: "Không có báo cáo mặt hàng nào",
      });
    } else {
      ExportSaleReport(data, "SaleReport.xlsx");
    }
  };

  return (
    <div>
      <div>
        <ReportHeader
          title="Báo cáo mặt hàng"
          firstAction="Xem báo cáo"
          secondAction="Tải excel"
          onClick={onGetSale}
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
                  percent: 1,
                },
                {
                  percent: 1,
                },
                {
                  percent: 5,
                },
                {
                  percent: 1,
                },
                {
                  percent: 1,
                },
              ]}
            ></TableSkeleton>
          ) : (
            <SaleReportTable
              report={data}
              data={
                data == undefined || data == null
                  ? []
                  : (data!.details as SaleReportDetail[])
              }
            />
          )}
        </Card>
      </div>
    </div>
  );
};

export default SaleReport;
