import { apiKey } from "@/constants";
import { Supplier } from "@/types";
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
    .then((json) => json.data);

export default function getSupplier(idSupplier: string) {
  const { data, error, isLoading } = useSWR(
    `http://localhost:8080/v1/suppliers/${idSupplier}`,
    fetcher
  );

  return {
    data: data as Supplier,
    isLoading,
    isError: error,
  };
}
