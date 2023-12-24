import { FormValues } from "@/app/sale/page";
import { useState } from "react";
import {
  Control,
  FieldArrayWithId,
  UseFieldArrayRemove,
  UseFormRegister,
  UseFormReset,
  UseFormSetValue,
  UseFormWatch,
  useWatch,
} from "react-hook-form";
import { Card, CardContent } from "../ui/card";
import { FiTrash2 } from "react-icons/fi";
import { HiPlus, HiMinus } from "react-icons/hi";
import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { toVND } from "@/lib/utils";
import { Label } from "../ui/label";
import { RadioGroup, RadioGroupItem } from "../ui/radio-group";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "../ui/dialog";
import { toast } from "../ui/use-toast";
import { AiOutlineClose } from "react-icons/ai";

const AddUp = ({
  control,
  index,
}: {
  control: Control<FormValues>;
  index: number;
}) => {
  const formValues = useWatch({
    name: `details.${index}`,
    control,
  });
  const addUp = formValues.sellPrice * formValues.qty;
  return <span className="text-sm font-bold">{toVND(addUp)}</span>;
};

export const Total = ({ control }: { control: Control<FormValues> }) => {
  const formValues = useWatch({
    name: "details",
    control,
  });
  const total = formValues.reduce(
    (acc, current) => acc + (current.sellPrice || 0) * (current.qty || 0),
    0
  );
  const totalQuantity = formValues.reduce(
    (acc, current) => acc + 1 * (current.qty || 0),
    0
  );
  return (
    <div className="flex gap-2 items-center">
      <span>Tổng tiền</span>
      <div className="border rounded-lg px-2 py-1">{totalQuantity}</div>
      <h1 className="text-sm">{toVND(total)}</h1>
    </div>
  );
};

const BillTab = ({
  fields,
  setValue,
  register,
  watch,
  control,
  remove,
  reset,
  onPayClick,
  isSheet,
}: {
  fields: FieldArrayWithId<FormValues, "details", "id">[];
  setValue: UseFormSetValue<FormValues>;
  register: UseFormRegister<FormValues>;
  watch: UseFormWatch<FormValues>;
  control: Control<FormValues, any>;
  remove: UseFieldArrayRemove;
  reset: UseFormReset<FormValues>;
  onPayClick: () => void;
  isSheet?: boolean;
}) => {
  const [open, setOpen] = useState(false);
  const invoices = watch("details");
  return (
    <Card className="sticky right-0 top-0 h-[86vh] overflow-hidden">
      <CardContent
        className={`flex flex-col p-0 overflow-hidden h-[86vh] ${
          isSheet ? "rounded-none" : ""
        }`}
      >
        <div className="flex items-center justify-between bg-white p-2 px-4 shadow-[0_2px_2px_-2px_rgba(0,0,0,0.2)]">
          <span className="font-medium text-primary">Hóa đơn</span>
          <Button
            variant={"ghost"}
            size={"icon"}
            className="rounded-full"
            onClick={() => reset()}
          >
            <AiOutlineClose />
          </Button>
        </div>
        <div className="flex flex-col gap-2  overflow-auto pt-4 flex-1">
          {fields.map((item, index) => {
            return (
              <div
                key={item.id}
                className={`flex ${
                  index === fields.length - 1 ? "" : "border-b"
                }  xl:px-4 px-2 pb-2 group gap-2`}
              >
                <span className="text-sm leading-6 self-center">
                  {index + 1}
                </span>
                <div className="flex flex-col flex-1">
                  {/* Name size price row */}
                  <div className="flex">
                    <div className="flex basis-[30%] self-center">
                      <h1 className="text-base font-medium">{item.name}</h1>
                    </div>

                    <div className="flex flex-wrap basis-[70%] items-center justify-between xl:gap-3 gap-2">
                      {/* Quantity */}
                      <div className="flex gap-2 items-center">
                        <Button
                          className="p-[2px] bg-primary hover:bg-primary/90 rounded-full cursor-pointer text-white invisible  group-hover:visible h-5 w-5"
                          onClick={() => {
                            const quantity = +invoices.at(index)?.qty!;
                            if (quantity === 1) {
                              remove(index);
                            } else {
                              setValue(`details.${index}.qty`, quantity - 1);
                            }
                          }}
                        >
                          <HiMinus />
                        </Button>
                        <Input
                          type="number"
                          className="px-1 w-10 text-center [&::-webkit-inner-spin-button]:appearance-none"
                          {...register(`details.${index}.qty` as const)}
                        ></Input>

                        <Button
                          className="p-[2px] bg-primary hover:bg-primary/90 rounded-full cursor-pointer text-white invisible group-hover:visible h-5 w-5"
                          onClick={() => {
                            if (
                              +invoices.at(index)?.qty! >=
                              +invoices.at(index)?.stock!
                            ) {
                              toast({
                                variant: "destructive",
                                title: "Có lỗi",
                                description:
                                  "Đã đạt số lượng tối đa của sản phẩm",
                              });
                            } else {
                              setValue(
                                `details.${index}.qty`,
                                +invoices.at(index)?.qty! + 1
                              );
                            }
                          }}
                        >
                          <HiPlus />
                        </Button>
                      </div>
                      <span className="text-sm text-right flex-1">
                        {toVND(item.sellPrice)}
                      </span>
                      <div className="text-right flex-1">
                        <AddUp control={control} index={index} />
                      </div>
                    </div>
                    <Button
                      variant={"ghost"}
                      className="h-8 p-0 ml-2 px-2 rounded-lg"
                      onClick={() => remove(index)}
                    >
                      <FiTrash2 className="opacity-50" />
                    </Button>
                  </div>
                </div>
              </div>
            );
          })}
        </div>
        <div className="flex flex-col gap-4 p-4 items-end shadow-[0_-2px_2px_-2px_rgba(0,0,0,0.2)] bg-white px-6">
          {/* Total */}
          <div className="ml-auto">
            <Total control={control} />
          </div>
          <RadioGroup defaultValue="all">
            <div className="flex items-center space-x-2">
              <RadioGroupItem value="all" id="r1" />
              <Label htmlFor="r1" className="font-normal">
                Thanh toán bằng tiền mặt
              </Label>
            </div>
          </RadioGroup>
          <Dialog open={open} onOpenChange={setOpen}>
            <DialogTrigger asChild>
              <Button>Thanh toán</Button>
            </DialogTrigger>
            <DialogContent>
              <DialogHeader>
                <DialogTitle>Xác nhận thanh toán</DialogTitle>
                <DialogDescription>
                  Bạn có chắc chắn muốn thanh toán
                </DialogDescription>
              </DialogHeader>
              <DialogFooter>
                <div className="flex gap-5 justify-end">
                  <Button variant={"outline"} onClick={() => setOpen(false)}>
                    Hủy
                  </Button>
                  <Button
                    onClick={() => {
                      onPayClick();
                      setOpen(false);
                    }}
                  >
                    Xác nhận
                  </Button>
                </div>
              </DialogFooter>
            </DialogContent>
          </Dialog>
        </div>
      </CardContent>
    </Card>
  );
};

export default BillTab;
