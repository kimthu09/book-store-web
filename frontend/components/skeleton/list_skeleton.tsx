import { Skeleton } from "../ui/skeleton";

type TableSkeletonProps = {
  numberRow: number;
};

const ListSkeleton: React.FC<TableSkeletonProps> = ({ numberRow }) => {
  return (
    <div className="flex flex-col gap-4">
      {[...Array(numberRow)].map((_, index) => (
        <Skeleton key={index} className={`w-full  h-6`} />
      ))}
    </div>
  );
};

export default ListSkeleton;
