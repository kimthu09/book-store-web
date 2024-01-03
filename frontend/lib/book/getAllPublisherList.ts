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
        data: json.data,
      };
    });
};

export default function getAllPublisherList() {
  const { data, error, isLoading, mutate } = useSWR(
    `${endPoint}/v1/publishers/all`,
    fetcher
  );

  return {
    publishers: data,
    isLoading,
    isError: error,
    mutate: mutate,
  };
}
