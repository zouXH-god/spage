import type { AxiosResponse } from "axios";

import type { BaseResponse } from "./base.models";
import client from "./client";
import type { CreateSiteReleaseResponse, Site, SiteRelease } from "./site.models";

export function createSite(data: Site): Promise<AxiosResponse<BaseResponse, unknown>> {
  return client.post<BaseResponse>("/site", data);
}

export function deleteSite(siteId: number): Promise<AxiosResponse<BaseResponse, unknown>> {
  return client.delete<BaseResponse>(`/site`, { data: { id: siteId } });
}

export function getSite(siteId: number): Promise<AxiosResponse<Site, unknown>> {
  return client.get<Site>(`/site/${siteId}`);
}

export function updateSite(data: Site): Promise<AxiosResponse<BaseResponse, unknown>> {
  return client.put<BaseResponse>("/site", data);
}

export function createSiteRelease(
  data: SiteRelease,
): Promise<AxiosResponse<CreateSiteReleaseResponse, unknown>> {
  return client.post<CreateSiteReleaseResponse>("/site/release", data);
}

export function deleteSiteRelease(
  releaseId: number,
): Promise<AxiosResponse<BaseResponse, unknown>> {
  return client.delete<BaseResponse>(`/site/release`, { data: { id: releaseId } });
}

// TODO: 后端未完成
export function getSiteReleases(
  releaseId: number,
): Promise<AxiosResponse<{ id: number; siteId: number; fileId: number; tag: string }, unknown>> {
  return client.get<{ id: number; siteId: number; fileId: number; tag: string }>(
    `/site/release/${releaseId}`,
  );
}
