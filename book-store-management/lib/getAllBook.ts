import { apiKey } from "@/constants";
export const dynamic = "force-dynamic";
export const revalidate = 0;

export default async function getAllBooks(page: number) {
  // Add a unique timestamp as a query parameter
  const uniqueParam = `cacheBuster=${new Date().getTime()}`;

  const res = await fetch(
    `http://localhost:8080/v1/booktitles?page=${page}&${uniqueParam}`,
    {
      headers: {
        accept: "application/json",
        Authorization: apiKey,
      },
      next: { revalidate: 0 },
    }
  );

  if (!res.ok) {
    throw new Error("Failed to fetch data");
  }

  return res.json().then((json) => {
    console.log(json);
    return {
      paging: json.paging,
      data: json.data,
    };
  });
}
