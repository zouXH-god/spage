import type { AxiosResponse } from "axios";

import client from "./client";
import type { GetMetaInfoResponse } from "./meta.models";

export function getMetaInfo(): Promise<AxiosResponse<GetMetaInfoResponse, unknown>> {
    return client.get<GetMetaInfoResponse>("/meta/info");
}
