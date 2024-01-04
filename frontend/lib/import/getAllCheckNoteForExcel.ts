import { endPoint } from "@/constants";
import { getApiKey } from "../auth/action";

export default async function getAllCheckNoteForExcel({
  limit,
  page,
}: {
  limit?: number;
  page: string;
}) {
  const url = `${endPoint}/v1/inventoryCheckNotes?page=${page}&limit=${
    limit ?? "10"
  }`;

  const token = await getApiKey();
  const res = await fetch(url, {
    headers: {
      accept: "application/json",
      Authorization: `Bearer ${token}`,
    },
  });

  if (!res.ok) {
    // throw new Error("Failed to fetch data");
    console.error(res);
    return res.json();
  }
  return res.json().then((json) => {
    return {
      paging: json.paging,
      data: json.data,
    };
  });
}
