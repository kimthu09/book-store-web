"use client";

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
import { CheckNoteDetail } from "@/types";

import { useState } from "react";
import { toVND } from "@/lib/utils";
import { useRouter } from "next/navigation";

export const columns: ColumnDef<CheckNoteDetail>[] = [
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
    size: 1,
  },
  {
    accessorKey: "id",
    accessorFn: (row) => row.book.id,
    header: () => {
      return <span className="font-semibold">Mã sách</span>;
    },
    cell: ({ row }) => <div className="leading-6">{row.original.book.id}</div>,
    size: 1,
  },
  {
    accessorKey: "name",
    header: () => {
      return <span className="font-semibold">Tên sách</span>;
    },
    cell: ({ row }) => (
      <div className="leading-6 flex flex-col">{row.original.book.name}</div>
    ),
    size: 5,
  },
  {
    accessorKey: "initial",
    header: ({ column }) => (
      <div className="flex justify-end whitespace-normal">
        <span className="font-semibold">Ban đầu</span>
      </div>
    ),
    cell: ({ row }) => {
      return (
        <div className="text-right font-medium">
          {row.original.initial.toLocaleString("vi-VN")}
        </div>
      );
    },
    size: 4,
  },
  {
    accessorKey: "difference",
    header: ({ column }) => (
      <div className="flex justify-end whitespace-normal">
        <span className="font-semibold">Chênh lệch</span>
      </div>
    ),
    cell: ({ row }) => {
      return (
        <div className="text-right font-medium">
          {row.original.difference.toLocaleString("vi-VN")}
        </div>
      );
    },
    size: 4,
  },
  {
    accessorKey: "final",
    header: ({ column }) => (
      <div className="flex justify-end whitespace-normal">
        <span className="font-semibold">Kiểm kê</span>
      </div>
    ),
    cell: ({ row }) => {
      return (
        <div className="text-right font-medium">
          {row.original.final.toLocaleString("vi-VN")}
        </div>
      );
    },
    size: 4,
  },
];
export function CheckDetailTable({
  details: data,
}: {
  details: CheckNoteDetail[];
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
    <div className="rounded-md border overflow-x-auto flex-1 min-w-full max-w-[50vw]">
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
                    {flexRender(cell.column.columnDef.cell, cell.getContext())}
                  </TableCell>
                ))}
              </TableRow>
            ))
          ) : (
            <TableRow>
              <TableCell colSpan={columns.length} className="h-24 text-center">
                Không tìm thấy bản ghi
              </TableCell>
            </TableRow>
          )}
        </TableBody>
      </Table>
    </div>
  );
}
