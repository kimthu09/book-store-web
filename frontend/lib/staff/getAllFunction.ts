import { endPoint } from "@/constants";
import { RoleFunction } from "@/types";
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
    .then((json) => json.data);
};

export default function getAllRoleFunction() {
  const { data, error, isLoading } = useSWR(`${endPoint}/v1/features`, fetcher);

  return {
    roleFunctions: data as RoleFunction[],
    isLoading,
    isError: error,
  };
}
