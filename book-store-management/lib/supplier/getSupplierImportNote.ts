import { apiKey } from "@/constants";
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

export default function getSupplierImportNote({
  idSupplier,
  page,
  limit,
}: {
  idSupplier: string;
  page: number;
  limit?: number;
}) {
  const { data, error, isLoading } = useSWR(
    `http://localhost:8080/v1/suppliers/${idSupplier}/importNotes?limit=${
      limit ?? 10
    }&page=${page}`,
    fetcher
  );
  return {
    data: data,
    isLoading,
    isError: error,
  };
}
