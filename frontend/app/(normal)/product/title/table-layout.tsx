"use client";

import CreateTitleDialog from "@/components/book-manage/create-title-dialog";
import { TitleTable } from "@/components/book-manage/title-table";
import { Button } from "@/components/ui/button";
import { endPoint } from "@/constants";
import { useCurrentUser } from "@/hooks/use-user";
import { getUser } from "@/lib/auth/action";
import { includesRoles } from "@/lib/utils";
import { useRouter, useSearchParams } from "next/navigation";
import { useSWRConfig } from "swr";

export const getFilterString = () => {
  const searchParams = useSearchParams();
  const minPrice = searchParams.get("minPrice") ?? undefined;
  const maxPrice = searchParams.get("maxPrice") ?? undefined;
  const createdAtFrom = searchParams.get("createdAtFrom") ?? undefined;
  const createdAtTo = searchParams.get("createdAtTo") ?? undefined;

  const search = searchParams.get("search") ?? undefined;

  let filters = [{ type: "", value: "" }];
  filters.pop();
  if (maxPrice) {
    filters = filters.concat({ type: "maxPrice", value: maxPrice });
  }
  if (minPrice) {
    filters = filters.concat({ type: "minPrice", value: minPrice });
  }
  if (search) {
    filters = filters.concat({ type: "search", value: search });
  }
  if (createdAtFrom) {
    filters = filters.concat({ type: "createdAtFrom", value: createdAtFrom });
  }
  if (createdAtTo) {
    filters = filters.concat({ type: "createdAtTo", value: createdAtTo });
  }
  let stringToFilter = "";
  filters.forEach((item) => {
    stringToFilter = stringToFilter.concat(`&${item.type}=${item.value}`);
  });
  return { stringToFilter: stringToFilter, filters: filters };
};

const TableLayout = () => {
  const router = useRouter();
  const searchParams = useSearchParams();
  const { mutate } = useSWRConfig();

  const { filters, stringToFilter } = getFilterString();
  const page = searchParams.get("page") ?? "1";
  const handleTitleAdded = async (titleId: string) => {
    mutate(
      `${endPoint}/v1/booktitles?page=${page ?? 1}&limit=10${
        stringToFilter ?? ""
      }`
    );
    router.refresh();
  };
  const { currentUser } = useCurrentUser();
  return (
    <div className="col">
      <div className="flex flex-row justify-between items-center">
        <h1>Danh sách đầu sách</h1>
        {currentUser &&
        includesRoles({
          currentUser: currentUser,
          allowedFeatures: ["BOOK_TITLE_CREATE"],
        }) ? (
          <CreateTitleDialog handleTitleAdded={handleTitleAdded}>
            <Button>Thêm đầu sách</Button>
          </CreateTitleDialog>
        ) : null}
      </div>

      <div className="my-4 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <TitleTable />
      </div>
    </div>
  );
};

export default TableLayout;
