"use client";
import { Book, ImportDetail, ImportNote, StatusActive } from "@/types";
import { saveAs } from "file-saver";

const fileType =
  "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet;charset=UTF-8";
export const ExportBookList = (excelData: Book[], fileName: string) => {
  const ExcelJS = require("exceljs");

  const workbook = new ExcelJS.Workbook();
  const sheet = workbook.addWorksheet("BookList");

  // set title cell
  sheet.mergeCells("A1", "F1");
  sheet.getCell("A1").value = "Book List";
  sheet.getCell("A1").alignment = {
    horizontal: "center",
    vertical: "middle",
  };
  sheet.getCell("A1").border = {
    top: { style: "thin" },
    left: { style: "thin" },
    bottom: { style: "thin" },
    right: { style: "thin" },
  };
  sheet.getCell("A1").font = {
    bold: true,
    size: 18,
  };
  sheet.getRow(1).height = 40;

  // set columns id
  sheet.columns = [
    { key: "id", width: 10 },
    { key: "name", width: 36 },
    { key: "nxb", width: 32 },
    { key: "quantity", width: 20 },
    { key: "price", width: 20 },
    { key: "status", width: 24 },
  ];

  // set column headers
  sheet.addRow(6).values = [
    "ID",
    "Tên sách",
    "Nhà xuất bản",
    "Số lượng",
    "Đơn giá",
    "Trạng thái",
  ];

  // add data
  excelData.forEach((book) => {
    sheet.addRow(book);
  });

  // set status col values
  sheet.getColumn(6).eachCell((cell: any, rowNumber: number) => {
    if (rowNumber > 2) {
      if (sheet.getCell(cell.address).value === true)
        sheet.getCell(cell.address).value = StatusActive.Active;
      else {
        sheet.getCell(cell.address).value = StatusActive.InActive;
      }
    }
  });

  // style header row
  sheet.getRow(2).eachCell({ includeEmpty: true }, function (cell: any) {
    sheet.getCell(cell.address).fill = {
      type: "pattern",
      pattern: "solid",
      fgColor: { argb: "cbdff2" },
      bgColor: { argb: "cbdff2" },
    };
    sheet.getCell(cell.address).border = {
      top: { style: "thin" },
      left: { style: "thin" },
      bottom: { style: "thin" },
      right: { style: "thin" },
    };
  });

  // sheet global font size
  sheet.eachRow((row: any) => {
    row.eachCell((cell: any) => {
      // default styles
      if (!cell.font?.size) {
        cell.font = Object.assign(cell.font || {}, { size: 13 });
      }
    });

    // row.getCell(6).value = "quao quao quao";
  });

  workbook.xlsx
    .writeBuffer()
    .then((buffer: any) => saveAs(new Blob([buffer]), fileName))
    .catch((err: any) => console.log("Error writing excel export", err));
};

export const ExportImportDetail = (
  note: ImportNote,
  details: ImportDetail[],
  fileName: string
) => {
  const ExcelJS = require("exceljs");

  const workbook = new ExcelJS.Workbook();
  const sheet = workbook.addWorksheet("Import Note");

  // set title cell
  sheet.mergeCells("A1", "E1");
  sheet.getCell("A1").value = `Import Note Id: ${note.id}`;
  sheet.getCell("A1").alignment = {
    horizontal: "center",
    vertical: "middle",
  };
  sheet.getCell("A1").border = {
    top: { style: "thin" },
    left: { style: "thin" },
    bottom: { style: "thin" },
    right: { style: "thin" },
  };
  sheet.getCell("A1").font = {
    bold: true,
    size: 18,
  };
  sheet.getRow(1).height = 40;

  // add general data
  sheet.addRow(5).values = [`Nhà cung cấp: ${note.supplierId}`];

  sheet.mergeCells("A3", "E3");
  sheet.getCell("A3").value = `Người lập phiếu ${
    note.createBy
  }, ngày ${note.createAt.toLocaleDateString("vi-VN")}`;

  // set columns id
  sheet.columns = [
    { key: "id", width: 10 },
    { key: "name", width: 36 },
    { key: "price", width: 32 },
    { key: "quantity", width: 20 },
    { key: "total", width: 20 },
  ];

  // set column headers
  sheet.addRow(5).values = [
    "Mã sách",
    "Tên sách",
    "Đơn giá",
    "Số lượng",
    "Thành tiền",
  ];

  const values = details.map((detail) => ({
    id: detail.book.id,
    name: detail.book.name,
    price: detail.book.price,
    quantity: detail.quantity,
    total: detail.book.price * detail.quantity,
  }));
  //add data
  values.forEach((row) => {
    sheet.addRow(row);
  });

  // style header row
  sheet.getRow(4).eachCell({ includeEmpty: true }, function (cell: any) {
    sheet.getCell(cell.address).fill = {
      type: "pattern",
      pattern: "solid",
      fgColor: { argb: "cbdff2" },
      bgColor: { argb: "cbdff2" },
    };
    sheet.getCell(cell.address).border = {
      top: { style: "thin" },
      left: { style: "thin" },
      bottom: { style: "thin" },
      right: { style: "thin" },
    };
  });

  // sheet global font size
  sheet.eachRow((row: any) => {
    row.eachCell((cell: any) => {
      // default styles
      if (!cell.font?.size) {
        cell.font = Object.assign(cell.font || {}, { size: 13 });
      }
    });

    // row.getCell(6).value = "quao quao quao";
  });

  workbook.xlsx
    .writeBuffer()
    .then((buffer: any) => saveAs(new Blob([buffer]), fileName))
    .catch((err: any) => console.log("Error writing excel export", err));
};
