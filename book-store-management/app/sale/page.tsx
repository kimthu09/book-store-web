"use client";
import ProductTab from "@/components/sale/product-tab";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { FaChevronUp } from "react-icons/fa6";
import {
  SubmitErrorHandler,
  SubmitHandler,
  useFieldArray,
  useForm,
} from "react-hook-form";
import { Sheet, SheetContent, SheetTrigger } from "@/components/ui/sheet";
import BillTab, { Total } from "@/components/sale/bill-tab";
import createInvoice from "@/lib/invoice/createInvoice";
import { toast } from "@/components/ui/use-toast";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
export type FormValues = {
  details: {
    bookId: string;
    qty: number;
    sellPrice: number;
    name: string;
    stock: number;
  }[];
};
const SaleScreen = () => {
  const form = useForm<FormValues>({
    defaultValues: {
      details: [],
    },
  });
  const { register, control, setValue, watch, handleSubmit } = form;

  const { fields, append, remove, update } = useFieldArray({
    control: control,
    name: "details",
  });
  const onErrors: SubmitErrorHandler<FormValues> = (data) => {
    toast({
      variant: "destructive",
      title: "Có lỗi",
      description: "Vui lòng thử lại sau",
    });
  };
  const onSubmit: SubmitHandler<FormValues> = async (data) => {
    if (fields.length < 1) {
      toast({
        variant: "destructive",
        title: "Có lỗi",
        description: "Vui lòng chọn ít nhất một sản phẩm",
      });
      return;
    }
    console.log(data);
    const response: Promise<any> = createInvoice({
      details: data.details.map((item) => {
        return {
          bookId: item.bookId,
          qty: +item.qty,
        };
      }),
    });
    const responseData = await response;
    if (responseData.hasOwnProperty("errorKey")) {
      toast({
        variant: "destructive",
        title: "Có lỗi",
        description: responseData.message,
      });
    } else {
      toast({
        variant: "success",
        title: "Thành công",
        description: "Thêm mới hóa đơn thành công",
      });
    }
  };
  return (
    <div className="flex gap-4 md:pb-0 pb-16">
      <div className="2xl:basis-3/5 xl:basis-1/2 md:basis-2/5  flex-1 ">
        <ProductTab append={append} fields={fields} update={update} />
      </div>
      <div className="2xl:basis-2/5 xl:basis-1/2 md:basis-3/5  md:block hidden">
        <BillTab
          onPayClick={handleSubmit(onSubmit)}
          fields={fields}
          setValue={setValue}
          register={register}
          watch={watch}
          control={control}
          remove={remove}
        />
      </div>
      <div className="fixed bottom-0 left-0 right-0">
        <Card className="md:hidden flex flex-col  h-16 bg-white rounded-none overflow-hidden">
          <div className="flex flex-1 justify-between items-center align-middle px-4">
            <Dialog>
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
                  <DialogClose className="flex gap-5 justify-end">
                    <Button variant={"outline"}>Hủy</Button>
                    <Button onClick={handleSubmit(onSubmit)}>Xác nhận</Button>
                  </DialogClose>
                </DialogFooter>
              </DialogContent>
            </Dialog>
            <div className="ml-auto">
              <Total control={control} />
            </div>
          </div>

          <Sheet>
            <SheetTrigger asChild>
              <Button className="w-8 h-8 absolute p-0 rounded-full top-[-14px] left-[50%]">
                <FaChevronUp className="w-5 h-5" />
              </Button>
            </SheetTrigger>
            <SheetContent side={"bottom"} className="w-full p-0 bg-white pt-10">
              <BillTab
                onPayClick={handleSubmit(onSubmit)}
                fields={fields}
                setValue={setValue}
                register={register}
                watch={watch}
                control={control}
                remove={remove}
                isSheet
              />
            </SheetContent>
          </Sheet>
        </Card>
      </div>
    </div>
  );
};

export default SaleScreen;
