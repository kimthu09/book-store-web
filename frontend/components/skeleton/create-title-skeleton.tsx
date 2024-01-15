import React from "react";
import DropdownSkeleton from "./dropdown-skeleton";

const CreateTitleSkeleton = () => {
  return (
    <div className="col items-center px-6">
      <div className="w-full py-6 pt-0">
        <div className="flex flex-col gap-4">
          <div className="flex-col flex gap-5">
            <div className="basis-1/3">
              <DropdownSkeleton />
            </div>
            <div className="flex-1">
              <DropdownSkeleton />
            </div>

            <DropdownSkeleton />
            <DropdownSkeleton />

            <div className="flex-1">
              <DropdownSkeleton />
            </div>
          </div>
          <div className="flex gap-4 py-4  justify-end">
            <DropdownSkeleton />
          </div>
        </div>
      </div>
    </div>
  );
};

export default CreateTitleSkeleton;
