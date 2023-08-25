import { request } from "@/utils/service"

export interface ApiDataBase {
  path: string
  api_group: string
  method: string
  description: string
}

export interface ApiData extends ApiDataBase {
  ID: number
}

export interface ApiDataPageInfo {
  list: ApiData[]
  total: number
  page: number
  pageSize: number
}

interface reqApis extends PageInfo {
  path?: string
  api_group?: string
  method?: string
  description?: string
  orderKey?: string
  desc?: boolean
}

// 获取所有api 分页
export function getApisApi(params: reqApis) {
  return request<ApiResponseData<ApiDataPageInfo>>({
    url: "/api",
    method: "get",
    params
  })
}

// 获取所有API分组
export function getApiGroups() {
  return request<ApiResponseData<ApiDataPageInfo>>({
    url: "/api/apiGroups",
    method: "get"
  })
}

interface children {
  key: string
  api_group: string
  path: string
  method: string
  description: string
}

export interface ApiTreeData {
  api_group: string
  children: children[]
}

interface ApiTreeAll {
  list: ApiTreeData[]
  checkedKey: string[]
}

// 获取所有api 不分页
export function getElTreeApisApi(params: reqId) {
  return request<ApiResponseData<ApiTreeAll>>({
    url: "/api/getElTreeApis",
    method: "get",
    params
  })
}

// 添加api
export function addApiApi(data: ApiDataBase) {
  return request<ApiResponseData<ApiData>>({
    url: "/api",
    method: "post",
    data
  })
}

// 删除api
export function deleteApiApi(id: number) {
  return request<ApiResponseData<null>>({
    url: `/api/${id}`,
    method: "delete"
  })
}

interface reqEdit extends ApiDataBase {
  id: number
}

// 编辑api
export function editApiApi(data: reqEdit) {
  return request<ApiResponseData<null>>({
    url: "/api",
    method: "put",
    data
  })
}
