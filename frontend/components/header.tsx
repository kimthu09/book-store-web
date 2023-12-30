"use client";

import Link from "next/link";
import Image from "next/image";
import { logOut } from "@/lib/auth/action";

const Header = () => {
  return (
    <div className="flex z-10 bg-white w-[100%] border-b border-gray-200">
      <div className="flex h-[47px] items-center justify-between px-4">
        <div className="flex items-center space-x-4">
          <Link
            href="/"
            className="flex flex-row space-x-3 items-center justify-center md:hidden"
          >
            <span className="font-semibold text-lg flex ">Book Store</span>
          </Link>
          <form className="self-end" action={logOut}>
            <button>Logout</button>
          </form>
        </div>
      </div>
    </div>
  );
};

export default Header;
