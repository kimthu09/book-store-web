"use client";

import { useEffect } from "react";
import Cookies from "js-cookie";
import { useRouter } from "next/navigation";

export default function Home() {
  const router = useRouter()

  useEffect(() => {
    const accessToken = Cookies.get('accessToken')
    // console.log(accessToken)
    if (typeof accessToken !== "string") {
      // console.log("ahihi" + accessToken)
      router.push('/login');
    }
  }, [])

  return <main className="flex">MainPage</main>;
}
