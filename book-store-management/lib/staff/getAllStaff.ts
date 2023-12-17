import { apiKey } from "@/constants";

export type FilterProps = {
  page: number;
  isActive?: string;
  search?: string;
};
export default async function getAllStaff({
  page,
  isActive,
  search,
}: FilterProps) {
  const isActiveString = isActive ?? "";
  const searchString = search ? `&search=${search}` : "";
  const url = `http://localhost:8080/v1/users?page=${page}${isActiveString}${searchString}`;
  console.log(url);

  const res = await fetch(url, {
    headers: {
      accept: "application/json",
      Authorization: apiKey,
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
