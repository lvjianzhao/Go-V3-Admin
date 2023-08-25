import { request } from "@/utils/service"

export interface MenusData {
  id: number
  pid: number
  name: string
  path: string
  redirect?: string
  component: string
  sort: number
  meta: {
    hidden?: boolean
    title: string
    elIcon?: string
    svgIcon?: string
    affix?: boolean
    keepAlive?: boolean
  }
  children?: MenusData[]
}

type MenusResponseData = ApiResponseData<MenusData[]>

// 获取动态路由
export function getMenus() {
  return request<MenusResponseData>({
    url: "/menu",
    method: "get"
  })
}

interface reqMenu {
  pid: number
  name?: string
  path: string
  redirect?: string
  component: string
  sort: number
  meta: {
    hidden?: boolean
    title?: string
    icon?: string
    affix?: boolean
    keepAlive?: boolean
  }
}

export function addMenuApi(data: reqMenu) {
  return request<ApiResponseData<null>>({
    url: "/menu",
    method: "post",
    data
  })
}

interface editReq extends reqMenu {
  id: number
}

export function editMenuApi(data: editReq) {
  return request<ApiResponseData<null>>({
    url: "/menu",
    method: "put",
    data
  })
}

export function deleteMenuApi(id: number) {
  return request<ApiResponseData<null>>({
    url: `menu/${id}`,
    method: "delete"
  })
}

interface allMenus {
  list: MenusData[]
  menuIds: number[]
}

export function getElTreeMenusApi(params: reqId) {
  return request<ApiResponseData<allMenus>>({
    url: "menu/getElTreeMenus",
    method: "get",
    params
  })
}
