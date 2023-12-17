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

const StaffList = ({ staff, setStaff }: StaffListProps) => {
  const [openRole, setOpenRole] = useState(false);
  const { staffs, isLoading, isError } = getAllStaff({ limit: 1000 });

  if (isError) return <div>Failed to load</div>;
  if (!staffs) {
    <Loading />;
  } else
    return (
      <DropdownMenu open={openRole} onOpenChange={setOpenRole}>
        <DropdownMenuTrigger asChild>
          <Button
            variant="outline"
            role="combobox"
            aria-expanded={openRole}
            className="justify-between w-[160px] min-w-0"
          >
            {staff
              ? staffs.find((item) => item.name === staff)?.name
              : "Chọn người tạo"}
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
                    setStaff(item.name);
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
