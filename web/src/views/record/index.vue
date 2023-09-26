<template>
  <div class="app-container">

    <!-- 表格 -->
    <vxe-grid ref="xGridDom" v-bind="xGridOpt">

      <template #form>
        <el-card v-loading="loading" shadow="never" class="search-wrapper">
          <el-form ref="searchFormRef" :inline="true" :model="searchFormData" @keyup.enter="crudStore.commitQuery">
            <el-form-item prop="menu">
              <el-select
                v-model="searchFormData.menu"
                filterable
                multiple
                collapse-tags
                collapse-tags-tooltip
                default-first-option
                :reserve-keyword="false"
                :max-collapse-tags="1"
                style="width: 230px"
                placeholder="操作菜单"
                :clearable="true"
              >
                <el-option v-for="item in menuOptions" :key="item.value" :label="item.label" :value="item.value"/>
              </el-select>
            </el-form-item>

            <el-form-item prop="username">
              <el-select
                v-model="searchFormData.username"
                filterable
                collapse-tags
                collapse-tags-tooltip
                :max-collapse-tags="1"
                placeholder="操作人"
                style="width: 230px"
                :clearable="true"
              >
                <el-option v-for="item in usernameOptions" :key="item.value" :label="item.label" :value="item.value"/>
              </el-select>
            </el-form-item>


            <el-form-item prop="date_range">
              <el-date-picker
                v-model="searchFormData.timeRange"
                type="datetimerange"
                :shortcuts="shortcuts"
                range-separator="~"
                start-placeholder="开始时间"
                end-placeholder="结束时间"
              />
            </el-form-item>

            <el-form-item prop="method">
              <el-select
                v-model="searchFormData.method"
                filterable
                multiple
                collapse-tags
                collapse-tags-tooltip
                :max-collapse-tags="2"
                placeholder="操作类型"
                style="width: 230px"
                :clearable="true"
              >
                <el-option v-for="item in methodOptions" :key="item.value" :label="item.label"
                           :value="item.value"/>
              </el-select>
            </el-form-item>

            <el-form-item prop="status">
              <el-select
                v-model="searchFormData.status"
                filterable
                multiple
                collapse-tags
                collapse-tags-tooltip
                default-first-option
                :reserve-keyword="false"
                :max-collapse-tags="2"
                placeholder="状态码"
                :clearable="true"
              >
                <el-option v-for="item in statusOptions" :key="item.value" :label="item.label" :value="item.value"/>
              </el-select>
            </el-form-item>


            <el-form-item>
              <el-button type="primary" icon="Search" @click="crudStore.commitQuery">查询</el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </template>


      <!--      展开行的内容-->
      <template #expand-content="{ row }">
        <el-tabs model-value="request" style="margin-left: 20px">
          <el-tab-pane label="请求信息" name="request">
            <div class="box-row">
              <vue-json-pretty :data="JSON.parse(row.reqParam)"
                               :showLine="true"
                               class="vjs-value"
                               virtual
                               showIcon
                               :height="300"
                               style="white-space: pre-wrap;"
              >
                <template #>

                </template>
              </vue-json-pretty>
            </div>
          </el-tab-pane>
          <el-tab-pane label="响应信息" name="response">
            <div class="box-row">
              <vue-json-pretty :data="JSON.parse(row.respData)"
                               :showLine="true"
                               virtual
                               showIcon
                               :height="300"
              />
            </div>
          </el-tab-pane>
        </el-tabs>
      </template>
    </vxe-grid>
  </div>
</template>

<script lang="ts" setup>
import {reactive, ref} from "vue"
import {formatDateTime, formatTimestamp} from "@/utils/index"
import TypeColumnSolts from "./tsx/TypeColumnSolts"
import {type VxeGridInstance, type VxeGridProps, type VxeGridPropTypes} from "vxe-table"
import {shortcuts} from "@/config/dateDict"
import {GetLicenseResponseData, RowMeta} from "@/api/license/license"
import {getRecordApi, getSearchOptionsApi} from "@/api/system/operationRecord";
import VueJsonPretty from "vue-json-pretty"
import "vue-json-pretty/lib/styles.css"
import {selectOption} from "@/api/customer";

defineOptions({
  name: "Record"
})


//#region searchfrom
const loading = ref(false)
const searchFormData = reactive({
  menu: [],
  username: "",
  method: [],
  status: [],
  timeRange: ""
})

const menuOptions = ref<selectOption[]>([])
const usernameOptions = ref<selectOption[]>([])
const methodOptions = [
  {value: "POST", label: "新增"},
  {value: "PUT", label: "修改"},
  {value: "DELETE", label: "删除"}
]
const statusOptions = ref<selectOption[]>([])

const getSearchOptions = async () => {
  try {
    const res = await getSearchOptionsApi()
    menuOptions.value = res.data.list?.menu
    usernameOptions.value = res.data.list?.user_name
    statusOptions.value = res.data.list?.status
  } catch (error) {
    console.log(error)
  }
}

getSearchOptions()

//#endregion

//#region vxe-grid
const xGridDom = ref<VxeGridInstance>()

const xGridOpt: VxeGridProps = reactive({
  /** 展开行 */
  expandConfig: {
    iconColumnIndex: 0, // 展开图标所在的列索引
    expandAll: false, // 是否默认展开所有行
    expandRowKeys: [], // 默认展开的行的 key
    accordion: false // 是否每次只能展开一行
  },
  loading: true,
  autoResize: true,
  /** 分页配置项 */
  pagerConfig: {
    align: "right"
  },
  /** 工具栏配置 */
  toolbarConfig: {
    refresh: true,
    custom: false,
  },
  /** 列配置 */
  columns: [
    {
      type: "expand", // 设置列类型为展开列
      width: 40, // 设置展开列的宽度
      slots: {content: "expand-content"}
    },
    {
      field: "menu",
      title: "操作菜单"
    },
    {
      field: "username",
      title: "操作人"
    },
    {
      field: "method",
      title: "操作类型",
      slots: TypeColumnSolts
    },
    {
      field: "uri",
      title: "uri"
    },
    {
      field: "status",
      title: "状态码",
    },
    {
      field: "responseTime",
      title: "响应时长(ms)",
    },
    {
      field: "CreatedAt",
      title: "操作时间",
      formatter: ({row}) => {
        return formatDateTime(row.CreatedAt)
      }
    },
  ],
  /** 数据代理配置项（基于 Promise API） */
  proxyConfig: {
    /** 启用动态序号代理 */
    seq: true,
    /** 是否代理表单 */
    form: true,
    /** 是否自动加载，默认为 true */
    // autoLoad: false,
    props: {
      total: "total"
    },
    ajax: {
      query: ({page, form}: VxeGridPropTypes.ProxyAjaxQueryParams) => {
        xGridOpt.loading = true
        crudStore.clearTable()
        return new Promise<any>((resolve: Function) => {
          let total = 0
          let result: RowMeta[] = []
          /** 加载数据 */
          const callback = (res: GetLicenseResponseData) => {
            if (res && res.data) {
              const resData = res.data
              // 总数
              if (Number.isInteger(resData.total)) {
                total = resData.total
              }
              // 分页数据
              if (Array.isArray(resData.list)) {
                result = resData.list
              }
            }
            xGridOpt.loading = false
            resolve({total, result})
          }

          let startTimestamp
          let endTimestamp
          if (
            searchFormData.timeRange != null &&
            typeof searchFormData.timeRange === "object" &&
            Object.keys(searchFormData.timeRange).length === 2
          ) {
            let s = searchFormData.timeRange[0]
            let e = searchFormData.timeRange[1]
            startTimestamp = formatTimestamp(s)
            endTimestamp = formatTimestamp(e)
          }

          /** 接口需要的参数 */
          const params = {
            method: searchFormData.method.join(",") || undefined,
            menu: searchFormData.menu.join(",") || undefined,
            username: searchFormData.username || undefined,
            startTime: startTimestamp || undefined,
            endTime: endTimestamp || undefined,
            status: searchFormData.status.join(",") || undefined,
            size: page.pageSize,
            currentPage: page.currentPage
          }
          /** 调用接口 */
          getRecordApi(params).then(callback).catch(callback)
        })
      }
    }
  }
})
//#endregion

//#region CRUD
const crudStore = reactive({
  /** 加载表格数据 */
  commitQuery: () => xGridDom.value?.commitProxy("query"),
  /** 清空表格数据 */
  clearTable: () => xGridDom.value?.reloadData([]),
})
//#endregion
</script>

<style lang="scss" scoped>
.box-row {
  height: 313px;
  background-color: #F9F9F9;
}

.search-wrapper {
  margin-bottom: 20px;

  :deep(.el-card__body) {
    padding-bottom: 2px;
  }
}
</style>

