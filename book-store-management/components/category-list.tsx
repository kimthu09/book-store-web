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
import { CategoryListProps } from "@/types";
import { categories } from "@/constants";

const CategoryList = ({
  category,
  setCategory,
  canAdd,
  readonly,
}: CategoryListProps) => {
  const [openCategory, setOpenCategory] = useState(false);
  const [newCategory, setNewCategory] = useState("");
  return (
    <DropdownMenu open={openCategory} onOpenChange={setOpenCategory}>
      <DropdownMenuTrigger asChild>
        <Button
          id="cateList"
          variant="outline"
          role="combobox"
          aria-expanded={openCategory}
          className={`justify-between w-full ${
            readonly ? "pointer-events-none" : ""
          }`}
        >
          {category}

          <LuChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent className=" DropdownMenuContent">
        <Command>
          <CommandInput
            autoFocus
            placeholder="Tìm điều kiện lọc"
            onValueChange={(str) => setNewCategory(str)}
          />
          <CommandEmpty className="py-2">
            {canAdd ? (
              <div className="flex">
                <Button variant="ghost" className="flex-1">
                  <div className="text-left flex-1 text-primary flex items-center gap-2">
                    <AiFillPlusCircle size={20} />
                    Thêm
                    {" " + newCategory}
                  </div>
                </Button>
              </div>
            ) : (
              <div className="text-sm">Không tìm thấy điều kiện lọc.</div>
            )}
          </CommandEmpty>
          <CommandGroup className="overflow-y-auto">
            {categories.map((item) => (
              <CommandItem
                value={item.name}
                key={item.id}
                onSelect={() => {
                  setCategory(item.name);
                  setOpenCategory(false);
                }}
              >
                <LuCheck
                  className={cn(
                    "mr-2 h-4 w-4",
                    item.name === category ? "opacity-100" : "opacity-0"
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

export default CategoryList;
