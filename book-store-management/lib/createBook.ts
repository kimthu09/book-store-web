import { apiKey } from "@/constants";
import axios from "axios";

export default async function createBook({
  id,
  name,
  desc,
  categoryIds,
  authorIds,
}: {
  id?: string;
  name: string;
  desc: string;
  categoryIds: string[];
  authorIds: string[];
}) {
  const url = "http://localhost:8080/v1/booktitles";

  const data = {
    id: id,
    name: name,
    desc: desc,
    categoryIds: categoryIds,
    authorIds: authorIds,
  };
  console.log(data);
  const headers = {
    accept: "application/json",
    "Content-Type": "application/json",

    // Add other headers as needed
  };

  // Make a POST request with headers
  const res = axios
    .post(url, data, { headers: headers })
    .then((response) => {
      if (response) return response.data;
    })
    .catch((error) => {
      console.error("Error:", error);
      return error.response.data;
    });
  return res;
}
