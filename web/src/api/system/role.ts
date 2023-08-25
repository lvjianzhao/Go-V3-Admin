import {request} from "@/utils/service"
import {type MenusData} from "./menu"

export interface roleData {
  ID: number
  roleName: string
  menus: MenusData[]
}

type RoleResponseData = ApiResponseData<roleData[]>

/** 获取用户详情 */
export function getRolesApi() {
  return request<RoleResponseData>({
    url: "/role",
    method: "get",
    data: {}
  })
}

export interface reqRole {
  roleName: string
}

export function addRoleApi(data: reqRole) {
  return request<ApiResponseData<roleData>>({
    url: "/role",
    method: "post",
    data: data
  })
}

export function deleteRoleApi(id: number) {
  return request<ApiResponseData<null>>({
    url: `/role/${id}`,
    method: "delete"
  })
}

interface reqEditRole {
  id: number
  roleName: string
}

export function editRoleApi(data: reqEditRole) {
  return request<ApiResponseData<null>>({
    url: "/role",
    method: "put",
    data
  })
}

interface reqEditRE {
  roleId: number
  ids: number[]
}

export function editRoleMenuApi(data: reqEditRE) {
  return request<ApiResponseData<null>>({
    url: "/role/roleMenu",
    method: "put",
    data
  })
}
