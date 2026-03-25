<template>
  <div>
    <ElCard :body-style="{ padding: '0' }">
      <template #header>
        <div class="flex justify-between items-center">
          <span>网站信息设置</span>
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
    
    <!-- 图片预览对话框 -->
    <ElDialog v-model="dialogVisible">
      <img w-full :src="dialogImageUrl" alt="Preview Image" class="w-full h-auto" />
    </ElDialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, h, onMounted } from 'vue'
import ArtForm from '@/components/core/forms/art-form/index.vue'
import { ElMessage, ElUpload, ElButton, ElIcon, ElDialog } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import type { UploadFile, UploadFiles, UploadUserFile } from 'element-plus'
import { API_URL } from '@/utils/constants'
import { fetchGet, fetchSave } from '@/api/common'

defineOptions({ name: 'WebsiteInfo' })

const formRef = ref()
const dialogVisible = ref(false)
const dialogImageUrl = ref('')
const loading = ref(false)

const initialFormData = {
  website_name: '',
  website_domain: '',
  website_title: '',
  website_keywords: '',
  website_description: '',
  website_logo: [] as UploadUserFile[],
  website_favicon: [] as UploadUserFile[],
}

const formData = ref({ ...initialFormData })

const formItems = computed(() => [
  {
    label: '网站名称',
    key: 'website_name',
    type: 'input',
    placeholder: '请输入网站名称',
    clearable: true,
    span: 24
  },
  {
    label: '网站域名',
    key: 'website_domain',
    type: 'input',
    placeholder: '请输入网站域名',
    clearable: true,
    span: 24
  },
  {
    label: '网站标题',
    key: 'website_title',
    type: 'input',
    placeholder: '请输入网站标题',
    clearable: true,
    span: 24
  },
  {
    label: '网站关键词',
    key: 'website_keywords',
    type: 'input',
    placeholder: '请输入网站关键词，多个用逗号分隔',
    clearable: true,
    span: 24
  },
  {
    label: '网站描述',
    key: 'website_description',
    type: 'input',
    props: {
      type: 'textarea',
      rows: 4,
      placeholder: '请输入网站描述'
    },
    span: 24
  },
  {
    label: '网站Logo',
    key: 'website_logo',
    span: 24,
    render: () => {
      return h(ElUpload, {
        accept: '.jpg,.jpeg,.png,.gif,.webp',
        limit: 1,
        action: API_URL.common.upload,
        autoUpload: true,
        showFileList: true,
        listType: 'picture-card',
        fileList: formData.value.website_logo,
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
          formData.value.website_logo = fileList as UploadUserFile[]
        },
        onSuccess: (response: any, file: UploadFile) => {
          const url = response.data?.url || response.url
          if (url) {
            file.url = url
          }
        },
        onRemove: () => {
          formData.value.website_logo = []
        },
        onPreview: (file: UploadFile) => {
          dialogImageUrl.value = file.url || ''
          dialogVisible.value = true
        }
      }, {
        default: () => [h(ElIcon, () => h(Plus))]
      })
    }
  },
  {
    label: '网站Favicon',
    key: 'website_favicon',
    span: 24,
    render: () => {
      return h(ElUpload, {
        accept: '.ico',
        limit: 1,
        action: API_URL.common.upload,
        autoUpload: true,
        showFileList: true,
        listType: 'picture-card',
        fileList: formData.value.website_favicon,
        beforeUpload: (file: File) => {
          const isICO = file.name.endsWith('.ico')
          const isLt100K = file.size / 1024 < 100
          if (!isICO) {
            ElMessage.error('只能上传 ICO 格式文件!')
            return false
          }
          if (!isLt100K) {
            ElMessage.error('文件大小不能超过 100KB!')
            return false
          }
          return true
        },
        onChange: (file: UploadFile, fileList: UploadFiles) => {
          formData.value.website_favicon = fileList as UploadUserFile[]
        },
        onSuccess: (response: any, file: UploadFile) => {
          const url = response.data?.url || response.url
          if (url) {
            file.url = url
          }
        },
        onRemove: () => {
          formData.value.website_favicon = []
        },
        onPreview: (file: UploadFile) => {
          dialogImageUrl.value = file.url || ''
          dialogVisible.value = true
        }
      }, {
        default: () => [h(ElIcon, () => h(Plus))]
      })
    }
  },
])

// 加载配置数据
const loadConfig = async () => {
  loading.value = true
  try {
    const res = await fetchGet(API_URL.config.get, {'key': 'website'})
    if (res) {
      const data = res
      Object.assign(formData.value, {
        website_name: data.website_name || '',
        website_domain: data.website_domain || '',
        website_title: data.website_title || '',
        website_keywords: data.website_keywords || '',
        website_description: data.website_description || '',
        website_logo: data.website_logo ? [{ url: data.website_logo }] : [],
        website_favicon: data.website_favicon ? [{ url: data.website_favicon }] : [],
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
      key: 'website',
      website_logo: formData.value.website_logo?.[0]?.url || '',
      website_favicon: formData.value.website_favicon?.[0]?.url || ''
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
    website_name: '',
    website_domain: '',
    website_title: '',
    website_keywords: '',
    website_description: '',
    website_logo: [],
    website_favicon: [],
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