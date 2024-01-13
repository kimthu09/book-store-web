import { Metadata } from "next";
import DetailLayout from "./detail-layout";

export const metadata: Metadata = {
  title: "Thiết lập cửa hàng",
};

const SettingScreen = () => {
  return <DetailLayout />;
};

export default SettingScreen;
