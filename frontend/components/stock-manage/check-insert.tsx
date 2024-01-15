import { useState } from "react";
import { Input } from "../ui/input";

import { Button } from "../ui/button";
import {
  Control,
  UseFormReturn,
  useFieldArray,
  useWatch,
} from "react-hook-form";

import { AiOutlineClose } from "react-icons/ai";
import { CiBoxes } from "react-icons/ci";
import { z } from "zod";
import { BookProps } from "@/types";
import { AutoComplete } from "../ui/autocomplete";
import getAllBookForSale from "@/lib/book/getAllBookForSale";
import { FormSchema } from "@/app/stockmanage/check/add/page";
import DropdownSkeleton from "../skeleton/dropdown-skeleton";

const Final = ({
  control,
  index,
}: {
  control: Control<z.infer<typeof FormSchema>>;
  index: number;
}) => {
  const formValues = useWatch({
    name: `details.${index}`,
    control,
  });
  const addUp = formValues.initial + +formValues.difference;
  return <p>{addUp}</p>;
};

const CheckInsert = ({
  form,
}: {
  form: UseFormReturn<z.infer<typeof FormSchema>, any, undefined>;
}) => {
  const {
    register,
    handleSubmit,
    control,
    watch,
    getValues,
    reset,
    formState: { errors },
  } = form;
  const {
    fields: fieldsBook,
    append: appendBook,
    remove: removeBook,
    replace,
  } = useFieldArray({
    control: control,
    name: "details",
  });
  const { books: data, isLoading, isError, mutate } = getAllBookForSale();
  const [value, setValue] = useState<BookProps>();
  const handleOnValueChange = (item: BookProps) => {
    if (!fieldsBook.find((book) => book.bookId === item.id)) {
      appendBook({
        bookId: item.id,
        initial: item.quantity,
        difference: 0,
      });
    }
  };
  if (isError) {
    return "Failed to fetch";
  } else if (isLoading || !data) {
    return <DropdownSkeleton />;
  } else {
    return (
      <div className="flex flex-col">
        <AutoComplete
          options={data.data}
          emptyMessage="Không có sách khớp với từ khóa"
          placeholder="Tìm sách"
          onValueChange={handleOnValueChange}
          value={value}
        />
        <div>
          <div className="grid grid-cols-4 lg:gap-3 gap-2 font-medium py-2 px-2 mt-2 rounded-t-md bg-[#a4c5ff]">
            <h2 className="col-span-1">Tên sách</h2>
            <h2 className=" text-right col-span-1">Ban đầu</h2>
            <h2 className=" text-right col-span-1"> Chênh lệch</h2>
            <h2 className=" text-right col-span-1 pr-12 ">Kiểm kê</h2>
          </div>
          <div className="border border-t-0 py-2 rounded-b-md">
            {fieldsBook.length < 1 ? (
              <div className="flex flex-col items-center gap-4 py-8 text-muted-foreground font-medium">
                <CiBoxes className="h-24 w-24 text-muted-foreground/40" />
                {errors.details?.root ? (
                  <span className="error___message">
                    {errors.details.root?.message}
                  </span>
                ) : (
                  "Chọn sản phẩm kiểm kho"
                )}
              </div>
            ) : null}
            {fieldsBook.map((book, index) => {
              const value = data.data.find((item) => item.id === book.bookId);

              if (value) {
                return (
                  <div
                    key={book.id}
                    className="grid grid-cols-4  p-2 lg:gap-3 gap-2 items-start"
                  >
                    <div className="flex flex-col col-span-1 ">
                      <h2 className="font-medium">{value?.name}</h2>
                      <span className="text-sm text-light">({value.id})</span>
                    </div>
                    <div className="relative p-1 col-span-1 text-right">
                      <p>{book.initial}</p>
                    </div>

                    <div className="relative p-1 col-span-1 flex-col flex items-end">
                      <Input
                        className="text-right p-2 max-w-[8rem]"
                        defaultValue={book.difference}
                        {...register(`details.${index}.difference` as const)}
                      ></Input>
                      {errors &&
                      errors.details &&
                      errors.details[index] &&
                      (errors.details[index]!.difference as
                        | { message: string }
                        | undefined) ? (
                        <span className="error___message">
                          {errors.details[index]!.difference!.message}
                        </span>
                      ) : null}
                    </div>
                    <div className="text-right flex justify-end gap-2 items-center col-span-1">
                      <div className="text-right">
                        <Final control={control} index={index} />
                      </div>
                      <Button
                        type="button"
                        variant={"ghost"}
                        className={`px-3`}
                        onClick={() => {
                          removeBook(index);
                        }}
                      >
                        <AiOutlineClose />
                      </Button>
                    </div>
                  </div>
                );
              }
            })}
          </div>
        </div>
      </div>
    );
  }
};

export default CheckInsert;
