import { apiKey, endPoint } from "@/constants";
import { Role } from "@/types";
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

export default function getAllRole() {
  const { data, error, isLoading } = useSWR(`${endPoint}/v1/roles`, fetcher);

  return {
    roles: data as Role[],
    isLoading,
    isError: error,
  };
}