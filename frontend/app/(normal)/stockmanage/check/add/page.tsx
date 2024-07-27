"use client";

import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { SubmitErrorHandler, SubmitHandler, useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { endPoint, required } from "@/constants";
import { Button } from "@/components/ui/button";
import { LuCheck } from "react-icons/lu";
import { FiTrash2 } from "react-icons/fi";
import { useState } from "react";
import { toast } from "@/components/ui/use-toast";
import { useSWRConfig } from "swr";
import CheckInsert from "@/components/stock-manage/check-insert";
import createCheckNote from "@/lib/check/createCheckNote";
import { useCurrentUser } from "@/hooks/use-user";
import { includesRoles } from "@/lib/utils";
import NoRole from "@/components/no-role";
import ImportSheet from "@/components/book-manage/import-sheet";
import getAllBookForSale from "@/lib/book/getAllBookForSale";
import { useLoading } from "@/hooks/loading-context";
import InventoryCheckNoteAddSkeleton from "@/components/skeleton/inventory-check-note-add-skeleton";

export const FormSchema = z.object({
  id: z.string().max(12, "Tối đa 12 ký tự"),
  details: z
    .array(
      z.object({
        bookId: z.string(),
        difference: z.coerce
          .number({
            invalid_type_error: "Chênh lệch phải là một số",
          })
          .refine((value) => Number.isInteger(value), {
            message: "Chênh lệch phải là số nguyên",
          })
          .refine((value) => value !== 0, {
            message: "Chênh lệch phải khác 0",
          }),
        initial: z.coerce.number(),
      })
    )
    .nonempty("Vui lòng chọn ít nhất một sách nhập"),
});

const AddNote = () => {
  const [openDialog, setOpenDialog] = useState(false);
  const { mutate } = useSWRConfig();

  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      id: "",
      details: [],
    },
  });
  const {
    register,
    handleSubmit,
    setValue,
    trigger,
    watch,
    reset,
    formState: { errors, isDirty },
  } = form;
  const { mutate: mutateAllBook } = useSWRConfig();
  const { showLoading, hideLoading } = useLoading();

  const onSubmit: SubmitHandler<z.infer<typeof FormSchema>> = async (data) => {
    const response: Promise<any> = createCheckNote({
      id: data.id,
      details: data.details.map((item) => {
        return {
          bookId: item?.bookId,
          difference: item?.difference,
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
      toast({
        variant: "success",
        title: "Thành công",
        description: "Thêm phiếu kiểm kho thành công",
      });
      reset({
        id: "",
        details: [],
      });
      mutateAllBook(`${endPoint}/v1/books/all`);
    }
  };
  const onError: SubmitErrorHandler<z.infer<typeof FormSchema>> = (data) => {
    if (data.hasOwnProperty("details")) {
      toast({
        variant: "destructive",
        title: "Có lỗi",
        description: data.details?.message,
      });
    }
  };
  const { books, isLoading, isError } = getAllBookForSale();

  const handleFile = (reader: FileReader) => {
    let importNote = {
      id: "",
      supplierId: "",
      details: [{}],
    };
    const ExcelJS = require("exceljs");
    const wb = new ExcelJS.Workbook();
    reader.onload = () => {
      const buffer = reader.result;
      wb.xlsx.load(buffer).then((workbook: any) => {
        workbook.eachSheet((sheet: any, id: any) => {
          sheet.eachRow((row: any, rowIndex: number) => {
            if (rowIndex === 1) {
              importNote.id =
                row.getCell(2).value === "Nhập mã phiếu dưới 12 ký tự"
                  ? ""
                  : row.getCell(2).value;
            }
            if (rowIndex > 2) {
              const idBook = row.getCell(1).value.toString();
              const oldQty = books?.data.find(
                (ingre) => ingre.id === idBook
              )?.quantity;

              if (oldQty) {
                const detail = {
                  bookId: idBook,
                  difference: row.getCell(3).value,
                  initial: oldQty,
                };

                importNote.details.push(detail);
              } else {
                toast({
                  variant: "destructive",
                  title: "Có lỗi",
                  description:
                    "Vui lòng kiểm tra thông tin điền vào file đúng mẫu",
                });
                return;
              }
            }
          });
          reset({
            id: importNote.id,
            details: importNote.details.filter(
              (value) => JSON.stringify(value) !== "{}"
            ),
          });
        });
      });
    };
  };
  const { currentUser } = useCurrentUser();
  if (!currentUser) {
    return <InventoryCheckNoteAddSkeleton />;
  } else if (
    currentUser &&
    !includesRoles({
      currentUser: currentUser,
      allowedFeatures: ["INVENTORY_NOTE_CREATE"],
    })
  ) {
    return <NoRole></NoRole>;
  } else
    return (
      <div className="col items-center">
        <div className="col xl:w-4/5 w-full px-0">
          <div className="flex justify-between gap-2">
            <h1 className="font-medium text-xxl self-start">
              Thêm phiếu kiểm kho
            </h1>
            <ImportSheet
              sampleFileLink="/check-sample.xlsx"
              handleFile={handleFile}
            />
          </div>

          <form onSubmit={handleSubmit(onSubmit, onError)}>
            <div className="flex flex-col gap-4">
              <div className="flex lg:flex-row flex-col gap-4">
                <Card className="flex-1">
                  <CardContent className="lg:p-6 p-4 flex lg:flex-row flex-col gap-4">
                    <div className="flex-1">
                      <Label>Mã phiếu</Label>
                      <Input
                        placeholder="Mã sinh tự động nếu để trống"
                        {...register("id")}
                      ></Input>
                      {errors.id && (
                        <span className="error___message">
                          {errors.id.message}
                        </span>
                      )}
                    </div>
                  </CardContent>
                </Card>
              </div>
              <Card>
                <CardContent className="lg:p-6 p-4">
                  <CheckInsert form={form} />
                </CardContent>
              </Card>
              <div className="flex md:justify-end justify-stretch gap-2">
                <Button
                  className="px-4 bg-white md:flex-none flex-1"
                  variant={"outline"}
                  type="button"
                  onClick={() => {
                    reset({
                      id: "",
                      details: [],
                    });
                  }}
                >
                  <div className="flex flex-wrap gap-2 items-center">
                    <FiTrash2 className="text-muted-foreground" />
                    Hủy
                  </div>
                </Button>
                <Button className="px-4 pl-2 md:flex-none  flex-1">
                  <div className="flex flex-wrap gap-2 items-center">
                    <LuCheck />
                    Thêm
                  </div>
                </Button>
              </div>
            </div>
          </form>
        </div>
      </div>
    );
};

export default AddNote;
