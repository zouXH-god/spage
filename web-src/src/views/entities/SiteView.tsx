"use client";

import { useParams } from "next/navigation";

export default function SiteView() {
  const params = useParams();
  return (
    <div>
      <h1>Owner: {params.owner}</h1>
      <h2>Project: {params.project}</h2>
      <h3>Site: {params.site}</h3>
    </div>
  );
}
