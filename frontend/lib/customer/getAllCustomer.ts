import { endPoint } from "@/constants";
import { Customer, Supplier } from "@/types";
import useSWR from "swr";
import { getApiKey } from "../auth/action";

export default function getAllCustomer() {
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
        return json.data as Customer[];
      });
  };
  const { data, error, isLoading, mutate } = useSWR(
    `${endPoint}/v1/customers/all`,
    fetcher
  );

  return {
    suppliers: data,
    isLoading,
    isError: error,
    mutate: mutate,
  };
}
