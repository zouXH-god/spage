import { BaseResponse } from "./base.models";

/**
 * 站点 Site
 */
export interface Site {
  description: string;
  domains: string[];
  id: number;
  name: string;
  projectId: number;
  subDomain: string;
  [property: string]: unknown;
}

/**
 * 站点发布 SiteRelease
 */
export interface SiteRelease {
  creatorId: number;
  fileId: number;
  id: number;
  siteId: number;
  tag: string;
  url: string;
  [property: string]: unknown;
}

export interface CreateSiteReleaseResponse extends BaseResponse {
  id: number;
}
