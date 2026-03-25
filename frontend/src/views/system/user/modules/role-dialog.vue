<template>
  <ElDialog
    v-model="dialogVisible"
    title="设置角色"
    width="40%"
    align-center
  >
    <ArtForm
      ref="formRef"
      v-model="formData"
      :items="formItems"
      label-width="80px"
      :span="24"
      :button-left-limit="0"
      :show-submit="true"
      @submit="handleSubmit"
      @reset="handleReset"
    />
  </ElDialog>
</template>

<script setup lang="ts">
  import type { FormInstance } from 'element-plus'
  import type { FormItem } from '@/components/core/forms/art-form/index.vue'
  import ArtForm from '@/components/core/forms/art-form/index.vue'
  import { ElMessage } from 'element-plus'
  import { API_URL } from '@/utils/constants'
  import { fetchGet, fetchSave } from '@/api/common'

  interface Props {
    visible: boolean
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

  const formRef = ref<FormInstance>()

  const formData = reactive({
    role_ids: [] as number[]
  })

  const roleList = ref<{ id: number; name: string }[]>([])

  const formItems = computed<FormItem[]>(() => [
    {
      label: '角色',
      key: 'role_ids',
      type: 'select',
      props: {
        multiple: true,
        placeholder: '请选择角色',
        options: roleList.value.map((role: any) => ({
          label: role.name,
          value: role.id
        }))
      }
    }
  ])

  watch(
    () => props.visible,
    async (visible) => {
      if (visible) {
        const res = await fetchGet(API_URL.common.roleList)
        roleList.value = Array.isArray(res) ? res : (res.data || [])
        
        const roles = props.data?.roles as any[] | undefined
        formData.role_ids = roles ? roles.map((role: any) => role.id) : []
      }
    },
    { immediate: true }
  )

  const handleReset = () => {
    dialogVisible.value = false
  }

  const handleSubmit = async () => {
    await fetchSave(`${API_URL.user.setRole}/${props.data?.id}`, {
      role_ids: formData.role_ids.join(',')
    })
    ElMessage.success('设置成功')
    dialogVisible.value = false
    emit('submit')
  }
</script>
