import { Skeleton } from "../ui/skeleton";
import { BiBox } from "react-icons/bi";
import { LuPhone } from "react-icons/lu";
import TableSkeleton from "./table-skeleton";

const ImportNoteDetailSkeleton = () => {
  return (
    <div className="flex flex-col xl:mx-[20%] gap-6">
      <div className="shadow-sm bg-white flex flex-col gap-6 md:px-8 px-4 pb-6">
        <div className="flex justify-between gap-6 font-bold text-lg border-b flex-1 py-2 pt-6">
          <div className="flex gap-4">
            <span className="font-light">Mã phiếu nhập</span>
            <Skeleton className="h-6 w-24" />
          </div>
        </div>
        <div className="grid grid-cols-2 text-sm">
          <div className="flex flex-col gap-4 w-fit">
            <div className="flex font-medium">
              <span className="w-16">Tạo</span>
              <div className="flex flex-col gap-2">
                <Skeleton className="h-4 w-16" />
                <Skeleton className="h-4 w-24" />
              </div>
            </div>
            <div className="flex font-medium w-fit">
              <span className="w-16">Đóng</span>
              <div className="flex flex-col gap-2">
                <Skeleton className="h-4 w-16" />
                <Skeleton className="h-4 w-24" />
              </div>
            </div>
          </div>
          <div className="flex flex-col items-end gap-4">
            <div className="w-fit">
              <div className="flex flex-col gap-2 font-medium">
                <div className="flex items-center gap-1">
                  <BiBox className="h-4 w-4" />
                  <Skeleton className="h-4 w-24" />
                </div>
                <div className="flex items-center gap-2">
                  <LuPhone className="h-4 w-4" />
                  <Skeleton className="h-4 w-24" />
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div className="shadow-sm bg-white flex flex-col gap-6 py-6 md:px-6 px-4">
        <TableSkeleton
          isHasExtensionAction={false}
          isHasFilter={false}
          isHasSearch={false}
          isHasChooseVisibleRow={false}
          isHasCheckBox={false}
          isHasPaging={false}
          numberRow={5}
          cells={[
            {
              percent: 2,
            },
            {
              percent: 2,
            },
            {
              percent: 5,
            },
            {
              percent: 2,
            },
            {
              percent: 2,
            },
            {
              percent: 2,
            },
          ]}
        ></TableSkeleton>
        <div className="flex justify-end space-x-2 pb-4 font-semibold">
          <span>Tổng tiền: </span>
          <Skeleton className="h-4 w-24" />
        </div>
      </div>
    </div>
  );
};

export default ImportNoteDetailSkeleton;
