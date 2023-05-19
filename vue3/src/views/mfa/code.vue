<template>
 <el-container>
  <el-main>
    <el-row>
        <el-row v-for="item in accounts" :key="item.secret">
          <!-- item -->
          <el-card>
            <!-- issuer -->
            <div class="item-issuer">{{item.issuer}}</div>

            <el-row type="flex" align="middle">
              <!-- counter -->
              <el-col :span="4">
                <el-progress type="circle" :percentage="item.percentage" :color="item.counterColor" :stroke-width="5" :width="36" :show-text="false" status="text">
                </el-progress>
              </el-col>

              <!-- token -->
              <el-col :span="14">
                <div class="item-token">{{item.token}}</div>
              </el-col>

              <!-- copy -->
              <el-col :span="3">
                <el-button class="copy" type="info" size="mini" icon="el-icon-document-copy" circle @click="doCopy(item.token)" />
              </el-col>

              <!-- delete -->
              <el-col :span="3">
                <el-button type="danger" size="mini" icon="el-icon-delete" circle @click="doDelete(item)" />
              </el-col>
            </el-row>

            <!-- account -->
            <div class="item-account">{{item.account}}</div>
          </el-card>
        </el-row>
    </el-row>
  </el-main>
 </el-container>
</template>

<script>
export default {
    data: function () {
    let accounts = 
      [
        {
          account: "example@gmail.com",
          counter: 26,
          counterColor: "#E6A23C",
          digits: 6,
          issuer: "Google",
          percentage: 50,
          period: 30,
          token: 782908
        }
      ]

    return {
      accounts: accounts,
      timer: setInterval(() => {
        this.countersUpdate()
      }, 1000)
    }
  },
  methods: {
    countersUpdate () {
      const accounts = 
        [
          {
            account: "example@gmail.com",
            counter: 26,
            counterColor: "#E6A23C",
            digits: 6,
            issuer: "Google",
            percentage: 100,
            period: 30,
            token: 987902
          }
        ];
      this.accounts = accounts
    },
    doDelete: function (item) {
      this.$confirm('你确定删除该账号的动态令牌?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        const accounts = 
        [
          {
            account: "example@gmail.com",
            counter: 26,
            counterColor: "#E6A23C",
            digits: 6,
            issuer: "Google",
            percentage: 100,
            period: 30,
            token: 987902
          }
        ];
        //deleteAccount(item)
        this.accounts = accounts
        this.$message({
          type: 'success',
          message: '删除成功!'
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '取消删除'
        })
      })
    }
  },
  unmounted () {
    clearInterval(this.timer)
  }
}
</script>
<style>
.el-card {
  width: 360px;
}

.el-row {
  margin-bottom: 20px;
}

.item-issuer{
  color:#606266;
  margin-bottom: 15px;
}

.item-token{
  color: #409eff;
  font-size: 36pt;
}

.item-account{
  color:#909399;
}

.select{
  margin-top: 20px;
}

.select-tip{
  font-size: 10pt;
  color:#909399;
  padding-left: 5px;
  padding-right: 5px;
}
</style>