<template>
  <a-layout-header class="container">
    <div class="content">
      <div class="button-icon-container">
        <div @click="ToggleCollapsed()">
          <MenuUnfoldOutlined class="button-icon" v-if="collapsed" />
          <MenuFoldOutlined class="button-icon" v-else />
        </div>
        <a-breadcrumb>
          <a-breadcrumb-item
            class="button-breadcrumb"
            v-for="(item, index) in crumbs"
            :key="index"
            >{{ item }}
            {{
              index > 0 && index + 1 !== crumbs.length ? '/' : ''
            }}</a-breadcrumb-item
          >
        </a-breadcrumb>
      </div>
      <a-dropdown>
        <a class="ant-dropdown-link" @click="(e) => e.preventDefault()">
          <a-avatar :size="50">
            <template #icon><UserOutlined /></template>
          </a-avatar>
        </a>
        <template #overlay>
          <a-menu>
            <a-menu-item v-for="(item, index) in data" :key="index">
              <a @click="GoTo(item)">{{ item }}</a>
            </a-menu-item>
          </a-menu>
        </template>
      </a-dropdown>
    </div>
  </a-layout-header>
</template>
<script lang="ts">
import {
  UserOutlined,
  MenuUnfoldOutlined,
  MenuFoldOutlined,
} from '@ant-design/icons-vue'
import { defineComponent, ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { COLLAPSED } from '../../store/mutation-types'

export default defineComponent({
  name: 'CommonHeader',
  components: { UserOutlined, MenuUnfoldOutlined, MenuFoldOutlined },
  setup() {
    const data = ref<string[]>(['个人中心', '退出'])
    const router = useRouter()
    const store = useStore()
    function GoTo(val: string) {
      console.log(val)
      switch (val) {
        case '退出':
          localStorage.clear()
          router.replace('/login')
          break
      }
    }
    function ToggleCollapsed() {
      store.commit(COLLAPSED)
    }
    const crumbs = computed({
      get: () => {
        let r = []
        try {
          if (store.state.crumbs.list) {
            r = store.state.crumbs.list.split(',')
          }
        } catch (err) {
          console.log('err: ', err)
        }
        return r
      },
      set: () => {
        // do
      },
    })
    return {
      // data
      data,
      collapsed: computed(() => store.state.sys.collapsed),
      crumbs,
      // methods
      GoTo,
      ToggleCollapsed,
    }
  },
})
</script>
<style scoped lang="less">
@import 'header.less';
</style>
