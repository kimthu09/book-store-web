import { apiKey } from "@/constants";
import { Author } from "@/types";
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

export default function getAllAuthor() {
  const { data, error, isLoading } = useSWR(
    "http://localhost:8080/v1/authors",
    fetcher
  );

  return {
    authors: data as Author[],
    isLoading,
    isError: error,
  };
}
