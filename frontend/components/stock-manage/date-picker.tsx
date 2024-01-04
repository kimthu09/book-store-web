"use client";

import * as React from "react";
import { format } from "date-fns";
import { Calendar as CalendarIcon } from "lucide-react";

import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import { Calendar } from "@/components/ui/datepicker";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";

export type DatePickerProps = {
  handleDateSelected: (value: Date | undefined) => void;
  date: Date;
};
export function FilterDatePicker({
  handleDateSelected,
  date,
}: DatePickerProps) {
  // const [date, setDate] = React.useState<Date>();
  // const handleDate = (value: Date | undefined) => {
  //   setDate(value);
  //   handleDateSelected(value);
  // };
  return (
    <Popover>
      <PopoverTrigger asChild>
        <Button
          variant={"outline"}
          className={cn(
            "justify-start flex-1 text-left font-normal",
            !date && "text-muted-foreground"
          )}
        >
          <CalendarIcon className="mr-2 h-4 w-4" />
          {date ? (
            date.toLocaleDateString("vi-VN") === "Invalid Date" ? (
              <span>Chọn ngày</span>
            ) : (
              date.toLocaleDateString("vi-VN")
            )
          ) : (
            <span>Chọn ngày</span>
          )}
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-auto p-0">
        <Calendar
          mode="single"
          selected={date}
          onSelect={handleDateSelected}
          initialFocus
        />
      </PopoverContent>
    </Popover>
  );
}
