"use client";
import { Invoice } from "@/types";
import { Checkbox } from "../ui/checkbox";
import { toVND } from "@/lib/utils";
import { CaretSortIcon, ChevronDownIcon } from "@radix-ui/react-icons";
import { Button } from "../ui/button";
import Paging from "../paging";
import {
  ColumnDef,
  ColumnFiltersState,
  SortingState,
  VisibilityState,
  flexRender,
  getCoreRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  useReactTable,
} from "@tanstack/react-table";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { useRouter, useSearchParams } from "next/navigation";
import { useEffect, useState } from "react";
import { SubmitHandler, useFieldArray, useForm } from "react-hook-form";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../ui/select";
import { Popover, PopoverContent, PopoverTrigger } from "../ui/popover";
import { LuFilter } from "react-icons/lu";
import { Label } from "../ui/label";
import { Input } from "../ui/input";
import { AiOutlineClose } from "react-icons/ai";
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "../ui/dropdown-menu";
import StaffList from "../staff-list";
import { GiShamrock } from "react-icons/gi";

type FormValues = {
  filters: {
    type: string;
    value: string;
  }[];
};
function idToName(id: string) {
  if (id === "id") {
    return "Mã hóa đơn";
  } else if (id === "createdAt") {
    return "Ngày tạo";
  } else if (id === "createdBy") {
    return "Người tạo";
  } else if (id === "totalPrice") {
    return "Tổng tiền";
  } else if (id === "customer") {
    return "Khách hàng";
  } else if (id === "amountPriceUsePoint") {
    return "Giảm từ điểm tích luỹ";
  } else if (id === "amountReceived") {
    return "Thành tiền";
  } else {
    return id;
  }
}
export const columns: ColumnDef<Invoice>[] = [
  {
    accessorKey: "id",
    header: () => {
      return <span className="font-semibold">ID</span>;
    },
    cell: ({ row }) => <div>{row.getValue("id")}</div>,
  },
  {
    accessorKey: "customer",
    header: () => {
      return <div className="font-semibold">Khách hàng</div>;
    },
    cell: ({ row }) => {
      if (row.original.customer) {
        return (
          <div className="leading-6 flex flex-col text-left">
            <span>{row.original.customer.name}</span>
            <span className="font-light">{row.original.customer.phone}</span>
          </div>
        );
      } else {
        return <></>;
      }
    },
  },
  {
    accessorKey: "totalPrice",
    header: ({ column }) => (
      <div className=" flex justify-end">
        <Button
          className="p-1"
          variant={"ghost"}
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
        >
          <span className="font-semibold">Tổng tiền</span>

          <CaretSortIcon className="ml-1 h-4 w-4" />
        </Button>
      </div>
    ),
    cell: ({ row }) => {
      const amount = parseFloat(row.getValue("totalPrice"));
      const formatted = toVND(amount);
      return <div className="text-right font-medium">{formatted}</div>;
    },
  },
  {
    accessorKey: "amountPriceUsePoint",
    header: ({ column }) => (
      <div className="flex justify-end">
        <Button
          variant={"ghost"}
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
          className="p-1"
        >
          <CaretSortIcon className="h-4 w-4" />
          <span className="font-semibold">Dùng điểm</span>
        </Button>
      </div>
    ),
    cell: ({ row }) => {
      const amount = row.original.amountPriceUsePoint;

      // Format the amount as a dollar amount
      const formatted = new Intl.NumberFormat("vi-VN", {
        style: "currency",
        currency: "VND",
      }).format(amount);
      if (row.original.customer) {
        return (
          <div className="text-right font-medium flex flex-col items-end gap-1">
            -{formatted}
            <div className="flex items-center gap-1 text-rose-700">
              -{row.original.pointUse.toLocaleString("vi-VN")}{" "}
              <GiShamrock className="h-5 w-5" />
            </div>
          </div>
        );
      } else {
        return <></>;
      }
    },
  },
  {
    accessorKey: "amountReceived",
    header: ({ column }) => (
      <div className="flex justify-end">
        <Button
          variant={"ghost"}
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
          className="p-1"
        >
          <span className="font-semibold">Thành tiền</span>

          <CaretSortIcon className="h-4 w-4" />
        </Button>
      </div>
    ),
    cell: ({ row }) => {
      const amount = row.original.amountReceived;

      // Format the amount as a dollar amount
      const formatted = new Intl.NumberFormat("vi-VN", {
        style: "currency",
        currency: "VND",
      }).format(amount);

      return (
        <div className="text-right font-medium flex flex-col items-end gap-1">
          {formatted}
          <div className="flex items-center gap-1 text-green-700">
            {row.original.pointReceive.toLocaleString("vi-VN")}{" "}
            <GiShamrock className="h-5 w-5" />
          </div>
        </div>
      );
    },
  },
  {
    accessorKey: "createdAt",
    header: ({ column }) => {
      return (
        <div className="flex justify-end">
          <Button
            className="p-2 justify-end whitespace-normal"
            variant="ghost"
            onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
          >
            <CaretSortIcon className=" h-4 w-4" />
            <span className="font-semibold">Ngày tạo</span>
          </Button>
        </div>
      );
    },
    cell: ({ row }) => (
      <div className="leading-6 flex flex-col text-right">
        <span>
          {new Date(row.original.createdAt).toLocaleDateString("vi-VN")}
        </span>
        <span className="font-light">{row.original.createdBy.name}</span>
      </div>
    ),
    sortingFn: "datetime",
  },
];
const InvoiceTable = ({
  data,
  totalPage,
}: {
  data: Invoice[];
  totalPage: number;
}) => {
  const router = useRouter();
  const searchParams = useSearchParams();
  const page = searchParams.get("page") ?? "1";

  const [sorting, setSorting] = useState<SortingState>([]);
  const [columnFilters, setColumnFilters] = useState<ColumnFiltersState>([]);
  const [columnVisibility, setColumnVisibility] = useState<VisibilityState>({});
  const [rowSelection, setRowSelection] = useState({});
  const table = useReactTable({
    data,
    columns,
    onSortingChange: setSorting,
    onColumnFiltersChange: setColumnFilters,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    getSortedRowModel: getSortedRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    onColumnVisibilityChange: setColumnVisibility,
    onRowSelectionChange: setRowSelection,
    state: {
      sorting,
      columnFilters,
      columnVisibility,
      rowSelection,
    },
  });

  const [exportOption, setExportOption] = useState("all");
  const handleExport = () => {
    if (exportOption === "all") {
      // ExportSupplierList(data, "Suppliers.xlsx");
    }
    if (table.getFilteredSelectedRowModel().rows.length < 1) {
      //TODO: show notification
    } else {
      const values = table
        .getFilteredSelectedRowModel()
        .rows.map((row) => row.original);
      // ExportSupplierList(values, "Suppliers.xlsx");
    }
  };
  const [staff, setStaff] = useState("");
  const handleSetStaff = (staff: string) => {
    setStaff(staff);
    const index = fields.findIndex((item) => item.type === "createdBy");
    if (index > -1) {
      update(index, { type: "createdBy", value: staff });
    } else {
      append({ type: "createdBy", value: staff });
    }
  };
  const [latestFilter, setLatestFilter] = useState("");
  const filterValues = [
    { type: "search", name: "Từ khoá" },
    { type: "minPrice", name: "Tổng tiền nhỏ nhất" },
    { type: "maxPrice", name: "Tổng tiền lớn nhất" },
    { type: "createdBy", name: "Mã người tạo" },
  ];
  const maxPrice = searchParams.get("maxPrice") ?? undefined;
  const minPrice = searchParams.get("minPrice") ?? undefined;
  const search = searchParams.get("search") ?? undefined;
  const createdBy = searchParams.get("createdBy") ?? undefined;
  let filters = [{ type: "", value: "" }];
  filters.pop();
  if (maxPrice) {
    filters = filters.concat({ type: "maxPrice", value: maxPrice });
  }
  if (minPrice) {
    filters = filters.concat({ type: "minPrice", value: minPrice });
  }
  if (search) {
    filters = filters.concat({ type: "search", value: search });
  }
  if (createdBy) {
    filters = filters.concat({ type: "createdBy", value: createdBy });
  }
  let stringToFilter = "";
  filters.forEach((item) => {
    stringToFilter = stringToFilter.concat(`&${item.type}=${item.value}`);
  });
  const { register, handleSubmit, reset, control, getValues } =
    useForm<FormValues>({
      defaultValues: {
        filters: filters,
      },
    });
  const { fields, append, remove, update } = useFieldArray({
    control: control,
    name: "filters",
  });
  useEffect(() => {
    if (createdBy) {
      setStaff(createdBy);
    }
  }, [createdBy]);
  const onSubmit: SubmitHandler<FormValues> = async (data) => {
    let search = "";
    let minPrice = "";
    let maxPrice = "";
    let createdBy = "";
    data.filters.forEach((item) => {
      if (item.type === "minPrice") {
        minPrice = `&minPrice=${item.value}`;
      } else if (item.type === "maxPrice") {
        maxPrice = `&maxPrice=${item.value}`;
      } else if (item.type === "search") {
        search = `&search=${item.value}`;
      } else if (item.type === "createdBy") {
        createdBy = `&createdBy=${item.value}`;
      }
    });

    router.push(`/invoice?page=1${minPrice}${maxPrice}${search}${createdBy}`);
  };
  const [openFilter, setOpenFilter] = useState(false);

  return (
    <div className="w-full">
      <div className="flex items-start py-4 gap-2">
        {/* <DialogSupplierExport
          handleExport={handleExport}
          setExportOption={setExportOption}
        /> */}
        <div className="flex-1">
          <div className="flex gap-2">
            <Popover
              open={openFilter}
              onOpenChange={(open) => {
                setOpenFilter(open);
                reset({ filters: filters });
              }}
            >
              <PopoverTrigger asChild>
                <Button variant="outline" className="lg:px-3 px-2">
                  Lọc
                  <LuFilter className="ml-1 h-4 w-4" />
                </Button>
              </PopoverTrigger>
              <PopoverContent className="w-80">
                <form
                  className="flex flex-col gap-4"
                  onSubmit={handleSubmit(onSubmit)}
                >
                  <div className="space-y-2">
                    <p className="text-sm text-muted-foreground">
                      Hiển thị hóa đơn theo
                    </p>
                  </div>
                  <div className="flex flex-col gap-4">
                    {fields.map((item, index) => {
                      const name = filterValues.find(
                        (v) => v.type === item.type
                      );
                      return (
                        <div className="flex gap-2 items-center" key={item.id}>
                          <Label className="basis-1/4">{name?.name}</Label>
                          {item.type === "createdBy" ? (
                            <div className="flex-1">
                              <StaffList
                                staff={staff}
                                setStaff={handleSetStaff}
                              />
                            </div>
                          ) : item.type === "search" ? (
                            <Input
                              {...register(`filters.${index}.value`)}
                              className="flex-1"
                              type="text"
                              required
                            ></Input>
                          ) : (
                            <Input
                              {...register(`filters.${index}.value`)}
                              className="flex-1"
                              type="number"
                              required
                            ></Input>
                          )}

                          <Button
                            variant={"ghost"}
                            className={`px-3 `}
                            onClick={() => {
                              if (item.type === "createdBy") {
                                setStaff("");
                              }
                              remove(index);
                            }}
                          >
                            <AiOutlineClose />
                          </Button>
                        </div>
                      );
                    })}
                  </div>
                  {fields.length === filterValues.length ? null : (
                    <div className="flex justify-center">
                      <Select
                        value={latestFilter}
                        onValueChange={(value) => {
                          append({ type: value, value: "" });
                        }}
                      >
                        <SelectTrigger className="w-[160px] flex justify-center ml-8 px-3">
                          <SelectValue placeholder="Chọn điều kiện lọc" />
                        </SelectTrigger>
                        <SelectContent>
                          <SelectGroup>
                            {filterValues.map((item) => {
                              return fields.findIndex(
                                (v) => v.type === item.type
                              ) === -1 ? (
                                <SelectItem key={item.type} value={item.type}>
                                  {item.name}
                                </SelectItem>
                              ) : null;
                            })}
                          </SelectGroup>
                        </SelectContent>
                      </Select>
                    </div>
                  )}
                  <Button type="submit" className="self-end">
                    Lọc
                  </Button>
                </form>
              </PopoverContent>
            </Popover>
            <div className="flex-1">
              <Input
                placeholder="Tìm kiếm hóa đơn"
                value={
                  (table.getColumn("id")?.getFilterValue() as string) ?? ""
                }
                onChange={(event) =>
                  table.getColumn("id")?.setFilterValue(event.target.value)
                }
              />
            </div>
          </div>
          <div className="flex gap-2 mt-2">
            {filters.map((item, index) => {
              const name = filterValues.find((v) => v.type === item.type);
              return (
                <div
                  key={item.type}
                  className="rounded-xl flex self-start px-3 py-1 h-fit outline-none text-sm text-primary  bg-blue-100 items-center gap-1 group"
                >
                  <span>
                    {name?.name}
                    {": "}
                    {item.value}
                  </span>
                </div>
              );
            })}
          </div>
        </div>

        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="outline" className="lg:px-3 px-2">
              Cột hiển thị
              <ChevronDownIcon className="ml-1 h-4 w-4" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent className="DropdownMenuContent">
            {table
              .getAllColumns()
              .filter((column) => column.getCanHide())
              .map((column) => {
                return (
                  <DropdownMenuCheckboxItem
                    key={column.id}
                    className="capitalize"
                    checked={column.getIsVisible()}
                    onCheckedChange={(value) =>
                      column.toggleVisibility(!!value)
                    }
                  >
                    {idToName(column.id)}
                  </DropdownMenuCheckboxItem>
                );
              })}
          </DropdownMenuContent>
        </DropdownMenu>
      </div>
      <div className="rounded-md border overflow-x-auto min-w-full max-w-[50vw]">
        <Table>
          <TableHeader>
            {table.getHeaderGroups().map((headerGroup) => (
              <TableRow key={headerGroup.id}>
                {headerGroup.headers.map((header) => {
                  return (
                    <TableHead key={header.id}>
                      {header.isPlaceholder
                        ? null
                        : flexRender(
                            header.column.columnDef.header,
                            header.getContext()
                          )}
                    </TableHead>
                  );
                })}
              </TableRow>
            ))}
          </TableHeader>
          <TableBody>
            {table.getRowModel().rows?.length ? (
              table.getRowModel().rows.map((row) => (
                <TableRow
                  key={row.id}
                  data-state={row.getIsSelected() && "selected"}
                >
                  {row.getVisibleCells().map((cell) => (
                    <TableCell
                      key={cell.id}
                      onClick={() => {
                        if (!cell.id.includes("select")) {
                          router.push(`/invoice/${row.getValue("id")}`);
                        }
                      }}
                    >
                      {flexRender(
                        cell.column.columnDef.cell,
                        cell.getContext()
                      )}
                    </TableCell>
                  ))}
                </TableRow>
              ))
            ) : (
              <TableRow>
                <TableCell
                  colSpan={columns.length}
                  className="h-24 text-center"
                >
                  Không tìm thấy kết quả.
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </div>
      <div className="flex items-center justify-end space-x-2 py-4">
        <div className="flex-1 text-sm text-muted-foreground"></div>
        <Paging
          page={page}
          totalPage={totalPage}
          onNavigateBack={() =>
            router.push(`/invoice?page=${Number(page) - 1}${stringToFilter}`)
          }
          onNavigateNext={() =>
            router.push(`/invoice?page=${Number(page) + 1}${stringToFilter}`)
          }
          onPageSelect={(selectedPage) =>
            router.push(`/invoice?page=${selectedPage}${stringToFilter}`)
          }
          onNavigateFirst={() =>
            router.push(`/invoice?page=${1}${stringToFilter}`)
          }
          onNavigateLast={() =>
            router.push(`/invoice?page=${totalPage}${stringToFilter}`)
          }
        />
      </div>
    </div>
  );
};

export default InvoiceTable;
