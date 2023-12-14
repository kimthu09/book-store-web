import { apiKey } from "@/constants";
import axios, { isCancel, AxiosError } from "axios";
export default async function getAllSupplier(page: number) {
  const res = await fetch(`http://localhost:8080/v1/suppliers?page=${page}`, {
    headers: {
      accept: "application/json",
      Authorization: apiKey,
    },
  });

  if (!res.ok) {
    throw new Error("Failed to fetch data");
  }
  return res.json().then((json) => {
    return {
      paging: json.paging,
      data: json.data,
    };
  });
}
