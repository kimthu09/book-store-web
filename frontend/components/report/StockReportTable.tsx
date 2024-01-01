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
import { SaleReportDetail, StockReportDetail } from "@/types"


export const columns: ColumnDef<StockReportDetail>[] = [
    {
        accessorKey: "id",
        accessorFn: (row) => row.book.id,
        header: "Id",
        cell: ({ row }) => (
            <div className="capitalize">{row.original.book.id}</div>
        ),
    },
    {
        accessorKey: "name",
        accessorFn: (row) => row.book.name,
        header: "Tên sách",
        cell: ({ row }) => <div >{row.original.book.name}</div>,
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
                    <span className="font-semibold">Đầu kỳ</span>

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
                    <span className="font-semibold">Thêm</span>

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
                    <span className="font-semibold">Sửa</span>

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
                    <span className="font-semibold">Đã bán</span>

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
                    <span className="font-semibold">Cuối kỳ</span>

                    <CaretSortIcon className="h-4 w-4" />
                </Button>
            </div>
        ),
        cell: ({ row }) => {
            const amount = parseFloat(row.getValue("final"));

            return <div className="text-right font-medium">{amount}</div>;
        },
    }
]

export function StockReportTable(
    { data }: { data: StockReportDetail[] }
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
