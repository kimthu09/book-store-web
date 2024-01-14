import { RoleListProps } from "@/types";
import { useState } from "react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "../ui/dropdown-menu";
import { Command, CommandGroup, CommandItem } from "../ui/command";
import { LuCheck, LuChevronsUpDown } from "react-icons/lu";
import { Button } from "../ui/button";
import { cn } from "@/lib/utils";
import getAllRole from "@/lib/staff/getAllRole";
import Loading from "../loading";
import ListSkeleton from "../skeleton/list_skeleton";

const RoleList = ({ role, setRole }: RoleListProps) => {
  const [openRole, setOpenRole] = useState(false);
  const { roles, isLoading, isError } = getAllRole();

  if (isError) return <div>Failed to load</div>;
  if (!roles) {
    <ListSkeleton numberRow={5} />;
  } else
    return (
      <DropdownMenu open={openRole} onOpenChange={setOpenRole}>
        <DropdownMenuTrigger asChild>
          <Button
            variant="outline"
            role="combobox"
            aria-expanded={openRole}
            className="justify-between w-full"
          >
            {role
              ? roles.find((item) => item.id === role)?.name
              : "Chọn phân quyền"}
            <LuChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent className=" DropdownMenuContent">
          <Command>
            <CommandGroup>
              {roles.map((item) => (
                <CommandItem
                  value={item.name}
                  key={item.id}
                  onSelect={() => {
                    setRole(item.id);
                    setOpenRole(false);
                  }}
                >
                  <LuCheck
                    className={cn(
                      "mr-2 h-4 w-4",
                      item.id === role ? "opacity-100" : "opacity-0"
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

export default RoleList;
