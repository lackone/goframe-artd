<template>
  <ElDialog
    v-model="dialogVisible"
    :title="dialogType === 'add' ? '添加用户' : '编辑用户'"
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
  import type { UploadFile, UploadFiles, UploadUserFile } from 'element-plus'
  import ArtForm from '@/components/core/forms/art-form/index.vue'
  import { ElMessage, ElUpload, ElIcon } from 'element-plus'
  import { Plus } from '@element-plus/icons-vue'
  import { STATUS_OPTIONS, GENDER_OPTIONS, API_URL } from '@/utils/constants'
  import { fetchSave } from '@/api/common'

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
    account: '',
    password: '',
    nick_name: '',
    real_name: '',
    avatar: [] as UploadUserFile[],
    sex: 0,
    phone: '',
    tel: '',
    email: '',
    weixin: '',
    is_super: -1,
    status: 1,
    address: '',
    remark: ''
  }

  const formData = reactive({ ...initialFormData })

  const rules: FormRules = {
    account: [
      { required: true, message: '请输入账号', trigger: 'blur' },
      { min: 2, max: 32, message: '长度在 2 到 32 个字符', trigger: 'blur' }
    ],
    nick_name: [{ max: 32, message: '最多 32 个字符', trigger: 'blur' }],
    real_name: [{ max: 32, message: '最多 32 个字符', trigger: 'blur' }],
    phone: [{ pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号格式', trigger: 'blur' }],
    tel: [{ max: 32, message: '最多 32 个字符', trigger: 'blur' }],
    email: [{ type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }],
    weixin: [{ max: 32, message: '最多 32 个字符', trigger: 'blur' }],
    address: [{ max: 255, message: '最多 255 个字符', trigger: 'blur' }],
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
      label: '账号',
      key: 'account',
      type: 'input',
      props: { placeholder: '请输入账号' }
    },
    {
      label: '昵称',
      key: 'nick_name',
      type: 'input',
      props: { placeholder: '请输入昵称' }
    },
    {
      label: '真实姓名',
      key: 'real_name',
      type: 'input',
      props: { placeholder: '请输入真实姓名' }
    },
    {
      label: '头像',
      key: 'avatar',
      span: 12,
      render: () => {
        return h(ElUpload, {
          accept: '.jpg,.jpeg,.png,.gif,.webp',
          limit: 1,
          action: API_URL.common.upload,
          autoUpload: true,
          showFileList: true,
          listType: 'picture-card',
          fileList: formData.avatar,
          beforeUpload: (file: File) => {
            const isImage = file.type.startsWith('image/')
            const isLt2M = file.size / 1024 / 1024 < 2
            if (!isImage) {
              ElMessage.error('只能上传图片文件!')
              return false
            }
            if (!isLt2M) {
              ElMessage.error('图片大小不能超过 2MB!')
              return false
            }
            return true
          },
          onChange: (file: UploadFile, fileList: UploadFiles) => {
            formData.avatar = fileList as UploadUserFile[]
          },
          onSuccess: (response: any, file: UploadFile) => {
            const url = response.data?.url || response.url
            if (url) {
              file.url = url
            }
          },
          onRemove: () => {
            formData.avatar = []
          }
        }, {
          default: () => [h(ElIcon, { type: 'primary' }, () => h(Plus))]
        })
      }
    },
    {
      label: '密码',
      key: 'password',
      type: 'input',
      props: { type: 'password', placeholder: '请输入密码' }
    },
    {
      label: '性别',
      key: 'sex',
      type: 'select',
      props: {
        placeholder: '请选择性别',
        options: GENDER_OPTIONS
      }
    },
    {
      label: '手机号',
      key: 'phone',
      type: 'input',
      props: { placeholder: '请输入手机号', maxlength: 11 }
    },
    {
      label: '座机',
      key: 'tel',
      type: 'input',
      props: { placeholder: '请输入座机' }
    },
    {
      label: '邮箱',
      key: 'email',
      type: 'input',
      props: { placeholder: '请输入邮箱' }
    },
    {
      label: '微信',
      key: 'weixin',
      type: 'input',
      props: { placeholder: '请输入微信' }
    },
    {
      label: '超级管理员',
      key: 'is_super',
      type: 'select',
      props: {
        placeholder: '请选择',
        options: [
          { label: '否', value: -1 },
          { label: '是', value: 1 }
        ]
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
      label: '地址',
      key: 'address',
      type: 'input',
      props: { placeholder: '请输入地址' }
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
          account: isEdit && row ? (row.account as string) || '' : '',
          password: '',
          nick_name: isEdit && row ? (row.nick_name as string) || '' : '',
          real_name: isEdit && row ? (row.real_name as string) || '' : '',
          avatar: isEdit && row ? (row.avatar ? [{ url: row.avatar }] : []) as UploadUserFile[] : [],
          sex: isEdit && row ? (row.sex as number) || 0 : 0,
          phone: isEdit && row ? (row.phone as string) || '' : '',
          tel: isEdit && row ? (row.tel as string) || '' : '',
          email: isEdit && row ? (row.email as string) || '' : '',
          weixin: isEdit && row ? (row.weixin as string) || '' : '',
          is_super: isEdit && row ? (row.is_super as number) || -1 : -1,
          status: isEdit && row ? (row.status as number) || 1 : 1,
          address: isEdit && row ? (row.address as string) || '' : '',
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

    const mergedData = { ...initialFormData, ...params }
    
    const submitData: Record<string, any> = {
      ...mergedData,
      avatar: params.avatar?.[0]?.url || ''
    }
    
    const id = submitData.id
    const url = id ? `${API_URL.user.save}/${id}` : API_URL.user.save
    
    await fetchSave(url, submitData)
    
    ElMessage.success(dialogType.value === 'add' ? '添加成功' : '更新成功')
    dialogVisible.value = false
    emit('submit')
  }
</script>
