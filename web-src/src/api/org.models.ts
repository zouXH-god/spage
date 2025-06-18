import type { BaseResponse } from "./base.models";
import type { Project } from "./project.models";

/**
 * 组织 Org
 */
export interface Org {
  avatarUrl?: string;
  creatorId: number;
  description?: string;
  displayName?: string;
  email?: string;
  id: number;
  members: number[];
  name: string;
  oidcMembersGroups: string[];
  owners: number[];
  projectLimit: number;
  [property: string]: unknown;
}

export interface OrgProjectsResponse extends BaseResponse {
  projects: Project[];
}
