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
  })
    .then((res) => {
      return res.json();
    })
    .then((json) => json.data);
};

export default function getAllStaff() {
  const { data, error, isLoading } = useSWR(
    `${endPoint}/v1/users/all`,
    fetcher
  );

  return {
    staffs: data as Staff[],
    isLoading,
    isError: error,
  };
}
