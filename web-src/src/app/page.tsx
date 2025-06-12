"use client";
import Nav from "@/components/Nav";
import "@/utils/i18n";

export default function Home() {
  console.log(process.env.NEXT_PUBLIC_API_BASE_URL)
  return (
    <div className="grid grid-rows-[auto_1fr_20px] ...">
      <Nav />
      {/* 其他内容 */}
    </div>
  );
}
