import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export const toVND = (money: number) => {
  const formatted = new Intl.NumberFormat("vi-VN", {
    style: "currency",
    currency: "VND",
  }).format(money);
  return formatted;
};

export const includesRoles = ({
  currentUser,
  allowedFeatures,
}: {
  currentUser:
    | {
        name?: string | null | undefined;
        email?: string | null | undefined;
        image?: string | null | undefined;
      }
    | undefined;
  allowedFeatures: string[];
}) => {
  try {
    const json = JSON.stringify(currentUser);
    const user = JSON.parse(json);
    const features = user.data.role.features.map((item: any) => item.featureId);
    if (
      currentUser &&
      allowedFeatures.every((item) => features.includes(item))
    ) {
      return true;
    } else {
      return false;
    }
  } catch (error) {
    throw new Error("Có lỗi xảy ra");
  }
};
