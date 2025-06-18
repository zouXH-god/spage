"use client";
import React from "react";

export default function MainLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="flex flex-col h-screen text-slate-800 dark:text-slate-200">
          {children}
    </div>
  );
}