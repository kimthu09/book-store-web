import { endPoint } from "@/constants";
import { getApiKey } from "../auth/action";

export type FilterProps = {
  page: number;
  maxPoint?: string;
  minPoint?: string;
  search?: string;
};
export default async function getListCustomer({
  page,
  maxPoint,
  minPoint,
  search,
}: FilterProps) {
  const maxString = maxPoint ? `&maxPoint=${maxPoint}` : "";
  const minString = minPoint ? `&minPoint=${minPoint}` : "";
  const searchString = search ? `&search=${search}` : "";
  const url = `${endPoint}/v1/customers?page=${page}${maxString}${minString}${searchString}`;

  const token = await getApiKey();
  const res = await fetch(url, {
    headers: {
      accept: "application/json",
      Authorization: `Bearer ${token}`,
    },
    cache: "no-store",
  });

  if (!res.ok) {
    throw new Error("Failed to fetch data");
  }
  return res.json().then((json) => {
    return {
      paging: json.paging,
      data: json.data,
    };
  });
}
