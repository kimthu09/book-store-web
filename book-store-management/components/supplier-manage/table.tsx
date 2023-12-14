"use client";

import * as React from "react";
import {
  CaretSortIcon,
  ChevronDownIcon,
  DotsHorizontalIcon,
} from "@radix-ui/react-icons";
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
  DropdownMenuItem,
  DropdownMenuSeparator,
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
// import FilterSheet from "./filter-sheet";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogOverlay,
  DialogTitle,
  DialogTrigger,
} from "../ui/dialog";
import { useState } from "react";
import { Input } from "../ui/input";
import { ExportSupplierList } from "./excel-export";
import { Label } from "../ui/label";
import { RadioGroup, RadioGroupItem } from "../ui/radio-group";
import Link from "next/link";
import Paging from "../paging";
import { useRouter, useSearchParams } from "next/navigation";

// const data: Supplier[] = suppliers;

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
  // {
  //   id: "actions",
  //   enableHiding: false,
  //   cell: ({ row }) => {
  //     const supplier = row.original;
  //     return (
  //       <div className="flex justify-end">
  //         <Link
  //           href={{
  //             pathname: "supplier/detail",
  //             query: {
  //               id: supplier.id,
  //             },
  //           }}
  //         >
  //           <Button variant="ghost" className="h-8 w-8 p-0">
  //             <DotsHorizontalIcon className="h-4 w-4" />
  //           </Button>
  //         </Link>
  //       </div>
  //     );
  //   },
  //   size: 1,
  //   maxSize: 2,
  // },
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
  return (
    <div className="w-full">
      <div className="flex items-center py-4 gap-2">
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="outline" className="w-40 justify-between">
              Chọn thao tác <ChevronDownIcon className="ml-2 h-4 w-4" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent className="max-h-44 w-40">
            <Dialog>
              <DialogTrigger asChild>
                <DropdownMenuItem onSelect={(e) => e.preventDefault()}>
                  Xuất danh sách
                </DropdownMenuItem>
              </DialogTrigger>
              <DialogOverlay>
                <DialogContent className="p-0">
                  <DialogTitle className="p-6 pb-0">
                    Xuất file danh sách nhà cung cấp
                  </DialogTitle>
                  <div className="flex flex-col border-y-[1px] p-6 gap-4">
                    <Label>Giới hạn kết quả xuất</Label>
                    <RadioGroup
                      defaultValue="all"
                      onValueChange={(e: string) => setExportOption(e)}
                    >
                      <div className="flex items-center space-x-2">
                        <RadioGroupItem value="all" id="r1" />
                        <Label htmlFor="r1" className="font-normal">
                          Tất cả các nhà cung cấp
                        </Label>
                      </div>
                      <div className="flex items-center space-x-2">
                        <RadioGroupItem value="comfortable" id="r2" />
                        <Label className="font-normal" htmlFor="r2">
                          Các nhà cung cấp được chọn
                        </Label>
                      </div>
                    </RadioGroup>
                  </div>

                  <DialogClose className="ml-auto p-6 pt-0">
                    <div className="flex gap-4">
                      <Button type="button" variant={"outline"}>
                        Thoát
                      </Button>

                      <Button type="button" onClick={() => handleExport()}>
                        Hoàn tất
                      </Button>
                    </div>
                  </DialogClose>
                </DialogContent>
              </DialogOverlay>
            </Dialog>

            <Dialog>
              <DialogTrigger asChild>
                <DropdownMenuItem onSelect={(e) => e.preventDefault()}>
                  Chuyển trạng thái
                </DropdownMenuItem>
              </DialogTrigger>
              <DialogOverlay>
                <DialogContent className="p-0">
                  <DialogTitle className="p-6 pb-0">
                    Chuyển trạng thái nhà cung cấp
                  </DialogTitle>
                  <div className="flex flex-col border-y-[1px] p-6">
                    <p>Chọn trạng thái muốn chuyển</p>
                    <div className="mt-4 flex-1 ">
                      {/* <CategoryList
                        category={category}
                        setCategory={setCategory}
                      /> */}
                    </div>
                  </div>

                  <DialogClose className="ml-auto p-6 pt-0">
                    <Button type="submit">Hoàn tất</Button>
                  </DialogClose>
                </DialogContent>
              </DialogOverlay>
            </Dialog>
          </DropdownMenuContent>
        </DropdownMenu>

        <div className="flex-1">
          <Input
            placeholder="Tìm kiếm nhà cung cấp"
            value={(table.getColumn("name")?.getFilterValue() as string) ?? ""}
            onChange={(event) =>
              table.getColumn("name")?.setFilterValue(event.target.value)
            }
          />
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

        <div className="ml-auto">{/* <FilterSheet /> */}</div>
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
            router.push(`/supplier?page=${Number(page) - 1}`)
          }
          onNavigateNext={() =>
            router.push(`/supplier?page=${Number(page) + 1}`)
          }
          onPageSelect={(selectedPage) =>
            router.push(`/supplier?page=${selectedPage}`)
          }
        />
      </div>
    </div>
  );
}
