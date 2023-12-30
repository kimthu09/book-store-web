import { getUser } from "@/lib/auth/action";

export default async function Home() {
  const test = await getUser()
  console.log(test)
  return <main className="flex">MainPage</main>;
}
