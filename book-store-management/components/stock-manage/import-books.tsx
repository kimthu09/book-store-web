import { FormValues } from "@/app/stock/import/new/page";
import { Label } from "@radix-ui/react-dropdown-menu";
import {
  Control,
  UseFormReturn,
  useFieldArray,
  useWatch,
} from "react-hook-form";
import {
  CommandDialog,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from "../ui/command";
import { Button } from "../ui/button";
import { books } from "@/constants";
import { AiOutlineClose } from "react-icons/ai";
import { Input } from "../ui/input";
import { useState } from "react";
import { Checkbox } from "../ui/checkbox";
import { DialogClose } from "@radix-ui/react-dialog";

const Total = ({ control }: { control: Control<FormValues> }) => {
  const formValues = useWatch({
    name: "books",
    control,
  });
  const total = formValues.reduce(
    (acc, current) => acc + (current.price || 0) * (current.quantity || 0),
    0
  );
  const formatted = new Intl.NumberFormat("vi-VN", {
    style: "currency",
    currency: "VND",
  }).format(total);
  return <p>{formatted}</p>;
};

const AddUp = ({
  control,
  index,
}: {
  control: Control<FormValues>;
  index: number;
}) => {
  const formValues = useWatch({
    name: `books.${index}`,
    control,
  });
  const addUp = formValues.price * formValues.quantity;
  console.log(addUp);
  return <p>{addUp}</p>;
};

const ImportBooks = ({
  form,
}: {
  form: UseFormReturn<FormValues, any, undefined>;
}) => {
  const { register, handleSubmit, control, watch, getValues } = form;
  const {
    fields: fieldsBooks,
    append: appendBooks,
    remove: removeBooks,
  } = useFieldArray({
    control: control,
    name: "books",
  });
  const [openIngre, setOpenIngre] = useState(false);
  const [checkedIngre, setCheckedIngre] = useState(
    new Array(books.length).fill(false)
  );
  const handleOnChecked = (position: number) => {
    const updateCheckedState = checkedIngre.map((item, index) =>
      index === position ? !item : item
    );

    setCheckedIngre(updateCheckedState);
  };

  const resetCheckedIngre = () => {
    setCheckedIngre(new Array(books.length).fill(false));
  };

  const handleIngreConfirm = () => {
    setOpenIngre(false);
    checkedIngre.forEach((element, index) => {
      const book = books.at(index);
      if (element === true) {
        if (!fieldsBooks.find((item) => item.idBook === book?.id)) {
          appendBooks({
            idBook: book?.id!,
            quantity: 1,
            price: book?.price!,
          });
        }
      }
    });
  };
  return (
    <div className="flex flex-col">
      <div className="flex sm:gap-4 gap-2">
        <Input
          className="mb-4 flex-1"
          placeholder="Tìm tên sách"
          onClick={() => {
            setOpenIngre((open) => !open);
            resetCheckedIngre();
          }}
        />
      </div>

      <div className="flex pr-12 font-medium py-2 mb-2 bg-orange-100 gap-4">
        <h2 className="w-12 justify-center sm:flex hidden">STT</h2>

        <h2 className="flex-1 sm:ml-0 ml-1">Tên sách</h2>

        <h2 className="flex-1 text-center">Đơn giá</h2>
        <h2 className="flex-1 text-center">Số lượng</h2>
        <h2 className="flex-1 text-right">Thành tiền</h2>
      </div>
      <div>
        {fieldsBooks.length < 1 ? (
          <div className="text-center py-4">Chọn sản phẩm nhập kho</div>
        ) : null}
        {fieldsBooks.map((book, index) => {
          const value = books.find((item) => item.id === book.idBook);

          return (
            <div
              key={book.id}
              className="flex items-center py-2 sm:gap-4 gap-2"
            >
              <h2 className="justify-center w-12 sm:flex hidden">
                {index + 1}
              </h2>

              <div className="flex-1 flex flex-col">
                <h2 className="flex-1">{value?.name}</h2>
              </div>
              <Input
                className="flex-1"
                type="number"
                min={1}
                max={1000}
                placeholder="Nhập đơn giá"
                defaultValue={book.quantity}
                {...register(`books.${index}.price` as const)}
              ></Input>
              <Input
                className="flex-1 "
                type="number"
                min={1}
                max={1000}
                placeholder="Nhập số lượng"
                defaultValue={book.quantity}
                {...register(`books.${index}.quantity` as const)}
              ></Input>

              <div className="flex-1 text-right">
                <AddUp control={control} index={index} />
              </div>
              <Button
                type="button"
                variant={"ghost"}
                className={`sm:px-3 px-2`}
                onClick={() => {
                  removeBooks(index);
                }}
              >
                <AiOutlineClose />
              </Button>
            </div>
          );
        })}
      </div>
      <div className="flex justify-end py-2 pr-12 font-medium ">
        <h2 className="w-1/4">Tổng cộng</h2>
        <div className="flex">
          <Total control={control} />
        </div>
      </div>
      <CommandDialog open={openIngre} onOpenChange={setOpenIngre}>
        <CommandInput placeholder="Tìm tên sách" />
        <CommandList className="h-96">
          <CommandEmpty>
            <div>
              <p>Không tìm thấy sách khớp với từ khoá.</p>
              <p>Bạn có thể thêm mới sách ở trang quản lý sách</p>
            </div>
          </CommandEmpty>
          <CommandGroup heading="Sách" className="overflow-y-auto pb-20 ">
            {books.map((item, index) => (
              <CommandItem
                value={item.name}
                key={item.id}
                onSelect={() => {
                  handleOnChecked(index);
                }}
              >
                <div className="px-1 blur-none flex items-center gap-2 flex-1">
                  <Checkbox
                    id={item.name}
                    checked={checkedIngre[index]}
                  ></Checkbox>
                  <div className="flex-1">
                    <Label>{item.name}</Label>
                  </div>

                  <div className="ml-auto flex">
                    <p className="text-muted-foreground">
                      Tồn:{" "}
                      <span className="text-blue-600">{item.quantity}</span> |
                      Giá nhập: <span className="text-black">{item.price}</span>
                    </p>
                  </div>
                </div>
              </CommandItem>
            ))}
          </CommandGroup>
        </CommandList>
        <DialogClose>
          <div className="p-4 flex w-full justify-end fixed bottom-0 left-0 bg-white border-t">
            <Button onClick={handleIngreConfirm}>Thêm</Button>
          </div>
        </DialogClose>
      </CommandDialog>
    </div>
  );
};

export default ImportBooks;
