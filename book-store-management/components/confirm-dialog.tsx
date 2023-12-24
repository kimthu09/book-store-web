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
          <DialogTitle>{title}</DialogTitle>
          <DialogDescription>{description}</DialogDescription>
        </DialogHeader>
        <DialogFooter>
          <div className="flex gap-5 sm:justify-end justify-stretch">
            <Button
              variant={"outline"}
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
              className="sm:flex flex-1 whitespace-nowrap"
              onClick={() => {
                handleYes();
                setOpen(false);
              }}
            >
              Xác nhận
            </Button>
          </div>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
};

export default ConfirmDialog;
