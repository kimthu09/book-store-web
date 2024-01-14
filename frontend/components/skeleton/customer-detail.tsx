import { Card, CardContent } from "@/components/ui/card";
import DropdownSkeleton from "./dropdown-skeleton";

const CustomerDetailSkeleton = () => {
  return (
    <Card>
      <CardContent className="p-6 flex flex-col gap-4">
        <div className="flex gap-4  flex-col">
          <div className="flex gap-4 lg:flex-row flex-col">
            <div className="basis-2/3">
              <DropdownSkeleton />
            </div>
            <div className="basis-1/3">
              <DropdownSkeleton />
            </div>
          </div>
          <div className="flex gap-4 lg:flex-row flex-col">
            <div className="basis-2/3">
              <DropdownSkeleton />
            </div>
            <div className="basis-1/3">
              <DropdownSkeleton />
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  );
};

export default CustomerDetailSkeleton;
