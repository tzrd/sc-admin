<template>
  <div>

    <common-search
      v-model:text="text"
      placeholder="请输入账号"
      @search="SearchRtn"
    />
    <a-table
      style="margin-top: 15px"
      :columns="tableData.columns"
      :data-source="tableData.data"
      :loading="tableData.loading"
      rowKey="id"
      :pagination="{
        total: tableData.total,
      }"
      @change="Change"
    >
      <template #period="{ text }">
        <div :style="{color: 'red'}">
          {{ text }}
        </div>
      </template>
      <template #depart="{ record }">
        {{ record.depart_info.name }}
      </template>
      <template #action="{ record }">
        <a-button
          v-bt-auth:power
          type="primary"
          style="margin-right: 15px"
          @click="PowerAllocation(record)"
        />

        <a-button
          type="primary"
          style="margin-right: 15px"
          v-bt-auth:edit
          @click="Editor(record)"
        />
        <a-popconfirm
          title="确定删除吗?"
          ok-text="删除"
          cancel-text="取消"
          @confirm="Del(record)"
        >
          <a-button type="danger" v-bt-auth:del />
        </a-popconfirm>
      </template>
    </a-table>

    <common-drawer
      :title="commonDrawerData.title"
      :visible="commonDrawerData.visible"
      :loading="commonDrawerData.loading"
      ok-text="确定"
      cancel-text="取消"
      @on-ok="Submit()"
      @on-close="commonDrawerData.visible = false"
    >
      <a-form :label-col="{ span: 4 }" :wrapper-col="{ span: 20 }">
        <a-form-item label="动态令牌" v-bind="validateInfos.role_name">
          <a-input v-model:value="modelRef.rid"></a-input>
        </a-form-item>
        <a-form-item label="名称">
          <a-input v-model:value="modelRef.token"></a-input>
        </a-form-item>
      </a-form>
    </common-drawer>
    <common-tree
      :visible="allocationTree.visible"
      :data="allocationTree.data"
      :loading="allocationTree.loading"
      @on-close="Close"
      @on-submit="SubmitOk"
    />
  </div>
</template>
<script lang="ts">
import { defineComponent, onMounted, onBeforeMount, reactive, toRaw } from 'vue'
import { useForm } from '@ant-design-vue/use'
import {
  CommonTreeSelectKeys,
  TableDataType,
  TablePaginType,
} from '../../types/type'
import { OtpType } from '../../types/mfa'
import { MenuType } from '../../types/sys'
import { http } from '../../utils/request'
import CommonDrawer, { DrawerProps } from '../../components/Drawer/Drawer.vue'
import CommonSearch from '../../components/Search/Search.vue'
import CommonTree from '../../views/common/tree.vue'
import { message } from 'ant-design-vue'
import { Method } from 'axios'
import { ref } from 'vue'

export interface AllocateType {
  visible: boolean
  loading: boolean
  data?: string[]
  allocateId: string
}

const OptInfo = defineComponent({
  name: 'OptInfo',
  components: {
    CommonSearch,
    CommonDrawer,
    CommonTree,
  },
  setup() {
    const modelRef = reactive<OtpType>({
      id : '',
      rid: '',
      issuer: '',
      account: '',
      period: 0,
      token: '',
    })
    const editId = reactive<{ id: string | undefined }>({ id: '' })
    const commonDrawerData = reactive<DrawerProps>({
      title: '添加',
      loading: false,
      visible: false,
    })
    const allocationTree = reactive<AllocateType>({
      visible: false,
      loading: false,
      data: [],
      allocateId: '',
    })
    const rulesRef = reactive({
      role_name: [
        {
          required: true,
          message: '请输入角色名称',
        },
      ],
    })
    const tableData = reactive<TableDataType<OtpType>>({
      data: [],
      page: 1,
      page_size: 10,
      total: 0,
      loading: false,
      columns: [
        {
          title: '编号',
          key: 'rid',
          dataIndex: 'rid',
        },
        {
          title: '平台',
          key: 'issuer',
          dataIndex: 'issuer',
        },
        {
          title: '账号',
          key: 'account',
          dataIndex: 'account',
        },
        {
          title: '动态口令',
          key: 'token',
          dataIndex: 'token',
        },
        {
          title: '计时(s)',
          key: 'period',
          dataIndex: 'period',
          customRender: () => {
            var date = new Date();
            let scd =  date.getSeconds();
            if (scd > 30) {
              scd = 60 - scd;
            } else {
              scd = 30 - scd;
            }

            return scd;
          }
        },
        {
          title: '操作',
          key: 'action',
          slots: {
            customRender: 'action',
          },
        },
      ],
    })
    
    let timer

    onMounted(() => {
      getList()
      addSetInterval()
    })

    onBeforeMount(()=> {
      clearInterval(timer);
    })

    const { resetFields, validate, validateInfos } = useForm(modelRef, rulesRef)

    let text = ref(null)
    function getList() {
      tableData.loading = true
      http<OtpType>({
        url: 'otp',
        method: 'GET',
        params: {
          q: text.value,
          page: tableData.page,
          page_size: tableData.page_size,
        },
      }).then((res) => {
        tableData.loading = false
        tableData.page = res.page
        tableData.page_size = res.page_size
        tableData.total = res.total
        tableData.data = res.list
      })
    }
    function Submit() {
      validate().then(() => {
        let url = 'otp'
        let method: Method = 'POST'
        const body = toRaw(modelRef)
        commonDrawerData.loading = true
        if (editId.id) {
          url = `otp/${editId.id}`
          method = 'PUT'
        }
        http({
          url,
          method,
          body,
        }).then(() => {
          message.success(`${commonDrawerData.title}成功`)
          commonDrawerData.loading = false
          commonDrawerData.visible = false
          getList()
        })
      })
    }
    function ChangAdd() {
      resetFields()
      commonDrawerData.visible = true
      editId.id = ''
    }
    
    function SearchRtn(value) {
      tableData.loading = true
      http<OtpType>({
        url: 'otp',
        method: 'GET',
        params: {
          q: value, 
          page: tableData.page,
          page_size: tableData.page_size,
        },
      }).then((res) => {
        tableData.loading = false
        tableData.page = res.page
        tableData.page_size = res.page_size
        tableData.total = res.total
        tableData.data = res.list
      })
    }
    function Editor(record: OtpType) {
      if (record.id) {
        editId.id = record.id
        commonDrawerData.title = '修改'
        commonDrawerData.visible = true
        modelRef.rid = record.rid
        modelRef.period = record.period
        modelRef.token = record.token
      }
    }
    function Del(record: OtpType) {
      http({ url: 'otp/' + record.id, method: 'delete' }).then(() => {
        message.success('删除成功');
        getList();
      })
    }
    function Change(pagin: TablePaginType) {
      tableData.page = pagin.current
      getList();
    }
    function PowerAllocation(record: OtpType) {
      allocationTree.loading = true
      allocationTree.visible = true
      if (record.id != null) {
        allocationTree.allocateId = record.id
      }
      http<MenuType>({
        url: '/role/permissions/' + record.id,
        method: 'get',
      }).then((res) => {
        const list: string[] = []
        res.list.forEach((item) => {
          return list.push(item.id || '')
        })
        allocationTree.data = list
        allocationTree.loading = false;
      })
    }
    function Close() {
      allocationTree.visible = false;
    }
    function SubmitOk(val: CommonTreeSelectKeys) {
      const data = {
        role_id: allocationTree.allocateId,
        menu_id: val.checked.join(','),
      }
      allocationTree.loading = true
      http<MenuType>({
        url: '/role/permissions',
        method: 'post',
        body: data,
      })
        .then(() => {
          message.success('更新成功')
          allocationTree.loading = false
          allocationTree.visible = false
        })
        .catch(() => {
          allocationTree.loading = false
        })
    }
    function addSetInterval() {
        if(timer) {
          clearInterval(timer);
        } else {
          timer = setInterval(()=> {
            let scd = new Date().getSeconds();
            if (scd > 30) {
              scd = 60 - scd;
            } else {
              scd = 30 - scd;
            }
            if (scd == 0 || scd == 30) {
              getList();
            } else {
              if (tableData.data) {
                for(var i = 0; i < tableData.data.length; i++) {
                  tableData.data[i]['period'] = scd;
                }
              }
            }
          }, 1000);
        }
    }
    return {
      //data
      text,
      tableData,
      commonDrawerData,
      modelRef,
      allocationTree,

      // methods
      ChangAdd,
      SearchRtn,
      Submit,
      Editor,
      Del,
      Change,
      PowerAllocation,
      Close,
      SubmitOk,

      // form
      validateInfos,
    }
  }
})

export default OptInfo
</script>
