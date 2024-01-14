"use client";
import { useCallback, useEffect, useRef, useState } from "react";
import Image from "next/image";
import { Input } from "../ui/input";
import {
  FieldArrayWithId,
  UseFieldArrayAppend,
  UseFieldArrayUpdate,
} from "react-hook-form";
import { toVND } from "@/lib/utils";
import { AspectRatio } from "@radix-ui/react-aspect-ratio";
import Loading from "../loading";
import { BookProps } from "@/types";
import getAllBookForSale from "@/lib/book/getAllBookForSale";
import { FormValues } from "@/app/sale/page-layout";
import getAllCategoryList from "@/lib/book/getAllCategoryList";
import ProductSkeleton from "../skeleton/product-skeleton";
const ProductTab = ({
  append,
  fields,
  update,
}: {
  append: UseFieldArrayAppend<FormValues, "details">;
  fields: FieldArrayWithId<FormValues, "details", "id">[];
  update: UseFieldArrayUpdate<FormValues, "details">;
}) => {
  const { categories, isLoading, isError, mutate } = getAllCategoryList();
  const {
    books,
    isLoading: isLoadingBook,
    isError: isErrorBook,
    mutate: mutateBook,
  } = getAllBookForSale();
  const [cateList, setCateList] = useState<
    {
      id: string;
      name: string;
      isSelected: boolean;
    }[]
  >();
  const [all, setAll] = useState(true);
  const [prodList, setProdList] = useState<Array<BookProps>>();
  const handleAllSelected = () => {
    if (!all) {
      setAll((prev) => !prev);
      setCateList(
        categories?.data.map((item: any) => {
          return { id: item.id, name: item.name, isSelected: false };
        })
      );
      setProdList(books?.data.filter((item: BookProps) => item.quantity > 0));
    }
  };

  const handleCateSelected = (id: string) => {
    if (all) {
      setAll(false);
      setProdList(new Array());
    }
    const newCateList = cateList?.map((item: any) => {
      return {
        id: item.id,
        name: item.name,
        isSelected: item.id === id ? !item.isSelected : item.isSelected,
      };
    });
    setCateList(newCateList);

    if (newCateList?.every((item) => !item.isSelected)) {
      handleAllSelected();
    } else {
      const categorySet = new Set(
        newCateList
          ?.filter((item: any) => item.isSelected === true)
          .map((value) => value.id)
      );
      const newProdList = new Array<BookProps>();
      books?.data.forEach((book: BookProps) => {
        for (let element of book.bookTitle.categories) {
          if (categorySet.has(element.id)) {
            if (book.quantity > 0) {
              newProdList.push(book);
              break;
            }
          }
        }
      });
      setProdList(newProdList);
    }
  };

  const [inputValue, setInputValue] = useState<string>("");
  const [filteredList, setFilteredList] = useState<Array<BookProps>>();

  // Search Handler
  const searchHandler = useCallback(() => {
    const filteredData = prodList?.filter((prod) => {
      return prod.name.toLowerCase().includes(inputValue.toLowerCase());
    });
    setFilteredList(filteredData);
  }, [prodList, inputValue]);

  // EFFECT: Search Handler
  useEffect(() => {
    // Debounce search handler
    const timer = setTimeout(() => {
      searchHandler();
    }, 500);

    // Cleanup
    return () => {
      clearTimeout(timer);
    };
  }, [searchHandler]);
  const inputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    document.addEventListener("keydown", detectKeyDown, true);
  }, []);
  const detectKeyDown = (e: any) => {
    if (e.key === "F2") {
      inputRef.current?.focus();
    }
    return () => {
      document.removeEventListener("keydown", detectKeyDown);
    };
  };
  useEffect(() => {
    if (categories) {
      setCateList(
        categories?.data.map((item: any) => {
          return { id: item.id, name: item.name, isSelected: false };
        })
      );
    }
  }, [categories]);

  useEffect(() => {
    if (books) {
      setProdList(books?.data.filter((item: BookProps) => item.quantity > 0));
    }
  }, [books]);
  if (isError || isErrorBook) return <div>Failed to load</div>;
  if (isLoading || isLoadingBook) {
    return <ProductSkeleton />;
  } else
    return (
      <div className="flex flex-col gap-6">
        <div className="flex items-end">
          <Input
            ref={inputRef}
            value={inputValue}
            onChange={(e) => {
              setInputValue(e.target.value);
            }}
            className=" bg-white rounded-xl"
            placeholder="Tìm kiếm sản phẩm (F2)"
          ></Input>
        </div>

        {/* Category list */}
        <div className="flex flex-wrap gap-2">
          <div
            className={` rounded-xl flex self-start px-3 py-1 border-gray-200 border text-sm font-medium cursor-pointer ${
              all
                ? "bg-blue-50 border-primary text-slate-800 font-medium"
                : "bg-white text-muted-foreground"
            }`}
            onClick={handleAllSelected}
          >
            Tất cả
          </div>
          {cateList?.map((item) => (
            <div
              key={item.id}
              className={`rounded-xl flex self-start px-3 py-1 border-gray-200 border text-sm font-medium cursor-pointer
            ${
              item.isSelected
                ? "bg-blue-50 border-primary text-brown "
                : "bg-white text-muted-foreground"
            }`}
              onClick={() => handleCateSelected(item.id)}
            >
              {item.name}
            </div>
          ))}
        </div>
        <h1 className="text-lg">Sản phẩm</h1>

        {/* Product list */}
        <div className="grid 2xl:grid-cols-5 xl:grid-cols-4 lgr:grid-cols-3 md:grid-cols-2 sm:grid-cols-4 grid-cols-3 gap-4">
          {filteredList?.map((prod) => {
            return (
              <div
                key={prod.id}
                className="bg-white shadow-sm rounded-xl  overflow-hidden cursor-pointer hover:shadow-md"
                onClick={() => {
                  const index = fields.findIndex(
                    (value) => value.bookId === prod.id
                  );
                  if (index > -1) {
                    update(index, {
                      bookId: prod.id,
                      qty: +fields.at(index)?.qty! + 1,
                      sellPrice: prod.sellPrice,
                      name: prod.name,
                      stock: prod.quantity,
                    });
                  } else {
                    append({
                      bookId: prod.id,
                      qty: 1,
                      sellPrice: prod.sellPrice,
                      name: prod.name,
                      stock: prod.quantity,
                    });
                  }
                }}
              >
                <AspectRatio ratio={1 / 1}>
                  <Image
                    className=" object-cover"
                    src={
                      prod.img ??
                      "https://img.freepik.com/free-psd/feel-nature-flyer-template_23-2148607911.jpg?w=826&t=st=1702908240~exp=1702908840~hmac=a15dbebf7250428d66641b47b0de66f0964dde722046fd9cac2378f4ee44b2da"
                    }
                    alt="image"
                    fill
                    sizes="(max-width: 768px) 33vw, 20vw"
                  ></Image>
                </AspectRatio>
                <div className="px-1">
                  <h1 className="text-base font-medium text-center">
                    {prod.name}
                  </h1>
                  <h1 className="font-light text-sm text-center">{prod.id}</h1>
                  <h1 className="text-base font-semibold text-primary text-center pb-1">
                    {toVND(prod.sellPrice)}
                  </h1>
                </div>
              </div>
            );
          })}
        </div>
      </div>
    );
};

export default ProductTab;
