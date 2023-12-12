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
  DropdownMenuLabel,
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
import { Book } from "@/types";
import FilterSheet from "./filter-sheet";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogOverlay,
  DialogTitle,
  DialogTrigger,
} from "../ui/dialog";
import { useState } from "react";
import CategoryList from "../category-list";
import Link from "next/link";
import { ExportBookList } from "../excel-export";
import { LuChevronLeft, LuChevronRight } from "react-icons/lu";
import deleteBook from "@/lib/deleteBook";
import { useRouter, useSearchParams } from "next/navigation";
import { toast } from "../ui/use-toast";

function idToName(id: string) {
  if (id === "name") {
    return "Tên";
  } else if (id === "price") {
    return "Giá";
  } else if (id === "status") {
    return "Trạng thái";
  }
  return id;
}

export function BookTable({
  data,
  totalPage,
}: {
  data: Book[];
  totalPage: number;
}) {
  // const data: Book[] = books;
  const router = useRouter();
  const searchParams = useSearchParams();
  const page = searchParams.get("page") ?? "1";

  const [sorting, setSorting] = useState<SortingState>([]);
  const [columnFilters, setColumnFilters] = useState<ColumnFiltersState>([]);
  const [columnVisibility, setColumnVisibility] = useState<VisibilityState>({});
  const [rowSelection, setRowSelection] = useState({});
  const [category, setCategory] = useState("");

  const columns: ColumnDef<Book>[] = [
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
            <span className="font-semibold">Tên sản phẩm</span>

            <CaretSortIcon className="ml-1 h-4 w-4" />
          </Button>
        );
      },
      cell: ({ row }) => (
        <div className="capitalize">{row.getValue("name")}</div>
      ),
    },
    // {
    //   accessorKey: "publisherId",
    //   header: () => {
    //     return <span className="font-semibold">NXB</span>;
    //   },
    //   cell: ({ row }) => (
    //     <div className="capitalize">{row.getValue("publisherId")}</div>
    //   ),
    // },
    // {
    //   accessorKey: "sellPrice",
    //   header: ({ column }) => (
    //     <Button
    //       className="p-2"
    //       variant={"ghost"}
    //       onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
    //     >
    //       <span className="font-semibold">Giá</span>

    //       <CaretSortIcon className="ml-1 h-4 w-4" />
    //     </Button>
    //   ),
    //   cell: ({ row }) => {
    //     const amount = parseFloat(row.getValue("sellPrice"));

    //     // Format the amount as a dollar amount
    //     const formatted = new Intl.NumberFormat("vi-VN", {
    //       style: "currency",
    //       currency: "VND",
    //     }).format(amount);

    //     return <div className="text-left font-medium">{formatted}</div>;
    //   },
    // },
    // {
    //   accessorKey: "quantity",
    //   header: () => {
    //     return <div className="font-semibold flex justify-end">Số lượng</div>;
    //   },
    //   cell: ({ row }) => (
    //     <div className="text-right">{row.getValue("quantity")}</div>
    //   ),
    // },
    {
      accessorKey: "isActive",
      header: () => {
        return (
          <div className="font-semibold flex justify-center">Trạng thái</div>
        );
      },
      cell: ({ row }) => {
        const status = row.getValue("isActive");
        return (
          <div
            className={`lg:max-w-[16rem] max-w-[3rem] truncate ${
              status === 1 ? "text-green-600" : "text-red-600"
            } text-center`}
          >
            {status === 1 ? "Đang bán" : "Ngừng bán"}
          </div>
        );
      },
    },
    {
      id: "actions",
      enableHiding: false,
      cell: ({ row }) => {
        const book = row.original;

        return (
          <div className="flex justify-end">
            <DropdownMenu>
              <DropdownMenuTrigger asChild>
                <Button variant="ghost" className="h-8 w-8 p-0">
                  <span className="sr-only">Open menu</span>
                  <DotsHorizontalIcon className="h-4 w-4" />
                </Button>
              </DropdownMenuTrigger>
              <DropdownMenuContent align="end">
                <DropdownMenuLabel></DropdownMenuLabel>
                <DropdownMenuItem
                  onClick={() => navigator.clipboard.writeText(book.id)}
                >
                  Sao chép mã sách
                </DropdownMenuItem>
                <DropdownMenuSeparator />
                {/* <DropdownMenuItem>
                <Link
                  href={{
                    pathname: "books/edit",
                    query: {
                      id: book.id,
                    },
                  }}
                >
                  Chỉnh sửa sách
                </Link>
              </DropdownMenuItem> */}
                <DropdownMenuItem
                  onClick={async () => {
                    const response: Promise<any> = deleteBook(book.id);
                    const responseData = await response;
                    console.log(responseData);
                    if (responseData.data) {
                      toast({
                        title: "Thành công",
                        description: "Sản phẩm đã có trạng thái ngừng bán",
                      });
                      router.refresh();
                    }
                  }}
                >
                  Ngừng bán
                </DropdownMenuItem>
              </DropdownMenuContent>
            </DropdownMenu>
          </div>
        );
      },
    },
  ];

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

  return (
    <div className="w-full">
      <div className="flex items-center py-4 gap-2">
        <div>
          {/* <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="outline" className="w-40 justify-between">
                Chọn thao tác <ChevronDownIcon className="ml-2 h-4 w-4" />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent className="max-h-44 w-40">
              <DropdownMenuItem
                onSelect={() => {
                  if (table.getFilteredSelectedRowModel().rows.length < 1) {
                    //TODO: show notification
                  } else {
                    const values = table
                      .getFilteredSelectedRowModel()
                      .rows.map((row) => row.original);
                    ExportBookList(values, "BookList.xlsx");
                  }
                }}
              >
                Xuất danh sách
              </DropdownMenuItem>
              <Dialog>
                <DialogTrigger asChild>
                  <DropdownMenuItem onSelect={(e) => e.preventDefault()}>
                    Chuyển danh mục
                  </DropdownMenuItem>
                </DialogTrigger>
                <DialogOverlay>
                  <DialogContent className="p-0">
                    <DialogTitle className="p-6 pb-0">
                      Chuyển mặt hàng tới danh mục khác
                    </DialogTitle>
                    <div className="flex flex-col border-y-[1px] p-6">
                      <p>Chọn 1 danh mục muốn chuyển tới</p>
                      <div className="mt-4 flex-1 ">
                        <CategoryList
                          category={category}
                          setCategory={setCategory}
                        />
                      </div>
                    </div>

                    <DialogClose className="ml-auto p-6 pt-0">
                      <Button type="submit">Hoàn tất</Button>
                    </DialogClose>
                  </DialogContent>
                </DialogOverlay>
              </Dialog>

              <Dialog>
                <DialogTrigger asChild>
                  <DropdownMenuItem onSelect={(e) => e.preventDefault()}>
                    Ngừng bán
                  </DropdownMenuItem>
                </DialogTrigger>
                <DialogContent>
                  <DialogTitle>Ngừng bán sách được chọn</DialogTitle>
                </DialogContent>
              </Dialog>
            </DropdownMenuContent>
          </DropdownMenu> */}
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
        <div className="space-x-2">
          <Button
            variant="outline"
            size="icon"
            onClick={() => router.push(`/books?page=${Number(page) - 1}`)}
            disabled={Number(page) <= 1}
          >
            <LuChevronLeft className="h-4 w-4" />
          </Button>
          <Button
            variant="outline"
            size="icon"
            onClick={() => router.push(`/books?page=${Number(page) + 1}`)}
            disabled={Number(page) >= totalPage}

            // disabled={!table.getCanNextPage()}
          >
            <LuChevronRight className="h-4 w-4" />
          </Button>
        </div>
      </div>
    </div>
  );
}
