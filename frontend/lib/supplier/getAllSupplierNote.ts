import { endPoint } from "@/constants";
import { getApiKey } from "../auth/action";

export default async function getAllSupplierNote({
  idSupplier,
}: {
  idSupplier: string;
}) {
  const token = await getApiKey();
  const res = await fetch(
    `${endPoint}/v1/suppliers/${idSupplier}/importNotes?limit=${1000}`,
    {
      headers: {
        accept: "application/json",
        Authorization: `Bearer ${token}`,
      },
    }
  );
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
