import { request } from "@/utils/service"
import { type MenusData } from "./menu"

export interface roleData {
  ID: number
  roleName: string
  menus: MenusData[]
}

type RoleResponseData = IApiResponseData<roleData[]>

/** 获取用户详情 */
export function getRolesApi() {
  return request<RoleResponseData>({
    url: "/role/getRoles",
    method: "post",
    data: {}
  })
}

export interface reqRole {
  roleName: string
}

export function addRoleApi(data: reqRole) {
  return request<IApiResponseData<roleData>>({
    url: "/role/addRole",
    method: "post",
    data: data
  })
}

export function deleteRoleApi(data: reqId) {
  return request<IApiResponseData<null>>({
    url: "/role/deleteRole",
    method: "delete",
    data
  })
}

interface reqEditRole {
  id: number
  roleName: string
}

export function editRoleApi(data: reqEditRole) {
  return request<IApiResponseData<null>>({
    url: "/role/editRole",
    method: "post",
    data
  })
}

export function getRoleMenusApi(data: reqId) {
  return request<IApiResponseData<MenusData[]>>({
    url: "/role/getRoleMenus",
    method: "post",
    data
  })
}
