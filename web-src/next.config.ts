const isDev = process.env.NODE_ENV === "development";

const nextConfig = isDev ?
  {
    async rewrites() {
      const backendUrl = (process.env.NEXT_PUBLIC_API_BASE_URL || "http://localhost:8888")
      console.log("Using development API base URL:", backendUrl);
      return [
        {
          source: '/api/:path*',
          destination: backendUrl + '/api/:path*',
        },
      ]
    },
  } :
  {
    output: "export",
    images: {
      unoptimized: true,
    },
  }

module.exports = nextConfig