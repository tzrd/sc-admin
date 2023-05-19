<template>
  <div>
    <common-button
      v-bt-auth:add="{ title: true }"
      title="添加"
      icon-name="add"
      @change="ChangAdd"
    />
    <common-search
      v-model:text="text"
      placeholder="请输入账号ID或账号"
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
      <template #status="{ text }">
        <div :style="text === 0 ? { color: 'red' } : { color: 'green' }">
          {{ formatStatus(text) }}
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
          @click="Allocate(record)"
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
          <a-button v-bt-auth:del type="danger"/>
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
        <a-form-item label="Qr识别">
          <a-upload
            action="#"
            list-type="picture"
            :default-file-list="fileList"
            :beforeUpload="beforeUpload"
            @change="handleChange"
          >
            <a-button> 上传Qr图片 </a-button>
          </a-upload>
        </a-form-item>
        <a-form-item label="平台" v-bind="validateInfos.issuer">
          <a-input v-model:value="infoRef.issuer"></a-input>
        </a-form-item>
         <a-form-item label="主账号ID" v-bind="validateInfos.uid">
          <a-input v-model:value="infoRef.uid"></a-input>
        </a-form-item>       
        <a-form-item label="账号" v-bind="validateInfos.account">
          <a-input v-model:value="infoRef.account"></a-input>
        </a-form-item>
        <a-form-item label="位数" v-bind="validateInfos.digits">
          <a-input v-model:value="infoRef.digits"></a-input>
        </a-form-item>
        <a-form-item label="周期" v-bind="validateInfos.period">
          <a-input v-model:value="infoRef.period"></a-input>
        </a-form-item>        
        <a-form-item label="名称" v-bind="validateInfos.name">
          <a-input v-model:value="infoRef.name"></a-input>
        </a-form-item>
        <a-form-item label="QR码" v-bind="validateInfos.qrcode">
          <a-input v-model:value="infoRef.qrcode"></a-input>
        </a-form-item>
        <a-form-item label="密钥" v-bind="validateInfos.secret">
          <a-input v-model:value="infoRef.secret"></a-input>
        </a-form-item>
        <a-form-item label="备注" v-bind="validateInfos.desc">
          <a-input v-model:value="infoRef.desc"></a-input>
        </a-form-item> 
        <a-form-item label="状态" v-bind="validateInfos.status">
          <a-radio-group name="radioGroup" v-model:value="infoRef.status">
            <a-radio :value="0"> 失效 </a-radio>
            <a-radio :value="1"> 有效 </a-radio>
          </a-radio-group>
        </a-form-item>
      </a-form>
    </common-drawer>
    <common-drawer
      title="MFA列表"
      ok-text="确定"
      cancel-text="取消"
      :visible="allocationTree.visible"
      :loading="allocationTree.loading"
      @on-close="Close"
      @on-ok="SubmitOk"
    >
      <a-spin :spinning="allocationTree.loading">
        <a-checkbox-group
          v-model:value="selectkeys"
          v-if="roleList.length > 0"
          class="checkoutContainer">
          <div v-for="(item, index) in roleList" :key="index">
            <a-checkbox :value="item.id"> {{ item.remark }} </a-checkbox>
          </div>
        </a-checkbox-group>
        <a-empty v-else />
      </a-spin>
    </common-drawer>
  </div>
</template>
<script lang="ts">
import { defineComponent, onMounted, reactive, toRaw, ref } from 'vue'
import { useForm } from '@ant-design-vue/use'
import { TableDataType, TablePaginType } from '../../types/type'
import { DepartType, RoleType, UserType } from '../../types/sys'
import { MfaType, MfaInfoType } from '../../types/mfa'
import { http } from '../../utils/request'
import { ParseQrPictureFile } from '../../utils/qrcode'
import CommonDrawer, { DrawerProps } from '../../components/Drawer/Drawer.vue'
import CommonButton from '../../components/Button/Button.vue'
import CommonSearch from '../../components/Search/Search.vue'
import { message } from 'ant-design-vue'
import { Method } from 'axios'
import { AllocateType } from '../../views/sys/role.vue'

interface UserAndRole {
  user_id: string
  role_id: string
}
const MfaInfo = defineComponent({
  name: 'MfaInfo',
  components: {
    CommonSearch,
    CommonButton,
    CommonDrawer,
  },
  setup() {
    const selectkeys = ref<string[]>([])
    const roleList = ref<RoleType[]>([])
    const allocationTree = reactive<AllocateType>({
      visible: false,
      loading: false,
      allocateId: '',
    })
    const modelRef = reactive<MfaType>({
      id: '',
      uid: '',
      issuer: '',
      account: '',
      name: '',
      desc: '',
      status: 1,
    })
    const infoRef = reactive<MfaInfoType>({
      id: '',
      uid: '',
      issuer: '',
      account: '',
      digits: 0,
      period: 0,
      qrcode: '',
      secret: '',
      name: '',
      desc: '',
      status: 1,
    })
    const issuerSelectData = [
          {"name": "阿里云", "value": "阿里云"}, 
          {"name": "腾讯云", "value": "腾讯云"}, 
          {"name": "亚马逊", "value": "亚马逊"}, 
          {"name": "谷歌云", "value": "谷歌云"}, 
          {"name": "优刻得", "value": "优刻得"},
          {"name": "Cloudflare", "value": "Cloudflare"},
    ];
    const treeOptions = reactive<{ options: DepartType[] }>({ options: [] })
    const editId = reactive<{ id: string | undefined }>({ id: '' })
    const commonDrawerData = reactive<DrawerProps>({
      title: '添加',
      loading: false,
      visible: false,
    })
    const uploadImageData = reactive({
      fileList: [],
      loading: false,
      imageUrl: '',
      action: '',
    })
    const rulesRef = reactive({
      uid: [
        {
          required: true,
          message: '请输入主账号ID',
        },
      ],      
      issuer: [
        {
          required: true,
          message: '请输入云平台',
        },
      ],
      account: [
        {
          required: true,
          message: '请输入账号',
        },
      ],
      name: [
        {
          required: true,
          message: '请输入名称',
        },
      ],
      secret: [
        {
          required: true,
          message: '请输入密钥',
        },
      ],
    })
    const tableData = reactive<TableDataType<MfaType>>({
      data: [],
      page: 1,
      page_size: 10,
      total: 0,
      loading: false,
      columns: [
        {
          title: '云平台',
          key: 'issuer',
          dataIndex: 'issuer',
        },
        {
          title: '账号ID',
          key: 'uid',
          dataIndex: 'uid',
        },        
        {
          title: '账号',
          key: 'account',
          dataIndex: 'account',
        },
        {
          title: '名称',
          key: 'name',
          dataIndex: 'name',
        },
        {
          title: '描述',
          key: 'desc',
          dataIndex: 'desc',
        },
        {
          title: '状态',
          key: 'status',
          dataIndex: 'status',
          slots: {
            customRender: 'status',
          },
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

let text = ref(null)
    function getList() {
      tableData.loading = true
      http<MfaType>({
        url: 'mfa',
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

    function getRoleList() {
      http<RoleType>({
        url: '/role',
        method: 'get',
        params: {
          page: 1,
          page_size: 1000,
        },
      }).then((res) => {
        roleList.value = res.list
      })
    }
    onMounted(() => {
      getList()
      getRoleList()
    })

    function formatStatus(text: number): string {
      let result = ''
      switch (text) {
        case 0:
          result = '失效'
          break
        case 1:
          result = '有效'
          break
        default:
          result = '未知'
      }
      return result
    }

    const { resetFields, validate, validateInfos } = useForm(infoRef, rulesRef)
    function Submit() {
      validate().then(() => {
        let url = 'mfa'
        let method: Method = 'POST'
        const data = toRaw(infoRef)
        commonDrawerData.loading = true
        if (editId.id) {
          url = `mfa/${editId.id}`
          method = 'PUT'
        }
        http({
          url,
          method,
          body: data,
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
      http<MfaType>({
        url: 'mfa',
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

    function beforeUpload(file) {
      const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png'
      if (!isJpgOrPng) {
        message.error('请上传jpeg或png格式图片')
      }
      const isLt2M = file.size / 1024 / 1024 < 2
      if (!isLt2M) {
        message.error('图片上传最大为2M')
      }

      return false
    }

    function handleChange(info) {      
      let fileList = [...info.fileList];
      fileList = fileList.slice(-1);
      if(fileList.length > 0) {
        ParseQrPictureFile(fileList[0], setAddForm);
      }
    }

    function setAddForm(qr) {
      if (qr && !qr['error']) {
        infoRef.issuer = qr['data']['issuer'];
        infoRef.account = qr['data']['account'];
        infoRef.digits = qr['data']['digits'];
        infoRef.period = qr['data']['period'];
        infoRef.qrcode = qr['data']['qrcode'];
        infoRef.secret = qr['data']['secret'];
      }
    }

    function Editor(record: MfaInfoType) {
      if (record.id) {
        editId.id = record.id
        commonDrawerData.title = '修改'
        commonDrawerData.visible = true
        infoRef.uid = record.uid
        infoRef.issuer = record.issuer
        infoRef.account = record.account
        infoRef.name = record.name
        infoRef.desc = record.desc
        infoRef.status = record.status
        
        if (record.digits) {
          infoRef.digits = record.digits
        }

        if (record.period) {
          infoRef.period = record.period
        }

        if (record.qrcode) {
          infoRef.qrcode = record.qrcode
        }

        if (record.secret) {
          infoRef.secret = record.secret
        }
      }
    }

    function Del(record: MfaType) {
      http({ url: 'mfa/' + record.id, method: 'delete' }).then(() => {
        message.success('删除成功')
        getList()
      })
    }

    function Change(pagin: TablePaginType) {
      tableData.page = pagin.current
      getList()
    }

    function Allocate(record: MfaType) {
      allocationTree.loading = true
      allocationTree.visible = true
      if (record.id != null) {
        allocationTree.allocateId = record.id
      }
      http<RoleType>({
        url: '/user/permissions/' + record.id,
        method: 'get',
      }).then((res) => {
        const list: string[] = []
        res.list.forEach((item) => {
          if (item.id != null) {
            list.push(item.id)
          }
        })
        selectkeys.value = list
        allocationTree.loading = false
      })
    }

    function Close() {
      allocationTree.visible = false
    }

    function SubmitOk() {
      const data: UserAndRole = {
        user_id: allocationTree.allocateId,
        role_id: selectkeys.value.join(','),
      }
      allocationTree.loading = true
      http<UserType>({
        url: '/user/permissions',
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
    
    return {
      //data
      text,
      tableData,
      commonDrawerData,
      modelRef,
      infoRef,
      treeOptions,
      uploadImageData,
      allocationTree,
      selectkeys,
      roleList,
      issuerSelectData,

      // methods
      ChangAdd,
      SearchRtn,
      formatStatus,
      Submit,
      beforeUpload,
      handleChange,
      Editor,
      Del,
      Change,
      Allocate,
      Close,
      SubmitOk,
      // form
      validateInfos,
    }
  },
})

export default MfaInfo
</script>
<style scoped lang="less">
.checkoutContainer {
  display: flex;
  flex: 1;
  flex-direction: row;
  flex-wrap: wrap;
}
</style>
