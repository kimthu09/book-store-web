import { apiKey } from "@/constants";

export default async function getAllBooks() {
  const res = await fetch("http://localhost:8080/v1/books?page=1", {
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
    return json.data;
  });
}
