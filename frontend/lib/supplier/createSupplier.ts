import { endPoint } from "@/constants";
import axios from "axios";
import { getApiKey } from "../auth/action";

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
  debt: number;
}) {
  const url = `${endPoint}/v1/suppliers`;
  const data = {
    email: email,
    id: id,
    name: name,
    phone: phone,
    debt: +debt,
  };

  const token = await getApiKey();
  const headers = {
    "Content-Type": "application/json",
    Authorization: `Bearer ${token}`,
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
