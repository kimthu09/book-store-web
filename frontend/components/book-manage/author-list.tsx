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
import { AuthorListProps } from "@/types";
import { Checkbox } from "../ui/checkbox";
import CreateAuthor from "./create-author";
import { FaPlus } from "react-icons/fa";
import { AiOutlineClose } from "react-icons/ai";
import { useCurrentUser } from "@/hooks/use-user";
import { includesRoles } from "@/lib/utils";
import getAllAuthorList from "@/lib/book/getAllAuthorList";

const AuthorList = ({
  checkedAuthor,
  onCheckChanged,
  canAdd,
  readonly,
  isEdit,
  onRemove,
}: AuthorListProps) => {
  const [openAuthor, setOpenAuthor] = useState(false);
  const { authors, isLoading, isError, mutate } = getAllAuthorList();
  const handleAuthorAdded = async (idAuthor: string) => {
    await mutate();
    onCheckChanged(idAuthor);
  };
  const { currentUser } = useCurrentUser();
  if (isError) return <div>Failed to load</div>;
  if (!authors) {
  } else
    return (
      <div className="flex flex-col">
        <div className="flex gap-1">
          <DropdownMenu open={openAuthor} onOpenChange={setOpenAuthor}>
            <DropdownMenuTrigger asChild>
              <Button
                variant="outline"
                role="combobox"
                aria-expanded={openAuthor}
                className="justify-between w-full bg-white"
              >
                Chọn tác giả
                <LuChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent className=" DropdownMenuContent">
              <Command>
                <CommandInput placeholder="Tìm điều kiện lọc" />
                <CommandEmpty className="py-2">
                  {<div className="text-sm">Không tìm thấy tác giả.</div>}
                </CommandEmpty>
                <CommandGroup className="overflow-y-auto max-h-48">
                  {authors.data.map((item) => (
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
                        checked={checkedAuthor.includes(item.id)}
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
              allowedFeatures: ["AUTHOR_CREATE"],
            }) ? (
              <CreateAuthor handleAuthorAdded={handleAuthorAdded}>
                <Button type="button" size={"icon"} className="px-3">
                  <FaPlus />
                </Button>
              </CreateAuthor>
            ) : null
          ) : null}
        </div>
        {isEdit !== null ? (
          <div className="flex flex-wrap gap-2 mt-1">
            {checkedAuthor.map((author, index) => (
              <div
                key={author}
                className="rounded-xl flex  px-3 py-1 h-fit outline-none text-sm text-primary  bg-blue-100 items-center gap-1 group"
              >
                {authors?.data.find((item: any) => item.id === author)?.name}
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

export default AuthorList;
