"use client";

import * as React from "react";
import { CaretSortIcon, ChevronDownIcon } from "@radix-ui/react-icons";
import { LuFilter } from "react-icons/lu";
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

import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Supplier } from "@/types";
import { useState } from "react";
import { Input } from "../ui/input";
import { ExportSupplierList } from "./excel-export";
import { Label } from "../ui/label";

import Paging from "../paging";
import { useRouter, useSearchParams } from "next/navigation";
import { Popover, PopoverContent, PopoverTrigger } from "../ui/popover";
import DialogSupplierExport from "./dialog-supplier-export";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../ui/select";
import { SubmitHandler, useFieldArray, useForm } from "react-hook-form";
import { AiOutlineClose } from "react-icons/ai";

export const columns: ColumnDef<Supplier>[] = [
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
    accessorKey: "name",
    header: ({ column }) => {
      return (
        <Button
          className="p-2"
          variant="ghost"
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
        >
          <span className="font-semibold">Nhà cung cấp</span>

          <CaretSortIcon className="ml-1 h-4 w-4" />
        </Button>
      );
    },
    cell: ({ row }) => <div className="capitalize">{row.getValue("name")}</div>,
  },
  {
    accessorKey: "email",
    header: () => {
      return <div className="font-semibold">Email</div>;
    },
    cell: ({ row }) => (
      <div className="lg:max-w-[16rem] max-w-[3rem] truncate">
        {row.getValue("email")}
      </div>
    ),
  },
  {
    accessorKey: "phone",
    header: () => {
      return (
        <div className="font-semibold flex justify-end">Số điện thoại</div>
      );
    },
    cell: ({ row }) => (
      <div className="text-right">{row.getValue("phone")}</div>
    ),
  },
  {
    accessorKey: "debt",
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
      const amount = parseFloat(row.getValue("debt"));

      // Format the amount as a dollar amount
      const formatted = new Intl.NumberFormat("vi-VN", {
        style: "currency",
        currency: "VND",
      }).format(amount);

      return <div className="text-right font-medium">{formatted}</div>;
    },
  },
];

function idToName(id: string) {
  if (id === "name") {
    return "Tên";
  } else if (id === "email") {
    return "Email";
  } else if (id === "phone") {
    return "Điện thoại";
  } else if (id === "debt") {
    return "Tổng nợ";
  }
  return id;
}

type FormValues = {
  filters: {
    type: string;
    value: string;
  }[];
};
export function SupplierTable({
  data,
  totalPage,
}: {
  data: Supplier[];
  totalPage: number;
}) {
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
      ExportSupplierList(data, "Suppliers.xlsx");
    }
    if (table.getFilteredSelectedRowModel().rows.length < 1) {
      //TODO: show notification
    } else {
      const values = table
        .getFilteredSelectedRowModel()
        .rows.map((row) => row.original);
      ExportSupplierList(values, "Suppliers.xlsx");
    }
  };

  const [latestFilter, setLatestFilter] = useState("");
  const filterValues = [
    { type: "search", name: "Từ khoá" },
    { type: "minDebt", name: "Tổng nợ nhỏ nhất" },
    { type: "maxDebt", name: "Tổng nợ lớn nhất" },
  ];
  const maxDebt = searchParams.get("maxDebt") ?? undefined;
  const minDebt = searchParams.get("minDebt") ?? undefined;
  const search = searchParams.get("search") ?? undefined;
  let filters = [{ type: "", value: "" }];
  filters.pop();
  if (maxDebt) {
    filters = filters.concat({ type: "maxDebt", value: maxDebt });
  }
  if (minDebt) {
    filters = filters.concat({ type: "minDebt", value: minDebt });
  }
  if (search) {
    filters = filters.concat({ type: "search", value: search });
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

  const onSubmit: SubmitHandler<FormValues> = async (data) => {
    let search = "";
    let minDebt = "";
    let maxDebt = "";
    data.filters.forEach((item) => {
      if (item.type === "minDebt") {
        minDebt = `&minDebt=${item.value}`;
      } else if (item.type === "maxDebt") {
        maxDebt = `&maxDebt=${item.value}`;
      } else if (item.type === "search") {
        search = `&search=${item.value}`;
      }
    });

    router.push(`/supplier?page=1${minDebt}${maxDebt}${search}`);
  };
  const [openFilter, setOpenFilter] = useState(false);

  return (
    <div className="w-full">
      <div className="flex items-start py-4 gap-2">
        <DialogSupplierExport
          handleExport={handleExport}
          setExportOption={setExportOption}
        />
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
                          {item.type === "search" ? (
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
                placeholder="Tìm kiếm nhà cung cấp"
                value={
                  (table.getColumn("name")?.getFilterValue() as string) ?? ""
                }
                onChange={(event) =>
                  table.getColumn("name")?.setFilterValue(event.target.value)
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
        <Table className="min-w-full w-max">
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
                          router.push(`/supplier/${row.getValue("id")}`);
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
        <div className="flex-1 text-sm text-muted-foreground">
          {table.getFilteredSelectedRowModel().rows.length} trong{" "}
          {table.getFilteredRowModel().rows.length} dòng được chọn.
        </div>
        <Paging
          page={page}
          totalPage={totalPage}
          onNavigateBack={() =>
            router.push(`/supplier?page=${Number(page) - 1}${stringToFilter}`)
          }
          onNavigateNext={() =>
            router.push(`/supplier?page=${Number(page) + 1}${stringToFilter}`)
          }
          onPageSelect={(selectedPage) =>
            router.push(`/supplier?page=${selectedPage}${stringToFilter}`)
          }
          onNavigateFirst={() =>
            router.push(`/supplier?page=${1}${stringToFilter}`)
          }
          onNavigateLast={() =>
            router.push(`/supplier?page=${totalPage}${stringToFilter}`)
          }
        />
      </div>
    </div>
  );
}
