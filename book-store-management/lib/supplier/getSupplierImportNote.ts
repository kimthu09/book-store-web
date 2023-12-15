import { apiKey } from "@/constants";
import { ImportNote, PagingProps } from "@/types";
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

export default function getSupplierImportNote({
  idSupplier,
  page,
}: {
  idSupplier: string;
  page: number;
}) {
  const { data, error, isLoading } = useSWR(
    `http://localhost:8080/v1/suppliers/${idSupplier}/importNotes?page=${page}`,
    fetcher
  );
  return {
    data: data,
    isLoading,
    isError: error,
  };
}
