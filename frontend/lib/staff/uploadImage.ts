import { endPoint } from "@/constants";
import axios from "axios";
import { getApiKey } from "../auth/action";

export const imageUpload = async (formData: FormData) => {
  const token = await getApiKey();

  const res = axios
    .post(`${endPoint}/v1/upload`, formData, {
      headers: {
        "Content-Type": "multipart/form-data",
        accept: "application/json",
        Authorization: `Bearer ${token}`,
      },
    })
    .then((response) => {
      if (response) return response.data;
    })
    .catch((error) => {
      console.error("Error:", error);
      return error.response.data;
    });
  return res;
};
