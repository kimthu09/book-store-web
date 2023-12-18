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

export default function getAllCategory({
  page,
  limit,
}: {
  page?: string;
  limit?: number;
}) {
  const { data, error, isLoading } = useSWR(
    `${endPoint}/v1/categories?page=${page ?? 1}&limit=${limit ?? 10}`,
    fetcher
  );

  return {
    categories: data,
    isLoading,
    isError: error,
  };
}
