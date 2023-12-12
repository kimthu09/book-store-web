import { apiKey } from "@/constants";

export default async function getAllBooks(page: number) {
  const res = await fetch(`http://localhost:8080/v1/booktitles?page=${page}`, {
    headers: {
      accept: "application/json",
      Authorization: apiKey,
    },
  });
  if (!res.ok) {
    console.log("quao");

    throw new Error("Failed to fetch data");
  }
  return res.json().then((json) => {
    console.log(json);
    return {
      paging: json.paging,
      data: json.data,
    };
  });
}
