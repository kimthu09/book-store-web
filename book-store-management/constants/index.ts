import {
  Category,
  Book,
  SidebarItem,
  ImportNote,
  StatusString,
  Supplier,
} from "@/types";
import { GrBook } from "react-icons/gr";
import { MdOutlineWarehouse } from "react-icons/md";
import { GoPeople } from "react-icons/go";

export const books: Book[] = [
  {
    id: "1",
    name: "A meal a day",
    nxb: "NXB Thanh nien",
    quantity: 23,
    price: 21000,
    status: true,
    category: "Giao duc",
  },
  {
    id: "2",
    name: "Seven days a week",
    nxb: "NXB Tri thuc",
    quantity: 57,
    price: 18000,
    status: true,
    category: "Truyen tranh",
  },
  {
    id: "3",
    name: "Travel destinations",
    nxb: "NXB Tri thuc",
    quantity: 23,
    price: 46000,
    status: false,
    category: "Giao duc",
  },
  {
    id: "4",
    name: "Brain Workout",
    nxb: "NXB Sunflower",
    quantity: 23,
    price: 52000,
    status: false,
    category: "Van hoc",
  },
  {
    id: "5",
    name: "The Nightmare",
    nxb: "NXB Global Trees",
    quantity: 23,
    price: 98000,
    status: false,
    category: "Truyen tranh",
  },
  {
    id: "6",
    name: "David's journey to the land of wishes",
    nxb: "NXB Tourism",
    quantity: 23,
    price: 79000,
    status: false,
    category: "Van hoc",
  },
  {
    id: "7",
    name: "How to plan an awesome trip",
    nxb: "NXB Tourism",
    quantity: 23,
    price: 55000,
    status: false,
    category: "Van hoc",
  },
];

export const categories: Category[] = [
  {
    id: "1",
    name: "Van hoc",
  },
  {
    id: "2",
    name: "Giao duc",
  },
  {
    id: "3",
    name: "Truyen tranh",
  },
  {
    id: "4",
    name: "Trinh tham",
  },
];
export const importNotes: ImportNote[] = [
  {
    id: "NGAY1",
    supplierId: "DT01",
    totalPrice: 5060000,
    status: StatusString.Inprogress,
    createAt: new Date(),
    createBy: "NV002",
  },
  {
    id: "NGAY2",
    supplierId: "DT01",
    totalPrice: 3720000,
    status: StatusString.Done,
    createAt: new Date(2023, 9, 8),
    createBy: "NV002",
  },
  {
    id: "NGAY3",
    supplierId: "DT01",
    totalPrice: 4660000,
    status: StatusString.Cancel,
    createAt: new Date(2023, 10, 1),
    createBy: "NV002",
  },
];
export const suppliers: Supplier[] = [
  {
    id: "Ncc01",
    name: "Cong ty Sach Hoa Nghien",
  },
  {
    id: "Ncc02",
    name: "Cong ty Quang Nam",
  },
  {
    id: "Ncc03",
    name: "Cong ty Hoa Lan",
  },
  {
    id: "Ncc03",
    name: "Cong ty ForDream",
  },
];
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
    title: "Quản lý nhân viên",
    href: "/",
    icon: GoPeople,
  },
];
