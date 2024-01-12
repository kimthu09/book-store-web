"use client";
import { addDays, format } from "date-fns";
import { Button } from "../ui/button";
import { DateRange } from "react-day-picker";
import React from "react";
import { cn } from "@/lib/utils";
import { Popover, PopoverContent, PopoverTrigger } from "../ui/popover";
import { CalendarIcon } from "@radix-ui/react-icons";
import { Calendar } from "../ui/calendar";

const ReportHeader = (props: any) => {
  const { title, firstAction, secondAction, onClick, onExport } = props;

  const [date, setDate] = React.useState<DateRange | undefined>({
    from: addDays(new Date(), -30),
    to: new Date(),
  });

  return (
    <div className="flex justify-between flex-row">
      <h1 className="flex-1">{title}</h1>
      <div className="justify-end gap-2 flex xl:max-w-[550px] max-w-[400px] xl:flex-row flex-col ">
        <div className={cn("grid gap-2 flex-1")}>
          <Popover>
            <PopoverTrigger asChild>
              <Button
                id="date"
                variant={"outline"}
                className={cn(
                  "bg-white border border-primary hover:text-primary justify-start items-center text-left font-normal",
                  !date && "text-muted-foreground"
                )}
              >
                <CalendarIcon className="mr-2 h-4 w-4" />
                {date?.from ? (
                  date.to ? (
                    <>
                      {format(date.from, "dd/MM/yyyy")} -{" "}
                      {format(date.to, "dd/MM/yyyy")}
                    </>
                  ) : (
                    format(date.from, "dd/MM/yyyy")
                  )
                ) : (
                  <span>Pick a date</span>
                )}
              </Button>
            </PopoverTrigger>
            <PopoverContent className="w-auto p-0" align="start">
              <Calendar
                initialFocus
                mode="range"
                defaultMonth={date?.from}
                selected={date}
                onSelect={setDate}
                numberOfMonths={2}
              />
            </PopoverContent>
          </Popover>
        </div>
        {/* 1701388800 */}
        <div className="flex flex-row gap-2">
          <Button
            onClick={() =>
              onClick({
                timeFrom: date?.from
                  ? new Date(date.from.setHours(0, 0, 0, 0)).getTime() / 1000
                  : 0,
                timeTo: date?.to
                  ? new Date(date.to.setHours(23, 59, 59, 999)).getTime() / 1000
                  : 0,
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

export default ReportHeader;
