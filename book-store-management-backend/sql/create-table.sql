create table Author
(
    id   varchar(12) not null
        primary key,
    name text        not null
);

create table Book
(
    id         varchar(12)          not null,
    bookInfoId varchar(12)          not null,
    edition    int                  not null,
    qty        float      default 0 null,
    price      float                not null,
    salePrice  float                not null,
    isActive   tinyint(1) default 1 not null,
    primary key (id, bookInfoId)
);

create index bookInfoId
    on Book (bookInfoId);

create table BookChangeHistory
(
    id         varchar(12)                       not null,
    bookId     varchar(12)                       not null,
    amount     float                             not null,
    amountLeft float                             not null,
    type       enum ('Sell', 'Import', 'Modify') not null,
    primary key (id, bookId)
);

create table BookInfo
(
    bookId      varchar(12) not null,
    name        text        not null,
    authorIds   text        not null,
    categoryIds text        not null,
    publisherId varchar(12) not null,
    `desc`      text        null,
    primary key (bookId, publisherId),
    constraint BookInfo_ibfk_1
        foreign key (bookId) references Book (bookInfoId)
);

create index publisherId
    on BookInfo (publisherId);

create table Category
(
    id   varchar(12) not null
        primary key,
    name varchar(50) not null
);

create table Feature
(
    id          varchar(12) not null
        primary key,
    description text        null
);

create table ImportNote
(
    id         varchar(12)                                                not null
        primary key,
    supplierId varchar(12)                                                not null,
    totalPrice float                                 default 0            null,
    status     enum ('InProgress', 'Done', 'Cancel') default 'InProgress' null,
    createBy   varchar(12)                                                not null,
    closeBy    varchar(12)                                                null,
    createAt   datetime                              default (now())      null,
    closeAt    datetime                                                   null
);

create table ImportNoteDetail
(
    importNoteId varchar(12)     not null,
    bookId       varchar(12)     not null,
    price        float           not null,
    qtyImport    float default 0 null,
    primary key (importNoteId, bookId)
);

create table InventoryCheckNote
(
    id             varchar(12)              not null
        primary key,
    qtyDifferent   float                    not null,
    qtyAfterAdjust float                    not null,
    createBy       varchar(12)              not null,
    createAt       datetime default (now()) null
);

create table InventoryCheckNoteDetail
(
    inventoryCheckNoteId varchar(12) not null,
    bookId               varchar(12) not null,
    initial              float       not null,
    difference           float       not null,
    final                float       not null,
    primary key (inventoryCheckNoteId, bookId)
);

create table Invoice
(
    id          varchar(13)              not null
        primary key,
    totalPrice  float                    not null,
    qtyReceived float                    not null,
    createAt    datetime default (now()) null,
    createBy    varchar(13)              not null
);

create table InvoiceDetail
(
    invoiceId varchar(13) not null,
    bookId    varchar(13) not null,
    qty       float       not null,
    unitPrice float       not null,
    primary key (invoiceId, bookId)
);

create table MUser
(
    id       varchar(12)          not null
        primary key,
    name     text                 not null,
    phone    varchar(13)          not null,
    address  text                 not null,
    email    text                 not null,
    password text                 not null,
    salt     text                 not null,
    roleId   varchar(12)          not null,
    isActive tinyint(1) default 1 not null
);

create table Publisher
(
    id     varchar(12) not null,
    bookId varchar(12) not null,
    name   varchar(50) not null,
    primary key (id, bookId),
    constraint Publisher_ibfk_1
        foreign key (id) references BookInfo (publisherId)
);

create table Role
(
    id   varchar(13) not null
        primary key,
    name text        null
);

create table RoleFeature
(
    roleId    varchar(12) not null,
    featureId varchar(12) not null,
    primary key (roleId, featureId)
);

create table ShopGeneral
(
    id      varchar(12) not null
        primary key,
    name    varchar(12) not null,
    email   float       not null,
    phone   float       not null,
    address text        null
);

create table StockReport
(
    id    varchar(12) not null
        primary key,
    year  int         not null,
    month int         not null
);

create table StockReportDetail
(
    reportId varchar(12) not null,
    bookId   varchar(12) not null,
    initial  float       not null,
    sell     float       not null,
    import   float       not null,
    modify   float       not null,
    final    float       not null,
    primary key (reportId, bookId)
);

create table Supplier
(
    id    varchar(12)     not null
        primary key,
    name  text            not null,
    email text            not null,
    phone varchar(11)     not null,
    debt  float default 0 null,
    constraint phone
        unique (phone)
);

create table SupplierDebt
(
    id         varchar(12)              not null,
    supplierId varchar(12)              not null,
    qty        float                    not null,
    qtyLeft    float                    not null,
    type       enum ('Debt', 'Pay')     not null,
    createAt   datetime default (now()) null,
    createBy   varchar(9)               not null,
    primary key (id, supplierId)
);

create table SupplierDebtDetail
(
    reportId   varchar(12) not null,
    supplierId varchar(12) not null,
    initial    float       not null,
    arise      float       not null,
    final      float       not null,
    primary key (reportId, supplierId)
);

create table SupplierDebtReport
(
    id    varchar(12) not null
        primary key,
    year  int         not null,
    month int         not null
);

