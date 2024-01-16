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
import { ImportNote, StatusNote } from "@/types";
import { useState } from "react";
import Paging, { PagingProps } from "../paging";
import getSupplierImportNote from "@/lib/supplier/getSupplierImportNote";
import Loading from "../loading";
import ExportDialog from "./export-dialog";
import { ExportImportNote } from "./export-import-note";
import getAllSupplierNote from "@/lib/supplier/getAllSupplierNote";
import { toast } from "../ui/use-toast";
import { useLoading } from "@/hooks/loading-context";
import TableSkeleton from "../skeleton/table-skeleton";

export const columns: ColumnDef<ImportNote>[] = [
  {
    accessorKey: "createAt",
    accessorFn: (row) => {
      return new Date(row.createdAt).toLocaleDateString("vi-VN");
    },
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
        {row.getValue("createAt")}
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
    accessorKey: "status",
    header: () => {
      return (
        <div className="font-semibold whitespace-nowrap flex justify-center">
          Trạng thái
        </div>
      );
    },
    cell: ({ row }) => {
      const status = row.getValue("status");
      return (
        <div className="flex justify-center min-w-0">
          <div
            className={`leading-5 rounded-full text-center px-2 w-[80%] max-w-[6rem] p-1 truncate whitespace-nowrap ${
              status === StatusNote.Done
                ? "bg-green-100 text-green-700 "
                : status === StatusNote.Inprogress
                ? "bg-blue-200 text-blue-600 "
                : "bg-rose-100 text-rose-600 "
            }`}
          >
            {status === StatusNote.Done
              ? "Hoàn thành"
              : status === StatusNote.Inprogress
              ? "Đang xử lý"
              : "Đã hủy"}
          </div>
        </div>
      );
    },
  },
];
export function ImportTable({ supplierId }: { supplierId: string }) {
  const [pageIndex, setPageIndex] = useState(1);

  const {
    data: importNotes,
    isLoading,
    isError,
  } = getSupplierImportNote({
    idSupplier: supplierId,
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

  const [exportOption, setExportOption] = useState("all");
  const { showLoading, hideLoading } = useLoading();
  const handleExport = async () => {
    if (exportOption === "all") {
      const importNoteData: Promise<{
        data: ImportNote[];
        paging: PagingProps;
      }> = getAllSupplierNote({ idSupplier: supplierId });
      showLoading();
      const notesToExport = await importNoteData;
      hideLoading();
      if (notesToExport.data.length < 1) {
        toast({
          variant: "destructive",
          title: "Có lỗi",
          description: "Không có phiếu nhập nào",
        });
      } else {
        ExportImportNote(
          notesToExport.data,
          `NCC${supplierId} Danh sách phiếu nhập.xlsx`
        );
      }
    } else {
      if (data.length < 1) {
        toast({
          variant: "destructive",
          title: "Có lỗi",
          description: "Không có phiếu nhập nào",
        });
      } else {
        ExportImportNote(data, `NCC${supplierId} Danh sách phiếu nhập.xlsx`);
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
        ]}
      ></TableSkeleton>
    );
  } else {
    return (
      <div className="flex flex-col">
        <ExportDialog
          handleExport={handleExport}
          setExportOption={setExportOption}
          isImport
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
