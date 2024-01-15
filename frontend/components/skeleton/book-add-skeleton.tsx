import { Card, CardContent } from "../ui/card";
import { Skeleton } from "../ui/skeleton";

const BookAddSkeleton = () => {
  return (
    <div className="col items-center">
      <div className="2xl:w-4/5 w-full 2xl:px-0 md:px-6 px-0 ">
        <div className="flex flex-row justify-between">
          <h1 className="font-medium text-xxl self-start">Thêm sách mới</h1>
        </div>
        <div>
          <div className="flex flex-col flex-1 gap-4 xl:flex-row">
            <Card className="flex-1">
              <CardContent className="flex-col flex gap-1 mt-5">
                <div className="flex flex-row gap-1">
                  <Skeleton className="flex-1 w-full h-8" />
                  <Skeleton className="w-8 h-8" />
                </div>
              </CardContent>
            </Card>
            <Card className="flex-1">
              <CardContent className="flex-col flex gap-5 mt-5">
                <div className="basis-1/3">
                  <Skeleton className="w-full h-6" />
                </div>

                <div className="flex lg:gap-4 gap-3 xl:flex-col sm:flex-row flex-col">
                  <div className="flex flex-1 flex-row gap-1">
                    <Skeleton className="flex-1 w-full h-8" />
                    <Skeleton className="w-8 h-8" />
                  </div>
                  <div className="flex-1">
                    <Skeleton className="w-full h-6" />
                  </div>
                </div>

                <div className="flex lg:gap-4 gap-3 xl:flex-col sm:flex-row flex-col">
                  <div className="flex-1">
                    <Skeleton className="w-full h-6" />
                  </div>
                  <div className="flex-1">
                    <Skeleton className="w-full h-6" />
                  </div>
                </div>
              </CardContent>
            </Card>
          </div>
        </div>
      </div>
    </div>
  );
};

export default BookAddSkeleton;
