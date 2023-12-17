import { apiKey } from "@/constants";
import axios from "axios";

export default async function createStaff({ staff }: { staff: {} }) {
  const url = "http://localhost:8080/v1/users";

  // const data = {
  //   name: name,
  //   features: features,
  // };
  const headers = {
    accept: "application/json",
    "Content-Type": "application/json",
    Authorization: apiKey,

    // Add other headers as needed
  };

  // Make a POST request with headers
  const res = axios
    .post(url, staff, { headers: headers })
    .then((response) => {
      if (response) return response.data;
    })
    .catch((error) => {
      console.error("Error:", error);
      return error.response.data;
    });
  return res;
}
