import { CharComponent } from "@/types";
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  ScriptableContext,
  Filler,
} from "chart.js";
import { Line } from "react-chartjs-2";
import { toDateString, toLocalDateTime } from "@/lib/utils";

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  Filler
);

export const options = {
  responsive: true,
  resizeDelay: 0,
  maintainAspectRatio: false,
  plugins: {
    legend: {
      position: "top" as const,
    },
    title: {
      display: false,
    },
  },
};

const DashboardChart = (props: any) => {
  const { type, price, profit, timeFrom, timeTo } = props;

  let labels: string[] = [];
  let priceAmount: number[] = [];
  let costAmount: number[] = [];

  if (type == "day") {
    labels = [
      "0h",
      "1h",
      "2h",
      "3h",
      "4h",
      "5h",
      "6h",
      "7h",
      "8h",
      "9h",
      "10h",
      "11h",
      "12h",
      "13h",
      "14h",
      "15h",
      "16h",
      "17h",
      "18h",
      "19h",
      "20h",
      "21h",
      "22h",
      "23h",
    ];

    priceAmount = groupAndSumByHour(price);
    costAmount = groupAndSumByHour(profit);
  } else if (type == "month") {
    for (
      let currentDate = toLocalDateTime(timeFrom)!;
      currentDate <= toLocalDateTime(timeTo)!;
      currentDate.setDate(currentDate.getDate() + 1)
    ) {
      let dayString = toDateString(currentDate);
      labels.push(dayString);
    }

    priceAmount = groupAndSumByDate(price, labels);
    costAmount = groupAndSumByDate(profit, labels);
  } else if (type == "month-day") {
    labels = ["CN", "T2", "T3", "T4", "T5", "T6", "T7"];
    priceAmount = groupAndSumByDayOfWeek(price);
    costAmount = groupAndSumByDayOfWeek(profit);
  }

  const data = () => {
    return {
      labels,
      datasets: [
        {
          label: "Doanh thu bán được",
          data: priceAmount,
          fill: "start",
          borderColor: "rgb(53, 162, 235)",
          backgroundColor: (context: ScriptableContext<"line">) => {
            const ctx = context.chart.ctx;
            const gradient = ctx.createLinearGradient(0, 0, 0, 450);
            gradient.addColorStop(0, "rgba(53, 162, 235, 1)");
            gradient.addColorStop(1, "rgba(53, 162, 235, 0)");
            return gradient;
          },
          tension: 0.3,
        },
        {
          label: "Số tiền lời",
          data: costAmount,
          fill: "start",
          borderColor: "rgb(255, 99, 132)",
          backgroundColor: (context: ScriptableContext<"line">) => {
            const ctx = context.chart.ctx;
            const gradient = ctx.createLinearGradient(0, 0, 0, 450);
            gradient.addColorStop(0, "rgba(255, 99, 132, 1)");
            gradient.addColorStop(1, "rgba(255, 99, 132, 0)");
            return gradient;
          },
          tension: 0.3,
        },
      ],
    };
  };

  return <Line options={options} data={data()} className="aspect-square" />;
};

function groupAndSumByHour(charComponents: CharComponent[]): number[] {
  const groupedByHour = new Map<string, number>();

  // Initialize all hours with a value of 0
  for (let hour = 0; hour <= 23; hour++) {
    const key = `${hour}h`;
    groupedByHour.set(key, 0);
  }

  // Update values based on actual CharComponent objects
  charComponents.forEach((component) => {
    const dt = toLocalDateTime(component.time);
    const roundedHour = dt!.getHours();
    const key = `${roundedHour}h`;

    groupedByHour.set(key, groupedByHour.get(key)! + component.value);
  });

  // Convert the Map to an array of values sorted by hour
  return Array.from(groupedByHour.values());
}

function groupAndSumByDate(
  charComponents: CharComponent[],
  labels: string[]
): number[] {
  const groupedByDate = new Map<string, number>();

  charComponents.forEach((component) => {
    const dt = toDateString(component.time);

    if (groupedByDate.has(dt)) {
      groupedByDate.set(dt, groupedByDate.get(dt)! + component.value);
    } else {
      groupedByDate.set(dt, component.value);
    }
  });

  // Create an array of values sorted by date
  const sortedValues = Array.from(groupedByDate.entries())
    .sort(([date1], [date2]) => date1.localeCompare(date2))
    .map(([date, totalValue]) => totalValue);

  // If you want to ensure there are values for all dates within the range, you can fill the gaps
  const startDate = new Date(
    Math.min(
      ...charComponents.map((component) =>
        toLocalDateTime(component.time)!.getTime()
      )
    )
  );
  const endDate = new Date(
    Math.max(
      ...charComponents.map((component) =>
        toLocalDateTime(component.time)!.getTime()
      )
    )
  );

  const allDates: number[] = [];
  labels.forEach((label) => {
    allDates.push(groupedByDate.get(label) ?? 0);
  });

  return allDates;
}

function groupAndSumByDayOfWeek(charComponents: CharComponent[]): number[] {
  const groupedByDayOfWeek = new Map<number, number>();

  charComponents.forEach((component) => {
    const dt = toLocalDateTime(component.time);
    const dayOfWeek = dt!.getDay();

    if (groupedByDayOfWeek.has(dayOfWeek)) {
      groupedByDayOfWeek.set(
        dayOfWeek,
        groupedByDayOfWeek.get(dayOfWeek)! + component.value
      );
    } else {
      groupedByDayOfWeek.set(dayOfWeek, component.value);
    }
  });

  // Create an array of values sorted by day of the week (0: Sunday, 1: Monday, ..., 6: Saturday)
  const sortedValues = Array.from(groupedByDayOfWeek.entries())
    .sort(([day1], [day2]) => day1 - day2)
    .map(([day, totalValue]) => totalValue);

  // If you want to ensure there are values for all days of the week, you can fill the gaps
  const result = Array.from(
    { length: 7 },
    (_, index) => groupedByDayOfWeek.get(index) || 0
  );

  return result;
}

export default DashboardChart;
