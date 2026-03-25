export const STATUS_CONFIG = {
  '-1': { type: 'danger' as const, text: '禁用' },
  '1': { type: 'success' as const, text: '启用' }
} as const

export const STATUS_OPTIONS = [
  { label: '请选择', value: '' },
  { label: '禁用', value: -1 },
  { label: '启用', value: 1 }
]

export const getStatusConfig = (status: string) => {
  return (
    STATUS_CONFIG[status as keyof typeof STATUS_CONFIG] || {
      type: 'info' as const,
      text: '未知'
    }
  )
}

export const GENDER_CONFIG = {
  '0': { type: 'info' as const, text: '未知' },
  '1': { type: '' as const, text: '男' },
  '2': { type: '' as const, text: '女' }
} as const

export const GENDER_OPTIONS = [
  { label: '未知', value: 0 },
  { label: '男', value: 1 },
  { label: '女', value: 2 }
]

export const getGenderConfig = (gender: string) => {
  return (
    GENDER_CONFIG[gender as keyof typeof GENDER_CONFIG] || {
      type: 'info' as const,
      text: '未知'
    }
  )
}

export const formatDateTime = (value: string | number | undefined | null): string => {
  if (!value && value !== 0) return '-'

  let date: Date

  if (typeof value === 'string') {
    if (/^\d+$/.test(value)) {
      date = new Date(Number(value) * 1000)
    } else {
      date = new Date(value)
    }
  } else if (typeof value === 'number') {
    date = new Date(value * 1000)
  } else {
    return '-'
  }

  if (isNaN(date.getTime())) return '-'

  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')
  const hours = String(date.getHours()).padStart(2, '0')
  const minutes = String(date.getMinutes()).padStart(2, '0')
  const seconds = String(date.getSeconds()).padStart(2, '0')

  return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`
}
