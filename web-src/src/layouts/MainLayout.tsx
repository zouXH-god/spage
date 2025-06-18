"use client";
import Nav from "@/components/Nav";
import React from "react";

export default function MainLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="flex flex-col h-screen text-slate-800 dark:text-slate-200">
      <div className="flex-shrink-0">
        <Nav />
      </div>
      <main className="flex-1 relative">
        <div className="absolute inset-0 overflow-y-auto overflow-x-hidden scrollbar-thin scrollbar-thumb-gray-300 scrollbar-track-transparent dark:scrollbar-thumb-gray-600">
          {children}
        </div>
      </main>
    </div>
  );
}