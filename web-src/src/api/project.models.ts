/**
 * 项目 Project
 */
export interface Project {
  description: string;
  displayName?: string;
  id: number;
  name: string;
  ownerId: number;
  ownerType: "org" | "user";
  owners: number[];
  siteLimit: number;
  [property: string]: unknown;
}
