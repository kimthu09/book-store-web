"use client";

import {
  Sheet,
  SheetClose,
  SheetContent,
  SheetFooter,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";
import { Button } from "../ui/button";
import { Label } from "../ui/label";
import { Input } from "../ui/input";
import { useState } from "react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "../ui/dropdown-menu";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
} from "../ui/command";
import { LuCheck, LuChevronsUpDown } from "react-icons/lu";
import { noteStatus } from "@/constants";
import { cn } from "@/lib/utils";
import { FiFilter } from "react-icons/fi";
import RangeSlider from "../range-slider";
import { DateRange } from "react-day-picker";
import { addDays, format } from "date-fns";
import { Popover, PopoverContent, PopoverTrigger } from "../ui/popover";
import { Calendar } from "../ui/calendar";
import { CalendarIcon } from "@radix-ui/react-icons";

const ImportFilter = () => {
  const [openStatus, setOpenStatus] = useState(false);
  const [status, setStatus] = useState("");

  const [date, setDate] = useState<DateRange | undefined>({
    from: new Date(2022, 0, 20),
    to: addDays(new Date(2022, 0, 20), 20),
  });
  return (
    <Sheet>
      <SheetTrigger asChild>
        <Button variant="outline" role="combobox">
          Lọc danh sách
          <FiFilter className="ml-4 h-4 w-4 shrink-0 opacity-50" />
        </Button>
      </SheetTrigger>
      <SheetContent side={"right"} className="w-[480px] m-auto bg-white">
        <SheetHeader>
          <SheetTitle>Lọc danh sách phiếu nhập</SheetTitle>
        </SheetHeader>
        <form action="submit"></form>
        <div className="grid gap-4 py-4">
          <div className="grid grid-cols-3 items-center gap-4">
            <Label htmlFor="id" className="text-right col-span-1">
              Mã phiếu
            </Label>
            <Input id="id" className="col-span-2" />
          </div>
          <div className="grid grid-cols-3 items-center gap-4">
            <Label className="text-right col-span-1">Người tạo</Label>
            <Input className="col-span-2" />
          </div>
          <div className="grid grid-cols-3 items-center gap-4">
            <Label htmlFor="sta" className="text-right ml-auto p-0">
              Trạng thái
            </Label>
            <div className="col-span-2">
              <DropdownMenu open={openStatus} onOpenChange={setOpenStatus}>
                <DropdownMenuTrigger asChild>
                  <Button
                    id="sta"
                    variant="outline"
                    role="combobox"
                    aria-expanded={openStatus}
                    className="w-full flex justify-between"
                  >
                    {status
                      ? noteStatus.find((item) => item.label === status)?.label
                      : "Chọn trạng thái"}
                    {/* <FiFilter className="ml-2 h-4 w-4 shrink-0 opacity-50" /> */}
                    <LuChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
                  </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent className="p-0 w-full">
                  <Command>
                    <CommandInput placeholder="Tìm điều kiện lọc" />
                    <CommandEmpty>Không tìm thấy điều kiện lọc</CommandEmpty>
                    <CommandGroup>
                      {noteStatus.map((item) => (
                        <CommandItem
                          value={item.label}
                          key={item.label}
                          onSelect={() => {
                            setStatus(item.label);
                            setOpenStatus(false);
                          }}
                        >
                          <LuCheck
                            className={cn(
                              "mr-2 h-4 w-4",
                              item.label === status
                                ? "opacity-100"
                                : "opacity-0"
                            )}
                          />
                          {item.label === "" ? "Chọn trạng thái" : item.label}
                        </CommandItem>
                      ))}
                      <CommandItem
                        key={""}
                        onSelect={() => {
                          setStatus("");
                          setOpenStatus(false);
                        }}
                      >
                        <LuCheck
                          className={cn(
                            "mr-2 h-4 w-4",
                            "" === status ? "opacity-100" : "opacity-0"
                          )}
                        />
                        {"Chọn trạng thái"}
                      </CommandItem>
                    </CommandGroup>
                  </Command>
                </DropdownMenuContent>
              </DropdownMenu>
            </div>
          </div>

          {/* Date picker */}
          <div className="flex gap-4 flex-col items-center mt-4">
            <Popover>
              <PopoverTrigger asChild>
                <Button
                  id="date"
                  variant={"outline"}
                  className={cn(
                    "w-[300px] justify-start text-left font-normal",
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
                      format(date.from, "LLL dd, y")
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
            <Label className="text-right">Khoảng thời gian</Label>
          </div>

          {/* Price range */}
          <div className="mt-4">
            <div className="flex justify-center">
              <RangeSlider
                initialMin={0}
                initialMax={100}
                min={0}
                max={500}
                step={1}
                priceCap={1}
              />
            </div>
            <p className="mt-1 text-center text-sm">
              Đơn vị tính là ngàn đồng.
            </p>
          </div>
        </div>
        <SheetFooter>
          <SheetClose asChild>
            <Button type="submit">Lọc</Button>
          </SheetClose>
        </SheetFooter>
      </SheetContent>
    </Sheet>
  );
};

export default ImportFilter;
