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

export default function getAllCategory({
  page,
  limit,
}: {
  page?: string;
  limit?: number;
}) {
  const { data, error, isLoading, mutate } = useSWR(
    `${endPoint}/v1/categories?page=${page ?? 1}&limit=${limit ?? 10}`,
    fetcher
  );

  return {
    categories: data,
    isLoading,
    isError: error,
    mutate: mutate,
  };
}
