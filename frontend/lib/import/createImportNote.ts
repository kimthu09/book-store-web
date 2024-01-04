import { endPoint } from "@/constants";
import axios from "axios";
import { getApiKey } from "../auth/action";

export default async function createImportNote({
  details,
  id,
  supplierId,
}: {
  details: {
    bookId: string;
    qtyImport: number;
    price: number;
    isReplacePrice?: boolean;
  }[];
  id?: string;
  supplierId: string;
}) {
  const url = `${endPoint}/v1/importNotes`;
  const token = await getApiKey();

  const data = {
    details,
    id: id,
    supplierId: supplierId,
  };

  const headers = {
    accept: "application/json",
    "Content-Type": "application/json",
    Authorization: `Bearer ${token}`,

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
