"use client";
import { format } from "date-fns";
import { Button } from "../ui/button";
import { DateRange } from "react-day-picker";
import React, { useEffect } from "react";
import { cn } from "@/lib/utils";
import { Popover, PopoverContent, PopoverTrigger } from "../ui/popover";
import { CalendarIcon } from "@radix-ui/react-icons";
import { Calendar } from "../ui/calendar";

const DashboardHeader = (props: any) => {
  const { onClick } = props;

  const [date, setDate] = React.useState<DateRange | undefined>({
    from: new Date(),
    to: new Date(),
  });
  const [selectedDateRangeType, setSelectedDateRangeType] =
    React.useState<string>("Today");

  useEffect(() => {
    onClick(getTimeRangeParams(date)); // Call the fetchData function on component mount
  }, []);

  const getTimeRangeParams = (range: DateRange | undefined) => {
    return {
      timeFrom: range?.from
        ? new Date(range.from.setHours(0, 0, 0, 0)).getTime() / 1000
        : 0,
      timeTo: range?.to
        ? new Date(range.to.setHours(23, 59, 59, 999)).getTime() / 1000
        : 0,
    };
  };
  const handleTodayClick = () => {
    const today = {
      from: new Date(),
      to: new Date(),
    };
    setDate(today);
    setSelectedDateRangeType("Today");
    onClick(getTimeRangeParams(today));
  };

  const handleThisMonthClick = () => {
    const startOfMonth = new Date();
    startOfMonth.setDate(1); // Set to the first day of the month
    const thisMonth = {
      from: startOfMonth,
      to: new Date(),
    };
    setDate(thisMonth);
    setSelectedDateRangeType("ThisMonth");
    onClick(getTimeRangeParams(thisMonth));
  };

  const handleLastMonthClick = () => {
    const startOfLastMonth = new Date();
    startOfLastMonth.setDate(1);
    startOfLastMonth.setMonth(startOfLastMonth.getMonth() - 1);
    const endOfLastMonth = new Date(
      startOfLastMonth.getFullYear(),
      startOfLastMonth.getMonth() + 1,
      0
    );

    const lastMonth = {
      from: startOfLastMonth,
      to: endOfLastMonth,
    };
    setDate(lastMonth);
    setSelectedDateRangeType("LastMonth");
    onClick(getTimeRangeParams(lastMonth));
  };

  const handlePickDateClick = () => {
    setSelectedDateRangeType("PickDate");
  };

  return (
    <div className="flex justify-between flex-row">
      <div className="gap-2 flex flex-row flex-1 items-center">
        <div className="flex gap-2 flex-wrap rounded-full bg-white p-1">
          <Button
            variant={"ghost"}
            onClick={handleTodayClick}
            className={`bg-white whitespace-nowrap rounded-full ${
              selectedDateRangeType === "Today"
                ? "bg-blue-200 hover:bg-blue-200/90 text-primary hover:text-primary"
                : ""
            }`}
          >
            Hôm nay
          </Button>
          <Button
            variant={"ghost"}
            onClick={handleThisMonthClick}
            className={`bg-white whitespace-nowrap rounded-full ${
              selectedDateRangeType === "ThisMonth"
                ? "bg-blue-200 hover:bg-blue-200/90 text-primary hover:text-primary"
                : ""
            }`}
          >
            Tháng này
          </Button>
          <Button
            variant={"ghost"}
            onClick={handleLastMonthClick}
            className={`bg-white whitespace-nowrap rounded-full ${
              selectedDateRangeType === "LastMonth"
                ? "bg-blue-200 hover:bg-blue-200/90 text-primary hover:text-primary"
                : ""
            }`}
          >
            Tháng trước
          </Button>
          <Button
            variant={"ghost"}
            onClick={handlePickDateClick}
            className={`bg-white whitespace-nowrap rounded-full ${
              selectedDateRangeType === "PickDate"
                ? "bg-blue-200 hover:bg-blue-200/90 text-primary hover:text-primary"
                : ""
            }`}
          >
            Chọn thời điểm
          </Button>
        </div>
        {selectedDateRangeType === "PickDate" ? (
          <div className="flex flex-row gap-2">
            <Popover>
              <PopoverTrigger asChild>
                <Button
                  id="date"
                  variant={"outline"}
                  className={cn(
                    `bg-white whitespace-nowrap border border-primary hover:text-primary justify-start items-center text-left font-normal",
                  !date && "text-muted-foreground`
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
            <div className="flex flex-row gap-2">
              <Button
                onClick={() => {
                  onClick({
                    timeFrom: date?.from
                      ? new Date(date.from.setHours(0, 0, 0, 0)).getTime() /
                        1000
                      : 0,
                    timeTo: date?.to
                      ? new Date(date.to.setHours(23, 59, 59, 999)).getTime() /
                        1000
                      : 0,
                  });
                }}
                className="px-5"
              >
                Xem
              </Button>
            </div>
          </div>
        ):(<div></div>)}
      </div>
    </div>
  );
};

export default DashboardHeader;
