import { endPoint } from "@/constants";
import { Supplier } from "@/types";
import useSWR from "swr";
import { getApiKey } from "../auth/action";

export default function getAllSupplier() {

  const fetcher = async (url: string) =>
    {
      const token=await getApiKey()
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
          return json.data as Supplier[];
        });
    }
  const { data, error, isLoading, mutate } = useSWR(
    `${endPoint}/v1/suppliers/all`,
    fetcher
  );

  return {
    suppliers: data ,
    isLoading,
    isError: error,
    mutate: mutate,
  };
}
