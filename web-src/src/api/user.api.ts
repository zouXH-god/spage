import type { AxiosResponse } from "axios";

import { PAGE_LIMIT } from "@/consts";

import type { BaseResponse } from "./base.models";
import client from "./client";
import type {
  CaptchaConfig,
  LoginRequest,
  LoginResponse,
  OidcConfig,
  RegisterRequest,
  RegisterResponse,
  User,
  UserOrganizationsResponse,
  UserProjectsResponse,
  UserResponse,
} from "./user.models";

export function login(data: LoginRequest): Promise<AxiosResponse<LoginResponse, unknown>> {
  return client.post<LoginResponse>("/user/login", data);
}

export function logout(): Promise<AxiosResponse<BaseResponse, unknown>> {
  return client.post("/user/logout");
}

export function register(data: RegisterRequest): Promise<AxiosResponse<RegisterResponse, unknown>> {
  return client.post<RegisterResponse>("/user/register", data);
}

export function getUser(
  userId: number | null = null,
): Promise<AxiosResponse<UserResponse, unknown>> {
  return client.get<UserResponse>(userId ? `/user-info/${userId}` : "/user-info");
}

export function updateUser(data: User): Promise<AxiosResponse<BaseResponse, unknown>> {
  return client.put<BaseResponse>("/user-info", data);
}

export function getUserProjects(
  userId: number,
  page: number = 1,
  limit: number = PAGE_LIMIT,
): Promise<AxiosResponse<UserProjectsResponse, unknown>> {
  return client.get<UserProjectsResponse>(`/user-info/${userId}/projects?page=${page}&limit=${limit}`);
}

export function getUserOrganizations(
  userId: number,
): Promise<AxiosResponse<UserOrganizationsResponse, unknown>> {
  return client.get<UserOrganizationsResponse>(`/user-info/${userId}/orgs`);
}

export function getCaptchaConfig(): Promise<AxiosResponse<CaptchaConfig, unknown>> {
  return client.get<CaptchaConfig>("/user/captcha");
}

export function getOidcConfig(): Promise<AxiosResponse<{ oidcConfigs: OidcConfig[] }, unknown>> {
  return client.get<{ oidcConfigs: OidcConfig[] }>("/user/oidc/config");
}
