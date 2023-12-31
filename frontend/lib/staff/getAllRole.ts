import { endPoint } from "@/constants";
import { Role } from "@/types";
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

export default function getAllRole() {
  const { data, error, isLoading } = useSWR(`${endPoint}/v1/roles`, fetcher);

  return {
    roles: data as Role[],
    isLoading,
    isError: error,
  };
}
