"use client"
import { cn } from '@/lib/utils';
import { DatePicker } from 'antd'
import { RangePickerProps } from 'antd/es/date-picker';
import dayjs from 'dayjs';
import { useState } from 'react';
import { Button } from '../ui/button';

const MonthReportHeader = (props: any) => {
    const { title, firstAction, secondAction, onClick } = props

    const { RangePicker } = DatePicker
    const monthFormat = 'MM/YYYY';

    const disabledDate: RangePickerProps['disabledDate'] = (current) => {
        return current > dayjs().endOf('day')
            || dayjs('2023-01-01') > current;
    };

    const [from, setFrom] = useState(dayjs('2023-12-1'))
    const [to, setTo] = useState(dayjs('2024-01-01'))
    return (
        <div className='flex justify-between items-center'>
            <h1>{title}</h1>
            <div className='flex-[0.9] justify-end flex gap-[15px] items-center'>
                <div className={cn("grid gap-2 ")}>
                    <div>
                        <RangePicker
                            className='w-[300px] bg-white border border-primary hover:text-primary justify-start items-center text-left font-normal'
                            picker='month'
                            format={monthFormat}
                            placeholder={["Từ", "Đến"]}
                            value={[from, to]}
                            onChange={(values, string) => {
                                setFrom(values![0]!)
                                setTo(values![1]!)
                            }}
                            disabledDate={disabledDate}
                        />
                    </div>
                </div>
                {/* 1701388800 */}
                <Button onClick={() => onClick({
                    timeFrom: from.valueOf() / 1000,
                    timeTo: to.valueOf() / 1000
                })}
                    className='px-5'>{firstAction}</Button>
                <Button
                    type="button"
                    className='pl-[20px] pr-[20px] bg-white border border-primary text-primary hover:text-primary' variant={"outline"}
                >{secondAction}
                </Button>
            </div>
        </div>
    );
}

export default MonthReportHeader;