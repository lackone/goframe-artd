<template>
  <div class="page-content">
    <div class="p-4">
      <ElTabs v-model="activeTab" type="card" class="mb-4">
        <ElTabPane label="网站信息" name="info">
          <WebsiteInfo />
        </ElTabPane>
        <ElTabPane label="网站备案" name="filing">
          <WebsiteFiling />
        </ElTabPane>
        <ElTabPane label="政策协议" name="policy">
          <PolicyAgreement />
        </ElTabPane>
      </ElTabs>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import WebsiteInfo from './modules/website-info.vue'
import WebsiteFiling from './modules/website-filing.vue'
import PolicyAgreement from './modules/policy-agreement.vue'

defineOptions({ name: 'Website' })

const activeTab = ref('info')
</script>

<style scoped>
/* 自定义样式 */
</style>