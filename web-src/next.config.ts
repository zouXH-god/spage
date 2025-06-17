const isDev = process.env.NODE_ENV === "development";

const nextConfig = isDev ?
  {
    async rewrites() {
      const baseUrl = (process.env.NEXT_PUBLIC_API_BASE_URL || isDev ? "http://localhost:8888" : "")
      console.log("Using development API base URL:", baseUrl);
      return [
        {
          source: '/api/:path*',
          destination: baseUrl + '/api/:path*',
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