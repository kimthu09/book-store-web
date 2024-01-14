import { endPoint } from "@/constants";
import { Supplier } from "@/types";
import useSWR from "swr";
import { getApiKey } from "../auth/action";

const fetcher = async (url: string) => {
  const token = await getApiKey();
  return fetch(url, {
    headers: {
      accept: "application/json",
      Authorization: `Bearer ${token}`,
    },
    cache: "no-store",
  })
    .then((res) => {
      return res.json();
    })
    .then((json) => json.data);
};

export default function getSupplier(idSupplier: string) {
  const { data, error, isLoading, mutate } = useSWR(
    `${endPoint}/v1/suppliers/${idSupplier}`,
    fetcher
  );

  return {
    data: data as Supplier,
    isLoading,
    isError: error,
    mutate,
  };
}
