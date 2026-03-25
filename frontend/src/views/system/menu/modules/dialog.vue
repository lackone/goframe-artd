<template>
  <ElDialog
    :title="dialogTitle"
    :model-value="visible"
    @update:model-value="handleCancel"
    width="860px"
    align-center
    class="menu-dialog"
    @closed="handleClosed"
  >
    <ArtForm
      ref="formRef"
      v-model="form"
      :items="formItems"
      :rules="rules"
      :span="width > 640 ? 12 : 24"
      :gutter="20"
      label-width="100px"
      :show-reset="false"
      :show-submit="false"
      :sanitizeOutput="{ removeEmptyString: false }"
    >
      <template #menuType>
        <ElRadioGroup v-model="form.menuType" :disabled="disableMenuType">
          <ElRadioButton value="menu" label="menu">菜单</ElRadioButton>
          <ElRadioButton value="button" label="button">按钮</ElRadioButton>
        </ElRadioGroup>
      </template>
    </ArtForm>

    <template #footer>
      <span class="dialog-footer">
        <ElButton @click="handleCancel">取 消</ElButton>
        <ElButton type="primary" @click="handleSubmit">确 定</ElButton>
      </span>
    </template>
  </ElDialog>
</template>

<script setup lang="ts">
  import type { FormRules } from 'element-plus'
  import { ElIcon, ElTooltip } from 'element-plus'
  import { QuestionFilled } from '@element-plus/icons-vue'
  import { formatMenuTitle } from '@/utils/router'
  import type { AppRouteRecord } from '@/types/router'
  import type { FormItem } from '@/components/core/forms/art-form/index.vue'
  import ArtForm from '@/components/core/forms/art-form/index.vue'
  import { useWindowSize } from '@vueuse/core'
  import { API_URL } from '@/utils/constants'
  import { fetchSave } from '@/api/common'

  const { width } = useWindowSize()

  /**
   * 创建带 tooltip 的表单标签
   * @param label 标签文本
   * @param tooltip 提示文本
   * @returns 渲染函数
   */
  const createLabelTooltip = (label: string, tooltip: string) => {
    return () =>
      h('span', { class: 'flex items-center' }, [
        h('span', label),
        h(
          ElTooltip,
          {
            content: tooltip,
            placement: 'top'
          },
          () => h(ElIcon, { class: 'ml-0.5 cursor-help' }, () => h(QuestionFilled))
        )
      ])
  }

  interface MenuFormData {
    id: number
    type: number
    title: string
    path: string
    name: string
    component: string
    icon: string
    is_enable: boolean
    sort: number
    keepalive: boolean
    is_hide: boolean
    is_hide_tab: boolean
    link: string
    is_iframe: boolean
    show_badge: boolean
    show_text_badge: string
    fixed_tab: boolean
    active_path: string
    is_full_page: boolean
    pid: number
  }

  interface Props {
    visible: boolean
    editData?: AppRouteRecord | any
    type?: 'menu' | 'button'
    lockType?: boolean
  }

  interface Emits {
    (e: 'update:visible', value: boolean): void
    (e: 'submit', data: any): void
  }

  const props = withDefaults(defineProps<Props>(), {
    visible: false,
    type: 'menu',
    lockType: false
  })

  const emit = defineEmits<Emits>()

  const formRef = ref()
  const isEdit = ref(false)

  const initialFormData = {
    menuType: 'menu',
    id: 0,
    type: 1,
    title: '',
    path: '',
    name: '',
    component: '',
    icon: '',
    is_enable: true,
    sort: 1,
    keepalive: false,
    is_hide: false,
    is_hide_tab: false,
    link: '',
    is_iframe: false,
    show_badge: false,
    show_text_badge: '',
    fixed_tab: false,
    active_path: '',
    is_full_page: false,
    pid: 0
  }

  const form = reactive<typeof initialFormData>({ ...initialFormData })

  const rules = reactive<FormRules>({
    title: [
      { required: true, message: '请输入菜单/权限名称', trigger: 'blur' },
      { min: 2, max: 20, message: '长度在 2 到 20 个字符', trigger: 'blur' }
    ],
    path: [{ required: true, message: '请输入路由地址', trigger: 'blur' }],
    name: [{ required: true, message: '请输入权限标识', trigger: 'blur' }],
  })

  /**
   * 表单项配置
   */
  const formItems = computed<FormItem[]>(() => {
    const baseItems: FormItem[] = [{ label: '菜单类型', key: 'menuType', span: 24 }]

    // Switch 组件的 span：小屏幕 12，大屏幕 6
    const switchSpan = width.value < 640 ? 12 : 6

    if (form.menuType === 'menu') {
      return [
        ...baseItems,
        { label: '菜单名称', key: 'title', type: 'input', props: { placeholder: '菜单名称' }, span: 24 },
        {
          label: createLabelTooltip(
            '路由地址',
            '一级菜单：以 / 开头的绝对路径（如 /dashboard）\n二级及以下：相对路径（如 console、user）'
          ),
          key: 'path',
          type: 'input',
          props: { placeholder: '如：/dashboard 或 console' },
          span: 24
        },
        { label: '权限标识', key: 'name', type: 'input', props: { placeholder: '如：User' }, span: 24 },
        {
          label: createLabelTooltip(
            '组件路径',
            '一级父级菜单：填写 /index/index\n具体页面：填写组件路径（如 /system/user）\n目录菜单：留空'
          ),
          key: 'component',
          type: 'input',
          props: { placeholder: '如：/system/user 或留空' },
          span: 24
        },
        { label: '图标', key: 'icon', type: 'input', props: { placeholder: '如：ri:user-line' }, span: 24 },
        {
          label: '菜单排序',
          key: 'sort',
          type: 'number',
          props: { min: 1, controlsPosition: 'right', style: { width: '100%' } },
          span: 24
        },
        {
          label: '外部链接',
          key: 'link',
          type: 'input',
          props: { placeholder: '如：https://www.example.com' },
          span: 24
        },
        {
          label: '文本徽章',
          key: 'show_text_badge',
          type: 'input',
          props: { placeholder: '如：New、Hot' },
          span: 24
        },
        {
          label: createLabelTooltip(
            '激活路径',
            '用于详情页等隐藏菜单，指定高亮显示的父级菜单路径\n例如：用户详情页高亮显示"用户管理"菜单'
          ),
          key: 'active_path',
          type: 'input',
          props: { placeholder: '如：/system/user' },
          span: 24
        },
        { label: '是否启用', key: 'is_enable', type: 'switch', span: switchSpan },
        { label: '页面缓存', key: 'keepalive', type: 'switch', span: switchSpan },
        { label: '隐藏菜单', key: 'is_hide', type: 'switch', span: switchSpan },
        { label: '是否内嵌', key: 'is_iframe', type: 'switch', span: switchSpan },
        { label: '显示徽章', key: 'show_badge', type: 'switch', span: switchSpan },
        { label: '固定标签', key: 'fixed_tab', type: 'switch', span: switchSpan },
        { label: '标签隐藏', key: 'is_hide_tab', type: 'switch', span: switchSpan },
        { label: '全屏页面', key: 'is_full_page', type: 'switch', span: switchSpan }
      ]
    } else {
      return [
        ...baseItems,
        {
          label: '权限名称',
          key: 'title',
          type: 'input',
          props: { placeholder: '如：新增、编辑、删除' },
          span: 24
        },
        {
          label: '权限标识',
          key: 'name',
          type: 'input',
          props: { placeholder: '如：add、edit、delete' },
          span: 24
        },
        {
          label: '权限排序',
          key: 'sort',
          type: 'number',
          props: { min: 1, controlsPosition: 'right', style: { width: '100%' } },
          span: 24
        }
      ]
    }
  })

  const dialogTitle = computed(() => {
    const type = form.menuType === 'menu' ? '菜单' : '按钮'
    return isEdit.value ? `编辑${type}` : `新建${type}`
  })

  /**
   * 是否禁用菜单类型切换
   */
  const disableMenuType = computed(() => {
    if (isEdit.value) return true
    if (!isEdit.value && form.menuType === 'menu' && props.lockType) return true
    return false
  })

  /**
   * 重置表单数据
   */
  const resetForm = (): void => {
    formRef.value?.reset()
    form.menuType = 'menu'
    form.pid = 0
  }

  /**
   * 加载表单数据（编辑模式）
   */
  const loadFormData = (): void => {
    if (!props.editData) return

    isEdit.value = true

    if (form.menuType === 'menu') {
      const row = props.editData
      form.id = row.id || 0
      form.type = row.meta?.type || 1
      form.title = formatMenuTitle(row.meta?.title || '')
      form.path = row.path || ''
      form.name = row.name || ''
      form.component = row.component || ''
      form.icon = row.meta?.icon || ''
      form.sort = row.meta?.sort || 1
      form.keepalive = row.meta?.keepAlive ?? false
      form.is_hide = row.meta?.isHide ?? false
      form.is_hide_tab = row.meta?.isHideTab ?? false
      form.is_enable = row.meta?.isEnable ?? true
      form.link = row.meta?.link || ''
      form.is_iframe = row.meta?.isIframe ?? false
      form.show_badge = row.meta?.showBadge ?? false
      form.show_text_badge = row.meta?.showTextBadge || ''
      form.fixed_tab = row.meta?.fixedTab ?? false
      form.active_path = row.meta?.activePath || ''
      form.is_full_page = row.meta?.isFullPage ?? false
      form.pid = row.meta?.pid || 0
    } else {
      const row = props.editData
      form.id = row.id || 0
      form.title = row.title || ''
      form.name = row.authMark || ''
      form.sort = row.sort || 1
      form.pid = row.pid || 0
    }
  }

  /**
   * 将布尔值转换为后端需要的数值
   * true -> 1, false -> -1
   */
  const convertBoolToNum = (value: boolean): number => (value ? 1 : -1)

  /**
   * 提交表单
   */
  const handleSubmit = async (): Promise<void> => {
    if (!formRef.value) return

    try {
      await formRef.value.validate()

      // 用初始数据填补 ArtForm 删除的字段
      const mergedData = { ...initialFormData, ...form }
      const submitData: Record<string, any> = { ...mergedData }

      submitData.is_enable = convertBoolToNum(form.is_enable)
      submitData.keepalive = convertBoolToNum(form.keepalive)
      submitData.is_hide = convertBoolToNum(form.is_hide)
      submitData.is_hide_tab = convertBoolToNum(form.is_hide_tab)
      submitData.is_iframe = convertBoolToNum(form.is_iframe)
      submitData.show_badge = convertBoolToNum(form.show_badge)
      submitData.fixed_tab = convertBoolToNum(form.fixed_tab)
      submitData.is_full_page = convertBoolToNum(form.is_full_page)

      submitData.id = isEdit.value ? form.id : 0
      submitData.type = form.menuType === 'menu' ? 1 : 2

      const url = submitData.id ? `${API_URL.auth.save}/${submitData.id}` : API_URL.auth.save
      await fetchSave(url, submitData)

      emit('submit', submitData)
      ElMessage.success(`${isEdit.value ? '编辑' : '新增'}成功`)
      handleCancel()
    } catch {
      ElMessage.error('表单校验失败，请检查输入')
    }
  }

  /**
   * 取消操作
   */
  const handleCancel = (): void => {
    emit('update:visible', false)
  }

  /**
   * 对话框关闭后的回调
   */
  const handleClosed = (): void => {
    resetForm()
    isEdit.value = false
  }

  /**
   * 监听对话框显示状态
   */
  watch(
    () => props.visible,
    (newVal) => {
      if (newVal) {
        form.menuType = props.type
        if (props.editData) {
          if (props.editData.id) {
            nextTick(() => {
              loadFormData()
            })
          } else {
            form.pid = props.editData.pid || 0
          }
        } else {
          form.pid = 0
        }
      }
    }
  )

  /**
   * 监听菜单类型变化
   */
  watch(
    () => props.type,
    (newType) => {
      if (props.visible) {
        form.menuType = newType
      }
    }
  )
</script>
