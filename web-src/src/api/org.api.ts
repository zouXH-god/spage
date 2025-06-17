import client from "./client";
import type { BaseResponse } from "./base.models";
import type { Org } from "./org.models";
import type { AxiosResponse } from "axios";


export function createOrg(data: Org): Promise<AxiosResponse<BaseResponse, unknown>> {
    return client.post<BaseResponse>('/org', data);
}

export function deleteOrg(orgId: number): Promise<AxiosResponse<BaseResponse, unknown>> {
    return client.delete<BaseResponse>(`/org`, { data: { id: orgId } });
}

export function updateOrg(data: Org): Promise<AxiosResponse<BaseResponse, unknown>> {
    return client.put<BaseResponse>('/org', data);
}

export function getOrg(orgId: number): Promise<AxiosResponse<Org, unknown>> {
    return client.get<Org>(`/org/${orgId}`);
}

export function getOrgProjects(orgId: number): Promise<AxiosResponse<BaseResponse, unknown>> {
    return client.get<BaseResponse>(`/org/${orgId}/projects`);
}