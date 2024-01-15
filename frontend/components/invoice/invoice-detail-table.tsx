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
import { Checkbox } from "@/components/ui/checkbox";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";

import { useState } from "react";
import { Input } from "../ui/input";
import { toVND } from "@/lib/utils";
import { useRouter } from "next/navigation";
import { InvoiceDetailProps } from "@/app/invoice/[invoiceId]/page";

export const columns: ColumnDef<InvoiceDetailProps>[] = [
  {
    id: "stt",
    header: ({ table }) => (
      <div className="flex justify-center font-semibold">STT</div>
    ),
    cell: ({ row }) => (
      <div className="flex justify-center">{row.index + 1}</div>
    ),
    enableSorting: false,
    enableHiding: false,
    size: 4,
  },
  {
    accessorKey: "id",
    accessorFn: (row) => row.book.id,
    header: () => {
      return <span className="font-semibold">Mã sách</span>;
    },
    cell: ({ row }) => <div className="leading-6">{row.original.book.id}</div>,
    size: 4,
  },
  {
    accessorKey: "name",
    header: () => {
      return <span className="font-semibold ">Tên sách</span>;
    },
    cell: ({ row }) => (
      <div className="leading-6 flex flex-col 2xl:max-w-[20rem] xl:max-w-[14rem] max-w-[8rem]">
        {row.original.book.name}
      </div>
    ),
    size: 4,
  },
  {
    accessorKey: "amountImport",
    header: ({ column }) => (
      <div className="flex justify-end whitespace-normal">
        <span className="font-semibold">SL</span>
      </div>
    ),
    cell: ({ row }) => {
      return (
        <div className="text-right font-medium">
          {row.original.qty.toLocaleString("vi-VN")}
        </div>
      );
    },
    size: 4,
  },
  {
    accessorKey: "price",
    header: ({ column }) => (
      <div className="flex justify-end whitespace-normal">
        <span className="font-semibold">Đơn giá</span>
      </div>
    ),
    cell: ({ row }) => {
      const amount = row.original.unitPrice;

      return <div className="text-right font-medium">{toVND(amount)}</div>;
    },
    size: 4,
  },
  {
    accessorKey: "totalUnit",
    accessorFn: (row) => row.qty * row.unitPrice,
    header: ({ column }) => (
      <div className="flex justify-end whitespace-normal">
        <span className="font-semibold">Thành tiền</span>
      </div>
    ),
    cell: ({ row }) => {
      return (
        <div className="text-right font-medium">
          {toVND(row.original.qty * row.original.unitPrice)}
        </div>
      );
    },
    size: 4,
  },
];
export function InvoiceDetailTable({
  details: data,
}: {
  details: InvoiceDetailProps[];
}) {
  const router = useRouter();
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
    initialState: {
      pagination: {
        pageSize: 1000,
      },
    },
  });
  return (
    <div className="rounded-md border overflow-x-auto min-w-full max-w-[40vw]">
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
            table.getRowModel().rows.map((row, index) => (
              <TableRow
                key={row.id}
                data-state={row.getIsSelected() && "selected"}
              >
                {row.getVisibleCells().map((cell) => (
                  <TableCell key={cell.id}>
                    {flexRender(cell.column.columnDef.cell, cell.getContext())}
                  </TableCell>
                ))}
              </TableRow>
            ))
          ) : (
            <TableRow>
              <TableCell colSpan={columns.length} className="h-24 text-center">
                Không tìm thấy kết quả.
              </TableCell>
            </TableRow>
          )}
        </TableBody>
      </Table>
    </div>
  );
}
