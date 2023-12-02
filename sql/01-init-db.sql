use bookstoremanagement;

create table Author
(
    id        varchar(12)                          not null
        primary key,
    name      text                                 not null,
    createdAt datetime   default CURRENT_TIMESTAMP null,
    updatedAt datetime   default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    isActive  tinyint(1) default 1                 null
);

create table Book
(
    id          varchar(12)                          not null
        primary key,
    name        varchar(100)                         not null,
    `desc`      text                                 null,
    edition     int                                  not null,
    qty         int        default 0                 null,
    price       float                                not null,
    salePrice   float                                not null,
    publisherId varchar(12)                          null,
    authorIds   text                                 not null,
    categoryIds text                                 not null,
    isActive    tinyint(1) default 1                 not null,
    createdAt   datetime   default CURRENT_TIMESTAMP null,
    updatedAt   datetime   default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    constraint Book_pk2
        unique (id)
);

create table BookChangeHistory
(
    id         varchar(12)                          not null,
    bookId     varchar(12)                          not null,
    amount     float                                not null,
    amountLeft float                                not null,
    type       enum ('Sell', 'Import', 'Modify')    not null,
    createdAt  datetime   default CURRENT_TIMESTAMP null,
    updatedAt  datetime   default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    isActive   tinyint(1) default 1                 null,
    primary key (id, bookId)
);

create table Category
(
    id        varchar(12)                          not null
        primary key,
    name      varchar(50)                          not null,
    createdAt datetime   default CURRENT_TIMESTAMP null,
    updatedAt datetime   default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    isActive  tinyint(1) default 1                 null
);

create table Feature
(
    id          varchar(12)                          not null
        primary key,
    description text                                 null,
    createdAt   datetime   default CURRENT_TIMESTAMP null,
    updatedAt   datetime   default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    isActive    tinyint(1) default 1                 null
);

create table ImportNote
(
    id         varchar(12)                                                     not null
        primary key,
    supplierId varchar(12)                                                     not null,
    totalPrice float                                 default 0                 null,
    status     enum ('InProgress', 'Done', 'Cancel') default 'InProgress'      null,
    createBy   varchar(12)                                                     not null,
    closeBy    varchar(12)                                                     null,
    createAt   datetime                              default (now())           null,
    closeAt    datetime                                                        null,
    createdAt  datetime                              default CURRENT_TIMESTAMP null,
    updatedAt  datetime                              default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    isActive   tinyint(1)                            default 1                 null
);

create table ImportNoteDetail
(
    importNoteId varchar(12)                          not null,
    bookId       varchar(12)                          not null,
    price        float                                not null,
    qtyImport    float      default 0                 null,
    createdAt    datetime   default CURRENT_TIMESTAMP null,
    updatedAt    datetime   default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    isActive     tinyint(1) default 1                 null,
    primary key (importNoteId, bookId)
);

create table InventoryCheckNote
(
    id             varchar(12)                          not null
        primary key,
    qtyDifferent   float                                not null,
    qtyAfterAdjust float                                not null,
    createBy       varchar(12)                          not null,
    createAt       datetime   default (now())           null,
    createdAt      datetime   default CURRENT_TIMESTAMP null,
    updatedAt      datetime   default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    isActive       tinyint(1) default 1                 null
);

create table InventoryCheckNoteDetail
(
    inventoryCheckNoteId varchar(12)                          not null,
    bookId               varchar(12)                          not null,
    initial              float                                not null,
    difference           float                                not null,
    final                float                                not null,
    createdAt            datetime   default CURRENT_TIMESTAMP null,
    updatedAt            datetime   default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    isActive             tinyint(1) default 1                 null,
    primary key (inventoryCheckNoteId, bookId)
);

create table Invoice
(
    id          varchar(13)                          not null
        primary key,
    totalPrice  float                                not null,
    qtyReceived float                                not null,
    createAt    datetime   default (now())           null,
    createBy    varchar(13)                          not null,
    createdAt   datetime   default CURRENT_TIMESTAMP null,
    updatedAt   datetime   default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    isActive    tinyint(1) default 1                 null
);

create table InvoiceDetail
(
    invoiceId varchar(13)                          not null,
    bookId    varchar(13)                          not null,
    qty       float                                not null,
    unitPrice float                                not null,
    createdAt datetime   default CURRENT_TIMESTAMP null,
    updatedAt datetime   default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    isActive  tinyint(1) default 1                 null,
    primary key (invoiceId, bookId)
);

create table MUser
(
    id        varchar(12)                          not null
        primary key,
    name      text                                 not null,
    phone     varchar(13)                          not null,
    address   text                                 not null,
    email     text                                 not null,
    password  text                                 not null,
    salt      text                                 not null,
    roleId    varchar(12)                          not null,
    isActive  tinyint(1) default 1                 not null,
    createdAt datetime   default CURRENT_TIMESTAMP null,
    updatedAt datetime   default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);

create table Publisher
(
    id        varchar(12)                          not null
        primary key,
    name      varchar(50)                          not null,
    createdAt datetime   default CURRENT_TIMESTAMP null,
    updatedAt datetime   default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    isActive  tinyint(1) default 1                 null
);

create table Role
(
    id        varchar(13)                          not null
        primary key,
    name      text                                 null,
    createdAt datetime   default CURRENT_TIMESTAMP null,
    updatedAt datetime   default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    isActive  tinyint(1) default 1                 null
);

create table RoleFeature
(
    roleId    varchar(12)                          not null,
    featureId varchar(30)                          not null,
    createdAt datetime   default CURRENT_TIMESTAMP null,
    updatedAt datetime   default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    isActive  tinyint(1) default 1                 null,
    primary key (roleId, featureId)
);

create table ShopGeneral
(
    id        varchar(12)                          not null
        primary key,
    name      varchar(12)                          not null,
    email     float                                not null,
    phone     float                                not null,
    address   text                                 null,
    createdAt datetime   default CURRENT_TIMESTAMP null,
    updatedAt datetime   default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    isActive  tinyint(1) default 1                 null
);

create table StockReport
(
    id        varchar(12)                          not null
        primary key,
    year      int                                  not null,
    month     int                                  not null,
    createdAt datetime   default CURRENT_TIMESTAMP null,
    updatedAt datetime   default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    isActive  tinyint(1) default 1                 null
);

create table StockReportDetail
(
    reportId  varchar(12)                          not null,
    bookId    varchar(12)                          not null,
    initial   float                                not null,
    sell      float                                not null,
    import    float                                not null,
    modify    float                                not null,
    final     float                                not null,
    createdAt datetime   default CURRENT_TIMESTAMP null,
    updatedAt datetime   default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    isActive  tinyint(1) default 1                 null,
    primary key (reportId, bookId)
);

create table Supplier
(
    id        varchar(12)                          not null
        primary key,
    name      text                                 not null,
    email     text                                 not null,
    phone     varchar(11)                          not null,
    debt      float      default 0                 null,
    createdAt datetime   default CURRENT_TIMESTAMP null,
    updatedAt datetime   default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    isActive  tinyint(1) default 1                 null,
    constraint phone
        unique (phone)
);

create table SupplierDebt
(
    id         varchar(12)                          not null,
    supplierId varchar(12)                          not null,
    qty        float                                not null,
    qtyLeft    float                                not null,
    type       enum ('Debt', 'Pay')                 not null,
    createAt   datetime   default (now())           null,
    createBy   varchar(9)                           not null,
    createdAt  datetime   default CURRENT_TIMESTAMP null,
    updatedAt  datetime   default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    isActive   tinyint(1) default 1                 null,
    primary key (id, supplierId)
);

create table SupplierDebtDetail
(
    reportId   varchar(12)                          not null,
    supplierId varchar(12)                          not null,
    initial    float                                not null,
    arise      float                                not null,
    final      float                                not null,
    createdAt  datetime   default CURRENT_TIMESTAMP null,
    updatedAt  datetime   default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    isActive   tinyint(1) default 1                 null,
    primary key (reportId, supplierId)
);

create table SupplierDebtReport
(
    id        varchar(12)                          not null
        primary key,
    year      int                                  not null,
    month     int                                  not null,
    createdAt datetime   default CURRENT_TIMESTAMP null,
    updatedAt datetime   default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    isActive  tinyint(1) default 1                 null
);

