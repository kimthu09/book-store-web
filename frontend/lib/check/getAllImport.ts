import { endPoint } from "@/constants";
import useSWR from "swr";
import { getApiKey } from "../auth/action";

export default function getAllCheckNote({
  page,
  limit,
  filterString,
}: {
  page: string;
  limit?: number;
  filterString?: string;
}) {
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
          paging: json.paging,
          data: json.data,
        };
      });
  };
  const { data, error, isLoading, mutate, isValidating } = useSWR(
    `${endPoint}/v1/inventoryCheckNotes?page=${page}&limit=${limit ?? "10"}${
      filterString ?? ""
    }`,
    fetcher
  );
  return {
    data: data,
    isLoading,
    isError: error,
    mutate: mutate,
    isValidating: isValidating,
  };
}
