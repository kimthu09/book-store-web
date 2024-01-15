import { Skeleton } from "../ui/skeleton";

type TableSkeletonCell = {
  percent: number;
};

type TableSkeletonProps = {
  isHasExtensionAction: boolean;
  isHasFilter: boolean;
  isHasSearch: boolean;
  isHasChooseVisibleRow: boolean;
  isHasCheckBox: boolean;
  isHasPaging: boolean;
  numberRow: number;
  cells: TableSkeletonCell[];
};

const TableSkeleton: React.FC<TableSkeletonProps> = ({
  isHasExtensionAction,
  isHasFilter,
  isHasSearch,
  isHasChooseVisibleRow,
  isHasCheckBox,
  isHasPaging,
  numberRow,
  cells,
}) => {
  return (
    <div className="flex flex-col w-full gap-4">
      <div className="flex flex-row items-center gap-2">
        {isHasExtensionAction && <Skeleton className="w-24 h-8" />}
        {isHasFilter && <Skeleton className="w-16 h-8" />}
        {isHasSearch && <Skeleton className="w-full h-8 flex-1" />}
        {isHasChooseVisibleRow && <Skeleton className="w-24 h-8" />}
      </div>
      <div className="flex flex-col rounded-md border overflow-x-auto">
        {[...Array(numberRow + 1)].map((_, index) => (
          <div
            key={index}
            className={`flex flex-row gap-2 ${
              index !== numberRow && "border-b"
            } p-2`}
          >
            {isHasCheckBox && <Skeleton key={0} className="w-6 h-6" />}
            <div className="flex flex-1 flex-row gap-2">
              {cells.map((cell, cellIndex) => (
                <Skeleton
                  key={cellIndex + 1}
                  className={`${
                    cell.percent != 1 ? `flex-[${cell.percent}]` : `flex-1`
                  }  h-6`}
                />
              ))}
            </div>
          </div>
        ))}
      </div>
      <div className="flex items-center justify-end space-x-2 mb-2">
        <div className="flex-1">
          {isHasCheckBox && <Skeleton className="w-56 h-8" />}
        </div>
        {isHasPaging && (
          <div className="flex flex-row gap-2">
            <Skeleton className="w-8 h-8" />
            <Skeleton className="w-8 h-8" />
            <Skeleton className="flex-1 h-8" />
            <Skeleton className="w-8 h-8" />
            <Skeleton className="w-8 h-8" />
          </div>
        )}
      </div>
    </div>
  );
};

export default TableSkeleton;
