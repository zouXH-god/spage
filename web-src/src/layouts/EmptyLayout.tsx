"use client";

import React from "react";
import { Outlet } from "react-router-dom";

export default function MainLayout() {
  return (
    <div className="flex flex-col h-screen text-slate-800 dark:text-slate-200">
      <Outlet />
    </div>
  );
}
