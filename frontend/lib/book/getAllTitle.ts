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
