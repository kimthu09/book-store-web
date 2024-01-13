import { toLocalTime } from "@/lib/utils";
import { StockReport} from "@/types";
import { saveAs } from "file-saver";

export const ExportStockReport = (excelData: StockReport, fileName: string) => {
  const ExcelJS = require("exceljs");

  const workbook = new ExcelJS.Workbook();
  const sheet = workbook.addWorksheet("Stocks");

  // set title cell
  sheet.mergeCells("A1", "G1");
  sheet.getCell("A1").value = "Báo cáo tồn kho";
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

  sheet.getCell("A2").value = "Từ";
  sheet.mergeCells("B2", "G2");
  sheet.getCell("B2").value = toLocalTime(excelData.timeFrom);

  sheet.getCell("A3").value = "Đến";
  sheet.mergeCells("B3", "G3");
  sheet.getCell("B3").value = toLocalTime(excelData.timeTo);

  // set columns id
  sheet.columns = [
    { key: "id", width: 10 },
    { key: "name", width: 36 },
    { key: "initial", width: 20 },
    { key: "import", width: 20 },
    { key: "modify", width: 20 },
    { key: "sell", width: 20 },
    { key: "final", width: 20 },
  ];

  sheet.addRow(1)
  
  // set column headers
  sheet.addRow(5).values = [
    "ID",
    "Tên sách",
    "Tồn đầu",
    "Nhập",
    "Kiểm kho",
    "Bán",
    "Tồn cuối",
  ];

  // add data
  excelData.details.forEach((detail) => {
    let columnDetail={
      id: detail.book.id,
      name: detail.book.name,
      initial: detail.initial,
      import: detail.import,
      modify: detail.modify,
      sell: detail.sell,
      final: detail.final,
    }
    sheet.addRow(columnDetail);
  });

  //add footer
  let addressNow = 1 + 2 + excelData.details.length + 3
  sheet.mergeCells("A" + addressNow, "B"+ addressNow);
  sheet.getCell("A" + addressNow).value = "Tổng cộng";

  sheet.getCell("C" + addressNow).value = excelData.initial;
  sheet.getCell("D" + addressNow).value = excelData.import;
  sheet.getCell("E" + addressNow).value = excelData.modify;
  sheet.getCell("F" + addressNow).value = excelData.sell;
  sheet.getCell("G" + addressNow).value = excelData.final;

  // style header row
  sheet.getRow(5).eachCell({ includeEmpty: true }, function (cell: any) {
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
  });

  workbook.xlsx
    .writeBuffer()
    .then((buffer: any) => saveAs(new Blob([buffer]), fileName))
    .catch((err: any) => console.log("Error writing excel export", err));
};
