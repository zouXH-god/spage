import type { AxiosResponse } from "axios";

import type { BaseResponse } from "./base.models";
import client from "./client";
import type { Project } from "./project.models";

export function createProject(data: Project): Promise<AxiosResponse<BaseResponse, unknown>> {
  return client.post<BaseResponse>("/project", data);
}

export function deleteProject(projectId: number): Promise<AxiosResponse<BaseResponse, unknown>> {
  return client.delete<BaseResponse>(`/project`, { data: { id: projectId } });
}

export function updateProject(data: Project): Promise<AxiosResponse<BaseResponse, unknown>> {
  return client.put<BaseResponse>("/project", data);
}

export function getProject(projectId: number): Promise<AxiosResponse<Project, unknown>> {
  return client.get<Project>(`/project/${projectId}`);
}
