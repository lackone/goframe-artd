declare namespace Api {
  namespace Common {
    interface PaginationParams {
      current: number
      size: number
      total: number
    }

    type CommonSearchParams = {
      current?: number
      size?: number
      total?: number
      [key: string]: any
    }

    interface PaginatedResponse<T = any> {
      records: T[]
      current: number
      size: number
      total: number
    }

    interface ListItem {
      [key: string]: any
    }

    type DataList = PaginatedResponse<ListItem>

    type EnableStatus = '1' | '2'
  }
}
