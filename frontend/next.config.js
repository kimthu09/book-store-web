/** @type {import('next').NextConfig} */
const nextConfig = {
  images: {
    remotePatterns: [
      {
        protocol: "https",
        hostname: "*",
        pathname: "**",
      },
    ],
    domains: ["localhost"],
  },
};

module.exports = nextConfig;
