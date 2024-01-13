"use client";

import * as React from "react";
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
import { StockReport, StockReportDetail } from "@/types";
import { unknown } from "zod";

export const columns: ColumnDef<StockReportDetail>[] = [
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
    header: ({ column }) => (
      <div className="flex-1 justify-start">
        <Button
          variant={"ghost"}
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
          className="p-1"
        >
          <span className="font-semibold">Id</span>

          <CaretSortIcon className="h-4 w-4" />
        </Button>
      </div>
    ),
    cell: ({ row }) => <div className="capitalize">{row.original.book.id}</div>,
  },
  {
    accessorKey: "name",
    accessorFn: (row) => row.book.name,
    header: ({ column }) => (
      <div className="flex-[5] justify-start">
        <Button
          variant={"ghost"}
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
          className="p-1"
        >
          <span className="font-semibold">Tên sách</span>

          <CaretSortIcon className="h-4 w-4" />
        </Button>
      </div>
    ),
    cell: ({ row }) => <div>{row.original.book.name}</div>,
  },
  {
    accessorKey: "initial",
    header: ({ column }) => (
      <div className="flex justify-end">
        <Button
          variant={"ghost"}
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
          className="p-1"
        >
          <span className="font-semibold">Tồn đầu</span>

          <CaretSortIcon className="h-4 w-4" />
        </Button>
      </div>
    ),
    cell: ({ row }) => {
      const amount = parseFloat(row.getValue("initial"));

      return <div className="text-right font-medium">{amount}</div>;
    },
  },
  {
    accessorKey: "import",
    header: ({ column }) => (
      <div className="flex justify-end">
        <Button
          variant={"ghost"}
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
          className="p-1"
        >
          <span className="font-semibold">Nhập</span>

          <CaretSortIcon className="h-4 w-4" />
        </Button>
      </div>
    ),
    cell: ({ row }) => {
      const amount = parseFloat(row.getValue("import"));

      return <div className="text-right font-medium">{amount}</div>;
    },
  },
  {
    accessorKey: "modify",
    header: ({ column }) => (
      <div className="flex justify-end">
        <Button
          variant={"ghost"}
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
          className="p-1"
        >
          <span className="font-semibold">Kiểm kho</span>

          <CaretSortIcon className="h-4 w-4" />
        </Button>
      </div>
    ),
    cell: ({ row }) => {
      const amount = parseFloat(row.getValue("modify"));

      return <div className="text-right font-medium">{amount}</div>;
    },
  },
  {
    accessorKey: "sell",
    header: ({ column }) => (
      <div className="flex justify-end">
        <Button
          variant={"ghost"}
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
          className="p-1"
        >
          <span className="font-semibold">Bán</span>

          <CaretSortIcon className="h-4 w-4" />
        </Button>
      </div>
    ),
    cell: ({ row }) => {
      const amount = parseFloat(row.getValue("sell"));

      return <div className="text-right font-medium">{amount}</div>;
    },
  },
  {
    accessorKey: "final",
    header: ({ column }) => (
      <div className="flex justify-end">
        <Button
          variant={"ghost"}
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
          className="p-1"
        >
          <span className="font-semibold">Tồn cuối</span>

          <CaretSortIcon className="h-4 w-4" />
        </Button>
      </div>
    ),
    cell: ({ row }) => {
      const amount = parseFloat(row.getValue("final"));

      return <div className="text-right font-medium">{amount}</div>;
    },
  },
];

export function StockReportTable({
  report,
  data,
}: {
  report: StockReport | undefined;
  data: StockReportDetail[];
}) {
  const [sorting, setSorting] = React.useState<SortingState>([]);
  const [columnFilters, setColumnFilters] = React.useState<ColumnFiltersState>(
    []
  );
  const [columnVisibility, setColumnVisibility] =
    React.useState<VisibilityState>({});

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
    state: {
      sorting,
      columnFilters,
      columnVisibility,
    },
    initialState: {
      pagination: {
        pageSize: data.length,
      },
    },
  });

  return (
    <div className="w-full">
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
            {report != undefined ? (
              <TableRow key={"subHeaderRow"}>
                <TableHead key={"header-final"} className="col-span-2">
                  {flexRender(
                    <div className="font-semibold">Tổng cộng:</div>,
                    unknown
                  )}
                </TableHead>
                <TableHead key={"header-empty-1"}></TableHead>
                <TableHead key={"header-empty-2"}></TableHead>
                <TableHead key={"header-initial"}>
                  {flexRender(
                    <div className="text-right font-semibold">
                      {report.initial}
                    </div>,
                    unknown
                  )}
                </TableHead>
                <TableHead key={"header-import"}>
                  {flexRender(
                    <div className="text-right font-semibold">
                      {report.import}
                    </div>,
                    unknown
                  )}
                </TableHead>
                <TableHead key={"header-modify"}>
                  {flexRender(
                    <div className="text-right font-semibold">
                      {report.modify}
                    </div>,
                    unknown
                  )}
                </TableHead>
                <TableHead key={"header-sell"}>
                  {flexRender(
                    <div className="text-right font-semibold">
                      {report.sell}
                    </div>,
                    unknown
                  )}
                </TableHead>
                <TableHead key={"header-final"}>
                  {flexRender(
                    <div className="text-right font-semibold">
                      {report.final}
                    </div>,
                    unknown
                  )}
                </TableHead>
              </TableRow>
            ) : null}
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
                  Không có kết quả.
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </div>
    </div>
  );
}
