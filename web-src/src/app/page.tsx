"use client";

import dynamic from 'next/dynamic';

// 动态导入应用组件，禁用 SSR
const App = dynamic(() => import('@/components/App'), { ssr: false });

export default function HomePage() {
  return <App />;
}