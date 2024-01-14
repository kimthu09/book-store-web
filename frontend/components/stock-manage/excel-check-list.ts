import { CheckNote } from "@/types";
import { saveAs } from "file-saver";

export const ExportCheckNote = (excelData: CheckNote[], fileName: string) => {
  const ExcelJS = require("exceljs");

  const workbook = new ExcelJS.Workbook();
  const sheet = workbook.addWorksheet("Danh sách phiếu nhập hàng");

  // set title cell
  sheet.mergeCells("A1", "E1");
  sheet.getCell("A1").value = `Danh sách phiếu kiểm kho`;
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
    { key: "qtyDifferent", width: 32 },
    { key: "qtyAfterAdjust", width: 24 },
    { key: "createdBy", width: 24 },
    { key: "createdAt", width: 24 },
  ];

  // set column headers
  sheet.addRow(5).values = [
    "ID",
    "Số lượng thay đổi",
    "Tổng số lượng sau thay đổi",
    "Người tạo",
    "Ngày tạo",
  ];
  const values = excelData.map((note) => ({
    id: note.id,
    qtyDifferent: note.qtyDifferent,
    qtyAfterAdjust: note.qtyAfterAdjust,
    createdAt: new Date(note.createdAt).toLocaleDateString("vi-VN"),
    createdBy: note.createdBy.name,
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

      // Set cell border regardless of content
      sheet.getCell(cell.address).border = {
        top: { style: "thin" },
        left: { style: "thin" },
        bottom: { style: "thin" },
        right: { style: "thin" },
      };
    });
  });

  workbook.xlsx
    .writeBuffer()
    .then((buffer: any) => saveAs(new Blob([buffer]), fileName))
    .catch((err: any) => console.log("Error writing excel export", err));
};
