"use client";

import React from "react";
import { Outlet } from "react-router-dom";

import Nav from "@/components/nav/Nav";

export default function MainLayout() {
  return (
    <div className="flex flex-col h-screen text-slate-800 dark:text-slate-200">
      <div className="flex-shrink-0">
        <Nav />
      </div>
      <main className="flex-1 relative">
        <div className="absolute inset-0 overflow-y-auto overflow-x-hidden scrollbar-thin scrollbar-thumb-gray-300 scrollbar-track-transparent dark:scrollbar-thumb-gray-600">
          <Outlet />
        </div>
      </main>
    </div>
  );
}
