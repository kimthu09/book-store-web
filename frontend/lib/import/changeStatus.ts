import { endPoint } from "@/constants";
import { StatusNote } from "@/types";
import axios from "axios";
import { getApiKey } from "../auth/action";

export default async function updateStatus({
  id: id,
  status,
}: {
  id: string;
  status: StatusNote;
}) {
  const url = `${endPoint}/v1/importNotes/${id}`;
  const data = {
    status: status,
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
