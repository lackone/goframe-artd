import request from '@/utils/http'

export function fetchGet<T = any>(url: string, params?: any) {
  return request.get<T>({
    url,
    params
  })
}

export function fetchDelete(url: string, ids: number[]) {
  return request.post({
    url,
    params: { ids: ids.join(',') }
  })
}

export function fetchSave(url: string, data: any) {
  return request.post({
    url,
    data
  })
}
