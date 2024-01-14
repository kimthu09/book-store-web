import { StaffListProps } from "@/types";
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
import Loading from "./loading";
import getAllStaff from "@/lib/staff/getAllStaffClient";
import DropdownSkeleton from "./skeleton/dropdown-skeleton";

const StaffList = ({ staff, setStaff }: StaffListProps) => {
  const [openRole, setOpenRole] = useState(false);
  const { staffs, isLoading, isError } = getAllStaff();

  if (isError) return <div>Failed to load</div>;
  if (!staffs) {
    return <DropdownSkeleton />;
  } else
    return (
      <DropdownMenu open={openRole} onOpenChange={setOpenRole}>
        <DropdownMenuTrigger asChild>
          <Button
            variant="outline"
            role="combobox"
            aria-expanded={openRole}
            className="justify-between w-full min-w-0"
          >
            {staff
              ? staffs.find((item) => item.id === staff)?.name
              : "Chọn nhân viên"}
            <LuChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent className="DropdownMenuContent">
          <Command>
            <CommandGroup>
              {staffs.map((item) => (
                <CommandItem
                  value={item.name}
                  key={item.id}
                  onSelect={() => {
                    setStaff(item.id);
                    setOpenRole(false);
                  }}
                >
                  <LuCheck
                    className={cn(
                      "mr-2 h-4 w-4",
                      item.id === staff ? "opacity-100" : "opacity-0"
                    )}
                  />
                  {item.name}
                </CommandItem>
              ))}
            </CommandGroup>
          </Command>
        </DropdownMenuContent>
      </DropdownMenu>
    );
};

export default StaffList;
