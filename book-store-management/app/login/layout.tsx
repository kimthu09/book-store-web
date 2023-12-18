import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "../globals.css";
import Sidebar from "@/components/sidebar";
import Header from "@/components/header";
import HeaderMobile from "@/components/header-mobile";
import { Toaster } from "@/components/ui/toaster";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
    title: "Nhà sách",
    description: "Book store management app",
    icons: {
        icon: ["/favicon.ico?v=4"],
        apple: ["/apple-touch-icon.png?v=4"],
        shortcut: ["apple-touch-icon.png"],
    },
    manifest: "/site.webmanifest",
};

export default function RootLayout({
    children,
}: {
    children: React.ReactNode;
}) {
    return (
        <html lang="en">
            <body>
                <div className="flex w-full flex-col">{children}</div>
            </body>
        </html>
    );
}
export const revalidate = 0;
