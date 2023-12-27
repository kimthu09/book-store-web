import { apiKey, endPoint } from "@/constants";
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
