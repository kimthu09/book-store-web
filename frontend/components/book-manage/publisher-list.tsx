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
import { cn, includesRoles } from "@/lib/utils";
import Loading from "../loading";
import CreatePublisher from "./create-publisher";
import { FaPlus } from "react-icons/fa";
import getAllPublisherList from "@/lib/book/getAllPublisherList";
import { useCurrentUser } from "@/hooks/use-user";
import ListSkeleton from "../skeleton/list_skeleton";

const PublisherList = ({
  publisherId,
  setPublisherId,
  canAdd,
  readOnly,
}: PublisherListProps) => {
  const [open, setOpen] = useState(false);
  const { publishers, isLoading, isError, mutate } = getAllPublisherList();
  const handlePublisherAdded = async (publisherId: string) => {
    await mutate();
    setPublisherId(publisherId);
  };
  const { currentUser } = useCurrentUser();

  if (isError) return <div>Failed to load</div>;
  if (!publishers) {
    <ListSkeleton numberRow={5} />;
  } else
    return (
      <div className="flex gap-1">
        <DropdownMenu open={open} onOpenChange={setOpen}>
          <DropdownMenuTrigger asChild>
            <Button
              disabled={readOnly}
              variant="outline"
              role="combobox"
              aria-expanded={open}
              className="justify-between w-full pl-2"
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
              <CommandGroup className="max-h-48 overflow-y-auto">
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
          currentUser &&
          includesRoles({
            currentUser: currentUser,
            allowedFeatures: ["PUBLISHER_CREATE"],
          }) ? (
            <CreatePublisher handlePublisherAdded={handlePublisherAdded}>
              <Button type="button" size={"icon"} className="px-3">
                <FaPlus />
              </Button>
            </CreatePublisher>
          ) : null
        ) : null}
      </div>
    );
};

export default PublisherList;
