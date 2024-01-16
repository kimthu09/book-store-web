"use client";

import * as React from "react";
import { Check, ChevronsUpDown } from "lucide-react";

import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import { Command, CommandGroup, CommandItem } from "@/components/ui/command";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import DashboardChart from "./dashboard-chart";
import { Card } from "../ui/card";
import { typeCharts } from "@/constants";
import { Skeleton } from "../ui/skeleton";


const DashboardChartContainer = (props: any) => {
  const {
    price,
    profit,
    timeFrom,
    timeTo,
    chartType,
    setChartType,
    isLoading,
  } = props;
  const [open, setOpen] = React.useState(false);
  const [value, setValue] = React.useState(chartType);

  return (
    <Card className="p-4 flex-[2] w-full flex flex-col gap-4">
      <div className="flex flex-row gap-4">
        <div className="flex-1"></div>
        {isLoading ? (
          <Skeleton className="h-8 w-32" />
        ) : (
          <Popover open={open} onOpenChange={setOpen}>
            <PopoverTrigger asChild>
              <Button
                variant="outline"
                role="combobox"
                aria-expanded={open}
                className="w-[200px] justify-between"
              >
                {value
                  ? typeCharts.find((type) => type.value === value)?.label
                  : "Chọn loại biểu đồ"}
                <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
              </Button>
            </PopoverTrigger>
            <PopoverContent className="w-[200px] p-0">
              <Command>
                <CommandGroup>
                  {typeCharts.map((type) => (
                    <CommandItem
                      key={type.value}
                      value={type.value}
                      onSelect={(currentValue) => {
                        setValue(currentValue === value ? "" : currentValue);
                        setOpen(false);
                      }}
                    >
                      <Check
                        className={cn(
                          "mr-2 h-4 w-4",
                          value === type.value ? "opacity-100" : "opacity-0"
                        )}
                      />
                      {type.label}
                    </CommandItem>
                  ))}
                </CommandGroup>
              </Command>
            </PopoverContent>
          </Popover>
        )}

        {isLoading ? (
          <Skeleton className="h-8 w-16" />
        ) : (
          <Button onClick={() => setChartType(value)} className="px-5">
            Xem
          </Button>
        )}
      </div>
      <div className="flex-1 overflow-x-auto min-w-full max-w-[20vw] min-h-[28rem]">
        {isLoading ? (
          <Skeleton className="w-full aspect-square"/>
        ) : (
          <DashboardChart
            type={chartType}
            price={price}
            profit={profit}
            timeFrom={timeFrom}
            timeTo={timeTo}
          />
        )}
      </div>
    </Card>
  );
};

export default DashboardChartContainer;
