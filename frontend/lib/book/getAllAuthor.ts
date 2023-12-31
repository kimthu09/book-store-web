import { endPoint } from "@/constants";
import { Author, PagingProps } from "@/types";
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
        paging: json.paging as PagingProps,
        data: json.data as Author[],
      };
    });
};

export default function getAllAuthor({
  limit,
  page,
}: {
  limit?: number;
  page?: string;
}) {
  const { data, error, isLoading, mutate } = useSWR(
    `${endPoint}/v1/authors?page=${page ?? 1}&limit=${limit ?? 10}`,
    fetcher
  );

  return {
    authors: data,
    isLoading,
    isError: error,
    mutate: mutate,
  };
}
