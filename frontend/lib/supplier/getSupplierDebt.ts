import { endPoint } from "@/constants";
import useSWR from "swr";
import { getApiKey } from "../auth/action";

const fetcher = async (url: string) => {
  const token = await getApiKey();
  return fetch(url, {
    headers: {
      accept: "application/json",
      Authorization: `Bearer ${token}`,
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
};

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
