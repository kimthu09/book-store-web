import { ImportNote, StatusNote } from "@/types";
import { saveAs } from "file-saver";

export const ExportImportNote = (excelData: ImportNote[], fileName: string) => {
  const ExcelJS = require("exceljs");

  const workbook = new ExcelJS.Workbook();
  const sheet = workbook.addWorksheet("Danh sách phiếu nhập hàng");

  // set title cell
  sheet.mergeCells("A1", "G1");
  sheet.getCell("A1").value = `Danh sách phiếu nhập hàng`;
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
    { key: "totalPrice", width: 32 },
    { key: "status", width: 24 },
    { key: "createdBy", width: 24 },
    { key: "createdAt", width: 24 },
    { key: "closedBy", width: 24 },
    { key: "closedAt", width: 24 },
  ];

  // set column headers
  sheet.addRow(7).values = [
    "ID",
    "Tổng tiền",
    "Trạng thái",
    "Người tạo",
    "Ngày tạo",
    "Người đóng",
    "Ngày đóng",
  ];
  const values = excelData.map((importNote) => ({
    id: importNote.id,
    totalPrice: importNote.totalPrice,
    status:
      importNote.status === StatusNote.Done
        ? "Hoàn thành"
        : importNote.status === StatusNote.Inprogress
        ? "Đang xử lý"
        : "Đã hủy",
    createdAt: new Date(importNote.createdAt).toLocaleDateString("vi-VN"),
    createdBy: importNote.createdBy.name,
    closedAt: importNote.closedAt
      ? new Date(importNote.closedAt).toLocaleDateString("vi-VN")
      : "",
    closedBy: importNote.closedBy?.name,
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
