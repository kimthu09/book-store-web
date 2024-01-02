import { StatusNote, SupplierDebt } from "@/types";
import { saveAs } from "file-saver";

export const ExportDebtNote = (excelData: SupplierDebt[], fileName: string) => {
  const ExcelJS = require("exceljs");

  const workbook = new ExcelJS.Workbook();
  const sheet = workbook.addWorksheet("Danh sách phiếu nợ");

  // set title cell
  sheet.mergeCells("A1", "F1");
  sheet.getCell("A1").value = `Danh sách phiếu nợ`;
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
    { key: "id", width: 20 },
    { key: "qty", width: 32 },
    { key: "qtyLeft", width: 24 },
    { key: "type", width: 24 },
    { key: "createdAt", width: 24 },
    { key: "createdBy", width: 24 },
  ];

  // set column headers
  sheet.addRow(6).values = [
    "ID",
    "Giá trị",
    "Còn lại",
    "Loại phiếu",
    "Ngày tạo",
    "Người tạo",
  ];
  const values = excelData.map((debtNote) => ({
    id: debtNote.id,
    qty: debtNote.qty,
    qtyLeft: debtNote.qtyLeft,
    type: debtNote.type === "Pay" ? "Trả nợ" : "Nhập",
    createdAt: new Date(debtNote.createdAt).toLocaleDateString("vi-VN"),
    createdBy: debtNote.createdBy.name,
  }));
  //add data
  values.forEach((row) => {
    sheet.addRow(row);
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
