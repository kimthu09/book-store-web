import { SidebarItem } from "@/types";
import { GrBook } from "react-icons/gr";
import { MdOutlineWarehouse } from "react-icons/md";
import { GoPeople, GoPerson } from "react-icons/go";
import { AiOutlineLineChart } from "react-icons/ai";
import { LuClipboardList } from "react-icons/lu";
import { FaRegHandshake } from "react-icons/fa";
import { z } from "zod";

export const apiKey =
  "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJJZCI6ImczVzIxQTdTUiIsInJvbGUiOiIifSwiZXhwIjoxNzA1MTk3NDIxLCJpYXQiOjE3MDI2MDU0MjF9.SJxvTYBk6fNodXd_8M0vqM4hjJpCg5MkonTtbLjOtNU";
export const endPoint = "http://103.57.221.113:8080";
export const required = z.string().min(1, "Không để trống trường này");
export const phoneRegex = new RegExp(/(0[3|5|7|8|9])+([0-9]{8})\b/g);
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
    title: "Báo cáo",
    href: "/report",
    icon: AiOutlineLineChart,
    submenu: true,
    subMenuItems: [{ title: "Báo cáo nợ", href: "/report/debt" }],
  },
  {
    title: "Hóa đơn",
    href: "/invoice",
    icon: LuClipboardList,
    submenu: false,
  },
  {
    title: "Quản lý sản phẩm",
    href: "/product",
    icon: GrBook,
    submenu: true,
    subMenuItems: [
      { title: "Danh sách sách", href: "/product/books" },
      { title: "Thể loại", href: "/product/categories" },
      { title: "Tác giả", href: "/product/authors" },
      { title: "Nhà xuất bản", href: "/product/publishers" },
    ],
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
