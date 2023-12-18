import { apiKey, endPoint } from "@/constants";

export default async function getAllSupplierNote({
  idSupplier,
}: {
  idSupplier: string;
}) {
  const res = await fetch(
    `${endPoint}/v1/suppliers/${idSupplier}/debts?limit=${1000}`,
    {
      headers: {
        accept: "application/json",
        Authorization: apiKey,
      },
      next: {
        revalidate: 0,
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
