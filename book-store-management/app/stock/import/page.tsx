import { ImportTable } from "@/components/stock-manage/import-table";
import Link from "next/link";
import React from "react";

const Import = () => {
  return (
    <div className="col">
      <div className="flex flex-row justify-between items-center">
        <h1>Danh sách phiếu nhập kho</h1>
        <Link
          href="/stock/import/new"
          className="inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 bg-primary text-primary-foreground shadow hover:bg-primary/90 h-9 px-4 py-2"
        >
          Tạo phiếu nhập
        </Link>
      </div>

      <div className="my-4 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <ImportTable />
      </div>
    </div>
  );
};

export default Import;
