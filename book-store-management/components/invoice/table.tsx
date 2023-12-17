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
  }
  return id;
}
export const columns: ColumnDef<Invoice>[] = [
  {
    id: "select",
    header: ({ table }) => (
      <Checkbox
        checked={table.getIsAllPageRowsSelected()}
        onCheckedChange={(value) => table.toggleAllPageRowsSelected(!!value)}
        aria-label="Select all"
      />
    ),
    cell: ({ row }) => (
      <Checkbox
        checked={row.getIsSelected()}
        onCheckedChange={(value) => row.toggleSelected(!!value)}
        aria-label="Select row"
      />
    ),
    enableSorting: false,
    enableHiding: false,
  },
  {
    accessorKey: "id",
    header: () => {
      return <span className="font-semibold">ID</span>;
    },
    cell: ({ row }) => <div>{row.getValue("id")}</div>,
  },
  {
    accessorKey: "createdAt",
    accessorFn: (row) => new Date(row.createdAt).toLocaleDateString("vi-VN"),

    header: ({ column }) => {
      return <span className="font-semibold">Ngày tạo</span>;
    },
    cell: ({ row }) => <div>{row.getValue("createdAt")}</div>,
  },
  {
    accessorKey: "createdBy",
    accessorFn: (row) => row.createdBy.name,
    header: () => {
      return <div className="font-semibold">Người tạo</div>;
    },
    cell: ({ row }) => <span>{row.getValue("createdBy")}</span>,
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
          <span className="font-semibold">Tổng nợ</span>

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
    { type: "min", name: "Tổng tiền nhỏ nhất" },
    { type: "max", name: "Tổng tiền lớn nhất" },
    { type: "createdBy", name: "Người tạo" },
  ];
  const maxDebt = searchParams.get("maxPrice") ?? undefined;
  const minDebt = searchParams.get("minPrice") ?? undefined;
  const search = searchParams.get("search") ?? undefined;
  const createdBy = searchParams.get("createdBy") ?? undefined;
  let filters = [{ type: "", value: "" }];
  filters.pop();
  if (maxDebt) {
    filters = filters.concat({ type: "min", value: maxDebt });
  }
  if (minDebt) {
    filters = filters.concat({ type: "max", value: minDebt });
  }
  if (search) {
    filters = filters.concat({ type: "search", value: search });
  }
  if (createdBy) {
    filters = filters.concat({ type: "createdBy", value: createdBy });
  }
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
    console.log(data);
    let search = "";
    let minDebt = "";
    let maxDebt = "";
    let createdBy = "";
    data.filters.forEach((item) => {
      if (item.type === "min") {
        minDebt = `&minPrice=${item.value}`;
      } else if (item.type === "max") {
        maxDebt = `&maxPrice=${item.value}`;
      } else if (item.type === "search") {
        search = `&search=${item.value}`;
      } else if (item.type === "createdBy") {
        createdBy = `&createdBy=${item.value}`;
      }
    });

    router.push(
      `/invoice?page=${Number(page)}${minDebt}${maxDebt}${search}${createdBy}`
    );
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
                      Hiển thị nhà cung cấp theo
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
                            <StaffList
                              staff={staff}
                              setStaff={handleSetStaff}
                            />
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
                          console.log(value);
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
                placeholder="Tìm kiếm nhà cung cấp"
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
      <div className="rounded-md border">
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
                  No results.
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </div>
      <div className="flex items-center justify-end space-x-2 py-4">
        <div className="flex-1 text-sm text-muted-foreground">
          {table.getFilteredSelectedRowModel().rows.length} of{" "}
          {table.getFilteredRowModel().rows.length} row(s) selected.
        </div>
        <Paging
          page={page}
          totalPage={totalPage}
          onNavigateBack={() =>
            router.push(`/invoice?page=${Number(page) - 1}`)
          }
          onNavigateNext={() =>
            router.push(`/invoice?page=${Number(page) + 1}`)
          }
          onPageSelect={(selectedPage) =>
            router.push(`/invoice?page=${selectedPage}`)
          }
          onNavigateFirst={() => router.push(`/supplier?page=${1}`)}
          onNavigateLast={() => router.push(`/supplier?page=${totalPage}`)}
        />
      </div>
    </div>
  );
};

export default InvoiceTable;
