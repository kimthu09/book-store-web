"use client";

import * as React from "react";
import { CaretSortIcon, ChevronDownIcon } from "@radix-ui/react-icons";
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
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Staff } from "@/types";

import { useEffect, useState } from "react";
import { Input } from "../ui/input";
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "../ui/dropdown-menu";
import Paging from "../paging";
import { useRouter, useSearchParams } from "next/navigation";
import { Popover, PopoverContent, PopoverTrigger } from "../ui/popover";
import { LuFilter } from "react-icons/lu";
import { SubmitHandler, useFieldArray, useForm } from "react-hook-form";
import { Label } from "../ui/label";
import { AiOutlineClose } from "react-icons/ai";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../ui/select";
import StatusList from "../status-list";
import { boolean } from "zod";

function idToName(id: string) {
  if (id === "name") {
    return "Tên";
  } else if (id === "phone") {
    return "Số điện thoại";
  } else if (id === "address") {
    return "Địa chỉ";
  } else if (id === "role") {
    return "Phân quyền";
  } else {
    return id;
  }
}
type FormValues = {
  filters: {
    type: string;
    value: string;
  }[];
};
export const columns: ColumnDef<Staff>[] = [
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
          <span className="font-semibold">Tên nhân viên</span>

          <CaretSortIcon className="ml-2 h-4 w-4" />
        </Button>
      );
    },
    cell: ({ row }) => (
      <div className="capitalize pl-2 leading-6">{row.getValue("name")}</div>
    ),
  },
  {
    accessorKey: "email",
    header: () => {
      return <div className="font-semibold">Email</div>;
    },
    cell: ({ row }) => (
      <div className="lg:max-w-[24rem] max-w-[5rem] truncate">
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
    accessorKey: "address",
    header: () => {
      return <div className="font-semibold">Địa chỉ</div>;
    },
    cell: ({ row }) => (
      <div className="lg:max-w-[16rem] max-w-[3rem] truncate">
        {row.getValue("address")}
      </div>
    ),
  },
  {
    accessorKey: "role",
    accessorFn: (row) => row.role.name,
    header: () => {
      return <span className="font-semibold">Phân quyền</span>;
    },
    cell: ({ row }) => (
      <div className="lg:max-w-[24rem] max-w-[3rem] truncate">
        {row.getValue("role")}
      </div>
    ),
  },
];
export function StaffTable({
  data,
  totalPage,
}: {
  data: Staff[];
  totalPage: number;
}) {
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
  const router = useRouter();
  const searchParams = useSearchParams();
  const page = searchParams.get("page") ?? "1";

  const [latestFilter, setLatestFilter] = useState("");
  const filterValues = [
    { type: "search", name: "Từ khoá" },
    { type: "active", name: "Trạng thái" },
  ];
  const active = searchParams.get("active") ?? undefined;
  const search = searchParams.get("search") ?? undefined;
  let filters = [{ type: "", value: "" }];
  filters.pop();
  if (active) {
    filters = filters.concat({ type: "active", value: active });
  }
  if (search) {
    filters = filters.concat({ type: "search", value: search });
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

  const onSubmit: SubmitHandler<FormValues> = async (data) => {
    console.log(data);
    let search = "";
    let active = "";
    data.filters.forEach((item) => {
      if (item.type === "active") {
        active = `&active=${item.value}`;
      } else if (item.type === "search") {
        search = `&search=${item.value}`;
      }
    });
    router.push(`/staff?page=${Number(page)}${active}${search}`);
  };
  const [openFilter, setOpenFilter] = useState(false);
  const [status, setStatus] = useState<boolean>();
  const handleSetStatus = (value: boolean) => {
    setStatus(value);
    const index = fields.findIndex((item) => item.type === "active");
    if (index > -1) {
      update(index, { type: "active", value: value.toString() });
    } else {
      append({ type: "createdBy", value: value.toString() });
    }
  };
  const displayStatus = {
    trueText: "Đang hoạt động",
    falseText: "Ngừng hoạt động",
  };
  useEffect(() => {
    if (active === "true") {
      setStatus(true);
    } else if (active === "false") {
      setStatus(false);
    }
  }, [active]);
  return (
    <div className="w-full">
      <div className="flex items-start py-4 gap-2">
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
                            <StatusList
                              status={status}
                              setStatus={handleSetStatus}
                              display={displayStatus}
                            />
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
                className="flex-1"
                placeholder="Tìm kiếm nhân viên"
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
                    {item.type === "active"
                      ? item.value === "true"
                        ? displayStatus.trueText
                        : displayStatus.falseText
                      : item.value}
                  </span>
                </div>
              );
            })}
          </div>
        </div>
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="outline">
              Cột hiển thị <ChevronDownIcon className="ml-2 h-4 w-4" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent>
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
              table.getRowModel().rows.map((row, index) => (
                <TableRow
                  key={row.id}
                  data-state={row.getIsSelected() && "selected"}
                >
                  {row.getVisibleCells().map((cell) => (
                    <TableCell key={cell.id}>
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
          onNavigateBack={() => router.push(`/staff?page=${Number(page) - 1}`)}
          onNavigateNext={() => router.push(`/staff?page=${Number(page) + 1}`)}
          onPageSelect={(selectedPage) =>
            router.push(`/staff?page=${selectedPage}`)
          }
          onNavigateFirst={() => router.push(`/supplier?page=${1}`)}
          onNavigateLast={() => router.push(`/supplier?page=${totalPage}`)}
        />
      </div>
    </div>
  );
}
