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
import { Button } from "../ui/button";
import { AiFillPlusCircle } from "react-icons/ai";
import { cn } from "@/lib/utils";
import { CategoryListProps } from "@/types";
import { Checkbox } from "../ui/checkbox";
import getAllCategory from "@/lib/book/getAllCategory";
import Loading from "../loading";
import CreateCategory from "./create-category";
import { FaPlus } from "react-icons/fa";

const CategoryList = ({
  checkedCategory,
  onCheckChanged,
  canAdd,
  readonly,
}: CategoryListProps) => {
  const [openCategory, setOpenCategory] = useState(false);
  const { categories, isLoading, isError, mutate } = getAllCategory({
    limit: 1000,
  });
  const handleCategoryAdded = async (idCate: string) => {
    await mutate();
    onCheckChanged(idCate);
  };
  if (isError) return <div>Failed to load</div>;
  if (isLoading) {
    return <Loading />;
  } else
    return (
      <div className="flex gap-1">
        <DropdownMenu open={openCategory} onOpenChange={setOpenCategory}>
          <DropdownMenuTrigger asChild>
            <Button
              variant="outline"
              role="combobox"
              aria-expanded={openCategory}
              className="justify-between w-full"
            >
              Chọn thể loại
              <LuChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent className="DropdownMenuContent">
            <Command>
              <CommandInput placeholder="Tìm điều kiện lọc" />
              <CommandEmpty className="py-2">
                <div className="text-sm">Không tìm thấy thể loại</div>
              </CommandEmpty>
              <CommandGroup className="overflow-y-auto">
                {categories?.data.map((item: any) => (
                  <CommandItem
                    value={item.name}
                    key={item.id}
                    onSelect={() => {
                      onCheckChanged(item.id);
                      // setOpenCategory(false);
                    }}
                  >
                    <Checkbox
                      className="mr-2"
                      id={item.name}
                      checked={checkedCategory.includes(item.id)}
                    ></Checkbox>
                    {item.name}
                  </CommandItem>
                ))}
              </CommandGroup>
            </Command>
          </DropdownMenuContent>
        </DropdownMenu>
        {canAdd ? (
          <CreateCategory handleCategoryAdded={handleCategoryAdded}>
            <Button type="button" size={"icon"} className="px-3">
              <FaPlus />
            </Button>
          </CreateCategory>
        ) : null}
      </div>
    );
};

export default CategoryList;
