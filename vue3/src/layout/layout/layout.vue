<template>
  <a-layout id="components-layout-demo-fixed-sider">
    <a-layout-sider class="sider-height">
      <div class="logo">安全凭证管理后台</div>
      <common-menu />
    </a-layout-sider>
    <a-layout>
      <common-header />
      <a-layout-content class="layout-content">
        <common-tabs>
          <router-view />
        </common-tabs>
      </a-layout-content>
      <a-layout-footer :style="{ textAlign: 'center' }">
        <common-footer />
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>
<script lang="ts">
import { defineComponent } from 'vue'

import CommonFooter from '../footer/footer.vue'
import CommonMenu from '../menu/menu.vue'
import CommonHeader from '../header/header.vue'
import CommonTabs from '../tabs/index.vue'

export default defineComponent({
  name: 'Layout',
  setup() {
    const selectedKeys: string[] = ['4']
    return {
      selectedKeys,
    }
  },
  components: {
    CommonMenu,
    CommonFooter,
    CommonHeader,
    CommonTabs,
  },
})
</script>
<style lang="less" scoped>
@import 'layout.less';
</style>
