import { StatusNote } from "@/types";
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


export const statusNoteToString = (status: StatusNote) => {
  if (status === StatusNote.Inprogress) {
    return "Đang tiến hành";
  } else if (status === StatusNote.Done) {
    return "Đã hoàn thành";
  } else {
    return "Đã hủy";
  }
};
export const isAdmin = ({
  currentUser,
}: {
  currentUser:
    | {
        name?: string | null | undefined;
        email?: string | null | undefined;
        image?: string | null | undefined;
      }
    | undefined;
}) => {
  try {
    if (currentUser) {
      const json = JSON.stringify(currentUser);
      const user = JSON.parse(json);
      const roleId = user.data.role.id;
      if (roleId === "admin") {
        return true;
      } else {
        return false;
      }
    }
    return false;
  } catch (error) {
    throw new Error("Có lỗi xảy ra");
  }
};

export function toLocalTime(utcTime: Date | string): string {
  // Check if utcTime is a string, and try to parse it as a Date
  const parsedUtcTime = typeof utcTime === 'string' ? new Date(utcTime) : utcTime;

  // Check if parsing was successful and utcTime is a valid Date object
  if (!(parsedUtcTime instanceof Date) || isNaN(parsedUtcTime.getTime())) {
    // Handle the case where utcTime is not a valid Date
    console.error('Invalid UTC time value:', utcTime);
    return 'Invalid Time';
  }

  // Get local date and time in string format
  const options: Intl.DateTimeFormatOptions = {
    day: '2-digit',
    month: '2-digit',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
    hour12: false, // Use 24-hour format
    timeZone: 'Asia/Ho_Chi_Minh', // Set the timezone to Vietnam
  };

  // Format the date and time using Intl.DateTimeFormat
  const localTime = new Intl.DateTimeFormat('vi-VN', options).format(parsedUtcTime);

  return localTime;
}

export function toDateString(utcTime: Date | string): string {
  // Check if utcTime is a string, and try to parse it as a Date
  const parsedUtcTime = typeof utcTime === 'string' ? new Date(utcTime) : utcTime;

  // Check if parsing was successful and utcTime is a valid Date object
  if (!(parsedUtcTime instanceof Date) || isNaN(parsedUtcTime.getTime())) {
    // Handle the case where utcTime is not a valid Date
    console.error('Invalid UTC time value:', utcTime);
    return 'Invalid Time';
  }

  // Get local date and time in string format
  const options: Intl.DateTimeFormatOptions = {
    day: '2-digit',
    month: '2-digit',
    year: 'numeric',
    hour: undefined,
    minute: undefined,
    hour12: false, // Use 24-hour format
    timeZone: 'Asia/Ho_Chi_Minh', // Set the timezone to Vietnam
  };

  // Format the date and time using Intl.DateTimeFormat
  const localTime = new Intl.DateTimeFormat('vi-VN', options).format(parsedUtcTime);

  return localTime;
}

export function toLocalDateTime(utcTime: Date | string): Date {
  return new Date(utcTime);
}