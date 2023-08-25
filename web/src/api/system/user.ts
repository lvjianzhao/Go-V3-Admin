import { request } from "@/utils/service"

// type UserInfoResponseData = ApiResponseData<{ username: string; roles: string[] }>

/** 获取用户详情 */
export function getUserInfoApi() {
  return request<ApiResponseData<UsersResponse>>({
    url: "/user/detail",
    method: "get"
  })
}

export interface UsersResponse {
  createdAt: string
  ID: number
  username: string
  phone: string
  email: string
  active: boolean
  roleId: number
  role: string
}

export interface UsersResponsePageInfo {
  list: UsersResponse[]
  total: number
  page: number
  pageSize: number
}

export interface UsersRequestData {
  /** 当前页码 */
  currentPage?: number
  /** 查询条数 */
  size?: number
  /** 查询参数：用户名 */
  name?: string
}

type UsersResponseData = ApiResponseData<UsersResponsePageInfo>

/** 获取所有用户 */
export function getUsersApi(params: UsersRequestData) {
  return request<UsersResponseData>({
    url: "/user",
    method: "get",
    params
  })
}

// 删除用户
export function deleteUserApi(id: number) {
  return request<ApiResponseData<null>>({
    url: `/user/${id}`,
    method: "delete"
  })
}

export interface reqUser {
  username: string
  password: string
  phone: string
  email: string
  active: boolean
  roleId: number
}

// 添加用户
export function addUserApi(data: reqUser) {
  return request<ApiResponseData<null>>({
    url: "/user",
    method: "post",
    data
  })
}

interface reqEditUser {
  id: number
  username: string
  phone: string
  email: string
  active: boolean
  roleId: number
}

// 编辑用户
export function editUserApi(data: reqEditUser) {
  return request<ApiResponseData<UsersResponse>>({
    url: "/user",
    method: "put",
    data
  })
}

// 修改用户密码
interface reqModifyPass {
  id: number
  oldPassword: string
  newPassword: string
}

export function modifyPassApi(data: reqModifyPass) {
  return request<ApiResponseData<null>>({
    url: "/user/password",
    method: "put",
    data
  })
}

// 切换用户状态
interface reqSwitchActive {
  id: number
  active: boolean
}

export function SwitchActiveApi(data: reqSwitchActive) {
  return request<ApiResponseData<null>>({
    url: "/user/status",
    method: "put",
    data
  })
}
