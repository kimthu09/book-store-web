import { useState } from "react";
import { Button } from "../ui/button";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
} from "../ui/command";
import { books } from "@/constants";
import { BookListProps } from "@/types";
import { cn } from "@/lib/utils";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "../ui/dropdown-menu";
import { LuCheck, LuChevronsUpDown } from "react-icons/lu";
import { AiFillPlusCircle } from "react-icons/ai";

const BookList = ({ book, setBook, isNew, setIsNew }: BookListProps) => {
  const [openBookList, setOpenBookList] = useState(false);
  const [bookString, setBookString] = useState("");
  return (
    <div className="mb-4 flex-1">
      <DropdownMenu open={openBookList} onOpenChange={setOpenBookList}>
        <DropdownMenuTrigger asChild>
          <Button
            id="bookList"
            variant="outline"
            role="combobox"
            aria-expanded={openBookList}
            className="justify-between w-full"
          >
            {isNew
              ? book.name
              : book.id
              ? books.find((item) => item.id === book.id)?.name
              : "Chọn sách"}
            <LuChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent className=" DropdownMenuContent">
          <Command>
            <CommandInput
              placeholder="Tìm sách"
              onValueChange={(str) => setBookString(str)}
            />
            <CommandEmpty className="py-2">
              <div className="flex">
                <Button
                  variant="ghost"
                  className="flex-1"
                  onClick={() => {
                    setIsNew(true);
                    setBook({ name: bookString });
                    setOpenBookList(false);
                  }}
                >
                  <div className="text-left flex-1 text-primary flex items-center gap-2">
                    <AiFillPlusCircle size={20} />
                    Chọn
                    <span className="text-black">{bookString}</span>
                  </div>
                </Button>
              </div>
            </CommandEmpty>
            <CommandGroup className="overflow-y-auto  max-h-80">
              {books.map((item) => (
                <CommandItem
                  value={item.name}
                  key={item.id}
                  onSelect={() => {
                    setBook({
                      id: item.id,
                      name: item.name,
                      category: item.category,
                    });
                    setIsNew(false);
                    setOpenBookList(false);
                  }}
                >
                  <LuCheck
                    className={cn(
                      "mr-2 h-4 w-4",
                      item.id === book.id ? "opacity-100" : "opacity-0"
                    )}
                  />
                  {item.name}
                </CommandItem>
              ))}
            </CommandGroup>
          </Command>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  );
};

export default BookList;
