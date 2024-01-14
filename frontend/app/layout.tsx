import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import Sidebar from "@/components/sidebar";
import Header from "@/components/header";
import HeaderMobile from "@/components/header-mobile";
import { Toaster } from "@/components/ui/toaster";
import { auth } from "@/lib/auth/auth";
import { LoadingProvider } from "@/hooks/loading-context";

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

export default async function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const isAuthented = await auth().then((res) => res?.user);
  return (
    <html lang="en" className="h-full">
      <body className={`${inter.className} flex overflow-y-hidden h-full`}>
        <LoadingProvider>
          {isAuthented ? (
            <>
              <Sidebar />
              <main className="flex flex-1">
                <div className="flex w-full flex-col overflow-y-hidden">
                  <Header />
                  <HeaderMobile />
                  <div className="md:p-10 p-4 overflow-auto">{children}</div>
                  <Toaster />
                </div>
              </main>
            </>
          ) : (
            <main className="flex flex-1">
              {children}
              <Toaster />
            </main>
          )}
        </LoadingProvider>
      </body>
    </html>
  );
}
