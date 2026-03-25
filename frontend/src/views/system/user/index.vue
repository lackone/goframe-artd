<template>
  <div class="art-full-height">
    <!-- 搜索栏 -->
    <Search v-model="searchForm" @search="handleSearch" @reset="resetSearchParams"></Search>

    <ElCard class="art-table-card">
      <!-- 表格头部 -->
      <ArtTableHeader v-model:columns="columnChecks" :loading="loading" @refresh="refreshData">
        <template #left>
          <ElSpace wrap>
            <ElButton type="primary" v-auth="'add'" @click="showDialog('add')" v-ripple>
              <ElIcon>
                <Plus />
              </ElIcon>
              新增
            </ElButton>
          </ElSpace>
        </template>
      </ArtTableHeader>

      <!-- 表格 -->
      <ArtTable
        :loading="loading"
        :data="data"
        :columns="columns"
        :pagination="pagination"
        @selection-change="handleSelectionChange"
        @pagination:size-change="handleSizeChange"
        @pagination:current-change="handleCurrentChange"
      >
      </ArtTable>

      <!-- 弹窗 -->
      <Dialog
        v-model:visible="dialogVisible"
        :type="dialogType"
        :data="currentData"
        @submit="handleDialogSubmit"
      />
      <!-- 角色弹窗 -->
      <RoleDialog
        v-model:visible="roleDialogVisible"
        :data="currentData"
        @submit="refreshData"
      />
    </ElCard>
  </div>
</template>

<script setup lang="ts">
  import { Plus } from '@element-plus/icons-vue'
  import ArtButtonTable from '@/components/core/forms/art-button-table/index.vue'
  import { useTable } from '@/hooks/core/useTable'
  import { useAuth } from '@/hooks/core/useAuth'
  import { fetchGet, fetchDelete } from '@/api/common'
  import Search from './modules/search.vue'
  import Dialog from './modules/dialog.vue'
  import RoleDialog from './modules/role-dialog.vue'
  import { ElTag, ElMessageBox, ElImage, ElTooltip } from 'element-plus'
  import { DialogType } from '@/types'
  import { getStatusConfig, formatDateTime, API_URL } from '@/utils/constants'
  import defaultAvatar from '@/assets/images/avatar/avatar.png'

  defineOptions({ name: 'User' })

  type ListItem = Api.Common.ListItem

  // 弹窗相关
  const dialogType = ref<DialogType>('add')
  const dialogVisible = ref(false)
  const currentData = ref<Partial<ListItem>>({})

  // 角色弹窗相关
  const roleDialogVisible = ref(false)

  // 选中行
  const selectedRows = ref<ListItem[]>([])

  // 搜索表单
  const searchForm = ref({})

  const { hasAuth } = useAuth()

  const {
    columns,
    columnChecks,
    data,
    loading,
    pagination,
    getData,
    replaceSearchParams,
    resetSearchParams,
    handleSizeChange,
    handleCurrentChange,
    refreshData
  } = useTable({
    // 核心配置
    core: {
      apiFn: (params: any) => fetchGet<Api.Common.DataList>(API_URL.user.list, params),
      apiParams: {
        current: 1,
        size: 10,
        ...searchForm.value
      },
      paginationKey: {
        current: 'page',
        size: 'size'
      },
      columnsFactory: () => [
        { type: 'selection' }, // 勾选列
        { prop: 'id', width: 80, label: 'ID', sortable: true },
        { prop: 'account', label: '账号' },
        {
          prop: 'avatar',
          label: '头像/昵称/姓名',
          width: 280,
          formatter: (row) => {
            const avatarUrl = row.avatar || defaultAvatar
            return h('div', { class: 'flex-c' }, [
              h(ElImage, {
                class: 'size-9.5 rounded-md',
                src: avatarUrl,
                previewSrcList: [avatarUrl],
                previewTeleported: true
              }),
              h('div', { class: 'ml-2' }, [
                h('p', { class: '' }, row.nick_name),
                h('p', { class: '' }, row.real_name)
              ])
            ])
          }
        },
        {
          prop: 'roles',
          label: '角色',
          formatter: (row) => {
            const roleNames = (row.roles as any[])?.map((role: any) => role.name) || []
            const roleNamesStr = roleNames.join(',')
            return h('p', { class: '' }, roleNamesStr)
          }
        },
        {
          prop: 'is_super',
          label: '超级管理员',
          formatter: (row) => {
            const superAdmin = row.is_super == 1 ? '是' : '否'
            return h(ElTag, { type: row.is_super == 1 ? 'success' : 'danger' }, superAdmin)
          }
        },
        { prop: 'phone', label: '手机号' },
        { prop: 'email', label: '邮箱' },
        {
          prop: 'status',
          label: '状态',
          formatter: (row) => {
            const statusConfig = getStatusConfig(row.status)
            return h(ElTag, { type: statusConfig.type }, () => statusConfig.text)
          }
        },
        {
          prop: 'last_login_ip',
          label: '登录IP/时间',
          formatter: (row) => {
            return h('div', { class: '' }, [
              h('p', { class: '' }, row.last_login_ip || '-'),
              h('p', { class: '' }, formatDateTime(row.last_login_time))
            ]);
          }
        },
        {
          prop: 'created',
          label: '创建/更新时间',
          formatter: (row) => {
            return h('div', { class: '' }, [
              h('p', { class: '' }, formatDateTime(row.created)),
              h('p', { class: '' }, formatDateTime(row.updated))
            ]);
          }
        },
        {
          prop: 'operation',
          label: '操作',
          width: 180,
          fixed: 'right',
          formatter: (row) =>
            h('div', [
              hasAuth('edit') ? h(ElTooltip, { content: '编辑', placement: 'top' }, () => 
                h(ArtButtonTable, {
                  type: 'edit',
                  onClick: () => showDialog('edit', row)
                })
              ) : null,
              hasAuth('setRole') ? h(ElTooltip, { content: '设置角色', placement: 'top' }, () => 
                h(ArtButtonTable, {
                  icon: 'ri:user-settings-line',
                  iconClass: 'bg-info/12 text-info',
                  onClick: () => showRoleDialog(row)
                })
              ) : null,
              hasAuth('delete') ? h(ElTooltip, { content: '删除', placement: 'top' }, () => 
                h(ArtButtonTable, {
                  type: 'delete',
                  onClick: () => deleteData(row)
                })
              ) : null
            ])
        }
      ]
    },
    // 数据处理
    transform: {
      // 数据转换器 - 替换头像
      dataTransformer: (records) => {
        // 类型守卫检查
        if (!Array.isArray(records)) {
          console.warn('数据转换器: 期望数组类型，实际收到:', typeof records)
          return []
        }

        // 使用本地头像替换接口返回的头像
        return records.map((item, index: number) => {
          return {
            ...item,
          }
        })
      }
    }
  })

  /**
   * 搜索处理
   * @param params 参数
   */
  const handleSearch = (params: Api.Common.CommonSearchParams) => {
    const { start_end_time, ...filtersParams } = params
    const [start_time, end_time] = Array.isArray(start_end_time) ? start_end_time : [null, null]
    
    replaceSearchParams({ ...filtersParams, start_time, end_time })
    getData()
  }

  /**
   * 显示弹窗
   */
  const showDialog = (type: DialogType, row?: ListItem): void => {
    dialogType.value = type
    currentData.value = row || {}
    nextTick(() => {
      dialogVisible.value = true
    })
  }

  const showRoleDialog = (row: ListItem): void => {
    currentData.value = row
    nextTick(() => {
      roleDialogVisible.value = true
    })
  }

  /**
   * 删除
   */
  const deleteData = async (row: ListItem): Promise<void> => {
    try {
      await ElMessageBox.confirm(`确定要删除 ID:${row.id} 吗？`, '删除', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'error'
      })
      await fetchDelete(API_URL.user.delete, [row.id])
      ElMessage.success('删除成功')
      refreshData()
    } catch {
      // 取消删除
    }
  }

  /**
   * 处理弹窗提交事件
   */
  const handleDialogSubmit = async () => {
    try {
      dialogVisible.value = false
      currentData.value = {}
      refreshData()
    } catch (error) {
      console.error('提交失败:', error)
    }
  }

  /**
   * 处理表格行选择变化
   */
  const handleSelectionChange = (selection: ListItem[]): void => {
    selectedRows.value = selection
    console.log('选中行数据:', selectedRows.value)
  }
</script>
