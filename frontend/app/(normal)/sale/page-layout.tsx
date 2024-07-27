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
import { useRouter } from "next/navigation";
import { useState } from "react";
import { ToastAction } from "@/components/ui/toast";
import PrintInvoice from "@/components/invoice/print-invoice";
import { endPoint } from "@/constants";
import { useSWRConfig } from "swr";
import { useLoading } from "@/hooks/loading-context";
export type FormValues = {
  customer: {
    customerId: string;
    customerPoint: number;
  };
  isUsePoint: boolean;
  details: {
    bookId: string;
    qty: number;
    sellPrice: number;
    name: string;
    stock: number;
  }[];
};
const SaleScreen = () => {
  const { mutate } = useSWRConfig();
  const form = useForm<FormValues>({
    defaultValues: {
      customer: {},
      isUsePoint: false,
      details: [],
    },
  });
  const {
    register,
    control,
    setValue,
    watch,
    handleSubmit,
    reset,
    formState: { isDirty },
  } = form;
  const router = useRouter();
  const { fields, append, remove, update } = useFieldArray({
    control: control,
    name: "details",
  });
  const { showLoading, hideLoading } = useLoading();
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
    const response: Promise<any> = createInvoice({
      customerId: data.customer.customerId,
      isUsePoint: data.isUsePoint,
      details: data.details.map((item) => {
        return {
          bookId: item.bookId,
          qty: +item.qty,
        };
      }),
    });
    showLoading();
    const responseData = await response;
    hideLoading();

    if (responseData.hasOwnProperty("errorKey")) {
      toast({
        variant: "destructive",
        title: "Có lỗi",
        description: responseData.message,
      });
    } else {
      const id = responseData.data.id;
      toast({
        variant: "success",
        title: "Thành công",
        description: "Thêm mới hóa đơn thành công",
        action: (
          <ToastAction altText="print" className="p-0">
            <PrintInvoice responseData={responseData.data} onPrint={() => {}} />
          </ToastAction>
        ),
      });
      reset({
        customer: {},
        isUsePoint: false,
        details: [],
      });
      mutate(`${endPoint}/v1/customers/all`);
      mutate(`${endPoint}/v1/books/all`);
      router.refresh();
    }
  };
  const [open, setOpen] = useState(false);
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
          reset={reset}
          isDirty={isDirty}
        />
      </div>
      <div className="fixed bottom-0 left-0 right-0">
        <Card className="md:hidden flex flex-col  h-16 bg-white rounded-none overflow-hidden">
          <div className="flex flex-1 justify-between items-center align-middle px-4">
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
                        handleSubmit(onSubmit)();
                        setOpen(false);
                      }}
                    >
                      Xác nhận
                    </Button>
                  </div>
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
                reset={reset}
                isDirty={isDirty}
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
