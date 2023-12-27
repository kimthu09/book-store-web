import { apiKey, endPoint } from "@/constants";
import { PagingProps, Publisher } from "@/types";
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

export default function getAllPublisher({
  page,
  limit,
}: {
  page?: string;
  limit?: number;
}) {
  const { data, error, isLoading, mutate } = useSWR(
    `${endPoint}/v1/publishers?page=${page ?? 1}&limit=${limit ?? 10}`,
    fetcher
  );

  return {
    publishers: data,
    isLoading,
    isError: error,
    mutate: mutate,
  };
}
