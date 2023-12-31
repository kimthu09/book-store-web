import { StatusListProps } from "@/types";
import { useState } from "react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "./ui/dropdown-menu";
import { Command, CommandGroup, CommandItem } from "./ui/command";
import { LuCheck, LuChevronsUpDown } from "react-icons/lu";
import { Button } from "./ui/button";
import { cn } from "@/lib/utils";

const StatusList = ({ status, setStatus, display }: StatusListProps) => {
  const [open, setOpen] = useState(false);
  return (
    <DropdownMenu open={open} onOpenChange={setOpen}>
      <DropdownMenuTrigger asChild>
        <Button
          variant="outline"
          role="combobox"
          aria-expanded={open}
          className="justify-between w-full min-w-0 p-2"
        >
          {status === true
            ? display.trueText
            : status === false
            ? display.falseText
            : "Chọn trạng thái"}
          <LuChevronsUpDown className="ml-1 h-4 w-4 shrink-0 opacity-50" />
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent className="DropdownMenuContent p-0">
        <Command>
          <CommandGroup>
            <CommandItem
              className="p-1"
              value={display.trueText}
              onSelect={() => {
                setStatus(true);
                setOpen(false);
              }}
            >
              <LuCheck
                className={cn(
                  "mr-1 h-4 w-4",
                  status ? "opacity-100" : "opacity-0"
                )}
              />
              {display.trueText}
            </CommandItem>
            <CommandItem
              className="p-1"
              value={display.falseText}
              onSelect={() => {
                setStatus(false);
                setOpen(false);
              }}
            >
              <LuCheck
                className={cn(
                  "mr-1 h-4 w-4",
                  status === false ? "opacity-100" : "opacity-0"
                )}
              />
              {display.falseText}
            </CommandItem>
          </CommandGroup>
        </Command>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

export default StatusList;
