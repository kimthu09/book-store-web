import {
  Category,
  Book,
  SidebarItem,
  ImportNote,
  StatusNote,
  Supplier,
  ImportDetail,
  SupplierDebt,
  RoleFunction,
} from "@/types";
import { GrBook } from "react-icons/gr";
import { MdOutlineWarehouse } from "react-icons/md";
import { GoPeople, GoPerson } from "react-icons/go";
import { PiHandshake } from "react-icons/pi";
import { FaRegHandshake } from "react-icons/fa";
import { z } from "zod";

export const apiKey =
  "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJJZCI6ImczVzIxQTdTUiIsInJvbGUiOiIifSwiZXhwIjoxNzA1MTY1MjQzLCJpYXQiOjE3MDI1NzMyNDN9.bDwKkz9OZdwp15HDCESuQ0zoMOXHbCGWfN59u4Y9KkI";
export const required = z.string().min(1, "Không để trống trường này");

export const statuses = [
  {
    isActive: true,
    label: "Đang giao dịch",
  },
  {
    isActive: false,
    label: "Ngừng giao dịch",
  },
];

export const noteStatus = [
  {
    label: "Đang xử lý",
  },
  {
    label: "Đã nhập",
  },
  {
    label: "Đã huỷ",
  },
];
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
  {
    title: "Quản lý nhà cung cấp",
    href: "/supplier",
    icon: FaRegHandshake,
    submenu: false,
  },
  {
    title: "Quản lý nhân viên",
    href: "/staff",
    icon: GoPeople,
    submenu: true,
    subMenuItems: [
      { title: "Danh sách nhân viên", href: "/staff" },
      { title: "Phân quyền nhân viên", href: "/staff/role" },
    ],
  },
];
