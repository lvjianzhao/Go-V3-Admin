import {request} from "@/utils/service"

export interface GetRecordData {
  CreatedAt: string
  ip: string
  method: string
  path: string
  status: number
  reqParam: string
  respData: string
  respTime: number
  userName: string
}

export interface RowMeta extends GetRecordData {
  /** vxe-table 自动添加上去的属性 */
  _VXE_ID?: string
}

export type GetRecordResponseData = ApiResponseData<{
  list: GetRecordData[]
  total: number
}>


/** 查 */
export function getRecordApi(params: any) {
  return request<GetRecordResponseData>({
    url: "/record",
    method: "get",
    params
  })
}

export function getSearchOptionsApi() {
  return request<GetRecordResponseData>({
    url: "/record/options",
    method: "get",
  })
}
