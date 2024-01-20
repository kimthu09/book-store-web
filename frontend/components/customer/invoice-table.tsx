"use client";

import { CaretSortIcon } from "@radix-ui/react-icons";
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
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { CustomerInvoice, ImportNote, StatusNote } from "@/types";
import { useState } from "react";
import Paging, { PagingProps } from "../paging";
import getSupplierImportNote from "@/lib/supplier/getSupplierImportNote";
import Loading from "../loading";
import { GiShamrock } from "react-icons/gi";
import { toast } from "../ui/use-toast";
import getCustomerInvoice from "@/lib/customer/getCustomerInvoice";
import { Label } from "../ui/label";
import TableSkeleton from "../skeleton/table-skeleton";

export const columns: ColumnDef<CustomerInvoice>[] = [
  {
    accessorKey: "id",
    header: () => {
      return <span className="font-semibold ">Mã phiếu</span>;
    },
    cell: ({ row }) => <div className="leading-6">{row.getValue("id")}</div>,
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
    size: 1,
  },
  {
    accessorKey: "createdBy",
    accessorFn: (row) => row.createdBy.name,
    header: () => {
      return <div className="font-semibold flex justify-center">Người tạo</div>;
    },
    cell: ({ row }) => (
      <div className="leading-6 text-center">{row.getValue("createdBy")}</div>
    ),
  },
  {
    accessorKey: "totalPrice",
    header: ({ column }) => (
      <div className="flex justify-end">
        <Button
          variant={"ghost"}
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
          className="p-1"
        >
          <span className="font-semibold">Tổng tiền</span>

          <CaretSortIcon className="h-4 w-4" />
        </Button>
      </div>
    ),
    cell: ({ row }) => {
      const amount = parseFloat(row.getValue("totalPrice"));

      // Format the amount as a dollar amount
      const formatted = new Intl.NumberFormat("vi-VN", {
        style: "currency",
        currency: "VND",
      }).format(amount);

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

      return (
        <div className="text-right font-medium flex flex-col items-end gap-1">
          -{formatted}
          <div className="flex items-center gap-1 text-rose-700">
            -{row.original.pointUse.toLocaleString("vi-VN")}{" "}
            <GiShamrock className="h-5 w-5" />
          </div>
        </div>
      );
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
];
export function InvoiceTable({ customerId }: { customerId: string }) {
  const [pageIndex, setPageIndex] = useState(1);

  const {
    data: importNotes,
    isLoading,
    isError,
  } = getCustomerInvoice({
    idCustomer: customerId,
    page: pageIndex,
  });
  const data = importNotes?.data;
  const totalPage = Math.ceil(
    importNotes?.paging.total / importNotes?.paging.limit
  );

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

  if (isLoading) {
    return (
      <TableSkeleton
        isHasExtensionAction={false}
        isHasFilter={true}
        isHasSearch={true}
        isHasChooseVisibleRow={true}
        isHasCheckBox={false}
        isHasPaging={true}
        numberRow={5}
        cells={[
          {
            percent: 1,
          },
          {
            percent: 2,
          },
          {
            percent: 1,
          },
          {
            percent: 1,
          },
          {
            percent: 1,
          },
          {
            percent: 1,
          },
        ]}
      ></TableSkeleton>
    );
  } else {
    return (
      <div className="flex flex-col gap-4">
        <div className="flex items-center gap-2">
          <Label>Lịch sử mua hàng</Label>
        </div>
        <div className="rounded-md border overflow-x-auto flex-1 min-w-full max-w-[20vw]">
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
                  <TableRow key={row.id}>
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
                    Chưa có đơn nhập nào.
                  </TableCell>
                </TableRow>
              )}
            </TableBody>
          </Table>
        </div>
        <div className="flex items-center justify-end space-x-2 py-4">
          <Paging
            page={pageIndex.toString()}
            onNavigateNext={() => setPageIndex((prev) => prev + 1)}
            onNavigateBack={() => setPageIndex((prev) => prev - 1)}
            totalPage={totalPage}
            onPageSelect={(selectedPage) => {
              setPageIndex(+selectedPage);
            }}
            onNavigateFirst={() => setPageIndex(1)}
            onNavigateLast={() => setPageIndex(totalPage)}
          />
        </div>
      </div>
    );
  }
}
