import { apiKey, endPoint } from "@/constants";
export type FilterProps = {
  page: number;
  maxSellPrice?: string;
  minSellPrice?: string;
  search?: string;
  publisher?: string;
  categoryIds?: string;
  authorIds?: string;
};
export default async function getAllBooks({
  page,
  maxSellPrice,
  minSellPrice,
  search,
  publisher,
  categoryIds,
  authorIds,
}: FilterProps) {
  const maxString = maxSellPrice ? `&maxSellPrice=${maxSellPrice}` : "";
  const minString = minSellPrice ? `&minSellPrice=${minSellPrice}` : "";
  const searchString = search ? `&search=${search}` : "";
  const publisherString = publisher ? `&publisher=${publisher}` : "";
  const categoriesString = categoryIds ? `&categories=${categoryIds}` : "";
  const authorsString = authorIds ? `&authors=${authorIds}` : "";
  const url = `${endPoint}/v1/books?page=${page}${maxString}${minString}${searchString}${publisherString}${categoriesString}${authorsString}`;
  console.log(url);
  const res = await fetch(url, {
    headers: {
      accept: "application/json",
      Authorization: apiKey,
    },
    next: {
      revalidate: 0,
    },
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
