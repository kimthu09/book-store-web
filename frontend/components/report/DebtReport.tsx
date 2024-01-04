"use client"
import { Card } from '@/components/ui/card';
import ReportHeader from '@/components/report/ReportHeader';
import { DebtTable } from '@/components/supplier-manage/debt-table';
import { DebtReportTable } from '@/components/report/DebtReportTable';
import { DebtReportDetail } from '@/types';
import getReport from '@/lib/report/getReport';
import { useState } from 'react';
import MonthReportHeader from '@/components/report/MonthReportHeader';


const DebtReport = () => {
    const [data, setData] = useState<DebtReportDetail[]>([])

    const onGetDebt = async ({
        timeFrom,
        timeTo,
    }: {
        timeFrom: number,
        timeTo: number
    }) => {
        const debtReport = await getReport({
            timeFrom: timeFrom,
            timeTo: timeTo,
            type: "debt"
        });
        console.log(timeFrom, timeTo)
        console.log(debtReport)
        setData(debtReport.data.details)
    }

    return (
        <div>
            <div>
                {/* <ReportHeader
                    title="Báo cáo nợ"
                    firstAction="Xem báo cáo"
                    secondAction="Tải excel"
                    onClick={onGetDebt} /> */}
                <MonthReportHeader
                    title="Báo cáo nợ"
                    firstAction="Xem báo cáo"
                    secondAction="Tải excel"
                    onClick={onGetDebt} />
            </div>

            <div>
                <Card className='p-[10px] my-[22px]'>
                    <DebtReportTable data={data} />
                </Card>
            </div>
        </div>
    );
}

export default DebtReport;

// const testData: DebtReportDetail[] = [
//     {
//         debt: -40000,
//         final: 80000,
//         initial: 100000,
//         pay: 20000,
//         supplier: {
//             id: "123",
//             name: "Nguyễn Văn A",
//             phone: "0123456789"
//         }
//     },
//     {
//         debt: -40000,
//         final: 80000,
//         initial: 200000,
//         pay: 20000,
//         supplier: {
//             id: "1234",
//             name: "Nguyễn Văn B",
//             phone: "0123456789"
//         }
//     },
//     {
//         debt: -40000,
//         final: 80000,
//         initial: 500000,
//         pay: 20000,
//         supplier: {
//             id: "1235",
//             name: "Nguyễn Văn C",
//             phone: "0123456789"
//         }
//     },
//     {
//         debt: -40000,
//         final: 80000,
//         initial: 600000,
//         pay: 20000,
//         supplier: {
//             id: "126",
//             name: "Nguyễn Văn D",
//             phone: "0123456789"
//         }
//     },
//     {
//         debt: -40000,
//         final: 80000,
//         initial: 300000,
//         pay: 20000,
//         supplier: {
//             id: "128",
//             name: "Nguyễn Văn E",
//             phone: "0123456789"
//         }
//     },
// ]