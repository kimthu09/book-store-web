import { IconType } from "react-icons";

export type Book = {
  id: string;
  name: string;

  isActive: boolean;
  categoryIds: [];

  authorIds: [];
};

export type Category = {
  id: string;
  name: string;
};

export type Author = {
  id: string;
  name: string;
};
export interface CategoryListProps {
  checkedCategory: Array<string>;
  onCheckChanged: (idCate: string) => void;
  canAdd?: boolean;
  readonly?: boolean;
}
export interface AuthorListProps {
  checkedAuthor: Array<string>;
  onCheckChanged: (idAuthor: string) => void;
  canAdd?: boolean;
  readonly?: boolean;
}
export interface SupplierListProps {
  supplier: string;
  setSupplier: (supplier: string) => void;
  canAdd?: boolean;
}
export interface BookListProps {
  book: Partial<Book>;
  setBook: (book: Partial<Book>) => void;
  isNew: boolean;
  setIsNew: (isNew: boolean) => void;
}

export type SidebarItem = {
  title: string;
  href: string;
  icon?: IconType;
  submenu?: boolean;
  subMenuItems?: SidebarItem[];
};
