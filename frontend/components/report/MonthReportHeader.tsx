"use client";
import { cn } from "@/lib/utils";
import { DatePicker } from "antd";
import { RangePickerProps } from "antd/es/date-picker";
import dayjs from "dayjs";
import { SetStateAction, useState } from "react";
import { Button } from "../ui/button";

const MonthReportHeader = (props: any) => {
  const { title, firstAction, secondAction, onClick, onExport } = props;
  
  const monthFormat = "MM/YYYY";

  const disabledDate: RangePickerProps["disabledDate"] = (
    current: dayjs.Dayjs
  ) => {
    return current > dayjs().endOf("day") || dayjs("2023-01-01") > current;
  };

  const [from, setFrom] = useState(dayjs().startOf('month'));
  // const [to, setTo] = useState(dayjs('2024-01-01'))
  return (
    <div className="flex justify-between flex-row align-middle">
      <h1 className="flex-1">{title}</h1>
      <div className="justify-end align-middle flex gap-2 xl:max-w-[550px] max-w-[400px] flex-row">
        <div className={cn("grid gap-2 flex-1 flex")}>
          <DatePicker
            className={cn(
              "bg-white border border-primary hover:text-primary justify-start items-center text-left font-normal",
              !from && "text-muted-foreground flex-1"
            )}
            picker="month"
            format={monthFormat}
            placeholder={"Tháng"}
            value={from}
            onChange={(
              values: SetStateAction<dayjs.Dayjs>,
              _dateString: string
            ) => {
              setFrom(values!);
            }}
            disabledDate={disabledDate}
          />
        </div>
        <div className="flex flex-row gap-2">
          <Button
            onClick={() =>
              onClick({
                timeFrom: from.startOf("month").valueOf() / 1000,
                timeTo: from.endOf("month").valueOf() / 1000,
              })
            }
            className="px-5"
          >
            {firstAction}
          </Button>
          <Button
            onClick={onExport}
            type="button"
            className="pl-[20px] pr-[20px] bg-white border border-primary text-primary hover:text-primary"
            variant={"outline"}
          >
            {secondAction}
          </Button>
        </div>
      </div>
    </div>
  );
};

export default MonthReportHeader;
