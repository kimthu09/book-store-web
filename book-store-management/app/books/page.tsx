import { BookTable } from "@/components/book-manage/table";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import React from "react";
import { AiOutlinePlus } from "react-icons/ai";
import { FiDownload } from "react-icons/fi";

const BookManagement = () => {
  return (
    <div className="col">
      <div className="flex flex-row justify-between items-center">
        <h1>Tất cả sách</h1>
        <div className="flex gap-4">
          <Button className="p-2 hover:bg-blue-100" variant={"ghost"}>
            <div className="flex flex-wrap gap-1 items-center">
              <FiDownload />
              Xuất danh sách
            </div>
          </Button>
          <Link href="/books/insert">
            <Button className="p-2">
              <div className="flex flex-wrap gap-1 items-center">
                <AiOutlinePlus />
                Thêm sách mới
              </div>
            </Button>
          </Link>
        </div>
      </div>
      <div className="flex flex-row flex-wrap gap-2"></div>

      <div className="mb-4 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <BookTable />
      </div>
    </div>
  );
};

export default BookManagement;
