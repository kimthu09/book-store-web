import { endPoint } from "@/constants";
import { Author } from "@/types";
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
        data: json.data as Author[],
      };
    });
};

export default function getAllAuthorList() {
  const { data, error, isLoading, mutate } = useSWR(
    `${endPoint}/v1/authors/all`,
    fetcher
  );

  return {
    authors: data,
    isLoading,
    isError: error,
    mutate: mutate,
  };
}
