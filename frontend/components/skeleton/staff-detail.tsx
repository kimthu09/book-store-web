import { Card, CardContent } from "@/components/ui/card";
import { Skeleton } from "../ui/skeleton";
import DropdownSkeleton from "./dropdown-skeleton";

const StaffDetailSkeleton = () => {
  return (
    <div className="col items-center">
      <div className="col xl:w-4/5 w-full px-0">
        <div className="flex flex-col gap-4">
          <Card>
            <CardContent className="p-6 flex gap-6">
              <div className="relative h-min">
                <Skeleton className="h-[100px] w-[100px] rounded-full" />
              </div>
              <div className="flex-1 flex-col flex gap-4">
                <div>
                  <DropdownSkeleton />
                </div>
                <div className="flex flex-col gap-4 lg:flex-row">
                  <div className="basis-2/3">
                    <DropdownSkeleton />
                  </div>
                  <div className="basis-1/3">
                    <DropdownSkeleton />
                  </div>
                </div>
                <div>
                  <DropdownSkeleton />
                </div>
                <div className="flex gap-2 justify-end sm:flex-row flex-col">
                  <div className="flex gap-2 justify-center">
                    <DropdownSkeleton />
                  </div>
                  <DropdownSkeleton />
                </div>
              </div>
            </CardContent>
          </Card>
        </div>
        <div className="flex lg:flex-row flex-col gap-4">
          <Card className="flex-1">
            <CardContent className="p-6 flex items-end justify-between gap-2">
              <DropdownSkeleton />
            </CardContent>
          </Card>
          <Card className="flex-1">
            <CardContent className="p-6 flex items-end justify-between gap-2">
              <DropdownSkeleton />
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
};

export default StaffDetailSkeleton;
