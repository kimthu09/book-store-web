
import TableSkeleton from "./table-skeleton";
import { Card, CardContent } from "../ui/card";
import { Skeleton } from "../ui/skeleton";

const ImportNoteAddSkeleton = () => {
  return (
    <div className="col items-center">
      <div className="col xl:w-4/5 w-full px-0">
        <div className="flex justify-between gap-2">
          <h1 className="font-medium text-xxl self-start">Thêm phiếu nhập</h1>
        </div>

        <div className="flex flex-col gap-4">
          <div className="flex lg:flex-row flex-col gap-4">
            <Card className="basis-3/4 p-4 flex lg:flex-row flex-col gap-4">
              <Skeleton className="h-6 w-fill" />
              <Skeleton className="h-6 w-fill" />
            </Card>
            <Card className="basis-1/4 lg:p-6 p-4 flex justify-between gap-4">
              <Skeleton className="h-6 w-fill" />
            </Card>
          </div>
          <Card>
            <CardContent className="lg:p-6 p-4">
              <TableSkeleton
                isHasExtensionAction={false}
                isHasFilter={false}
                isHasSearch={true}
                isHasChooseVisibleRow={false}
                isHasCheckBox={false}
                isHasPaging={false}
                numberRow={5}
                cells={[
                  {
                    percent: 1,
                  },
                  {
                    percent: 2,
                  },
                  {
                    percent: 2,
                  },
                  {
                    percent: 1,
                  },
                ]}
              ></TableSkeleton>
            </CardContent>
          </Card>
          <div className="flex md:justify-end justify-stretch gap-2">
            <Skeleton className="h-6 flex-1" />
            <Skeleton className="h-6 flex-1" />
          </div>
        </div>
      </div>
    </div>
  );
};

export default ImportNoteAddSkeleton;
