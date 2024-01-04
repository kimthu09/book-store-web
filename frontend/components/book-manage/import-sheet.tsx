"use client";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Sheet,
  SheetClose,
  SheetContent,
  SheetFooter,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";
import { useState } from "react";
import { FiUpload } from "react-icons/fi";
import { toast } from "../ui/use-toast";

const ImportSheet = ({
  handleFile,
  sampleFileLink,
}: {
  handleFile: (reader: FileReader) => void;
  sampleFileLink: string;
}) => {
  const [file, setFile] = useState<any>();
  const fileType = [
    "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
  ];
  const handleChange = (e: any) => {
    let selectedFile = e.target.files[0];
    if (selectedFile) {
      if (selectedFile && fileType.includes(selectedFile.type)) {
        if (selectedFile.size > 2000000) {
          console.log("Dung lượng file không hợp lệ");
          toast({
            variant: "destructive",
            title: "Có lỗi",
            description: "Dung lượng file không hợp lệ",
          });
        } else {
          setFile(null);
          setFile(e.target.files[0]);
        }
      } else {
        setFile(null);
        console.log(selectedFile.type);

        toast({
          variant: "destructive",
          title: "Có lỗi",
          description: "File không hợp lệ",
        });
      }
    }
  };
  const handleImport = () => {
    if (!file) {
      return;
    }
    const reader = new FileReader();

    reader.readAsArrayBuffer(file);
    handleFile(reader);
  };
  return (
    <Sheet>
      <SheetTrigger asChild>
        <Button variant={"outline"} className="bg-white px-3">
          <div className="flex flex-nowrap gap-1 items-center">
            <FiUpload className="w-4 h-4 text-green-700" />
            Nhập danh sách
          </div>
        </Button>
      </SheetTrigger>
      <SheetContent side={"top"} className="w-[480px] sm:w-[540px] m-auto">
        <SheetHeader>
          <SheetTitle>Nhập danh sách</SheetTitle>
        </SheetHeader>
        <div className="text-sm text-muted-foreground mt-4">
          <p>- Chuyển đổi file nhập dưới dạng .XLSX trước khi tải dữ liệu</p>
          <span>
            - Tải file mẫu
            <Button variant={"link"} className="px-1">
              <a href={sampleFileLink} download="Sample.xlsx">
                tại đây
              </a>
            </Button>
          </span>
          <p>- File nhập có dung lượng tối đa là 2MB.</p>
        </div>
        <div className="grid gap-4 py-4">
          <div className="grid grid-cols-4 items-center gap-4">
            <Input
              id="file"
              type="file"
              onChange={handleChange}
              className="col-span-3"
            />
          </div>
        </div>
        <SheetFooter>
          <SheetClose asChild>
            <div className="gap-3 flex">
              <Button variant={"outline"}>Hủy</Button>
              <Button type="button" onClick={handleImport}>
                Nhập file
              </Button>
            </div>
          </SheetClose>
        </SheetFooter>
      </SheetContent>
    </Sheet>
  );
};

export default ImportSheet;
