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

import { useState } from "react";
import { Input } from "../ui/input";
import Loading from "../loading";
import Paging from "../paging";
import { useRouter } from "next/navigation";
import { Author } from "@/types";
import getAllAuthor from "@/lib/book/getAllAuthor";
import EditAuthor from "./edit-author";
import { FaPen } from "react-icons/fa";
import { useSWRConfig } from "swr";
import { endPoint } from "@/constants";
import { includesRoles } from "@/lib/utils";
import TableSkeleton from "../skeleton/table-skeleton";

export const columns: ColumnDef<Author>[] = [
  {
    accessorKey: "name",
    header: ({ column }) => {
      return (
        <Button
          className="p-2"
          variant="ghost"
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
        >
          <span className="font-semibold">Tác giả</span>

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
export function AuthorTable({
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
  const { authors, isLoading, isError } = getAllAuthor({
    page: page.toString(),
  });
  const data = authors?.data!;
  const totalPage = Math.ceil(authors?.paging.total! / authors?.paging.limit!);

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
    // initialState: {
    //   pagination: {
    //     pageSize: 15,
    //   },
    // },
    state: {
      sorting,
      columnFilters,
      columnVisibility,
      rowSelection,
    },
  });

  const { mutate } = useSWRConfig();
  const handleAuthorEdited = (name: string) => {
    mutate(`${endPoint}/v1/authors?page=${page ?? 1}&limit=10`);
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
            placeholder="Tìm kiếm tác giả"
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
                            allowedFeatures: ["AUTHOR_UPDATE"],
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
                            allowedFeatures: ["AUTHOR_UPDATE"],
                          }) ? (
                            <div className=" flex justify-end ">
                              <EditAuthor
                                author={row.original}
                                handleAuthorEdited={handleAuthorEdited}
                              >
                                <Button
                                  size={"icon"}
                                  variant={"ghost"}
                                  className="rounded-full bg-blue-200/60 hover:bg-blue-200/90 text-primary hover:text-primary"
                                >
                                  <FaPen />
                                </Button>
                              </EditAuthor>
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
              router.push(`/product/authors?page=${Number(page) - 1}`)
            }
            onNavigateNext={() =>
              router.push(`/product/authors?page=${Number(page) + 1}`)
            }
            onPageSelect={(selectedPage) =>
              router.push(`/product/authors?page=${selectedPage}`)
            }
            onNavigateLast={() =>
              router.push(`/product/authors?page=${totalPage}`)
            }
            onNavigateFirst={() => router.push(`/product/authors?page=${1}`)}
          />
        </div>
      </div>
    );
}
