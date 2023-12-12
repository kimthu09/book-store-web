import { apiKey } from "@/constants";
import axios from "axios";

export default async function createBook({
  id,
  name,
  desc,
  categoryIds,
}: {
  id?: string;
  name: string;
  desc: string;
  categoryIds: string[];
}) {
  const url = "http://localhost:8080/v1/booktitles";

  const data = {
    id: id,
    name: name,
    desc: desc,
    categoryIds: categoryIds,
    authorIds: ["tgnna"],
  };
  console.log(data);
  const headers = {
    accept: "application/json",
    "Content-Type": "application/json",
    Authorization: apiKey,

    // Add other headers as needed
  };

  // Make a POST request with headers
  const res = axios
    .post(url, data, { headers: headers })
    .then((response) => {
      return response.data;
    })
    .catch((error) => {
      console.error("Error:", error);
    });
  return res;
}
