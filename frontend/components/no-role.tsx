import React from "react";
import LoadingSpinner from "./ui/loading-spinner";
import { RiUserForbidLine } from "react-icons/ri";

const NoRole = () => {
  return (
    <div className="flex justify-center gap-5 flex-col items-center p-6">
      <RiUserForbidLine className={"h-12 w-12 text-primary"} />
      <p>Bạn không có quyền !</p>
    </div>
  );
};

export default NoRole;
