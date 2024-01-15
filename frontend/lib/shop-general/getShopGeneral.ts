import { endPoint } from "@/constants";
import { getApiKey } from "../auth/action";

export default async function getShopGeneral() {
  const url = `${endPoint}/v1/shop`;
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
    throw new Error("Có lỗi xảy ra");
    return res.json();
  }
  return res.json().then((json) => {
    return {
      data: json.data,
    };
  });
}
