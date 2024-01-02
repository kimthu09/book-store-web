"use client"

import * as React from "react"
import {
    CaretSortIcon,
} from "@radix-ui/react-icons"
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
} from "@tanstack/react-table"

import { Button } from "@/components/ui/button"
import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table"
import { DebtReportDetail } from "@/types"


export const columns: ColumnDef<DebtReportDetail>[] = [
    {
        accessorKey: "id",
        accessorFn: (row) => row.supplier.id,
        header: "Id",
        cell: ({ row }) => (
            <div className="capitalize">{row.original.supplier.id}</div>
        ),
    },
    {
        accessorKey: "name",
        accessorFn: (row) => row.supplier.name,
        header: "Tên nhà cung cấp",
        cell: ({ row }) => <div >{row.original.supplier.name}</div>,
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
                    <span className="font-semibold">Nợ đầu kì</span>

                    <CaretSortIcon className="h-4 w-4" />
                </Button>
            </div>
        ),
        cell: ({ row }) => {
            const amount = parseFloat(row.getValue("initial"));

            // Format the amount as a dollar amount
            const formatted = new Intl.NumberFormat("vi-VN", {
                style: "currency",
                currency: "VND",
            }).format(amount);

            return <div className="text-right font-medium">{formatted}</div>;
        },
    },
    {
        accessorKey: "pay",
        header: ({ column }) => (
            <div className="flex justify-end">
                <Button
                    variant={"ghost"}
                    onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
                    className="p-1"
                >
                    <span className="font-semibold">Chênh lệch</span>

                    <CaretSortIcon className="h-4 w-4" />
                </Button>
            </div>
        ),
        cell: ({ row }) => {
            const amount = parseFloat(row.getValue("pay"));

            // Format the amount as a dollar amount
            const formatted = new Intl.NumberFormat("vi-VN", {
                style: "currency",
                currency: "VND",
            }).format(amount);

            return <div className="text-right font-medium">{formatted}</div>;
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
                    <span className="font-semibold">Nợ cuối kì</span>

                    <CaretSortIcon className="h-4 w-4" />
                </Button>
            </div>
        ),
        cell: ({ row }) => {
            const amount = parseFloat(row.getValue("final"));

            // Format the amount as a dollar amount
            const formatted = new Intl.NumberFormat("vi-VN", {
                style: "currency",
                currency: "VND",
            }).format(amount);

            return <div className="text-right font-medium">{formatted}</div>;
        },
    },
]

export function DebtReportTable(
    { data }: { data: DebtReportDetail[] }
) {
    const [sorting, setSorting] = React.useState<SortingState>([])
    const [columnFilters, setColumnFilters] = React.useState<ColumnFiltersState>(
        []
    )
    const [columnVisibility, setColumnVisibility] =
        React.useState<VisibilityState>({})

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
    })

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
                                    )
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
                                    Không có kết quả.
                                </TableCell>
                            </TableRow>
                        )}
                    </TableBody>
                </Table>
            </div>
        </div>
    )
}
