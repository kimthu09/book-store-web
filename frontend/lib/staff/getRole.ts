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
    .then((json) => json.data);
};

export default function getRole(idRole: string) {
  const { data, error, isLoading } = useSWR(
    `${endPoint}/v1/roles/${idRole}`,
    fetcher
  );

  return {
    response: data,
    isLoading,
    isError: error,
  };
}
