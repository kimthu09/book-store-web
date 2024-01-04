import { Dialog, DialogTrigger } from "@radix-ui/react-dialog";
import React from "react";
import { Button } from "../ui/button";
import { DialogClose, DialogContent, DialogTitle } from "../ui/dialog";
import { RadioGroup, RadioGroupItem } from "../ui/radio-group";
import { Label } from "../ui/label";

const ExportDialog = ({
  handleExport,
  setExportOption,
  isImport,
  style,
}: {
  handleExport: () => void;
  setExportOption: (value: string) => void;
  isImport: boolean;
  style?: string;
}) => {
  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button
          variant={"outline"}
          className={style ? style : "absolute top-6 right-6"}
        >
          Xuất file
        </Button>
      </DialogTrigger>
      <DialogContent className="p-0">
        <DialogTitle className="p-6 pb-0">
          Xuất danh sách {isImport ? "phiếu nhập" : "phiếu nợ"}
        </DialogTitle>
        <div className="flex flex-col border-y-[1px] p-6 gap-4">
          <Label>Giới hạn kết quả xuất</Label>
          <RadioGroup
            defaultValue="all"
            onValueChange={(e: string) => setExportOption(e)}
          >
            <div className="flex items-center space-x-2">
              <RadioGroupItem value="all" id="r1" />
              <Label htmlFor="r1" className="font-normal">
                Tất cả các phiếu
              </Label>
            </div>
            <div className="flex items-center space-x-2">
              <RadioGroupItem value="comfortable" id="r2" />
              <Label className="font-normal" htmlFor="r2">
                Các phiếu trong trang này
              </Label>
            </div>
          </RadioGroup>
        </div>

        <DialogClose className="ml-auto p-6 pt-0">
          <div className="flex gap-4">
            <Button type="button" variant={"outline"}>
              Thoát
            </Button>

            <Button type="button" onClick={() => handleExport()}>
              Hoàn tất
            </Button>
          </div>
        </DialogClose>
      </DialogContent>
    </Dialog>
  );
};

export default ExportDialog;
