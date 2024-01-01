"use client"
import { Card } from '@/components/ui/card';
import ReportHeader from '@/components/report/ReportHeader';
import { SaleReportDetail } from '@/types';
import getReport from '@/lib/report/getReport';
import { useState } from 'react';
import { SaleReportTable } from '@/components/report/SaleReportTable';

const SaleReport = () => {
    const [data, setData] = useState<SaleReportDetail[]>([])
    const [total, setTotal] = useState<number>(0)

    const onGetDebt = async ({
        timeFrom,
        timeTo,
    }: {
        timeFrom: number,
        timeTo: number
    }) => {
        const report = await getReport({
            timeFrom: timeFrom,
            timeTo: timeTo,
            type: "sale"
        });
        setTotal(report.total)
        setData(report.data.details)
    }

    return (
        <div>
            <div>
                <ReportHeader
                    title="Báo cáo doanh thu"
                    firstAction="Xem báo cáo"
                    secondAction="Tải excel"
                    onClick={onGetDebt} />
            </div>

            <div>
                <Card className='p-[10px] my-[22px]'>
                    <SaleReportTable data={data} />
                    <div className='flex justify-between pt-[15px]'>
                        <p className='text-xl font-bold'>Tổng cộng</p>
                        <p className='text-xl font-bold'>{`${(new Intl.NumberFormat("vi-VN", {
                            style: "currency",
                            currency: "VND",
                        }).format(total))}`}</p>
                    </div>
                </Card>
            </div>
        </div>
    );
}

export default SaleReport;