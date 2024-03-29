import { endPoint } from "@/constants";
import axios from "axios";
import { getApiKey } from "../auth/action";

export default async function updateCustomer({
  idCustomer,
  name,
  email,
  phone,
}: {
  idCustomer: string;
  name: string;
  email: string;
  phone: string;
}) {
  const url = `${endPoint}/v1/customers/${idCustomer}`;
  const data = {
    email: email,
    name: name,
    phone: phone,
  };

  const token = await getApiKey();
  const headers = {
    "Content-Type": "application/json",
    Authorization: `Bearer ${token}`,
    accept: "application/json",
  };

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
