import { endPoint } from "@/constants";
import { getApiKey } from "../auth/action";

export type FilterProps = {
  page: number;
  maxDebt?: string;
  minDebt?: string;
  search?: string;
};
export default async function getListSupplier({
  page,
  maxDebt,
  minDebt,
  search,
}: FilterProps) {
  const maxString = maxDebt ? `&maxDebt=${maxDebt}` : "";
  const minString = minDebt ? `&minDebt=${minDebt}` : "";
  const searchString = search ? `&search=${search}` : "";
  const url = `${endPoint}/v1/suppliers?page=${page}${maxString}${minString}${searchString}`;

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
