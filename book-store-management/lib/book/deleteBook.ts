import { apiKey } from "@/constants";
import axios from "axios";

export default async function deleteBook(bookId: string) {
  const url = `http://localhost:8080/v1/booktitles/${bookId}`;

  const headers = {
    accept: "application/json",
    Authorization: apiKey,
    // Add other headers as needed
  };

  // Make a POST request with headers
  const res = axios
    .delete(url, { headers: headers })
    .then((response) => {
      return response.data;
    })
    .catch((error) => {
      console.error("Error:", error);
    });
  return res;
}
