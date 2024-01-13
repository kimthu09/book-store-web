import { endPoint } from "@/constants";
import { getApiKey } from "../auth/action";
import axios from "axios";

export default async function getDashboard({
  timeFrom,
  timeTo,
}: {
  timeFrom: number;
  timeTo: number;
}) {
  const url = `${endPoint}/v1/dashboard`;
  const data = {
    timeFrom: Math.floor(timeFrom),
    timeTo: Math.floor(timeTo),
  };

  const token = await getApiKey();

  const headers = {
    accept: "application/json",
    "Content-Type": "application/json",
    Authorization: `Bearer ${token}`,
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
