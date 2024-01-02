import { endPoint } from "@/constants";
import { Staff } from "@/types";
import useSWR from "swr";
import { getApiKey } from "../auth/action";

const fetcher = async (url: string) => {
  const token = await getApiKey();
  return fetch(url, {
    headers: {
      accept: "application/json",
      Authorization: `Bearer ${token}`,
    },
    cache: "no-store",
  })
    .then((res) => {
      return res.json();
    })
    .then((json) => json.data);
};

export default function getStaff(idStaff: string) {
  const { data, error, isLoading, mutate } = useSWR(
    `${endPoint}/v1/users/${idStaff}`,
    fetcher
  );

  return {
    data: data as Staff,
    isLoading,
    isError: error,
    mutate: mutate,
  };
}
