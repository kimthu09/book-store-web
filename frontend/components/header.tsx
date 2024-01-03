"use client";

import Link from "next/link";
import Image from "next/image";
import { logOut } from "@/lib/auth/action";
import { Button } from "./ui/button";
import { LuLogOut } from "react-icons/lu";

const Header = () => {
  return (
    <div className="flex z-10 bg-white w-[100%] border-b border-gray-200 px-4 pl-12">
      <div className="flex flex-1 h-[47px] items-center justify-between">
        <div className="flex items-center space-x-4 justify-between">
          <Link
            href="/"
            className="flex flex-row space-x-3 items-center justify-center md:hidden"
          >
            <span className="font-semibold text-lg flex ">Book Store</span>
          </Link>
        </div>
        <div className="self-center flex">
          <form action={logOut}>
            <Button variant={"link"}>
              <div className="flex gap-2 text-primary">
                Đăng xuất
                <LuLogOut className="h-5 w-5 " />
              </div>
            </Button>
          </form>
          <Button className="p-0 h-fit rounded-full overflow-clip">
            <div>
              <Image
                src="/no-image.jpg"
                className="h-9 w-9"
                alt="avatar"
                width={36}
                height={36}
              ></Image>
            </div>
          </Button>
        </div>
      </div>
    </div>
  );
};

export default Header;
