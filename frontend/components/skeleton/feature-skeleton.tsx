import { Skeleton } from "../ui/skeleton";
import { generateFields } from "./product-skeleton";

const FeatureSkeleton = () => {
  const features = generateFields(32);
  return (
    <div className="grid xl:grid-cols-3 lg:grid-cols-2 md:grid-cols-1 sm:grid-cols-2 grid-cols-1 gap-y-6 gap-x-4">
      {features.map((value) => {
        return (
          <div key={value}>
            <div className="flex gap-2">
              <Skeleton className="h-4 w-4" />
              <Skeleton className="h-4 w-40" />
            </div>
          </div>
        );
      })}
    </div>
  );
};

export default FeatureSkeleton;
