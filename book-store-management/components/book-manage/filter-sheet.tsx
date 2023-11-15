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
import { cn } from "@/lib/utils";
import RangeSlider from "./range-slider";
import { FiFilter } from "react-icons/fi";
import CategoryList from "../category-list";
import { statuses } from "@/constants";

const FilterSheet = () => {
  const [category, setCategory] = useState("");

  const [openStatus, setOpenStatus] = useState(false);
  const [status, setStatus] = useState("");

  return (
    <Sheet>
      <SheetTrigger asChild>
        <Button variant="outline" role="combobox">
          Lọc danh sách
          <FiFilter className="ml-4 h-4 w-4 shrink-0 opacity-50" />
        </Button>
      </SheetTrigger>
      <SheetContent
        side={"right"}
        className="w-[480px] sm:w-[540px] m-auto bg-white"
      >
        <SheetHeader>
          <SheetTitle>Lọc danh sách sách</SheetTitle>
        </SheetHeader>
        <form action="submit"></form>
        <div className="grid gap-4 py-4">
          <div className="grid grid-cols-3 items-center gap-6">
            <Label htmlFor="id" className="text-right col-span-1">
              Mã sách
            </Label>
            <Input id="id" className="col-span-2" />
          </div>
          <div className="grid grid-cols-3 items-center gap-6">
            <Label htmlFor="name" className="text-right col-span-1">
              Tên sách
            </Label>
            <Input id="name" className="col-span-2" />
          </div>
          <div className="flex items-center">
            <Label htmlFor="cate" className="text-right ml-auto p-0">
              Danh mục
            </Label>
            <div className="ml-auto w-[13.5rem]">
              <CategoryList category={category} setCategory={setCategory} />
            </div>
          </div>
          <div className="flex items-center">
            <Label htmlFor="sta" className="text-right ml-auto p-0">
              Trạng thái
            </Label>
            <div className="ml-auto">
              <DropdownMenu open={openStatus} onOpenChange={setOpenStatus}>
                <DropdownMenuTrigger asChild>
                  <Button
                    id="sta"
                    variant="outline"
                    role="combobox"
                    aria-expanded={openStatus}
                    className="justify-between w-[13.5rem]"
                  >
                    {status
                      ? statuses.find((item) => item.label === status)?.label
                      : "Chọn trạng thái"}
                    <LuChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
                  </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent className="w-[13.5rem] p-0">
                  <Command>
                    <CommandInput placeholder="Tìm điều kiện lọc" />
                    <CommandEmpty>Không tìm thấy điều kiện lọc</CommandEmpty>
                    <CommandGroup>
                      {statuses.map((item) => (
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

export default FilterSheet;
