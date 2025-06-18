import Nav from "@/components/Nav";
import React from "react";

export default function MainLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className=" text-slate-800 dark:text-slate-200">
      <Nav />
      <div className="flex-1 overflow-hidden">
        {children}
      </div>
    </div>
  );
}