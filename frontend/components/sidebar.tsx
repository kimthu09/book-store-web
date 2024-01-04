"use client";

import Image from "next/image";
import Link from "next/link";
import { RxHamburgerMenu } from "react-icons/rx";
import { useEffect, useState } from "react";

import { usePathname } from "next/navigation";
import { LuChevronDown } from "react-icons/lu";
import { adminSidebarItems, sidebarItems } from "@/constants";
import { SidebarItem } from "@/types";
import { useCurrentUser } from "@/hooks/use-user";
import { isAdmin } from "@/lib/utils";

export default function Sidebar() {
  const [isCollapse, toggleIsCollapse] = useState(false);

  const toggleSidebarHandler = () => {
    toggleIsCollapse((prev) => !prev);
  };
  const { currentUser } = useCurrentUser();

  const [items, setItems] = useState<SidebarItem[]>(sidebarItems);
  useEffect(() => {
    if (currentUser) {
      if (isAdmin({ currentUser: currentUser })) {
        setItems(adminSidebarItems);
      }
    }
  }, [currentUser]);
  return (
    <div className="md:flex hidden z-20">
      <aside
        className={`bg-white h-screen p-1 transition-all shadow-md overflow-auto ${
          isCollapse ? "w-[3.8rem]" : "w-64"
        }`}
      >
        <nav className="w-full">
          <div className={`flex items-center my-4 h-[64px]`}>
            <Link href="/">
              <div
                className={`flex align-middle justify-center items-center gap-4 h-[64px] w-[64px]  rounded-xl ${
                  isCollapse ? "hidden" : "flex"
                }`}
              >
                <Image
                  src="/android-chrome-192x192.png"
                  alt="logo"
                  width={50}
                  height={50}
                ></Image>
              </div>
            </Link>

            <Link href="/">
              <p
                className={`text-lg ml-2 font-semibold overflow-hidden whitespace-nowrap ${
                  isCollapse ? "hidden" : "block"
                }`}
              >
                Book Store
              </p>
            </Link>
            <div
              className={`hover:text-primary cursor-pointer ${
                isCollapse ? "m-auto p-1" : "ml-auto mr-1 p-1 "
              }`}
              onClick={toggleSidebarHandler}
            >
              <RxHamburgerMenu className="w-6 h-6 " />
            </div>
          </div>

          <ul className="sidebar__list">
            {items.map((item) => (
              <li className="sidebar__item" key={item.title}>
                <MenuItem item={item} isCollapse={isCollapse}></MenuItem>
              </li>
            ))}
          </ul>
        </nav>
      </aside>
    </div>
  );
}

const MenuItem = ({
  item,
  isCollapse,
}: {
  item: SidebarItem;
  isCollapse: boolean;
}) => {
  const pathname = usePathname();
  const [subMenuOpen, setSubMenuOpen] = useState(false);
  const toggleSubMenu = () => {
    setSubMenuOpen(!subMenuOpen);
  };

  return (
    <div>
      {item.submenu ? (
        <>
          <div onClick={toggleSubMenu}>
            <div
              className={`flex text-base no-underline text-black px-4 py-3 mb-2 rounded-md overflow-hidden max-h-15 hover:bg-blue-50 cursor-pointer ${
                pathname.includes(item.href) ? "bg-zinc-100" : ""
              }`}
            >
              {item.icon ? (
                <>
                  <span>
                    <item.icon className="sidebar__icon" />
                  </span>
                </>
              ) : null}

              <span
                className={`ml-2 text-lg overflow-hidden  whitespace-nowrap ${
                  isCollapse ? "hidden opacity-0" : "visible opacity-100"
                }`}
              >
                {item.title}
              </span>
              <div
                className={`ml-auto self-center ${
                  isCollapse ? "hidden" : "visible"
                } ${subMenuOpen && !isCollapse ? "rotate-180" : ""} flex`}
              >
                <LuChevronDown className="sidebar__icon" />
              </div>
            </div>
          </div>
          {subMenuOpen && !isCollapse && (
            <div className="my-2 ml-12 flex flex-col space-y-4">
              {item.subMenuItems?.map((subItem, idx) => {
                return (
                  <Link
                    key={idx}
                    href={subItem.href}
                    className={`flex text-base no-underline text-black rounded-md overflow-hidden hover:text-primary`}
                  >
                    {subItem.icon ? (
                      <>
                        <span>
                          <subItem.icon className="sidebar__icon" />
                        </span>
                      </>
                    ) : null}

                    <span
                      className={`ml-2 text-lg overflow-hidden  whitespace-nowrap ${
                        isCollapse ? "hidden" : "visible"
                      } ${
                        subItem.href === pathname
                          ? "text-primary font-medium"
                          : ""
                      }`}
                    >
                      {subItem.title}
                    </span>
                  </Link>
                );
              })}
            </div>
          )}
        </>
      ) : (
        <Link
          href={item.href}
          className={`flex text-base no-underline text-black px-4 py-3 mb-2 rounded-md overflow-hidden max-h-15 hover:bg-blue-50 ${
            item.href === pathname ? "bg-zinc-100" : ""
          }`}
        >
          {item.icon ? (
            <>
              <span>
                <item.icon className="sidebar__icon" />
              </span>
            </>
          ) : null}

          <span
            className={`ml-2 text-lg overflow-hidden  whitespace-nowrap ${
              isCollapse ? "hidden" : "visible"
            }`}
          >
            {item.title}
          </span>
        </Link>
      )}
    </div>
  );
};
