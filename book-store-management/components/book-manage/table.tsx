"use client";

import * as React from "react";
import {
  CaretSortIcon,
  ChevronDownIcon,
  DotsHorizontalIcon,
} from "@radix-ui/react-icons";
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
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Book } from "@/types";

import { useState } from "react";

import deleteBook from "@/lib/book/deleteBook";
import { useRouter, useSearchParams } from "next/navigation";
import { toast } from "../ui/use-toast";
import Paging from "../paging";
import { Input } from "../ui/input";
import { SubmitHandler, useFieldArray, useForm } from "react-hook-form";
import { Popover, PopoverContent, PopoverTrigger } from "../ui/popover";
import { LuFilter } from "react-icons/lu";
import { Label } from "../ui/label";
import { AiOutlineClose } from "react-icons/ai";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../ui/select";
import CategoryList from "./category-list";
import PublisherList from "./publisher-list";
import AuthorList from "./author-list";

function idToName(id: string) {
  if (id === "name") {
    return "Tên";
  } else if (id === "sellPrice") {
    return "Giá bán";
  } else if (id === "isActive") {
    return "Trạng thái";
  } else if (id === "authors") {
    return "Tác giả";
  } else if (id === "categories") {
    return "Thể loại";
  } else if (id === "publisher") {
    return "Nhà xuất bản";
  } else if (id === "quantity") {
    return "Số lượng";
  } else if (id === "listedPrice") {
    return "Giá niêm yết";
  }
  return id;
}
type FormValues = {
  filters: {
    type: string;
    value: string;
  }[];
};
export function BookTable({
  data,
  totalPage,
}: {
  data: Book[];
  totalPage: number;
}) {
  // const data: Book[] = books;
  const router = useRouter();
  const searchParams = useSearchParams();
  const page = searchParams.get("page") ?? "1";

  const [sorting, setSorting] = useState<SortingState>([]);
  const [columnFilters, setColumnFilters] = useState<ColumnFiltersState>([]);
  const [columnVisibility, setColumnVisibility] = useState<VisibilityState>({});
  const [rowSelection, setRowSelection] = useState({});
  const [category, setCategory] = useState("");

  const columns: ColumnDef<Book>[] = [
    {
      id: "select",
      header: ({ table }) => (
        <Checkbox
          checked={table.getIsAllPageRowsSelected()}
          onCheckedChange={(value) => table.toggleAllPageRowsSelected(!!value)}
          aria-label="Select all"
        />
      ),
      cell: ({ row }) => (
        <Checkbox
          checked={row.getIsSelected()}
          onCheckedChange={(value) => row.toggleSelected(!!value)}
          aria-label="Select row"
        />
      ),
      enableSorting: false,
      enableHiding: false,
    },
    {
      accessorKey: "id",
      header: () => {
        return <span className="font-semibold">ID</span>;
      },
      cell: ({ row }) => <div>{row.getValue("id")}</div>,
    },
    {
      accessorKey: "name",
      header: ({ column }) => {
        return (
          <Button
            className="p-2"
            variant="ghost"
            onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
          >
            <span className="font-semibold">Tên sản phẩm</span>

            <CaretSortIcon className="ml-1 h-4 w-4" />
          </Button>
        );
      },
      cell: ({ row }) => (
        <div className="capitalize">{row.getValue("name")}</div>
      ),
    },
    {
      accessorKey: "authors",
      accessorFn: (row) => {
        return row.bookTitle.authors.map((item) => item.name).join(", ");
      },
      header: ({ column }) => {
        return <span className="font-semibold">Tác giả</span>;
      },
      cell: ({ row }) => (
        <div className="capitalize">{row.getValue("authors")}</div>
      ),
    },
    {
      accessorKey: "categories",
      accessorFn: (row) => {
        return row.bookTitle.categories.map((item) => item.name).join(", ");
      },
      header: ({ column }) => {
        return <span className="font-semibold">Thể loại</span>;
      },
      cell: ({ row }) => (
        <div className="capitalize">{row.getValue("categories")}</div>
      ),
    },
    {
      accessorKey: "publisher",
      accessorFn: (row) => row.publisher.name,
      header: () => {
        return <span className="font-semibold">NXB</span>;
      },
      cell: ({ row }) => (
        <div className="capitalize">{row.getValue("publisher")}</div>
      ),
    },
    {
      accessorKey: "quantity",
      header: () => {
        return <div className="font-semibold flex justify-end">Số lượng</div>;
      },
      cell: ({ row }) => (
        <div className="text-right">{row.getValue("quantity")}</div>
      ),
    },
    {
      accessorKey: "listedPrice",
      header: ({ column }) => (
        <div className="flex justify-end">
          <Button
            className="p-1"
            variant={"ghost"}
            onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
          >
            <span className="font-semibold">Giá niêm yết</span>
            <CaretSortIcon className="ml-1 h-4 w-4" />
          </Button>
        </div>
      ),
      cell: ({ row }) => {
        const amount = parseFloat(row.getValue("listedPrice"));

        // Format the amount as a dollar amount
        const formatted = new Intl.NumberFormat("vi-VN", {
          style: "currency",
          currency: "VND",
        }).format(amount);

        return <div className="text-right font-medium">{formatted}</div>;
      },
    },
    {
      accessorKey: "sellPrice",
      header: ({ column }) => (
        <div className="flex justify-end">
          <Button
            className="p-1"
            variant={"ghost"}
            onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
          >
            <span className="font-semibold">Giá bán</span>
            <CaretSortIcon className="ml-1 h-4 w-4" />
          </Button>
        </div>
      ),
      cell: ({ row }) => {
        const amount = parseFloat(row.getValue("sellPrice"));

        // Format the amount as a dollar amount
        const formatted = new Intl.NumberFormat("vi-VN", {
          style: "currency",
          currency: "VND",
        }).format(amount);

        return <div className="text-right font-medium">{formatted}</div>;
      },
    },
    {
      accessorKey: "isActive",
      header: () => {
        return (
          <div className="font-semibold flex justify-center">Trạng thái</div>
        );
      },
      cell: ({ row }) => {
        const status = row.getValue("isActive");
        return (
          <div
            className={`lg:max-w-[16rem] max-w-[3rem] truncate ${
              status ? "text-green-600" : "text-red-600"
            } text-center`}
          >
            {status ? "Đang bán" : "Ngừng bán"}
          </div>
        );
      },
    },
  ];

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

  const [latestFilter, setLatestFilter] = useState("");
  const filterValues = [
    { type: "search", name: "Từ khoá" },
    { type: "minSellPrice", name: "Giá bán nhỏ nhất" },
    { type: "maxSellPrice", name: "Giá bán lớn nhất" },
    // { type: "publisher", name: "Nhà xuất bản" },
    // { type: "categories", name: "Thể loại" },
    // { type: "authors", name: "Tác giả" },
  ];

  const search = searchParams.get("search") ?? undefined;
  const minSellPrice = searchParams.get("minSellPrice") ?? undefined;
  const maxSellPrice = searchParams.get("maxSellPrice") ?? undefined;
  const categories = searchParams.get("categories") ?? undefined;
  const authors = searchParams.get("authors") ?? undefined;
  const publisher = searchParams.get("publisher") ?? undefined;

  let filters = [{ type: "", value: "" }];
  filters.pop();
  if (maxSellPrice) {
    filters = filters.concat({ type: "maxSellPrice", value: maxSellPrice });
  }
  if (minSellPrice) {
    filters = filters.concat({ type: "minSellPrice", value: minSellPrice });
  }
  if (search) {
    filters = filters.concat({ type: "search", value: search });
  }
  if (categories) {
    filters = filters.concat({ type: "categories", value: categories });
  }
  if (authors) {
    filters = filters.concat({ type: "authors", value: authors });
  }
  if (publisher) {
    filters = filters.concat({ type: "publisher", value: publisher });
  }

  const { register, handleSubmit, reset, control, getValues } =
    useForm<FormValues>({
      defaultValues: {
        filters: filters,
      },
    });
  const { fields, append, remove, update } = useFieldArray({
    control: control,
    name: "filters",
  });

  const onSubmit: SubmitHandler<FormValues> = async (data) => {
    console.log(data);
    let search = "";
    let minSellPrice = "";
    let maxSellPrice = "";
    data.filters.forEach((item) => {
      if (item.type === "minSellPrice") {
        minSellPrice = `&minSellPrice=${item.value}`;
      } else if (item.type === "maxSellPrice") {
        maxSellPrice = `&maxSellPrice=${item.value}`;
      } else if (item.type === "search") {
        search = `&search=${item.value}`;
      }
    });

    router.push(
      `/product/books?page=${Number(
        page
      )}${minSellPrice}${maxSellPrice}${search}`
    );
  };
  const [openFilter, setOpenFilter] = useState(false);
  return (
    <div className="w-full">
      <div className="flex items-start py-4 gap-2">
        <div>
          {/* <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="outline" className="w-40 justify-between">
                Chọn thao tác <ChevronDownIcon className="ml-2 h-4 w-4" />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent className="max-h-44 w-40">
              <DropdownMenuItem
                onSelect={() => {
                  if (table.getFilteredSelectedRowModel().rows.length < 1) {
                    //TODO: show notification
                  } else {
                    const values = table
                      .getFilteredSelectedRowModel()
                      .rows.map((row) => row.original);
                    ExportBookList(values, "BookList.xlsx");
                  }
                }}
              >
                Xuất danh sách
              </DropdownMenuItem>
              <Dialog>
                <DialogTrigger asChild>
                  <DropdownMenuItem onSelect={(e) => e.preventDefault()}>
                    Chuyển danh mục
                  </DropdownMenuItem>
                </DialogTrigger>
                <DialogOverlay>
                  <DialogContent className="p-0">
                    <DialogTitle className="p-6 pb-0">
                      Chuyển mặt hàng tới danh mục khác
                    </DialogTitle>
                    <div className="flex flex-col border-y-[1px] p-6">
                      <p>Chọn 1 danh mục muốn chuyển tới</p>
                      <div className="mt-4 flex-1 ">
                        <CategoryList
                          category={category}
                          setCategory={setCategory}
                        />
                      </div>
                    </div>

                    <DialogClose className="ml-auto p-6 pt-0">
                      <Button type="submit">Hoàn tất</Button>
                    </DialogClose>
                  </DialogContent>
                </DialogOverlay>
              </Dialog>

              <Dialog>
                <DialogTrigger asChild>
                  <DropdownMenuItem onSelect={(e) => e.preventDefault()}>
                    Ngừng bán
                  </DropdownMenuItem>
                </DialogTrigger>
                <DialogContent>
                  <DialogTitle>Ngừng bán sách được chọn</DialogTitle>
                </DialogContent>
              </Dialog>
            </DropdownMenuContent>
          </DropdownMenu> */}
        </div>

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
              <PopoverContent className="w-80">
                <form
                  className="flex flex-col gap-4"
                  onSubmit={handleSubmit(onSubmit)}
                >
                  <div className="space-y-2">
                    <p className="text-sm text-muted-foreground">
                      Hiển thị sách theo
                    </p>
                  </div>
                  <div className="flex flex-col gap-4">
                    {fields.map((item, index) => {
                      const name = filterValues.find(
                        (v) => v.type === item.type
                      );
                      if (item.type === "categories") {
                        return (
                          <div
                            className="flex gap-2 items-center"
                            key={item.id}
                          >
                            <Label className="basis-1/4">{name?.name}</Label>
                            <div className="flex-1">
                              <CategoryList
                                checkedCategory={[]}
                                onCheckChanged={(idCate) => {}}
                              />
                            </div>

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
                      } else if (item.type === "publisher") {
                        return (
                          <div
                            className="flex gap-2 items-center"
                            key={item.id}
                          >
                            <Label className="basis-1/4">{name?.name}</Label>
                            <PublisherList
                              publisherId={""}
                              setPublisherId={(id) => {}}
                            />
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
                      } else if (item.type === "authors") {
                        return (
                          <div
                            className="flex gap-2 items-center"
                            key={item.id}
                          >
                            <Label className="basis-1/4">{name?.name}</Label>
                            <div className="flex-1">
                              <AuthorList
                                checkedAuthor={[]}
                                onCheckChanged={(id) => {}}
                              />
                            </div>
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
                      } else
                        return (
                          <div
                            className="flex gap-2 items-center"
                            key={item.id}
                          >
                            <Label className="basis-1/4">{name?.name}</Label>
                            {item.type === "search" ? (
                              <Input
                                {...register(`filters.${index}.value`)}
                                className="flex-1"
                                type="text"
                                required
                              ></Input>
                            ) : (
                              <Input
                                {...register(`filters.${index}.value`)}
                                className="flex-1"
                                type="number"
                                required
                              ></Input>
                            )}
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
                    <div className="flex justify-center">
                      <Select
                        value={latestFilter}
                        onValueChange={(value) => {
                          console.log(value);
                          append({ type: value, value: "" });
                        }}
                      >
                        <SelectTrigger className="w-[160px] flex justify-center ml-8 px-3">
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
            <div className="flex-1">
              <Input
                placeholder="Tìm kiếm tên sách"
                value={
                  (table.getColumn("name")?.getFilterValue() as string) ?? ""
                }
                onChange={(event) =>
                  table.getColumn("name")?.setFilterValue(event.target.value)
                }
              />
            </div>
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
                    {item.value}
                  </span>
                </div>
              );
            })}
          </div>
        </div>
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="outline">
              Cột hiển thị <ChevronDownIcon className="ml-2 h-4 w-4" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent>
            {table
              .getAllColumns()
              .filter((column) => column.getCanHide())
              .map((column) => {
                return (
                  <DropdownMenuCheckboxItem
                    key={column.id}
                    className="capitalize"
                    checked={column.getIsVisible()}
                    onCheckedChange={(value) =>
                      column.toggleVisibility(!!value)
                    }
                  >
                    {idToName(column.id)}
                  </DropdownMenuCheckboxItem>
                );
              })}
          </DropdownMenuContent>
        </DropdownMenu>
      </div>
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
                  No results.
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </div>
      <div className="flex items-center justify-end space-x-2 py-4">
        <div className="flex-1 text-sm text-muted-foreground">
          {table.getFilteredSelectedRowModel().rows.length} of{" "}
          {table.getFilteredRowModel().rows.length} row(s) selected.
        </div>
        <Paging
          page={page}
          totalPage={totalPage}
          onNavigateBack={() =>
            router.push(`/product/books?page=${Number(page) - 1}`)
          }
          onNavigateNext={() =>
            router.push(`/product/books?page=${Number(page) + 1}`)
          }
          onPageSelect={(selectedPage) =>
            router.push(`/product/books?page=${selectedPage}`)
          }
          onNavigateLast={() => router.push(`/product/books?page=${totalPage}`)}
          onNavigateFirst={() => router.push(`/product/books?page=${1}`)}
        />
      </div>
    </div>
  );
}
