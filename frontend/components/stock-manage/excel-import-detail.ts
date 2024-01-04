import {
  CheckNote,
  CheckNoteDetail,
  ImportNote,
  ImportNoteDetail,
} from "@/types";
import { saveAs } from "file-saver";

export const ExportImportNoteDetail = (
  excelData: ImportNote,
  importDetails: ImportNoteDetail[],
  fileName: string
) => {
  const ExcelJS = require("exceljs");

  const workbook = new ExcelJS.Workbook();
  const sheet = workbook.addWorksheet(excelData.id);

  // set title cell
  sheet.mergeCells("A1", "E1");
  sheet.getCell("A1").value = `Phiếu nhập`;
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

  // set supplier name cell
  sheet.mergeCells("A2", "B2");
  sheet.getCell("A2").value = `Nhà cung cấp: ${excelData.supplier.name}`;
  sheet.getCell("A2").alignment = {
    vertical: "middle",
  };
  sheet.getCell("A2").border = {
    top: { style: "thin" },
    left: { style: "thin" },
    bottom: { style: "thin" },
    right: { style: "thin" },
  };
  // set date row
  sheet.mergeCells("C2", "E2");
  sheet.getCell("C2").value = `Ngày tạo: ${new Date(
    excelData.createdAt
  ).toLocaleDateString("vi-VN")}`;
  sheet.getCell("C2").alignment = {
    vertical: "middle",
  };
  sheet.getCell("C2").border = {
    top: { style: "thin" },
    left: { style: "thin" },
    bottom: { style: "thin" },
    right: { style: "thin" },
  };

  // set columns id
  sheet.columns = [
    { key: "id", width: 20 },
    { key: "name", width: 32 },
    { key: "quantity", width: 24 },
    { key: "price", width: 24 },
    { key: "totalUnit", width: 24 },
  ];

  // set column headers
  sheet.addRow(5).values = [
    "Mã sách",
    "Tên sách",
    "Số lượng",
    "Đơn giá",
    "Thành tiền",
  ];
  const values = importDetails.map((note) => ({
    id: note.book.id,
    name: note.book.name,
    price: note.price,
    quantity: note.qtyImport,
    totalUnit: note.price * note.qtyImport,
  }));
  values.forEach((row) => {
    sheet.addRow(row);
  });

  // style header row
  sheet.getRow(3).eachCell({ includeEmpty: true }, function (cell: any) {
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

  const totalRow = 4 + importDetails.length;
  sheet.mergeCells(`A${totalRow}`, `E${totalRow}`);
  sheet.getCell(`A${totalRow}`).value = `Tổng tiền: ${excelData.totalPrice}`;
  sheet.getCell(`A${totalRow}`).alignment = {
    vertical: "middle",
    horizontal: "right",
  };
  sheet.getCell(`A${totalRow}`).font = {
    bold: true,
  };
  // sheet global font size
  sheet.eachRow((row: any) => {
    row.eachCell((cell: any) => {
      // default styles
      if (!cell.font?.size) {
        cell.font = Object.assign(cell.font || {}, { size: 13 });
      }
      sheet.getCell(cell.address).border = {
        top: { style: "thin" },
        left: { style: "thin" },
        bottom: { style: "thin" },
        right: { style: "thin" },
      };
    });
    // row.getCell(6).value = "quao quao quao";
  });

  workbook.xlsx
    .writeBuffer()
    .then((buffer: any) => saveAs(new Blob([buffer]), fileName))
    .catch((err: any) => console.log("Error writing excel export", err));
};

export const ExportCheckNoteDetail = (
  excelData: CheckNote,
  importDetails: CheckNoteDetail[],
  fileName: string
) => {
  const ExcelJS = require("exceljs");

  const workbook = new ExcelJS.Workbook();
  const sheet = workbook.addWorksheet(excelData.id);

  // set title cell
  sheet.mergeCells("A1", "E1");
  sheet.getCell("A1").value = `Phiếu nhập`;
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

  // set date row
  sheet.mergeCells("A2", "E2");
  sheet.getCell("A2").value = `Ngày tạo: ${new Date(
    excelData.createdAt
  ).toLocaleDateString("vi-VN")}`;
  sheet.getCell("A2").alignment = {
    vertical: "middle",
  };
  sheet.getCell("A2").border = {
    top: { style: "thin" },
    left: { style: "thin" },
    bottom: { style: "thin" },
    right: { style: "thin" },
  };

  // set columns id
  sheet.columns = [
    { key: "id", width: 20 },
    { key: "name", width: 32 },
    { key: "initial", width: 24 },
    { key: "difference", width: 24 },
    { key: "final", width: 24 },
  ];

  // set column headers
  sheet.addRow(5).values = [
    "Mã sách",
    "Tên sách",
    "Ban đầu",
    "Chênh lệch",
    "Kiểm kê",
  ];
  const values = importDetails.map((note) => ({
    id: note.book.id,
    name: note.book.name,
    initial: note.initial,
    difference: note.difference,
    final: note.final,
  }));
  values.forEach((row) => {
    sheet.addRow(row);
  });

  // style header row
  sheet.getRow(3).eachCell({ includeEmpty: true }, function (cell: any) {
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
      sheet.getCell(cell.address).border = {
        top: { style: "thin" },
        left: { style: "thin" },
        bottom: { style: "thin" },
        right: { style: "thin" },
      };
    });
    // row.getCell(6).value = "quao quao quao";
  });

  workbook.xlsx
    .writeBuffer()
    .then((buffer: any) => saveAs(new Blob([buffer]), fileName))
    .catch((err: any) => console.log("Error writing excel export", err));
};
