<template>
  <div>
    <ElCard :body-style="{ padding: '0' }">
      <template #header>
        <div class="flex justify-between items-center">
          <span>网站备案信息</span>
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
import { ref, reactive, computed, onMounted } from 'vue'
import ArtForm from '@/components/core/forms/art-form/index.vue'
import { ElMessage } from 'element-plus'
import { API_URL } from '@/utils/constants'
import { fetchGet, fetchSave } from '@/api/common'

defineOptions({ name: 'WebsiteFiling' })

const formRef = ref()
const loading = ref(false)

const initialFormData = {
  filing_number: '',
  filing_license: '',
  filing_subject: '',
  filing_link: '',
  filing_remark: ''
}

const formData = ref({ ...initialFormData })

const formItems = computed(() => [
  {
    label: '备案号',
    key: 'filing_number',
    type: 'input',
    placeholder: '请输入ICP备案号',
    clearable: true,
    span: 24
  },
  {
    label: '备案许可证',
    key: 'filing_license',
    type: 'input',
    placeholder: '请输入备案许可证号',
    clearable: true,
    span: 24
  },
  {
    label: '备案主体',
    key: 'filing_subject',
    type: 'input',
    placeholder: '请输入备案主体名称',
    clearable: true,
    span: 24
  },
  {
    label: '备案链接',
    key: 'filing_link',
    type: 'input',
    placeholder: '请输入工信部备案查询链接',
    clearable: true,
    span: 24
  },
  {
    label: '备案备注',
    key: 'filing_remark',
    type: 'input',
    props: {
      type: 'textarea',
      rows: 3,
      placeholder: '请输入备案备注信息'
    },
    span: 24
  }
])

// 加载配置数据
const loadConfig = async () => {
  loading.value = true
  try {
    const res = await fetchGet(API_URL.config.get, {'key': 'filing'})
    if (res) {
      const data = res
      Object.assign(formData.value, {
        filing_number: data.filing_number || '',
        filing_license: data.filing_license || '',
        filing_subject: data.filing_subject || '',
        filing_link: data.filing_link || '',
        filing_remark: data.filing_remark || ''
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
      key: 'filing'
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
    filing_number: '',
    filing_license: '',
    filing_subject: '',
    filing_link: '',
    filing_remark: ''
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