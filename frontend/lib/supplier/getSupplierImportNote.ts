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
    `${endPoint}/v1/suppliers/${idSupplier}/importNotes?limit=${
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
