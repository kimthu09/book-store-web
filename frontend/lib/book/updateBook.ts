import { endPoint } from "@/constants";
import axios from "axios";
import { getApiKey } from "../auth/action";

export type Props = {
  bookTitleId: string;
  edition: number;
  image: string;
  listedPrice: number;
  publisherId: string;
  sellPrice: number;
};
export default async function updateBook(data: Props, bookId: string) {
  const url = `${endPoint}/v1/books/${bookId}/info`;

  const token = await getApiKey();
  const headers = {
    accept: "application/json",
    "Content-Type": "application/json",
    Authorization: `Bearer ${token}`,
    // Add other headers as needed
  };

  // Make a POST request with headers
  const res = axios
    .patch(url, data, { headers: headers })
    .then((response) => {
      if (response) return response.data;
    })
    .catch((error) => {
      console.error("Error:", error);
      return error.response.data;
    });
  return res;
}
