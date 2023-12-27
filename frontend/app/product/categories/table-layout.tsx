"use client";
import { CategoryTable } from "@/components/book-manage/category-table";
import CreateCategory from "@/components/book-manage/create-category";
import Loading from "@/components/loading";
import { Button } from "@/components/ui/button";
import { endPoint } from "@/constants";
import { useRouter } from "next/navigation";
import { Suspense } from "react";
import { useSWRConfig } from "swr";

const TableLayout = ({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) => {
  const { mutate } = useSWRConfig();

  const router = useRouter();
  const page = searchParams["page"] ?? "1";

  const handleCategoryAdded = (idAuthor: string) => {
    mutate(`${endPoint}/v1/categories?page=${page ?? 1}&limit=10`);
    router.refresh();
  };
  return (
    <div className="col">
      <div className="flex flex-row justify-between items-center">
        <h1>Thể loại</h1>
        <div className="flex gap-4">
          <CreateCategory handleCategoryAdded={handleCategoryAdded}>
            <Button>Thêm thể loại</Button>
          </CreateCategory>
        </div>
      </div>
      <div className="flex flex-row flex-wrap gap-2"></div>
      <div className="mb-4 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <Suspense fallback={<Loading />}>
          <CategoryTable searchParams={searchParams} />
        </Suspense>
      </div>
    </div>
  );
};

export default TableLayout;
