import request from '@/utils/http'
import { AppRouteRecord } from '@/types/router'

// 获取菜单列表
export function fetchGetMenuList(params?: any) {
  return request.get<AppRouteRecord[]>({
    url: '/api/auth/menu_list',
    params
  })
}