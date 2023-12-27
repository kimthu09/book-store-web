"use client";
import { AuthorTable } from "@/components/book-manage/author-table";
import CreateAuthor from "@/components/book-manage/create-author";
import { Button } from "@/components/ui/button";
import { endPoint } from "@/constants";
import { useRouter } from "next/navigation";
import { useSWRConfig } from "swr";

const TableLayout = ({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) => {
  const { mutate } = useSWRConfig();

  const router = useRouter();
  const page = searchParams["page"] ?? "1";

  const handleAuthorAdded = (idAuthor: string) => {
    mutate(`${endPoint}/v1/authors?page=${page ?? 1}&limit=10`);
    router.refresh();
  };
  return (
    <div className="col">
      <div className="flex flex-row justify-between items-center">
        <h1>Tác giả</h1>
        <CreateAuthor handleAuthorAdded={handleAuthorAdded}>
          <Button>Thêm tác giả</Button>
        </CreateAuthor>
      </div>
      <div className="flex flex-row flex-wrap gap-2"></div>
      <div className="mb-4 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <AuthorTable searchParams={searchParams} />
      </div>
    </div>
  );
};

export default TableLayout;
