"use client";

import { getOwnerByName } from "@/api/owner.api";
import { OrgPage } from "@/components/owner/OrgPage";
import { UserPage } from "@/components/owner/UserPage";
import { EntityTypeEnum } from "@/types/entity";
import { useParams } from "next/navigation";
import { useEffect, useState } from "react";

interface OwnerProps {
  type: string;
  id: number;
}

export default function OwnerPage() {
  const [ownerProps, setOwnerProps] = useState<OwnerProps | null>(null);
  const params = useParams();

  useEffect(() => {
    getOwnerByName({ name: params.owner as string }).then((owner) => {
      if (owner) {
        setOwnerProps({
          type: owner.data.type,
          id: owner.data.id
        });
      } else {
        setOwnerProps(null);
      }
    }
    ).catch((error) => {
      console.error("Error fetching owner:", error);
      setOwnerProps(null);
    }
    );
  }, [params.owner]);

  return (
    <div>
      {/* 这里可以根据 ownerProps.type 渲染不同的内容 */}
      {ownerProps ? (
        ownerProps.type === EntityTypeEnum.USER ? (
          <UserPage id={ownerProps.id} />
        ) : (
          <OrgPage id={ownerProps.id} />
        )
      ) : (
        <div className="text-center text-slate-700 dark:text-slate-100">加载中</div>
      )}
    </div>
  );
}
