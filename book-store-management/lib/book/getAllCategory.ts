import { apiKey } from "@/constants";
import { Category } from "@/types";
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
    .then((json) => json.data);

export default function getAllCategory() {
  const { data, error, isLoading } = useSWR(
    "http://localhost:8080/v1/categories",
    fetcher
  );

  return {
    categories: data as Category[],
    isLoading,
    isError: error,
  };
}
