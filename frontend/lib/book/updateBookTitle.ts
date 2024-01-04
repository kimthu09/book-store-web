import { endPoint } from "@/constants";
import axios from "axios";
import { getApiKey } from "../auth/action";

export default async function updateBookTitle({
  id,
  name,
  desc,
  categoryIds,
  authorIds,
}: {
  id: string;
  name: string;
  desc: string;
  categoryIds: string[];
  authorIds: string[];
}) {
  const url = `${endPoint}/v1/booktitles/${id}/info`;

  const data = {
    id: id,
    name: name,
    desc: desc,
    categoryIds: categoryIds,
    authorIds: authorIds,
  };

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
