import { StatusActive, Supplier } from "@/types";
import { saveAs } from "file-saver";

export const ExportSupplierList = (excelData: Supplier[], fileName: string) => {
  const ExcelJS = require("exceljs");

  const workbook = new ExcelJS.Workbook();
  const sheet = workbook.addWorksheet("Suppliers");

  // set title cell
  sheet.mergeCells("A1", "F1");
  sheet.getCell("A1").value = "Supplier List";
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
    { key: "email", width: 40 },
    { key: "phone", width: 20 },
    { key: "debt", width: 20 },
    { key: "status", width: 24 },
  ];

  // set column headers
  sheet.addRow(6).values = [
    "ID",
    "Nhà cung cấp",
    "Email",
    "Số điện thoại",
    "Tổng nợ",
    "Trạng thái",
  ];

  // add data
  excelData.forEach((supplier) => {
    sheet.addRow(supplier);
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
