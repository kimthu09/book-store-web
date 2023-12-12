import {
  Category,
  Book,
  SidebarItem,
  ImportNote,
  StatusNote,
  Supplier,
  ImportDetail,
  SupplierDebt,
} from "@/types";
import { GrBook } from "react-icons/gr";
import { MdOutlineWarehouse } from "react-icons/md";
import { GoPeople, GoPerson } from "react-icons/go";
import { PiHandshake } from "react-icons/pi";
import { FaRegHandshake } from "react-icons/fa";

export const apiKey =
  "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJJZCI6ImczVzIxQTdTUiIsInJvbGUiOiIifSwiZXhwIjoxNzA0OTEwNzI4LCJpYXQiOjE3MDIzMTg3Mjh9.orf7gFNsAMyYoOlVfcUt8hcE1GkVwgBiyoj3OwBIdYk";
// export const books: Book[] = [
//   {
//     id: "1",
//     name: "A meal a day",
//     nxb: "NXB Thanh nien",
//     quantity: 23,
//     price: 21000,
//     status: true,
//     category: "Giao duc",
//   },
//   {
//     id: "2",
//     name: "Seven days a week",
//     nxb: "NXB Tri thuc",
//     quantity: 57,
//     price: 18000,
//     status: true,
//     category: "Truyen tranh",
//   },
//   {
//     id: "3",
//     name: "Travel destinations",
//     nxb: "NXB Tri thuc",
//     quantity: 23,
//     price: 46000,
//     status: false,
//     category: "Giao duc",
//   },
//   {
//     id: "4",
//     name: "Brain Workout",
//     nxb: "NXB Sunflower",
//     quantity: 23,
//     price: 52000,
//     status: false,
//     category: "Van hoc",
//   },
//   {
//     id: "5",
//     name: "The Nightmare",
//     nxb: "NXB Global Trees",
//     quantity: 23,
//     price: 98000,
//     status: false,
//     category: "Truyen tranh",
//   },
//   {
//     id: "6",
//     name: "David's journey to the land of wishes",
//     nxb: "NXB Tourism",
//     quantity: 23,
//     price: 79000,
//     status: false,
//     category: "Van hoc",
//   },
//   {
//     id: "7",
//     name: "How to plan an awesome trip",
//     nxb: "NXB Tourism",
//     quantity: 23,
//     price: 55000,
//     status: false,
//     category: "Van hoc",
//   },
// ];

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
    supplierId: "Ncc01",
    totalPrice: 5060000,
    status: StatusNote.Inprogress,
    createAt: new Date(),
    createBy: "NV002",
  },
  {
    id: "NGAY2",
    supplierId: "Ncc01",
    totalPrice: 3720000,
    status: StatusNote.Done,
    createAt: new Date(2023, 9, 8),
    createBy: "NV002",
  },
  {
    id: "NGAY3",
    supplierId: "Ncc01",
    totalPrice: 4660000,
    status: StatusNote.Cancel,
    createAt: new Date(2023, 10, 1),
    createBy: "NV002",
  },
];

// export const importDetails: ImportDetail[] = [
//   {
//     book: books[0],
//     idNote: "NGAY1",
//     quantity: 20,
//     price: 21000,
//   },
//   {
//     book: books[1],
//     idNote: "NGAY1",
//     quantity: 30,
//     price: 52000,
//   },
//   {
//     book: books[2],
//     idNote: "NGAY1",
//     quantity: 10,
//     price: 23000,
//   },
//   {
//     book: books[3],
//     idNote: "NGAY2",
//     quantity: 30,
//     price: 46000,
//   },
//   {
//     book: books[4],
//     idNote: "NGAY2",
//     quantity: 15,
//     price: 18000,
//   },
//   {
//     book: books[5],
//     idNote: "NGAY3",
//     quantity: 20,
//     price: 73000,
//   },
//   {
//     book: books[6],
//     idNote: "NGAY3",
//     quantity: 30,
//     price: 53000,
//   },
// ];
// export const suppliers: Supplier[] = [
//   {
//     id: "Ncc01",
//     name: "Cong ty Sach Hoa Nghien",
//     phone: "0123456789",
//     debt: 2000000,
//     email: "ctyHoaNghien@gmail.com",
//   },
//   {
//     id: "Ncc02",
//     name: "Cong ty Quang Nam",
//     phone: "0123456789",
//     debt: 4600000,
//   },
//   {
//     id: "Ncc03",
//     name: "Cong ty Hoa Lan",
//     phone: "0123456789",
//     debt: 1200000,
//   },
//   {
//     id: "Ncc04",
//     name: "Cong ty ForDream",
//     phone: "0123456789",
//     debt: 2000000,
//   },
// ];
export const supplierDebts: SupplierDebt[] = [
  {
    id: "PN001",
    idSupplier: "Ncc01",
    amount: 23000000,
    amountLeft: 3000000,
    createBy: "NV01",
    createAt: new Date(2023, 4, 24),
  },
  {
    id: "PN002",
    idSupplier: "Ncc01",
    amount: 12000000,
    amountLeft: 500000,
    createBy: "NV01",
    createAt: new Date(2023, 4, 12),
  },
  {
    id: "PN003",
    idSupplier: "Ncc01",
    amount: 1050000,
    amountLeft: 50000,
    createBy: "NV01",
    createAt: new Date(2023, 4, 3),
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
