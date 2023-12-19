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
import { AuthorListProps } from "@/types";
import { Checkbox } from "../ui/checkbox";
import getAllAuthor from "@/lib/book/getAllAuthor";
import CreateAuthor from "./create-author";
import { FaPlus } from "react-icons/fa";

const AuthorList = ({
  checkedAuthor,
  onCheckChanged,
  canAdd,
  readonly,
}: AuthorListProps) => {
  const [openAuthor, setOpenAuthor] = useState(false);
  const { authors, isLoading, isError, mutate } = getAllAuthor({ limit: 1000 });
  const handleAuthorAdded = async (idAuthor: string) => {
    await mutate();
    onCheckChanged(idAuthor);
  };
  if (isError) return <div>Failed to load</div>;
  if (!authors) {
    console.log(authors);
  } else
    return (
      <div className="flex gap-1">
        <DropdownMenu open={openAuthor} onOpenChange={setOpenAuthor}>
          <DropdownMenuTrigger asChild>
            <Button
              variant="outline"
              role="combobox"
              aria-expanded={openAuthor}
              className="justify-between w-full"
            >
              Chọn tác giả
              <LuChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent className=" DropdownMenuContent">
            <Command>
              <CommandInput placeholder="Tìm điều kiện lọc" />
              <CommandEmpty className="py-2">
                {canAdd ? (
                  <div className="flex">
                    <Button variant="ghost" className="flex-1">
                      <div className="text-left flex-1 text-primary flex items-center gap-2">
                        <AiFillPlusCircle size={20} />
                        Thêm
                        {/* {" " + newCategory} */}
                      </div>
                    </Button>
                  </div>
                ) : (
                  <div className="text-sm">Không tìm thấy điều kiện lọc.</div>
                )}
              </CommandEmpty>
              <CommandGroup className="overflow-y-auto">
                {authors.data.map((item) => (
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
                      checked={checkedAuthor.includes(item.id)}
                    ></Checkbox>
                    {item.name}
                  </CommandItem>
                ))}
              </CommandGroup>
            </Command>
          </DropdownMenuContent>
        </DropdownMenu>
        {canAdd ? (
          <CreateAuthor handleAuthorAdded={handleAuthorAdded}>
            <Button type="button" size={"icon"} className="px-3">
              <FaPlus />
            </Button>
          </CreateAuthor>
        ) : null}
      </div>
    );
};

export default AuthorList;
