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
import { Checkbox } from "@/components/ui/checkbox";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Publisher } from "@/types";

import { useState } from "react";
import { Input } from "../ui/input";
import Loading from "../loading";
import Paging from "../paging";
import { useRouter } from "next/navigation";
import getAllPublisher from "@/lib/book/getAllPublisher";
import EditPublisher from "./edit-publisher";
import { FaPen } from "react-icons/fa";
import { useSWRConfig } from "swr";
import { endPoint } from "@/constants";
import { includesRoles } from "@/lib/utils";
import TableSkeleton from "../skeleton/table-skeleton";

export const columns: ColumnDef<Publisher>[] = [
  {
    accessorKey: "name",
    header: ({ column }) => {
      return (
        <Button
          className="p-2"
          variant="ghost"
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
        >
          <span className="font-semibold">Nhà xuất bản</span>

          <CaretSortIcon className="ml-2 h-4 w-4" />
        </Button>
      );
    },
    cell: ({ row }) => (
      <div className="capitalize pl-2 leading-6">{row.getValue("name")}</div>
    ),
  },
  {
    accessorKey: "actions",
    header: () => {
      return <div className="font-semibold flex justify-end">Thao tác</div>;
    },
    cell: ({ row }) => {
      return <></>;
    },
  },
];
export function PublisherTable({
  searchParams,
  currentUser,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
  currentUser:
    | {
        name?: string | null | undefined;
        email?: string | null | undefined;
        image?: string | null | undefined;
      }
    | undefined;
}) {
  const router = useRouter();
  const page = searchParams["page"] ?? "1";
  const { publishers, isLoading, isError } = getAllPublisher({
    page: page.toString(),
  });
  const data = publishers?.data;
  const totalPage = Math.ceil(
    publishers?.paging.total / publishers?.paging.limit
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

  const { mutate } = useSWRConfig();
  const handlePublisherEdited = (name: string) => {
    mutate(`${endPoint}/v1/publishers?page=${page ?? 1}&limit=10`);
  };
  if (isError) return <div>Failed to load</div>;
  if (isLoading) {
    return (
      <TableSkeleton
        isHasExtensionAction={false}
        isHasFilter={false}
        isHasSearch={true}
        isHasChooseVisibleRow={false}
        isHasCheckBox={false}
        isHasPaging={true}
        numberRow={5}
        cells={[
          {
            percent: 5,
          },
          {
            percent: 1,
          },
        ]}
      ></TableSkeleton>
    );
  } else
    return (
      <div className="w-full">
        <div className="flex items-center py-4 gap-2">
          <Input
            placeholder="Tìm kiếm nhà xuất bản"
            value={(table.getColumn("name")?.getFilterValue() as string) ?? ""}
            onChange={(event) =>
              table.getColumn("name")?.setFilterValue(event.target.value)
            }
          />
        </div>
        <div className="rounded-md border">
          <Table>
            <TableHeader>
              {table.getHeaderGroups().map((headerGroup) => (
                <TableRow key={headerGroup.id}>
                  {headerGroup.headers.map((header) => {
                    if (
                      header.id.includes("actions") &&
                      (currentUser ||
                        (currentUser &&
                          !includesRoles({
                            currentUser: currentUser,
                            allowedFeatures: ["PUBLISHER_UPDATE"],
                          })))
                    ) {
                      return null;
                    } else
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
                        {cell.id.includes("actions") ? (
                          currentUser &&
                          includesRoles({
                            currentUser: currentUser,
                            allowedFeatures: ["CATEGORY_UPDATE"],
                          }) ? (
                            <div className=" flex justify-end ">
                              <EditPublisher
                                publisher={row.original}
                                handlePublisherEdited={handlePublisherEdited}
                              >
                                <Button
                                  size={"icon"}
                                  variant={"ghost"}
                                  className="rounded-full bg-blue-200/60 hover:bg-blue-200/90 text-primary hover:text-primary"
                                >
                                  <FaPen />
                                </Button>
                              </EditPublisher>
                            </div>
                          ) : null
                        ) : (
                          flexRender(
                            cell.column.columnDef.cell,
                            cell.getContext()
                          )
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
                    Không tìm thấy kết quả.
                  </TableCell>
                </TableRow>
              )}
            </TableBody>
          </Table>
        </div>
        <div className="flex items-center justify-end space-x-2 py-4">
          <div className="flex-1 text-sm text-muted-foreground"></div>
          <Paging
            page={page.toString()}
            totalPage={totalPage}
            onNavigateBack={() =>
              router.push(`/product/publishers?page=${Number(page) - 1}`)
            }
            onNavigateNext={() =>
              router.push(`/product/publishers?page=${Number(page) + 1}`)
            }
            onPageSelect={(selectedPage) =>
              router.push(`/product/publishers?page=${selectedPage}`)
            }
            onNavigateLast={() =>
              router.push(`/product/publishers?page=${totalPage}`)
            }
            onNavigateFirst={() => router.push(`/product/publishers?page=${1}`)}
          />
        </div>
      </div>
    );
}
