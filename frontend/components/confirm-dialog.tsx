import { useState } from "react";
import { Button } from "./ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "./ui/dialog";

type DialogProps = {
  title: string;
  description: string;
  handleYes: () => void;
  handleNo?: () => void;
  children: React.ReactNode;
};
const ConfirmDialog = ({
  title,
  description,
  handleYes,
  handleNo,
  children,
}: DialogProps) => {
  const [open, setOpen] = useState(false);
  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>{children}</DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle className="pb-2">{title}</DialogTitle>
          <DialogDescription>{description}</DialogDescription>
        </DialogHeader>
        <div className="flex flex-col-reverse sm:flex-row sm:justify-end sm:space-x-2">
          <div className="flex gap-5 sm:justify-end justify-stretch">
            <Button
              variant={"outline"}
              type="button"
              onClick={() => {
                if (handleNo) {
                  handleNo();
                }
                setOpen(false);
              }}
              className="sm:flex flex-1 w-auto"
            >
              Hủy
            </Button>
            <Button
              type="button"
              className="sm:flex flex-1 whitespace-nowrap"
              onClick={() => {
                handleYes();
                setOpen(false);
              }}
            >
              Xác nhận
            </Button>
          </div>
        </div>
      </DialogContent>
    </Dialog>
  );
};

export default ConfirmDialog;
