import { useState } from "react";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "../ui/dialog";
import { Button } from "../ui/button";
import { FaPen } from "react-icons/fa";
import { Input } from "../ui/input";
import { toast } from "../ui/use-toast";

export type ImageProps = {
  currentImage: string;
  handleImageSelected: () => void;
  image: any;
  setImage: (value: any) => void;
};

const ChangeImage = ({
  image,
  setImage,
  currentImage,
  handleImageSelected,
}: ImageProps) => {
  const [imagePreviews, setImagePreviews] = useState<any>();
  const handleMultipleImage = (event: any) => {
    const file = event.target.files[0];
    if (file) {
      if (file && file.type.includes("image")) {
        setImage(file);
        console.log(file.type);
        const reader = new FileReader();
        reader.onload = () => {
          setImagePreviews(reader.result);
        };
        reader.readAsDataURL(file);
      } else {
        setImage(null);
        toast({
          variant: "destructive",
          title: "Có lỗi",
          description: "File không hợp lệ",
        });
        console.log("file không hợp lệ");
      }
    } else {
      setImage(null);
    }
  };
  const [open, setOpen] = useState(false);
  const handleOpen = (value: boolean) => {
    if (value) {
      setImage(null);
      setImagePreviews(null);
    }
    setOpen(value);
  };
  return (
    <Dialog open={open} onOpenChange={handleOpen}>
      <DialogTrigger asChild>
        <Button
          type="button"
          size={"icon"}
          variant={"outline"}
          className="rounded-full absolute bottom-0 right-0 bg-white"
        >
          <FaPen />
        </Button>
      </DialogTrigger>
      <DialogContent className="max-w-[472px] p-0 bg-white">
        <DialogHeader>
          <DialogTitle className="p-6 pb-0">Chọn hình ảnh</DialogTitle>
        </DialogHeader>
        <form>
          <div className="p-6 flex flex-col items-center gap-4 border-y-[1px]">
            <div className="relative rounded-full border overflow-clip">
              {image ? (
                <img
                  src={imagePreviews}
                  alt={`Preview`}
                  className="h-[100px] w-[100px] object-cover"
                />
              ) : (
                <img
                  src={currentImage}
                  className="h-[100px] w-[100px] object-cover"
                />
              )}
            </div>

            <Input
              className="w-56"
              id="img"
              type="file"
              onChange={handleMultipleImage}
            ></Input>
          </div>
          <div className="p-4 flex-1 flex justify-end">
            <div className="flex gap-4">
              <Button
                type="button"
                variant={"outline"}
                onClick={() => setOpen(false)}
              >
                Huỷ
              </Button>

              <Button
                type="button"
                onClick={() => {
                  handleImageSelected();
                  setOpen(false);
                }}
              >
                Thay đổi
              </Button>
            </div>
          </div>
        </form>
      </DialogContent>
    </Dialog>
  );
};

export default ChangeImage;
