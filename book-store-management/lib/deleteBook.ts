import { apiKey } from "@/constants";
import axios from "axios";

export default async function deleteBook(bookId: string) {
  const url = `http://localhost:8080/v1/books/:${bookId}`;

  const headers = {
    accept: "application/json",
    Authorization: apiKey,
    // Add other headers as needed
  };

  // Make a POST request with headers
  axios
    .delete(url, { headers: headers })
    .then((response) => {
      console.log("Response:", response.data);
    })
    .catch((error) => {
      console.error("Error:", error);
    });
}
