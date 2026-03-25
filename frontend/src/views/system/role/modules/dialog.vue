<template>
  <ElDialog
    v-model="dialogVisible"
    :title="dialogType === 'add' ? '添加角色' : '编辑角色'"
    width="50%"
    align-center
  >
    <ArtForm
      ref="formRef"
      v-model="formData"
      :items="formItems"
      :rules="rules"
      label-width="100px"
      :span="24"
      :show-submit="true"
      @submit="handleSubmit"
      @reset="handleReset"
    />
  </ElDialog>
</template>

<script setup lang="ts">
  import type { FormInstance, FormRules } from 'element-plus'
  import type { FormItem } from '@/components/core/forms/art-form/index.vue'
  import ArtForm from '@/components/core/forms/art-form/index.vue'
  import { ElMessage, ElUpload, ElIcon } from 'element-plus'
  import { STATUS_OPTIONS, GENDER_OPTIONS, API_URL } from '@/utils/constants'
  import { fetchGet, fetchSave } from '@/api/common'

  interface Props {
    visible: boolean
    type: string
    data?: Partial<Api.Common.ListItem>
  }

  interface Emits {
    (e: 'update:visible', value: boolean): void
    (e: 'submit'): void
  }

  const props = defineProps<Props>()
  const emit = defineEmits<Emits>()

  const dialogVisible = computed({
    get: () => props.visible,
    set: (value) => emit('update:visible', value)
  })

  const dialogType = computed(() => props.type)

  const formRef = ref<FormInstance>()

  const initialFormData = {
    id: 0,
    name: '',
    auth_ids: [] as any[],
    status: 1,
    remark: ''
  }

  const formData = reactive({ ...initialFormData })

  const treeSelectData = ref<any[]>([])

  const fetchAuthList = async () => {
    const res = await fetchGet(API_URL.common.authList)
    treeSelectData.value = Array.isArray(res) ? res : (res.data || [])
  }
  fetchAuthList()

  const rules: FormRules = {
    name: [
      { required: true, message: '请输入角色名称', trigger: 'blur' },
      { min: 2, max: 32, message: '长度在 2 到 32 个字符', trigger: 'blur' }
    ],
    remark: [{ max: 255, message: '最多 255 个字符', trigger: 'blur' }]
  }

  const formItems = computed<FormItem[]>(() => [
    {
      label: 'ID',
      key: 'id',
      type: 'input',
      hidden: true
    },
    {
      label: '角色名称',
      key: 'name',
      type: 'input',
      props: { placeholder: '请输入角色名称' }
    },
    {
      label: '权限',
      key: 'auth_ids',
      type: 'treeselect',
      props: {
        showCheckbox: true,
        multiple: true,
        clearable: true,
        defaultExpandAll: true,
        checkStrictly: true,
        props: {
          label: 'label',
          children: 'children',
          value: 'value'
        },
        data: treeSelectData
      }
    },
    {
      label: '状态',
      key: 'status',
      type: 'select',
      props: {
        placeholder: '请选择状态',
        options: STATUS_OPTIONS
      }
    },
    {
      label: '备注',
      key: 'remark',
      type: 'input',
      props: { placeholder: '请输入备注', type: 'textarea', rows: 3 }
    }
  ])

  watch(
    () => props.visible,
    (visible) => {
      if (!visible) {
        Object.assign(formData, initialFormData)
      } else {
        const isEdit = props.type === 'edit' && props.data
        const row = props.data

        Object.assign(formData, {
          id: isEdit && row ? (row.id as number) || 0 : 0,
          name: isEdit && row ? (row.name as string) || '' : '',
          auth_ids: isEdit && row && row.auth_ids 
            ? (row.auth_ids as string).split(',').filter(Boolean).map(Number) 
            : [],
          status: isEdit && row ? (row.status as number) || 1 : 1,
          remark: isEdit && row ? (row.remark as string) || '' : ''
        })

        nextTick(() => {
          ;(formRef.value as any).ref?.clearValidate()
        })
      }
    },
    { immediate: true }
  )

  const handleReset = () => {
    console.log('重置表单')
  }

  const handleSubmit = async (params: Record<string, any>) => {
    await formRef.value?.validate()
    
    // 用初始数据填补 ArtForm 删除的字段
    const mergedData = { ...initialFormData, ...params }
    
    const submitData: Record<string, any> = {
      ...mergedData,
      auth_ids: mergedData.auth_ids?.join(',') || ''
    }
    
    const id = submitData.id
    const url = id ? `${API_URL.role.save}/${id}` : API_URL.role.save
    
    await fetchSave(url, submitData)
    
    ElMessage.success(dialogType.value === 'add' ? '添加成功' : '更新成功')
    dialogVisible.value = false
    emit('submit')
  }
</script>
