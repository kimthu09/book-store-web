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

export default function getAllBookForSale() {
  const { data, error, isLoading, mutate } = useSWR(
    `${endPoint}/v1/books/all`,
    fetcher
  );

  return {
    books: data,
    isLoading,
    isError: error,
    mutate: mutate,
  };
}
