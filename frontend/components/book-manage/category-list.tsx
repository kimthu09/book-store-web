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
import { LuChevronsUpDown } from "react-icons/lu";
import { Button } from "../ui/button";
import { AiOutlineClose } from "react-icons/ai";
import { includesRoles } from "@/lib/utils";
import { CategoryListProps } from "@/types";
import { Checkbox } from "../ui/checkbox";
import Loading from "../loading";
import CreateCategory from "./create-category";
import { FaPlus } from "react-icons/fa";
import { useCurrentUser } from "@/hooks/use-user";
import getAllCategoryList from "@/lib/book/getAllCategoryList";
import ListSkeleton from "../skeleton/list_skeleton";

const CategoryList = ({
  checkedCategory,
  onCheckChanged,
  canAdd,
  readonly,
  isEdit,
  onRemove,
}: CategoryListProps) => {
  const [openCategory, setOpenCategory] = useState(false);
  const { categories, isLoading, isError, mutate } = getAllCategoryList();
  const handleCategoryAdded = async (idCate: string) => {
    await mutate();
    onCheckChanged(idCate);
  };
  const { currentUser } = useCurrentUser();
  if (isError) return <div>Failed to load</div>;
  if (isLoading) {
    return <ListSkeleton numberRow={5} />;
  } else
    return (
      <div className="flex flex-col">
        <div className="flex gap-1">
          <DropdownMenu open={openCategory} onOpenChange={setOpenCategory}>
            <DropdownMenuTrigger asChild>
              <Button
                variant="outline"
                role="combobox"
                aria-expanded={openCategory}
                className="justify-between w-full bg-white"
              >
                Chọn thể loại
                <LuChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent className="DropdownMenuContent">
              <Command>
                <CommandInput placeholder="Tìm tên thể loại" />
                <CommandEmpty className="py-2">
                  <div className="text-sm">Không tìm thấy thể loại</div>
                </CommandEmpty>
                <CommandGroup className="overflow-y-auto max-h-48">
                  {categories?.data.map((item: any) => (
                    <CommandItem
                      value={item.name}
                      key={item.id}
                      onSelect={() => {
                        if (isEdit !== false) {
                          onCheckChanged(item.id);
                        }
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
          {canAdd && (isEdit === true || isEdit === null) ? (
            currentUser &&
            includesRoles({
              currentUser: currentUser,
              allowedFeatures: ["CATEGORY_CREATE"],
            }) ? (
              <CreateCategory handleCategoryAdded={handleCategoryAdded}>
                <Button type="button" size={"icon"} className="px-3">
                  <FaPlus />
                </Button>
              </CreateCategory>
            ) : null
          ) : null}
        </div>
        {isEdit !== null ? (
          <div className="flex flex-wrap gap-2 mt-1">
            {checkedCategory.map((cate, index) => (
              <div
                key={cate}
                className="rounded-xl flex  px-3 py-1 h-fit outline-none text-sm text-primary  bg-blue-100 items-center gap-1 group"
              >
                {categories?.data.find((item: any) => item.id === cate)?.name}
                <div
                  className={`cursor-pointer w-4 ${
                    isEdit ? "block" : "hidden"
                  }`}
                >
                  <AiOutlineClose className="group-hover:hidden" />
                  <AiOutlineClose
                    color="red"
                    fill="red"
                    className="text-primary group-hover:flex hidden h-4 w-4"
                    onClick={() => {
                      if (onRemove) {
                        onRemove(index);
                      }
                    }}
                  />
                </div>
              </div>
            ))}
          </div>
        ) : null}
      </div>
    );
};

export default CategoryList;
