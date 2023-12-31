import { endPoint } from "@/constants";
import axios from "axios";
import { getApiKey } from "../auth/action";

export default async function deleteBookTitle(bookId: string) {
  const url = `${endPoint}/v1/booktitles/${bookId}`;
  const token = await getApiKey();
  const headers = {
    accept: "application/json",
    Authorization: `Bearer ${token}`, // Add other headers as needed
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
