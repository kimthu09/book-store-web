import { apiKey, endPoint } from "@/constants";
import { Staff } from "@/types";
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

export default function getAllStaff({
  page,
  limit,
}: {
  page?: string;
  limit?: number;
}) {
  const { data, error, isLoading } = useSWR(
    `${endPoint}/v1/users?page=${page ?? 1}&limit=${limit ?? 10}`,
    fetcher
  );

  return {
    staffs: data as Staff[],
    isLoading,
    isError: error,
  };
}
