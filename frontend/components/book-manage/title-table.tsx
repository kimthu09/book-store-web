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
import { Popover, PopoverContent, PopoverTrigger } from "../ui/popover";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../ui/select";
import { BookTitle, FilterValue } from "@/types";

import { Fragment, useEffect, useState } from "react";
import { Input } from "../ui/input";

import { useRouter, useSearchParams } from "next/navigation";
import Loading from "../loading";
import Paging, { PagingProps } from "../paging";
import {
  Controller,
  SubmitHandler,
  useFieldArray,
  useForm,
} from "react-hook-form";
import { LuFilter } from "react-icons/lu";
import { Label } from "../ui/label";
import { AiOutlineClose } from "react-icons/ai";
import getAllTitle from "@/lib/book/getAllTitle";
import { FilterDatePicker } from "../date-picker";
import { getFilterString } from "@/app/product/title/table-layout";
import TitleEditInline from "./title-edit-inline";
import { includesRoles } from "@/lib/utils";
import { useCurrentUser } from "@/hooks/use-user";
import TableSkeleton from "../skeleton/table-skeleton";

export const columns: ColumnDef<BookTitle>[] = [
  {
    accessorKey: "id",
    header: () => {
      return <span className="font-semibold">ID</span>;
    },
    cell: ({ row }) => <div className="leading-6">{row.getValue("id")}</div>,
    size: 4,
  },
  {
    accessorKey: "name",
    header: () => {
      return (
        <span className="font-semibold whitespace-normal">Tên đầu sách</span>
      );
    },
    cell: ({ row }) => <div className="capitalize">{row.getValue("name")}</div>,
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
      </div>
    ),

    sortingFn: "datetime",
    size: 5,
  },
];
export function TitleTable() {
  const searchParams = useSearchParams();
  const page = searchParams.get("page") ?? "1";

  const [latestFilter, setLatestFilter] = useState("");
  const { filters, stringToFilter } = getFilterString();
  const {
    mutate: mutate,
    titles: response,
    isLoading,
    isError,
  } = getAllTitle({
    page: page,
    filter: stringToFilter,
  });
  const data = response?.data;
  const totalPage = Math.ceil(response?.paging.total / response?.paging.limit);
  const router = useRouter();

  const filterValues = [
    { type: "search", name: "Từ khoá" },
    { type: "createdAtFrom", name: "Tạo từ ngày" },
    { type: "createdAtTo", name: "Tạo đến ngày" },
    // { type: "isActive", name: "Trạng thái" },
  ];

  const { register, handleSubmit, reset, control, getValues } =
    useForm<FilterValue>({
      defaultValues: {
        filters: filters,
      },
    });
  const { fields, append, remove, update } = useFieldArray({
    control: control,
    name: "filters",
  });
  const [openFilter, setOpenFilter] = useState(false);

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
  const onSubmit: SubmitHandler<FilterValue> = async (data) => {
    let filterString = "";
    data.filters.forEach((item) => {
      filterString = filterString.concat(`&${item.type}=${item.value}`);
    });
    router.push(`/product/title?page=1${filterString}`);
  };
  const { currentUser } = useCurrentUser();

  if (isError) return <div>Failed to load</div>;
  else if (isLoading) {
    return (
      <TableSkeleton
        isHasExtensionAction={false}
        isHasFilter={true}
        isHasSearch={true}
        isHasChooseVisibleRow={false}
        isHasCheckBox={false}
        isHasPaging={true}
        numberRow={5}
        cells={[
          {
            percent: 1,
          },
          {
            percent: 5,
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
        <div className="flex items-start py-4 gap-2">
          {/* <ExportDialog
            handleExport={handleExport}
            setExportOption={setExportOption}
            isImport
          /> */}
          <div className="flex-1">
            <div className="flex gap-2">
              <Popover
                open={openFilter}
                onOpenChange={(open) => {
                  setOpenFilter(open);
                  reset({ filters: filters });
                }}
              >
                <PopoverTrigger asChild>
                  <Button variant="outline" className="lg:px-3 px-2">
                    Lọc
                    <LuFilter className="ml-1 h-4 w-4" />
                  </Button>
                </PopoverTrigger>
                <PopoverContent className="w-96">
                  <form
                    className="flex flex-col gap-4"
                    onSubmit={handleSubmit(onSubmit)}
                  >
                    <div className="space-y-2">
                      <p className="text-sm text-muted-foreground">
                        Hiển thị đầu sách theo
                      </p>
                    </div>
                    <div className="flex flex-col gap-4">
                      {fields.map((item, index) => {
                        const name = filterValues.find(
                          (v) => v.type === item.type
                        );
                        return (
                          <div
                            className="flex gap-2 items-center"
                            key={item.id}
                          >
                            <Label className="basis-1/3">{name?.name}</Label>
                            {item.type === "search" ? (
                              <Input
                                {...register(`filters.${index}.value`)}
                                className="flex-1"
                                type="text"
                                required
                              ></Input>
                            ) : item.type.includes("createdAt") ? (
                              <Controller
                                control={control}
                                name={`filters.${index}.value`}
                                render={({ field }) => (
                                  <FilterDatePicker
                                    handleDateSelected={(date) => {
                                      const unixTimeMilliseconds =
                                        date?.getTime();
                                      const unixTimeSeconds = Math.floor(
                                        unixTimeMilliseconds! / 1000
                                      );
                                      if (item.type.includes("To")) {
                                        field.onChange(unixTimeSeconds + 86399);
                                      } else {
                                        field.onChange(unixTimeSeconds);
                                      }
                                    }}
                                    date={new Date(+field.value * 1000)}
                                  />
                                )}
                              />
                            ) : null}
                            <Button
                              variant={"ghost"}
                              className={`px-3 `}
                              onClick={() => {
                                remove(index);
                              }}
                            >
                              <AiOutlineClose />
                            </Button>
                          </div>
                        );
                      })}
                    </div>
                    {fields.length === filterValues.length ? null : (
                      <div className="flex justify-end pr-12">
                        <Select
                          value={latestFilter}
                          onValueChange={(value) => {
                            append({ type: value, value: "" });
                          }}
                        >
                          <SelectTrigger className="w-[180px] flex justify-center ml-8 px-3">
                            <SelectValue placeholder="Chọn điều kiện lọc" />
                          </SelectTrigger>
                          <SelectContent>
                            <SelectGroup>
                              {filterValues.map((item) => {
                                return fields.findIndex(
                                  (v) => v.type === item.type
                                ) === -1 ? (
                                  <SelectItem key={item.type} value={item.type}>
                                    {item.name}
                                  </SelectItem>
                                ) : null;
                              })}
                            </SelectGroup>
                          </SelectContent>
                        </Select>
                      </div>
                    )}
                    <Button type="submit" className="self-end">
                      Lọc
                    </Button>
                  </form>
                </PopoverContent>
              </Popover>
              <Input
                className="flex-1"
                placeholder="Tìm kiếm mã đầu sách"
                value={
                  (table.getColumn("id")?.getFilterValue() as string) ?? ""
                }
                onChange={(event) =>
                  table.getColumn("id")?.setFilterValue(event.target.value)
                }
              />
            </div>
            <div className="flex gap-2 mt-2">
              {filters.map((item, index) => {
                const name = filterValues.find((v) => v.type === item.type);
                return (
                  <div
                    key={item.type}
                    className="rounded-xl flex self-start px-3 py-1 h-fit outline-none text-sm text-primary  bg-blue-100 items-center gap-1 group"
                  >
                    <span>
                      {name?.name}
                      {": "}
                      {item.type.includes("createdAt")
                        ? new Date(+item.value * 1000).toLocaleDateString(
                            "vi-VN"
                          )
                        : item.value}
                    </span>
                  </div>
                );
              })}
            </div>
          </div>
        </div>

        <div className="rounded-md border overflow-x-auto min-w-full max-w-[50vw]">
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
                  <Fragment key={row.id}>
                    <TableRow data-state={row.getIsSelected() && "selected"}>
                      {row.getVisibleCells().map((cell) => (
                        <TableCell
                          key={cell.id}
                          onClick={() => {
                            if (
                              !cell.id.includes("select") &&
                              !cell.id.includes("actions")
                            ) {
                              table.resetExpanded();
                              if (!row.getIsExpanded()) {
                                row.toggleExpanded();
                              }
                            }
                          }}
                        >
                          {flexRender(
                            cell.column.columnDef.cell,
                            cell.getContext()
                          )}
                        </TableCell>
                      ))}
                    </TableRow>
                    {currentUser &&
                    includesRoles({
                      currentUser: currentUser,
                      allowedFeatures: ["BOOK_TITLE_VIEW"],
                    }) ? (
                      <TableRow>
                        <TableCell
                          colSpan={columns.length}
                          className={`w-full p-0 border-none ${
                            row.getIsExpanded() ? "table-cell" : "hidden"
                          }`}
                        >
                          <TitleEditInline
                            book={row.original}
                            handleTitleEdited={() => {
                              mutate();
                            }}
                          />
                        </TableCell>
                      </TableRow>
                    ) : null}
                  </Fragment>
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
            page={page}
            onNavigateNext={() =>
              router.push(`/product/title?page=${+page + 1}${stringToFilter}`)
            }
            onNavigateBack={() =>
              router.push(`/product/title?page=${+page - 1}${stringToFilter}`)
            }
            totalPage={totalPage}
            onPageSelect={(selectedPage) => {
              router.push(
                `/product/title?page=${selectedPage}${stringToFilter}`
              );
            }}
            onNavigateFirst={() =>
              router.push(`/product/title?page=${1}${stringToFilter}`)
            }
            onNavigateLast={() =>
              router.push(`/product/title?page=${totalPage}${stringToFilter}`)
            }
          />
        </div>
      </div>
    );
  }
}
