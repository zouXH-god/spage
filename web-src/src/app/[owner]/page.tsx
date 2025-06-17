"use client";
import { EntityTypeEnum } from "@/types/entity";
import { useParams } from "next/navigation";

interface OwnerInfo {
    type: EntityTypeEnum.USER | EntityTypeEnum.ORG;
    id: number;
}

export default function RepoPage() {
  const params = useParams();
  const { owner } = params;

  return (
    <div>
      <h1>Owner: {owner}</h1>
    </div>
  );
}