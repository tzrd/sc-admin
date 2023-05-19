<template>
    <a-input-search 
        class="inputSearch"
        :value="text"
        @input="$emit('update:text', $event.target.value)"
        :placeholder="placeholder"
        enter-button @search="onSearch" 
    />
</template>
<script lang="ts">
import { defineComponent } from 'vue'

const CommonSearch = defineComponent({
  name: 'CommonSearch',
  props: {
    placeholder: {
      type: String,
      required: true,
      default: '默认',
    },
    text: {
      type: String,
      default: ''
    }
  },
  setup(props, { emit }) {
    function onSearch() {
      emit('search')
    }
    return {
      CommonSearch,
    }
  },
})

export default CommonSearch
</script>
<style lang="less" scoped>
@import './Search.less';
</style>