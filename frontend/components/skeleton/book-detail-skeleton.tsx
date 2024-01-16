import { Card, CardContent } from "../ui/card";
import { Skeleton } from "../ui/skeleton";

const BookDetailSkeleton = () => {
  return (
    <div className="col items-center">
      <div className="col 2xl:w-4/5 w-full 2xl:px-0 px-0">
        <div className="flex flex-row justify-between">
          <h1 className="flex flex-row font-medium text-2xl self-start">
            Sách:{" "}
            <span>
              <Skeleton className="w-16 h-6" />
            </span>
          </h1>
        </div>
        <div>
          <div className="flex flex-col flex-1 gap-4 xl:flex-row">
            <Card className="flex-1">
              <CardContent className="flex-col flex gap-1 mt-5">
                <div className="flex flex-col gap-3">
                  <div className="flex gap-1">
                    <Skeleton className="w-full h-6" />
                  </div>
                  <div className="flex flex-col gap-3 text-sm font-medium mt-2">
                    <div className="flex">
                      <span className="flex basis-1/4 text-muted-foreground">
                        <div className="basis-3/4">
                          <Skeleton className="h-6 w-full" />
                        </div>
                      </span>
                      <span className="basis-3/4">
                        {<Skeleton className="w-24 h-6" />}
                      </span>
                    </div>
                    <div className="flex">
                      <span className="flex basis-1/4 text-muted-foreground">
                        <div className="basis-1/2">
                          <Skeleton className="h-6 w-full" />
                        </div>
                      </span>
                      <span className="basis-3/4">
                        {<Skeleton className="w-32 h-6" />}
                      </span>
                    </div>
                    <div className="flex">
                      <span className="flex basis-1/4 text-muted-foreground">
                        <div className="basis-1/2">
                          <Skeleton className="h-6 w-full" />
                        </div>
                      </span>
                      <span className="basis-3/4">
                        {<Skeleton className="w-24 h-6" />}
                      </span>
                    </div>
                    <div className="flex">
                      <span className="flex basis-1/4 text-muted-foreground">
                        <div className="basis-1/2">
                          <Skeleton className="h-6 w-full" />
                        </div>
                      </span>
                      <span className="basis-3/4 font-normal">
                        <div className="flex flex-col gap-2">
                          <Skeleton className="w-full h-6" />
                          <Skeleton className="w-full h-6" />
                          <Skeleton className="w-full h-6" />
                        </div>
                      </span>
                    </div>
                  </div>
                </div>
              </CardContent>
            </Card>
            <Card className="flex-1">
              <CardContent className=" flex gap-5 mt-5">
                <div className="flex flex-col items-center gap-4">
                  <div className="rounded-sm border overflow-clip w-fit">
                    <Skeleton className="h-[120px] w-[120px]"></Skeleton>
                  </div>
                  <Skeleton className="w-[120px] h-8" />
                </div>

                <div className="flex-col flex-1 flex gap-5">
                  <div className="flex lg:gap-4 gap-3 xl:flex-col sm:flex-row flex-col">
                    <div className="flex-1">
                      <Skeleton className="w-full h-6" />
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
                </div>
              </CardContent>
            </Card>
          </div>
        </div>
      </div>
    </div>
  );
};

export default BookDetailSkeleton;
