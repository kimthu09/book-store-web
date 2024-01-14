import { useState } from "react";
import Loading from "./loading";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "./ui/dropdown-menu";
import { Button } from "./ui/button";
import { LuCheck, LuChevronsUpDown } from "react-icons/lu";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
} from "./ui/command";
import { cn, includesRoles } from "@/lib/utils";
import getAllCustomer from "@/lib/customer/getAllCustomer";
import { useCurrentUser } from "@/hooks/use-user";
import CreateDialog from "./customer/create";
import { FaPlus } from "react-icons/fa";
import { IoPersonRemoveSharp } from "react-icons/io5";
import DropdownSkeleton from "./skeleton/dropdown-skeleton";
export interface SupplierListProps {
  customerId: string;
  setCustomerId: (id: string, point: number) => void;
  canAdd?: boolean;
  handleCustomerAdded?: (categoryId: string) => void;
  canRemove?: boolean;
  onRemove?: () => void;
}
const CustomerList = ({
  onRemove,
  canRemove,
  customerId,
  setCustomerId,
  canAdd,
  handleCustomerAdded,
}: SupplierListProps) => {
  const [open, setOpen] = useState(false);
  const { suppliers, isLoading, isError, mutate } = getAllCustomer();
  const { currentUser } = useCurrentUser();
  if (isError) return <div>Failed to load</div>;
  if (isLoading || !currentUser) {
    return <DropdownSkeleton />;
  } else
    return (
      <div className="flex gap-1">
        <DropdownMenu open={open} onOpenChange={setOpen}>
          <DropdownMenuTrigger asChild>
            <Button
              variant="outline"
              role="combobox"
              aria-expanded={open}
              className="justify-between w-full pl-2"
            >
              {customerId
                ? suppliers?.find((item: any) => item.id === customerId)?.name
                : "Chọn khách hàng"}
              <LuChevronsUpDown className="ml-1 h-4 w-4 shrink-0 opacity-50" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent className="DropdownMenuContent">
            <Command>
              <CommandInput
                placeholder="Tìm số điện thoại khách hàng"
                // onValueChange={(str) => setNewCategory(str)}
              />
              <CommandEmpty className="py-2 px-6">
                <div className="text-sm">Không tìm thấy</div>
              </CommandEmpty>
              <CommandGroup className="max-h-48 overflow-y-auto">
                {suppliers?.map((item: any) => (
                  <CommandItem
                    value={item.phone}
                    key={item.id}
                    onSelect={() => {
                      setCustomerId(item.id, item.point);
                      setOpen(false);
                    }}
                  >
                    <LuCheck
                      className={cn(
                        "mr-2 h-4 w-4",
                        item.id === customerId ? "opacity-100" : "opacity-0"
                      )}
                    />
                    <div className="flex flex-col">
                      {item.name}
                      <span className="text-muted-foreground">
                        {item.phone}
                      </span>
                    </div>
                  </CommandItem>
                ))}
              </CommandGroup>
            </Command>
          </DropdownMenuContent>
        </DropdownMenu>
        {canRemove ? (
          <Button
            type="button"
            size={"icon"}
            variant={"outline"}
            className="px-2"
            onClick={() => {
              if (onRemove) {
                onRemove();
              }
            }}
          >
            <IoPersonRemoveSharp className="h-5 w-5 text-muted-foreground" />
          </Button>
        ) : null}
        {canAdd ? (
          currentUser &&
          includesRoles({
            currentUser: currentUser,
            allowedFeatures: ["CUSTOMER_CREATE"],
          }) ? (
            <CreateDialog handleCustomerAdded={handleCustomerAdded}>
              <Button type="button" size={"icon"} className="px-3">
                <FaPlus />
              </Button>
            </CreateDialog>
          ) : null
        ) : null}
      </div>
    );
};

export default CustomerList;
