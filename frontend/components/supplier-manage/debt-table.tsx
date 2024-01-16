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
import { SupplierDebt } from "@/types";
import { useState } from "react";
import getSupplierDebt from "@/lib/supplier/getSupplierDebt";
import Loading from "../loading";
import Paging, { PagingProps } from "../paging";
import ExportDialog from "./export-dialog";
import getAllSupplierDebt from "@/lib/supplier/getAllSupplierDebt";
import { toast } from "../ui/use-toast";
import { ExportDebtNote } from "./export-debt-note";
import { useLoading } from "@/hooks/loading-context";
import TableSkeleton from "../skeleton/table-skeleton";

export const columns: ColumnDef<SupplierDebt>[] = [
  {
    accessorKey: "createdAt",
    accessorFn: (row) => new Date(row.createdAt).toLocaleDateString("vi-VN"),
    header: ({ column }) => {
      return (
        <div className="flex justify-end max-w-[8rem]">
          <Button
            className="p-1"
            variant="ghost"
            onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
          >
            <span className="font-semibold">Ngày tạo</span>

            <CaretSortIcon className="ml-2 h-4 w-4" />
          </Button>
        </div>
      );
    },
    cell: ({ row }) => (
      <div className="leading-6 flex justify-end max-w-[6rem]">
        {row.getValue("createdAt")}
      </div>
    ),
  },
  {
    accessorKey: "id",
    header: () => {
      return <span className="font-semibold ">Mã phiếu</span>;
    },
    cell: ({ row }) => <div className="leading-6">{row.getValue("id")}</div>,
  },
  {
    accessorKey: "qty",
    header: ({ column }) => (
      <div className="flex justify-end">
        <Button
          variant={"ghost"}
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
          className="p-1"
        >
          <span className="font-semibold">Giá trị</span>

          <CaretSortIcon className="h-4 w-4" />
        </Button>
      </div>
    ),
    cell: ({ row }) => {
      const amount = parseFloat(row.getValue("qty"));

      // Format the amount as a dollar amount
      const formatted = new Intl.NumberFormat("vi-VN", {
        style: "currency",
        currency: "VND",
      }).format(amount);

      return <div className="text-right font-medium">{formatted}</div>;
    },
  },
  {
    accessorKey: "qtyLeft",
    header: ({ column }) => (
      <div className="flex justify-end">
        <Button
          variant={"ghost"}
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
          className="p-1"
        >
          <span className="font-semibold">Còn lại</span>

          <CaretSortIcon className="h-4 w-4" />
        </Button>
      </div>
    ),
    cell: ({ row }) => {
      const amount = parseFloat(row.getValue("qtyLeft"));

      // Format the amount as a dollar amount
      const formatted = new Intl.NumberFormat("vi-VN", {
        style: "currency",
        currency: "VND",
      }).format(amount);

      return <div className="text-right font-medium">{formatted}</div>;
    },
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
    accessorKey: "type",
    header: () => {
      return (
        <div className="font-semibold whitespace-nowrap flex justify-center">
          Trạng thái
        </div>
      );
    },
    cell: ({ row }) => {
      const status = row.getValue("type");
      return (
        <div className="flex justify-center min-w-0">
          <div
            className={`leading-5 rounded-full text-center px-2 w-[80%] max-w-[5rem] truncate whitespace-nowrap p-1 ${
              status === "Pay"
                ? "bg-green-100 text-green-700"
                : "bg-blue-200  text-blue-600"
            }`}
          >
            {status === "Pay" ? "Trả nợ" : "Nhập"}
          </div>
        </div>
      );
    },
  },
];
export function DebtTable({
  supplierId,
  pageIndex,
  setPageIndex,
}: {
  supplierId: string;
  pageIndex: number;
  setPageIndex: (page: number) => void;
}) {
  const [sorting, setSorting] = useState<SortingState>([]);
  const [columnFilters, setColumnFilters] = useState<ColumnFiltersState>([]);
  const [columnVisibility, setColumnVisibility] = useState<VisibilityState>({});
  const [rowSelection, setRowSelection] = useState({});
  // const [pageIndex, setPageIndex] = useState(1);
  const {
    mutate: mutate,
    data: debts,
    isLoading,
    isError,
  } = getSupplierDebt({
    idSupplier: supplierId,
    page: pageIndex,
  });
  const data = debts?.data;
  const totalPage = Math.ceil(debts?.paging.total / debts?.paging.limit);
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
  const { showLoading, hideLoading } = useLoading();
  const [exportOption, setExportOption] = useState("all");
  const handleExport = async () => {
    if (exportOption === "all") {
      const debtNoteData: Promise<{
        data: SupplierDebt[];
        paging: PagingProps;
      }> = getAllSupplierDebt({ idSupplier: supplierId });
      showLoading();
      const notesToExport = await debtNoteData;
      hideLoading();
      if (notesToExport.data.length < 1) {
        toast({
          variant: "destructive",
          title: "Có lỗi",
          description: "Không có phiếu nợ nào",
        });
      } else {
        ExportDebtNote(
          notesToExport.data,
          `NCC${supplierId} Danh sách phiếu nợ.xlsx`
        );
      }
    } else {
      if (data.length < 1) {
        toast({
          variant: "destructive",
          title: "Có lỗi",
          description: "Không có phiếu nợ nào",
        });
      } else {
        ExportDebtNote(data, `NCC${supplierId} Danh sách phiếu nợ.xlsx`);
      }
    }
  };
  if (isLoading) {
    return (
      <TableSkeleton
        isHasExtensionAction={false}
        isHasFilter={false}
        isHasSearch={false}
        isHasChooseVisibleRow={false}
        isHasCheckBox={false}
        isHasPaging={true}
        numberRow={4}
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
      <div className="flex flex-col">
        <ExportDialog
          handleExport={handleExport}
          setExportOption={setExportOption}
          isImport={false}
        />
        <div className="flex items-center gap-2"></div>
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
                    Chưa có phiếu nợ nào.
                  </TableCell>
                </TableRow>
              )}
            </TableBody>
          </Table>
        </div>
        <div className="flex items-center justify-end space-x-2 py-4">
          <Paging
            page={pageIndex.toString()}
            onNavigateNext={() => setPageIndex(pageIndex + 1)}
            onNavigateBack={() => setPageIndex(pageIndex - 1)}
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
