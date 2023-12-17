import { RadioGroup, RadioGroupItem } from "../ui/radio-group";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogTitle,
  DialogTrigger,
} from "../ui/dialog";
import { Button } from "../ui/button";
import { Label } from "../ui/label";
const DialogSupplierExport = ({
  handleExport,
  setExportOption,
}: {
  handleExport: () => void;
  setExportOption: (value: string) => void;
}) => {
  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button className="lg:px-3 px-2" variant={"outline"}>
          Xuất danh sách
        </Button>
      </DialogTrigger>
      <DialogContent className="p-0">
        <DialogTitle className="p-6 pb-0">
          Xuất file danh sách nhà cung cấp
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
                Tất cả các nhà cung cấp
              </Label>
            </div>
            <div className="flex items-center space-x-2">
              <RadioGroupItem value="comfortable" id="r2" />
              <Label className="font-normal" htmlFor="r2">
                Các nhà cung cấp được chọn
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

export default DialogSupplierExport;
