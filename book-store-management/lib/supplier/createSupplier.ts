import { apiKey, endPoint } from "@/constants";
import axios from "axios";

export default async function createSupplier({
  id,
  name,
  email,
  phone,
  debt,
}: {
  id: string;
  name: string;
  email: string;
  phone: string;
  debt: string;
}) {
  const url = `${endPoint}/v1/suppliers`;
  const data = {
    email: email,
    id: id,
    name: name,
    phone: phone,
    debt: +debt,
  };
  console.log(data);
  const headers = {
    "Content-Type": "application/json",
    Authorization: apiKey,
    accept: "application/json",
  };

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
