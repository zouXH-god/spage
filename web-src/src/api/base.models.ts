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
 * base response
 */
export interface BaseResponse {
    message: string | null;
    [property: string]: unknown;
}
