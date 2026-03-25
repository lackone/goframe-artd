<template>
  <div>
    <ElCard :body-style="{ padding: '0' }">
      <template #header>
        <div class="flex justify-between items-center">
          <span>政策协议管理</span>
        </div>
      </template>
      
      <ArtForm
        ref="formRef"
        v-model="formData"
        :items="formItems"
        :labelWidth="120"
        @submit="handleSubmit"
        @reset="handleReset"
      />
    </ElCard>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, h, onMounted } from 'vue'
import ArtForm from '@/components/core/forms/art-form/index.vue'
import ArtWangEditor from '@/components/core/forms/art-wang-editor/index.vue'
import { ElMessage } from 'element-plus'
import { API_URL } from '@/utils/constants'
import { fetchGet, fetchSave } from '@/api/common'

defineOptions({ name: 'PolicyAgreement' })

const formRef = ref()
const loading = ref(false)

const initialFormData = {
  user_agreement: '',
  privacy_policy: '',
  service_terms: '',
  disclaimer: '',
  copyright_notice: '',
  cookie_policy: ''
}

const formData = ref({ ...initialFormData })

const formItems = computed(() => [
  {
    label: '用户协议',
    key: 'user_agreement',
    span: 24,
    render: () =>
      h(ArtWangEditor, {
        modelValue: formData.value.user_agreement,
        height: '300px',
        placeholder: '请输入用户协议内容...',
        'onUpdate:modelValue': (value: string) => {
          formData.value.user_agreement = value
        },
        uploadConfig: {
          server: API_URL.common.upload,
          maxFileSize: 2 * 1024 * 1024, // 2MB
          maxNumberOfFiles: 10,
          isCustomUpload: true
        }
      })
  },
  {
    label: '隐私政策',
    key: 'privacy_policy',
    span: 24,
    render: () =>
      h(ArtWangEditor, {
        modelValue: formData.value.privacy_policy,
        height: '300px',
        placeholder: '请输入隐私政策内容...',
        'onUpdate:modelValue': (value: string) => {
          formData.value.privacy_policy = value
        },
        uploadConfig: {
          server: API_URL.common.upload,
          maxFileSize: 2 * 1024 * 1024, // 2MB
          maxNumberOfFiles: 10,
          isCustomUpload: true
        }
      })
  },
  {
    label: '服务条款',
    key: 'service_terms',
    span: 24,
    render: () =>
      h(ArtWangEditor, {
        modelValue: formData.value.service_terms,
        height: '300px',
        placeholder: '请输入服务条款内容...',
        'onUpdate:modelValue': (value: string) => {
          formData.value.service_terms = value
        },
        uploadConfig: {
          server: API_URL.common.upload,
          maxFileSize: 2 * 1024 * 1024, // 2MB
          maxNumberOfFiles: 10,
          isCustomUpload: true
        }
      })
  },
  {
    label: '免责声明',
    key: 'disclaimer',
    span: 24,
    render: () =>
      h(ArtWangEditor, {
        modelValue: formData.value.disclaimer,
        height: '300px',
        placeholder: '请输入免责声明内容...',
        'onUpdate:modelValue': (value: string) => {
          formData.value.disclaimer = value
        },
        uploadConfig: {
          server: API_URL.common.upload,
          maxFileSize: 2 * 1024 * 1024, // 2MB
          maxNumberOfFiles: 10,
          isCustomUpload: true
        }
      })
  },
  {
    label: '版权声明',
    key: 'copyright_notice',
    span: 24,
    render: () =>
      h(ArtWangEditor, {
        modelValue: formData.value.copyright_notice,
        height: '200px',
        placeholder: '请输入版权声明内容...',
        'onUpdate:modelValue': (value: string) => {
          formData.value.copyright_notice = value
        },
        uploadConfig: {
          server: API_URL.common.upload,
          maxFileSize: 2 * 1024 * 1024, // 2MB
          maxNumberOfFiles: 10,
          isCustomUpload: true
        }
      })
  },
  {
    label: 'Cookie政策',
    key: 'cookie_policy',
    span: 24,
    render: () =>
      h(ArtWangEditor, {
        modelValue: formData.value.cookie_policy,
        height: '200px',
        placeholder: '请输入Cookie政策内容...',
        'onUpdate:modelValue': (value: string) => {
          formData.value.cookie_policy = value
        },
        uploadConfig: {
          server: API_URL.common.upload,
          maxFileSize: 2 * 1024 * 1024, // 2MB
          maxNumberOfFiles: 10,
          isCustomUpload: true
        }
      })
  }
])

// 加载配置数据
const loadConfig = async () => {
  loading.value = true
  try {
    const res = await fetchGet(API_URL.config.get, {'key': 'policy'})
    if (res) {
      const data = res
      Object.assign(formData.value, {
        user_agreement: data.user_agreement || '',
        privacy_policy: data.privacy_policy || '',
        service_terms: data.service_terms || '',
        disclaimer: data.disclaimer || '',
        copyright_notice: data.copyright_notice || '',
        cookie_policy: data.cookie_policy || ''
      })
    }
  } catch (error) {
    ElMessage.error('加载配置失败')
    console.error('加载配置失败:', error)
  } finally {
    loading.value = false
  }
}

const handleSubmit = async () => {
  loading.value = true
  try {
    const mergedData = { ...initialFormData, ...formData.value }
    
    const submitData = {
      ...mergedData,
      key: 'policy'
    }
    
    await fetchSave(API_URL.config.set, submitData)
    ElMessage.success('保存成功')
  } catch (error) {
    ElMessage.error('保存失败')
    console.error('保存失败:', error)
  } finally {
    loading.value = false
  }
}

const handleReset = () => {
  Object.assign(formData.value, {
    user_agreement: '',
    privacy_policy: '',
    service_terms: '',
    disclaimer: '',
    copyright_notice: '',
    cookie_policy: ''
  })
}

// 组件挂载时加载配置
onMounted(() => {
  loadConfig()
})
</script>

<style scoped>
/* 自定义样式 */
</style>