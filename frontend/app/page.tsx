import { getUser } from "@/lib/auth/action";
import { withAuth } from "@/lib/role/withAuth";
import { Sale } from "./sale/page";

// export default async function Home() {
//   const test = await getUser()
//   console.log(test)
//   return <main className="flex">MainPage</main>;
// }
export default withAuth(Sale, ["INVOICE_CREATE"]);
