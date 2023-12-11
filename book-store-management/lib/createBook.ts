import { apiKey } from "@/constants";
import axios from "axios";

export default async function createBook({
  name,
  desc,
  categoryIds,
}: {
  name: string;
  desc: string;
  categoryIds: string[];
}) {
  const url = "http://localhost:8080/v1/books";
  const data = {
    name: name,
    desc: desc,
    categoryIds: categoryIds,
    listedPrice: 75000,
    authorIds: ["tgnna"],
    edition: 1,
    publisherId: "nxbdk",
    quantity: 0,
    sellPrice: 80000,
  };
  console.log(data);
  const headers = {
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
