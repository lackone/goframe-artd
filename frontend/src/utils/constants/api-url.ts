export const API_URL = {
  common: {
    upload: '/api/common/upload',
    roleList: '/api/common/role_list',
    authList: '/api/common/auth_list',
  },
  user: {
    list: '/api/admin/list',
    delete: '/api/admin/delete',
    save: '/api/admin/save',
    setRole: '/api/admin/set_role',
  },
  role: {
    list: '/api/role/list',
    delete: '/api/role/delete',
    save: '/api/role/save'
  },
  auth: {
    list: '/api/auth/list',
    delete: '/api/auth/delete',
    save: '/api/auth/save'
  },
  config: {
    set: '/api/config/set',
    get: '/api/config/get'
  }
}
