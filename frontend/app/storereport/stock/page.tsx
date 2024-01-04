"use client"
import { Card } from '@/components/ui/card';
import ReportHeader from '@/components/report/ReportHeader';
import { StockReportDetail } from '@/types';
import getReport from '@/lib/report/getReport';
import { useState } from 'react';
import { StockReportTable } from '@/components/report/StockReportTable';
import MonthReportHeader from '@/components/report/MonthReportHeader';

const SaleReport = () => {
    const [data, setData] = useState<StockReportDetail[]>([])

    const onGetStock = async ({
        timeFrom,
        timeTo,
    }: {
        timeFrom: number,
        timeTo: number
    }) => {
        const stockReport = await getReport({
            timeFrom: timeFrom,
            timeTo: timeTo,
            type: "stock"
        });

        setData(stockReport.data.details)
    }

    return (
        <div>
            <div>
                <MonthReportHeader
                    title="Báo cáo tồn"
                    firstAction="Xem báo cáo"
                    secondAction="Tải excel"
                    onClick={onGetStock} />
            </div>

            <div>
                <Card className='p-[10px] my-[22px]'>
                    <StockReportTable data={data} />
                </Card>
            </div>
        </div>
    );
}

export default SaleReport;