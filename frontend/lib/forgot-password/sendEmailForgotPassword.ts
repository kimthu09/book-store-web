import { endPoint } from "@/constants";
import { getApiKey } from "../auth/action";
import axios from "axios";

export default async function sendEmailForgotPassword({
  email
}: {
  email: string;
}) {
  const url = `${endPoint}/v1/forgetPassword`;
  const data = {
    email: email
  };

  const token = await getApiKey();

  const headers = {
    accept: "application/json",
    "Content-Type": "application/json",
    Authorization: `Bearer ${token}`,
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
