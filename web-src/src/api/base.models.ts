/**
 * @file API 数据模型定义
 * @module api/models
 */

/**
 * 文件 File
 */
export interface File {
    creatorId: number;
    id: number;
    isPrivate: boolean;
    [property: string]: unknown;
}

/**
 * 标签 Label
 */
export interface Label {
    color?: string;
    name: string;
    ownerId: number;
    ownerType: "org" | "user";
    value?: string;
    [property: string]: unknown;
}


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
    [property: string]: unknown;
}

/**
 * base response
 */
export interface BaseResponse {
    message: string | null;
    [property: string]: unknown;
}
