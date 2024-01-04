import { endPoint } from "@/constants";
import { Book } from "@/types";
import useSWR from "swr";

export default function getAllBookList(token: string) {
  const fetcher = (url: string) =>
    fetch(url, {
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
          data: json.data as Book[],
        };
      });
  const { data, error, isLoading, mutate } = useSWR(
    `${endPoint}/books/all`,
    fetcher
  );

  return {
    data: data,
    isLoading,
    isError: error,
    mutate: mutate,
  };
}
