import { apiKey } from "@/constants";
import { Category } from "@/types";
import useSWR from "swr";

const fetcher = (url: string) =>
  fetch(url, {
    headers: {
      accept: "application/json",
    },
  })
    .then((res) => {
      return res.json();
    })
    .then((json) => json.data);

export default function getAllCategory() {
  const { data, error, isLoading } = useSWR(
    "http://localhost:8080/v1/categories",
    fetcher,
    {
      dedupingInterval: 1,
    }
  );

  return {
    categories: data as Category[],
    isLoading,
    isError: error,
  };
  // const res = await fetch("http://192.168.1.33:8080/v1/categories", {
  //   headers: {
  //     accept: "application/json",
  //     Authorization: apiKey,
  //   },
  // });
  // if (!res.ok) {
  //   throw new Error("Failed to fetch data");
  // }
  // return res.json().then((json) => json.data);
}
