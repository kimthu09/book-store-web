import React from "react";
import LoadingSpinner from "./ui/loading-spinner";

const Loading = () => {
  return (
    <div className="flex justify-center gap-5 flex-col items-center p-6">
      <LoadingSpinner className={"h-12 w-12 text-primary"} />
      <p>Vui lòng đợi trong giây lát</p>
    </div>
  );
};

export default Loading;
