import { apiKey, endPoint } from "@/constants";
import axios from "axios";

export default async function createRole({
  name,
  features,
}: {
  name: string;
  features: string[];
}) {
  const url = `${endPoint}/v1/roles`;

  const data = {
    name: name,
    features: features,
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
      if (response) return response.data;
    })
    .catch((error) => {
      console.error("Error:", error);
      return error.response.data;
    });
  return res;
}
