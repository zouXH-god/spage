"use client";

import { useParams } from "next/navigation";

export default function RepoPage() {
  const params = useParams();
  // params.owner 和 params.repo 都是 string
  return (
    <div>
      <h1>Owner: {params.owner}</h1>
      <h2>Project: {params.project}</h2>
      <h3>Site: {params.site}</h3>
      {/* 这里可以添加更多的内容或组件来展示项目详情 */}
    </div>
  );
}
