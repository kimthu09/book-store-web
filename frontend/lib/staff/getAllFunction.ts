import { apiKey, endPoint } from "@/constants";
import { RoleFunction } from "@/types";
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

export default function getAllRoleFunction() {
  const { data, error, isLoading } = useSWR(`${endPoint}/v1/features`, fetcher);

  return {
    roleFunctions: data as RoleFunction[],
    isLoading,
    isError: error,
  };
}
