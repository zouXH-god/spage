import client from "./client";
import type { BaseResponse } from "./base.models";
import type {
    CaptchaConfig,
    LoginRequest,
    LoginResponse,
    RegisterRequest,
    RegisterResponse,
    User,
    UserOrganizationsResponse,
    UserProjectsResponse,
    UserResponse
} from './user.models';
import type { AxiosResponse } from "axios";

export function login(data: LoginRequest): Promise<AxiosResponse<LoginResponse, unknown>> {
    console.log("Logging in with data:", data.captchaToken);
    return client.post<LoginResponse>("/user/login", data);
}

export function logout(): Promise<AxiosResponse<BaseResponse, unknown>> {
    return client.post("/user/logout");
}

export function register(data: RegisterRequest): Promise<AxiosResponse<RegisterResponse, unknown>> {
    return client.post<RegisterResponse>("/user/register", data);
}

export function getUser(userId: number | null = null): Promise<AxiosResponse<UserResponse, unknown>> {
    return client.get<UserResponse>(`/user/${userId}`);
}

export function updateUser(data: User): Promise<AxiosResponse<BaseResponse, unknown>> {
    return client.put<BaseResponse>("/user", data);
}

export function getUserProjects(userId: number): Promise<AxiosResponse<UserProjectsResponse, unknown>> {
    return client.get<UserProjectsResponse>(`/user/${userId}/projects`);
}

export function getUserOrganizations(userId: number): Promise<AxiosResponse<UserOrganizationsResponse, unknown>> {
    return client.get<UserOrganizationsResponse>(`/user/${userId}/organizations`);
}

export function getCaptchaConfig(): Promise<AxiosResponse<CaptchaConfig, unknown>> {
    return client.get<CaptchaConfig>("/user/captcha");
}