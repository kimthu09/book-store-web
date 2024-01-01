"use client"
import { addDays, format } from "date-fns";
import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { DateRange } from "react-day-picker";
import React from "react";
import { cn } from "@/lib/utils";
import { Popover, PopoverContent, PopoverTrigger } from "../ui/popover";
import { CalendarIcon } from "@radix-ui/react-icons";
import { Calendar } from "../ui/calendar";

const ReportHeader = (props: any) => {
    const { title, firstAction, secondAction, onClick } = props

    const [date, setDate] = React.useState<DateRange | undefined>({
        from: addDays(new Date(), -30),
        to: new Date(),
    })

    return (
        <div className='flex justify-between'>
            <h1>{title}</h1>
            <div className='flex-[0.9] justify-end flex gap-[15px]'>
                <div className={cn("grid gap-2 ")}>
                    <Popover>
                        <PopoverTrigger asChild>
                            <Button
                                id="date"
                                variant={"outline"}
                                className={cn(
                                    "w-[300px] bg-white border border-primary hover:text-primary justify-start items-center text-left font-normal",
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
                <Button onClick={() => onClick({
                    timeFrom: (date?.from?.getTime() ?? 0) / 1000,
                    timeTo: (date?.to?.getTime() ?? 0) / 1000
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

export default ReportHeader;