import { PublisherListProps } from "@/types";
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
import { cn } from "@/lib/utils";
import Loading from "../loading";
import getAllPublisher from "@/lib/book/getAllPublisher";
import CreatePublisher from "./create-publisher";
import { FaPlus } from "react-icons/fa";

const PublisherList = ({
  publisherId,
  setPublisherId,
  canAdd,
}: PublisherListProps) => {
  const [open, setOpen] = useState(false);
  const { publishers, isLoading, isError, mutate } = getAllPublisher({
    limit: 1000,
  });
  const handlePublisherAdded = async (publisherId: string) => {
    await mutate();
    setPublisherId(publisherId);
  };
  if (isError) return <div>Failed to load</div>;
  if (!publishers) {
    <Loading />;
  } else
    return (
      <div className="flex gap-1">
        <DropdownMenu open={open} onOpenChange={setOpen}>
          <DropdownMenuTrigger asChild>
            <Button
              variant="outline"
              role="combobox"
              aria-expanded={open}
              className="justify-between w-full pl-1"
            >
              {publisherId
                ? publishers.data.find((item: any) => item.id === publisherId)
                    ?.name
                : "Chọn nhà xuất bản"}
              <LuChevronsUpDown className="ml-1 h-4 w-4 shrink-0 opacity-50" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent className="DropdownMenuContent">
            <Command>
              <CommandInput
                placeholder="Tìm nhà xuất bản"
                // onValueChange={(str) => setNewCategory(str)}
              />
              <CommandEmpty className="py-2 px-6">
                <div className="text-sm">Không tìm thấy nhà xuất bản</div>
              </CommandEmpty>
              <CommandGroup>
                {publishers.data.map((item: any) => (
                  <CommandItem
                    value={item.name}
                    key={item.id}
                    onSelect={() => {
                      setPublisherId(item.id);
                      setOpen(false);
                    }}
                  >
                    <LuCheck
                      className={cn(
                        "mr-2 h-4 w-4",
                        item.id === publisherId ? "opacity-100" : "opacity-0"
                      )}
                    />
                    {item.name}
                  </CommandItem>
                ))}
              </CommandGroup>
            </Command>
          </DropdownMenuContent>
        </DropdownMenu>
        {canAdd ? (
          <CreatePublisher handlePublisherAdded={handlePublisherAdded}>
            <Button type="button" size={"icon"} className="px-3">
              <FaPlus />
            </Button>
          </CreatePublisher>
        ) : null}
      </div>
    );
};

export default PublisherList;
