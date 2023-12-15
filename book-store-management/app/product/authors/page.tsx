import { AuthorTable } from "@/components/book-manage/author-table";

const AuthorPage = ({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) => {
  return (
    <div className="col">
      <div className="flex flex-row justify-between items-center">
        <h1>Tác giả</h1>
        <div className="flex gap-4">{/* <CreateCategory /> */}</div>
      </div>
      <div className="flex flex-row flex-wrap gap-2"></div>
      <div className="mb-4 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <AuthorTable searchParams={searchParams} />
      </div>
    </div>
  );
};

export default AuthorPage;
