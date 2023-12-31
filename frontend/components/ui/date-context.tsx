import { useEffect, useState } from "react";

export const CurrentDate = () => {
  const locale = "vi-VN";
  const [today, setDate] = useState(new Date()); // Save the current date to be able to trigger an update

  useEffect(() => {
    const timer = setInterval(() => {
      // Creates an interval which will update the current data every minute
      // This will trigger a rerender every component that uses the useDate hook.
      setDate(new Date());
    }, 1000);
    return () => {
      clearInterval(timer); // Return a funtion to clear the timer so that it will stop being called on unmount
    };
  }, []);

  const day = today.toLocaleDateString(locale, { weekday: "long" });
  const date = `${day}, ${today.toLocaleDateString(locale)}\n\n`;

  const hour = today.getHours();

  const time = today.toLocaleTimeString(locale, {
    hour: "2-digit",
    minute: "2-digit",
  });

  return (
    <span className="p-2 text-sm text-muted-foreground">
      {time} {date}
    </span>
  );
};
