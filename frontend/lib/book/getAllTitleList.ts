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

export default function getAllTitleList() {
  const { data, error, isLoading, mutate } = useSWR(
    `${endPoint}/v1/booktitles/all`,
    fetcher
  );

  return {
    titles: data,
    isLoading,
    isError: error,
    mutate: mutate,
  };
}
