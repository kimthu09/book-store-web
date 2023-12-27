import { apiKey, endPoint } from "@/constants";
import { PagingProps, SupplierDebt } from "@/types";
import useSWR from "swr";

const fetcher = (url: string) =>
  fetch(url, {
    headers: {
      accept: "application/json",
      Authorization: apiKey,
    },
  })
    .then((res) => {
      return res.json();
    })
    .then((json) => {
      return {
        paging: json.paging,
        data: json.data,
      };
    });

export default function getSupplierDebt({
  idSupplier,
  page,
}: {
  idSupplier: string;
  page: number;
}) {
  const { data, error, isLoading, mutate, isValidating } = useSWR(
    `${endPoint}/v1/suppliers/${idSupplier}/debts?page=${page}`,
    fetcher
  );
  return {
    data: data,
    isLoading,
    isError: error,
    mutate: mutate,
    isValidating: isValidating,
  };
}
