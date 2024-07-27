"use client";

import CreatePublisher from "@/components/book-manage/create-publisher";
import { PublisherTable } from "@/components/book-manage/publisher-table";
import Loading from "@/components/loading";
import TableSkeleton from "@/components/skeleton/table-skeleton";
import { Button } from "@/components/ui/button";
import { endPoint } from "@/constants";
import { useCurrentUser } from "@/hooks/use-user";
import { includesRoles } from "@/lib/utils";
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

  const handlePublisherAdded = (idAuthor: string) => {
    mutate(`${endPoint}/v1/publishers?page=${page ?? 1}&limit=10`);
    router.refresh();
  };
  const { currentUser } = useCurrentUser();

  return (
    <div className="col">
      <div className="flex flex-row justify-between items-center">
        <h1>Nhà xuất bản</h1>
        {currentUser &&
        includesRoles({
          currentUser: currentUser,
          allowedFeatures: ["PUBLISHER_CREATE"],
        }) ? (
          <div className="flex gap-4">
            <CreatePublisher handlePublisherAdded={handlePublisherAdded}>
              <Button>Thêm nhà xuất bản</Button>
            </CreatePublisher>
          </div>
        ) : null}
      </div>
      <div className="flex flex-row flex-wrap gap-2"></div>
      <div className="mb-4 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <Suspense
          fallback={
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
            />
          }
        >
          <PublisherTable
            searchParams={searchParams}
            currentUser={currentUser}
          />
        </Suspense>
      </div>
    </div>
  );
};

export default TableLayout;
