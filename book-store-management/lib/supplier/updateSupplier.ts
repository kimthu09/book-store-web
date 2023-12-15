import { apiKey } from "@/constants";
import axios from "axios";

export default async function updateSupplier({
  idSupplier,
  name,
  email,
  phone,
}: {
  idSupplier: string;
  name: string;
  email: string;
  phone: string;
}) {
  const url = `http://localhost:8080/v1/suppliers/${idSupplier}`;
  const data = {
    email: email,
    name: name,
    phone: phone,
  };
  console.log(data);
  const headers = {
    "Content-Type": "application/json",
    Authorization: apiKey,
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
