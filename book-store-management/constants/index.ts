import { Category, Book, SidebarItem } from "@/types";
import { GrBook } from "react-icons/gr";
import { MdOutlineWarehouse } from "react-icons/md";
import { GoPeople, GoPerson } from "react-icons/go";
import { PiHandshake } from "react-icons/pi";
import { FaRegHandshake } from "react-icons/fa";

export const apiKey =
  "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJJZCI6ImczVzIxQTdTUiIsInJvbGUiOiIifSwiZXhwIjoxNzA1MDI4NjAyLCJpYXQiOjE3MDI0MzY2MDJ9.4KCH-72jMCwxtlyJgMRzT1QS46T5lbqLngSbeNs2rX0";

export const sidebarItems: SidebarItem[] = [
  {
    title: "Quản lý sách",
    href: "/books",
    icon: GrBook,
  },
  {
    title: "Quản lý kho",
    href: "/stock",
    icon: MdOutlineWarehouse,
    submenu: true,
    subMenuItems: [{ title: "Nhập kho", href: "/stock/import" }],
  },
  // {
  //   title: "Quản lý nhân viên",
  //   href: "/",
  //   icon: GoPeople,
  // },
  {
    title: "Quản lý nhà cung cấp",
    href: "/supplier",
    icon: FaRegHandshake,
    submenu: false,
  },
];
