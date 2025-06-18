import { EntityTypeEnum } from "@/types/entity";

export interface GetOwnerByNameResponse {
  id: number;
  type: EntityTypeEnum.USER | EntityTypeEnum.ORG;
}
