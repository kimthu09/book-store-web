"use client";
import { getUser } from "@/lib/auth/action";
import getShopGeneral from "@/lib/shop-general/getShopGeneral";
import { ShopGeneral } from "@/types";
import { useEffect, useState } from "react";

export const useShop = () => {
  const [shop, setShop] = useState<ShopGeneral | undefined>(undefined);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const user = await getShopGeneral();
        setShop(user.data);
      } catch (error) {
        console.error("Error fetching user data:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchUser();
  }, []);

  return { shop, loading };
};
