import type { AxiosResponse } from "axios";

import client from "./client";
import type { GetOwnerByNameResponse } from "./owner.model";

export function getOwnerByName(data: {
  name: string;
}): Promise<AxiosResponse<GetOwnerByNameResponse, unknown>> {
  return client.get<GetOwnerByNameResponse>(`/owner/${data.name}`);
}
