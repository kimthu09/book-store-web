import { useState } from "react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "../ui/dropdown-menu";
import { Button } from "../ui/button";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
} from "../ui/command";
import { LuCheck, LuChevronsUpDown } from "react-icons/lu";
import { cn, includesRoles } from "@/lib/utils";
import Loading from "../loading";
import { BookTitle, TitleListProps } from "@/types";
import { FaPlus } from "react-icons/fa";
import CreateTitleDialog from "./create-title-dialog";
import { useCurrentUser } from "@/hooks/use-user";
import getAllTitleList from "@/lib/book/getAllTitleList";
import ListSkeleton from "../skeleton/list_skeleton";

const BookTitleSelect = ({ handleTitleSet }: TitleListProps) => {
  const [open, setOpen] = useState(false);
  const [title, setTitle] = useState<BookTitle>();
  const { titles, isLoading, isError, mutate } = getAllTitleList();

  const onSetTitle = (title: BookTitle) => {
    setTitle(title);
    handleTitleSet(title.id);
  };
  const handleTitleAdded = async (titleId: string) => {
    const newTitleList = await mutate();
    setTitle(newTitleList?.data.find((item: BookTitle) => item.id === titleId));
    handleTitleSet(titleId);
  };
  const { currentUser } = useCurrentUser();
  if (isError) return <div>Failed to load</div>;
  if (!titles || isLoading) {
    <ListSkeleton numberRow={5} />;
  } else
    return (
      <div className="flex flex-col gap-3">
        <div className="flex gap-1">
          <DropdownMenu open={open} onOpenChange={setOpen}>
            <DropdownMenuTrigger asChild>
              <Button
                variant="outline"
                role="combobox"
                aria-expanded={open}
                className="justify-between w-full p-2"
              >
                {title
                  ? titles.data.find((item: BookTitle) => item.id === title.id)
                      ?.name
                  : "Chọn đầu sách"}
                <LuChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent className="DropdownMenuContent">
              <Command>
                <CommandInput placeholder="Tìm đầu sách" />
                <CommandEmpty className="py-2 px-6">
                  <div className="text-sm">Không tìm thấy đầu sách</div>
                </CommandEmpty>
                <CommandGroup className="max-h-48 overflow-y-auto">
                  {titles.data.map((item: BookTitle) => (
                    <CommandItem
                      value={item.name}
                      key={item.id}
                      onSelect={() => {
                        onSetTitle(item);
                        setOpen(false);
                      }}
                    >
                      <LuCheck
                        className={cn(
                          "mr-2 h-4 w-4",
                          title
                            ? item.id === title.id
                              ? "opacity-100"
                              : "opacity-0"
                            : "opacity-0"
                        )}
                      />
                      {item.name}
                    </CommandItem>
                  ))}
                </CommandGroup>
              </Command>
            </DropdownMenuContent>
          </DropdownMenu>
          {currentUser &&
          includesRoles({
            currentUser: currentUser,
            allowedFeatures: ["BOOK_TITLE_CREATE"],
          }) ? (
            <CreateTitleDialog handleTitleAdded={handleTitleAdded}>
              <Button type="button" size={"icon"} className="px-3">
                <FaPlus />
              </Button>
            </CreateTitleDialog>
          ) : null}
        </div>
        {title ? (
          <div className="flex flex-col gap-3 text-sm font-medium mt-2">
            <div className="flex">
              <span className="basis-1/4 text-muted-foreground">
                Mã đầu sách:
              </span>
              <span className="basis-3/4">{title?.id}</span>
            </div>
            <div className="flex">
              <span className="basis-1/4 text-muted-foreground">Tác giả: </span>
              <span className="basis-3/4">
                {title.authors.map((author) => author.name).join(", ")}
              </span>
            </div>
            <div className="flex">
              <span className="basis-1/4 text-muted-foreground">
                Thể loại:{" "}
              </span>
              <span className="basis-3/4">
                {title.categories.map((cate) => cate.name).join(", ")}
              </span>
            </div>
            <div className="flex">
              <span className="basis-1/4 text-muted-foreground">Mô tả: </span>
              <span className="basis-3/4 font-normal">
                {title?.desc && title.desc != "" ? title.desc : "Chưa có mô tả"}
              </span>
            </div>
          </div>
        ) : null}
      </div>
    );
};

export default BookTitleSelect;
