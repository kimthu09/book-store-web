import { SidebarItem } from "@/types";
import { GrBook } from "react-icons/gr";
import { MdOutlineWarehouse } from "react-icons/md";
import { GoPeople, GoPerson } from "react-icons/go";
import { AiOutlineLineChart } from "react-icons/ai";
import { LuClipboardList } from "react-icons/lu";
import { FaRegHandshake } from "react-icons/fa";
import { BsShop } from "react-icons/bs";
import { z } from "zod";

export const endPoint = "http://localhost:8080";

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

export const adminSidebarItems: SidebarItem[] = [
  {
    title: "Bán hàng",
    href: "/sale",
    icon: BsShop,
    submenu: false,
  },
  {
    title: "Báo cáo",
    href: "/storereport/stock",
    icon: AiOutlineLineChart,
    submenu: true,
    subMenuItems: [
      { title: "Báo cáo tồn kho", href: "/storereport/stock" },
      { title: "Báo cáo nợ", href: "/storereport/debt" },
      { title: "Báo cáo mặt hàng", href: "/storereport/sale" },
    ],
  },
  {
    title: "Quản lý hóa đơn",
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
      { title: "Danh sách đầu sách", href: "/product/title" },
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
    subMenuItems: [
      { title: "Nhập kho", href: "/stock/import" },
      { title: "Kiểm kho", href: "/stock/inventorycheck" },
    ],
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

export const sidebarItems: SidebarItem[] = [
  {
    title: "Bán hàng",
    href: "/sale",
    icon: BsShop,
    submenu: false,
  },
  {
    title: "Báo cáo",
    href: "/storereport/stock",
    icon: AiOutlineLineChart,
    submenu: true,
    subMenuItems: [
      { title: "Báo cáo tồn kho", href: "/storereport/stock" },
      { title: "Báo cáo nợ", href: "/storereport/debt" },
      { title: "Báo cáo mặt hàng", href: "/storereport/sale" },
    ],
  },
  {
    title: "Quản lý hóa đơn",
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
      { title: "Danh sách đầu sách", href: "/product/title" },
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
    subMenuItems: [
      { title: "Nhập kho", href: "/stock/import" },
      { title: "Kiểm kho", href: "/stock/inventorycheck" },
    ],
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
    submenu: false,
  },
];
