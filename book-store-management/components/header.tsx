"use client";

import Link from "next/link";
import Image from "next/image";

const Header = () => {
  return (
    <div className=" flex fixed right-0 z-10 bg-white w-[100%] border-b border-gray-200">
      <div className="flex h-[47px] items-center justify-between px-4">
        <div className="flex items-center space-x-4">
          <Link
            href="/"
            className="flex flex-row space-x-3 items-center justify-center md:hidden"
          >
            <span className="font-semibold text-lg flex ">Book Store</span>
          </Link>
        </div>
      </div>
    </div>
  );
};

export default Header;
