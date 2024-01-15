import { IconType } from "react-icons";

export type Book = {
  id: string;
  name: string;
  bookTitle: {
    id: string;
    name: string;
    desc: string;
    authors: {
      id: string;
      name: string;
    }[];
    categories: {
      id: string;
      name: string;
    }[];
  };
  publisher: {
    id: string;
    name: string;
  };
  edition: number;
  quantity: number;
  listedPrice: number;
  sellPrice: number;
  importPrice: number;
  isActive: boolean;
  image: string;
};
export type BookProps = {
  id: string;
  name: string;
  img?: string;
  bookTitle: {
    id: string;
    name: string;
    desc: string;
    authors: {
      id: string;
      name: string;
    }[];
    categories: {
      id: string;
      name: string;
    }[];
  };
  publisher: {
    id: string;
    name: string;
  };
  edition: number;
  quantity: number;
  listedPrice: number;
  sellPrice: number;
  importPrice: number;
};

export type BookTitle = {
  id: string;
  name: string;
  desc: string;
  isActive: boolean;
  createdAt: Date;
  categories: { id: string; name: string }[];
  authors: { id: string; name: string }[];
};
export type Invoice = {
  createdAt: Date;
  createdBy: {
    id: string;
    name: string;
  };
  id: string;
  totalPrice: number;
  customer: {
    id: string;
    name: string;
    phone: string;
  };
  amountReceived: number;
  amountPriceUsePoint: number;
  pointUse: number;
  pointReceive: number;
};
export type ImportNote = {
  id: string;
  totalPrice: number;
  status: StatusNote;
  closedAt?: Date;
  closedBy?: {
    id: string;
    name: string;
  };
  createdAt: Date;
  createdBy: {
    id: string;
    name: string;
  };
  supplier: {
    id: string;
    name: string;
    phone: string;
  };
};
export type ImportNoteDetail = {
  book: {
    id: string;
    name: string;
  };
  price: number;
  qtyImport: number;
};
export type CheckNoteDetail = {
  book: {
    id: string;
    name: string;
  };
  difference: number;
  final: number;
  initial: number;
};
export type CheckNote = {
  id: string;
  qtyAfterAdjust: number;
  qtyDifferent: number;
  createdAt: Date;
  createdBy: {
    id: string;
    name: string;
  };
};
export type Supplier = {
  id: string;
  name: string;
  email?: string;
  phone: string;
  debt: number;
};
export type SupplierDebt = {
  createdAt: Date;
  createdBy: {
    id: string;
    name: string;
  };
  id: string;
  qty: number;
  qtyLeft: number;
  supplierId: string;
  type: string;
};

export type DebtReport = {
  timeFrom: Date;
  timeTo: Date;
  initial: number;
  debt: number;
  pay: number;
  final: number;
  details: [
    {
      debt: number;
      final: number;
      initial: number;
      pay: number;
      supplier: {
        id: string;
        name: string;
        phone: string;
      };
    }
  ];
};
export type Customer = {
  id: string;
  name: string;
  email?: string;
  phone: string;
  point: number;
};
export type CustomerInvoice = {
  id: string;
  totalPrice: number;
  amountReceived: number;
  amountPriceUsePoint: number;
  pointUse: number;
  pointReceive: number;
  createdBy: {
    id: string;
    name: string;
  };
  createdAt: Date;
};
export type ShopGeneral = {
  name: string;
  email?: string;
  phone: string;
  address: string;
  wifiPass: string;
  accumulatePointPercent: number;
  usePointPercent: number;
};
export type DebtReportDetail = {
  debt: number;
  final: number;
  initial: number;
  pay: number;
  supplier: {
    id: string;
    name: string;
    phone: string;
  };
};

export type SaleReportDetail = {
  amount: number;
  book: {
    id: string;
    name: string;
  };
  totalSales: number;
};

export type SaleReport = {
  timeFrom: Date;
  timeTo: Date;
  total: number;
  amount: number;
  details: [
    {
      amount: number;
      book: {
        id: string;
        name: string;
      };
      totalSales: number;
    }
  ];
};

export type StockReportDetail = {
  book: {
    id: string;
    name: string;
  };
  final: number;
  initial: number;
  import: number;
  modify: number;
  sell: number;
};

export type StockReport = {
  id: string;
  timeFrom: Date;
  timeTo: Date;
  initial: number;
  sell: number;
  import: number;
  modify: number;
  final: number;
  details: [
    {
      book: {
        id: string;
        name: string;
      };
      initial: number;
      sell: number;
      import: number;
      modify: number;
      final: number;
    }
  ];
};

export enum StatusNote {
  Inprogress = "InProgress",
  Done = "Done",
  Cancel = "Cancel",
}
export enum StatusActive {
  Active = "Đang giao dịch",
  InActive = "Ngừng giao dịch",
}
export type Category = {
  id: string;
  name: string;
};

export type Author = {
  id: string;
  name: string;
};
export type Publisher = {
  id: string;
  name: string;
};

export type Staff = {
  address?: string;
  email: string;
  id: string;
  isActive: boolean;
  name: string;
  img: string;
  phone?: string;
  role: {
    id: string;
    name: string;
  };
};
export type Role = {
  id: string;
  name: string;
};
export type RoleFunction = {
  id: string;
  description: string;
  groupName: string;
};
export interface CategoryListProps {
  checkedCategory: Array<string>;
  onCheckChanged: (idCate: string) => void;
  canAdd?: boolean;
  readonly?: boolean;
  isEdit?: boolean;
  onRemove?: (index: number) => void;
}
export interface AuthorListProps {
  checkedAuthor: Array<string>;
  onCheckChanged: (idAuthor: string) => void;
  canAdd?: boolean;
  readonly?: boolean;
  isEdit?: boolean;
  onRemove?: (index: number) => void;
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
export interface RoleListProps {
  role: string;
  setRole: (role: string) => void;
}
export interface StaffListProps {
  staff: string;
  setStaff: (role: string) => void;
}
export interface TitleListProps {
  handleTitleSet: (titleId: string) => void;
}
export interface StatusListProps {
  status?: boolean;
  setStatus: (role: boolean) => void;
  display: { trueText: string; falseText: string };
}
export interface PublisherListProps {
  publisherId: string;
  setPublisherId: (nameId: string) => void;
  canAdd?: boolean;
  readOnly?: boolean;
}
export type PagingProps = {
  page: number;
  limit: number;
  total: number;
};
export type FilterValue = {
  filters: {
    type: string;
    value: string;
  }[];
};

export type Dashboard = {
  timeFrom: Date;
  timeTo: Date;
  totalSale: number;
  totalCustomer: number;
  totalSold: number;
  totalPoint: number;
  topSoldBooks: [{ id: string; name: string; qty: number; sale: number }];
  chartPriceComponents: [
    {
      time: Date;
      value: number;
    }
  ];
  chartProfitComponents: [
    {
      time: Date;
      value: number;
    }
  ];
};

export type CardDashboardInfo = {
  title: string;
  value: string;
  icon: any;
};

export type TopSoldFood = {
  id: string;
  name: string;
  qty: number;
  sale: number;
};

export type CharComponent = {
  time: Date;
  value: number;
};
