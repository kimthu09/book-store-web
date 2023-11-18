import { useState } from "react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "./ui/dropdown-menu";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
} from "./ui/command";
import { LuCheck, LuChevronsUpDown } from "react-icons/lu";
import { Button } from "./ui/button";
import { AiFillPlusCircle } from "react-icons/ai";
import { cn } from "@/lib/utils";
import { SupplierListProps } from "@/types";
import { suppliers } from "@/constants";

const SupplierList = ({ supplier, setSupplier, canAdd }: SupplierListProps) => {
  const [openSupplier, setOpenSupplier] = useState(false);
  const [newSupplier, setNewSupplier] = useState("");
  return (
    <DropdownMenu open={openSupplier} onOpenChange={setOpenSupplier}>
      <DropdownMenuTrigger asChild>
        <Button
          id="cateList"
          variant="outline"
          role="combobox"
          aria-expanded={openSupplier}
          className="justify-between w-full"
        >
          {supplier
            ? suppliers.find((item) => item.name === supplier)?.name
            : "Chọn nhà cung cấp"}
          <LuChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent className=" DropdownMenuContent">
        <Command>
          <CommandInput
            autoFocus
            placeholder="Tìm điều kiện lọc"
            onValueChange={(str) => setNewSupplier(str)}
          />
          <CommandEmpty className="py-2">
            {canAdd ? (
              <div className="flex">
                <Button variant="ghost" className="flex-1">
                  <div className="text-left flex-1 text-primary flex items-center gap-2">
                    <AiFillPlusCircle size={20} />
                    Thêm
                    {" " + newSupplier}
                  </div>
                </Button>
              </div>
            ) : (
              <div className="text-sm">Không tìm thấy điều kiện lọc.</div>
            )}
          </CommandEmpty>
          <CommandGroup className="overflow-y-auto">
            {suppliers.map((item) => (
              <CommandItem
                value={item.name}
                key={item.id}
                onSelect={() => {
                  setSupplier(item.name);
                  setOpenSupplier(false);
                }}
              >
                <LuCheck
                  className={cn(
                    "mr-2 h-4 w-4",
                    item.name === supplier ? "opacity-100" : "opacity-0"
                  )}
                />
                {item.name}
              </CommandItem>
            ))}
            <CommandItem
              key={""}
              onSelect={() => {
                setSupplier("");
                setOpenSupplier(false);
              }}
            >
              <LuCheck
                className={cn(
                  "mr-2 h-4 w-4",
                  "" === supplier ? "opacity-100" : "opacity-0"
                )}
              />
              {"Chọn nhà cung cấp"}
            </CommandItem>
          </CommandGroup>
        </Command>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

export default SupplierList;
