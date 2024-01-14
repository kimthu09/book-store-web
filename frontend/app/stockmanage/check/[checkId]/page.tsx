"use client";
import Loading from "@/components/loading";
import NoRole from "@/components/no-role";
import InventoryCheckNoteDetailSkeleton from "@/components/skeleton/inventory-check-note-detail-skeleton";
import { CheckDetailTable } from "@/components/stock-manage/check-detail-table";
import {
  ExportCheckNoteDetail,
  ExportImportNoteDetail,
} from "@/components/stock-manage/excel-import-detail";
import { Button } from "@/components/ui/button";
import { useCurrentUser } from "@/hooks/use-user";
import getCheckNoteDetail from "@/lib/import/getCheckDetail";
import { includesRoles } from "@/lib/utils";
import { useRouter } from "next/navigation";
import { FaRegFileExcel } from "react-icons/fa";
const CheckDetail = ({ params }: { params: { checkId: string } }) => {
  const router = useRouter();
  const { data, isLoading, isError, mutate } = getCheckNoteDetail({
    id: params.checkId,
  });
  const { currentUser } = useCurrentUser();
  if (isError) return <div>Failed to load</div>;
  else if (!currentUser || isLoading) {
    return <InventoryCheckNoteDetailSkeleton />;
  } else if (
    currentUser &&
    !includesRoles({
      currentUser: currentUser,
      allowedFeatures: ["INVENTORY_NOTE_VIEW"],
    })
  ) {
    return <NoRole></NoRole>;
  } else
    return (
      <div className="flex flex-col xl:mx-[20%] gap-6">
        <div className="shadow-sm bg-white flex flex-col gap-6 md:px-8 px-4 pb-6">
          <div className="flex justify-between gap-6 font-bold text-lg border-b flex-1 py-2 pt-6">
            <div className="flex gap-4">
              <span className="font-light">Mã phiếu kiểm kho</span>
              <span>{data.id}</span>
            </div>
            <div className="flex gap-2 flex-nowrap">
              <Button
                variant={"outline"}
                className="p-2"
                onClick={() => {
                  ExportCheckNoteDetail(data, data.details, "phieukiem.xlsx");
                }}
              >
                <FaRegFileExcel className="mr-1 h-5 w-5 text-green-700" />
                <span className="whitespace-nowrap">Xuất file</span>
              </Button>
            </div>
          </div>
          <div className="grid grid-cols-2 text-sm">
            <div className="flex flex-col gap-4 w-fit">
              <div className="flex font-medium">
                <span className="w-16">Tạo</span>
                <div className="flex flex-col">
                  <span>{new Date(data.createdAt).toLocaleDateString()}</span>
                  <span className="font-light">{data.createdBy.name}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div className="shadow-sm bg-white flex flex-col gap-6 py-6 md:px-6 px-4">
          <CheckDetailTable details={data.details} />
          <div className="flex justify-end space-x-2 pb-4 font-semibold">
            <span>Số lượng thay đổi: </span>
            <span>{data.qtyDifferent}</span>
          </div>
        </div>
      </div>
    );
};

export default CheckDetail;
