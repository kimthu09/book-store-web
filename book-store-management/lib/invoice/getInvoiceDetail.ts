import { apiKey, endPoint } from "@/constants";

export default async function getInvoiceDetail({
  idInvoice,
}: {
  idInvoice: string;
}) {
  const res = await fetch(`${endPoint}/v1/invoices/${idInvoice}`, {
    headers: {
      accept: "application/json",
      Authorization: apiKey,
    },
  });
  if (!res.ok) {
    // throw new Error("Failed to fetch data");
    return res.json();
  }
  return res.json().then((json) => json.data);
}
