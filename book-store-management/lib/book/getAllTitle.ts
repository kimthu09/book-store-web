import { apiKey, endPoint } from "@/constants";
import { Author, BookTitle, PagingProps } from "@/types";
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

export default function getAllTitle({
  limit,
  page,
  filter,
}: {
  limit?: number;
  page?: string;
  filter?: string;
}) {
  const { data, error, isLoading, mutate } = useSWR(
    `${endPoint}/v1/booktitles?page=${page ?? 1}&limit=${limit ?? 10}${
      filter ?? ""
    }`,
    fetcher
  );

  return {
    titles: data,
    isLoading,
    isError: error,
    mutate: mutate,
  };
}
