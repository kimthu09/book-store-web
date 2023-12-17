import { PublisherTable } from "@/components/book-manage/publisher-table";

import { Metadata } from "next";
export const metadata: Metadata = {
  title: "Nhà xuất bản",
};
const PublisherPage = ({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) => {
  return (
    <div className="col">
      <div className="flex flex-row justify-between items-center">
        <h1>Nhà xuất bản</h1>
        <div className="flex gap-4">{/* <CreateCategory /> */}</div>
      </div>
      <div className="flex flex-row flex-wrap gap-2"></div>
      <div className="mb-4 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <PublisherTable searchParams={searchParams} />
      </div>
    </div>
  );
};

export default PublisherPage;
