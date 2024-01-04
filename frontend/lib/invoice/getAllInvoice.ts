import { endPoint } from "@/constants";
import { getApiKey } from "../auth/action";

export type FilterProps = {
  page: number;
  maxPrice?: string;
  minPrice?: string;
  search?: string;
  createdBy?: string;
};
export default async function getAllInvoice({
  page,
  maxPrice,
  minPrice,
  search,
  createdBy,
}: FilterProps) {
  const maxString = maxPrice ? `&maxPrice=${maxPrice}` : "";
  const minString = minPrice ? `&minPrice=${minPrice}` : "";
  const searchString = search ? `&search=${search}` : "";
  const createdByString = createdBy ? `&createdBy=${createdBy}` : "";
  const url = `${endPoint}/v1/invoices?page=${page}${maxString}${minString}${searchString}${createdByString}`;

  const token = await getApiKey();
  const res = await fetch(url, {
    headers: {
      accept: "application/json",
      Authorization: `Bearer ${token}`,
    },
    cache: "no-store",
  });

  if (!res.ok) {
    console.error(res.json);
    return res.json();
  }

  return res.json().then((json) => {
    return {
      paging: json.paging,
      data: json.data,
    };
  });
}
