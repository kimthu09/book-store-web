import { AspectRatio } from "../ui/aspect-ratio";
import { Skeleton } from "../ui/skeleton";
import DropdownSkeleton from "./dropdown-skeleton";
export const generateFields = (length: number): string[] => {
  return Array.from({ length }, (_, index) => `Column ${index + 1}`);
};

const ProductSkeleton = () => {
  const products = generateFields(10);
  return (
    <div className="flex flex-col gap-6">
      <div className="flex items-end">
        <DropdownSkeleton />
      </div>

      {/* Category list */}
      <div className="flex flex-wrap gap-2">
        <Skeleton className="h-8 w-32" />
        <Skeleton className="h-8 w-24" />
        <Skeleton className="h-8 w-28" />
        <Skeleton className="h-8 w-32" />
        <Skeleton className="h-8 w-16" />
      </div>
      <h1 className="text-lg">Sản phẩm</h1>

      {/* Product list */}
      <div className="grid 2xl:grid-cols-5 xl:grid-cols-4 lgr:grid-cols-3 md:grid-cols-2 sm:grid-cols-4 grid-cols-3 gap-4">
        {products.map((item) => (
          <AspectRatio ratio={2 / 3} key={item}>
            <Skeleton className="h-full w-full" key={item} />
          </AspectRatio>
        ))}
      </div>
    </div>
  );
};

export default ProductSkeleton;
