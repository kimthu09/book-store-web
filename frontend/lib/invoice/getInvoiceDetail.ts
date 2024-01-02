import { endPoint } from "@/constants";
import { getApiKey } from "../auth/action";

export default async function getInvoiceDetail({
  idInvoice,
}: {
  idInvoice: string;
}) {
  const token = await getApiKey();
  const res = await fetch(`${endPoint}/v1/invoices/${idInvoice}`, {
    headers: {
      accept: "application/json",
      Authorization: `Bearer ${token}`,
    },
  });
  if (!res.ok) {
    // throw new Error("Failed to fetch data");
    return res.json();
  }
  return res.json().then((json) => json.data);
}
