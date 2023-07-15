<template>
  <div class="app-container">
    <el-card v-loading="loading" shadow="never" class="search-wrapper">
      <el-form ref="searchFormRef" :inline="true" :model="searchFormData">
        <el-form-item prop="path">
          <el-input v-model="searchFormData.path" placeholder="路径"/>
        </el-form-item>
        <el-form-item prop="group">
          <el-select
            v-model="searchFormData.api_group"
            filterable
            multiple
            collapse-tags
            collapse-tags-tooltip
            :max-collapse-tags="2"
            style="width: 230px"
            placeholder="API分组"
            :clearable="true"
          >
            <el-option v-for="item in ApiGroupOptions" :key="item" :label="String(item)" :value="item"/>
          </el-select>
        </el-form-item>
        <el-form-item prop="method">
          <el-select
            v-model="searchFormData.method"
            filterable
            multiple
            collapse-tags
            collapse-tags-tooltip
            :max-collapse-tags="1"
            placeholder="方法"
            :clearable="true"
          >
            <el-option v-for="item in methodOptions" :key="item.value" :label="item.label" :value="item.value"/>
          </el-select>
        </el-form-item>
        <el-form-item prop="description">
          <el-input v-model="searchFormData.description" placeholder="API描述"/>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="Search" @click="handleSearch">搜索</el-button>
        </el-form-item>
      </el-form>
    </el-card>
    <el-card v-loading="loading" shadow="never">
      <div class="toolbar-wrapper">
        <div>
          <el-button type="primary" icon="CirclePlus" @click="addDialog">新增</el-button>
        </div>
        <div>
          <el-tooltip content="刷新" effect="light">
            <el-button type="primary" icon="RefreshRight" circle plain @click="getTableData"/>
          </el-tooltip>
        </div>
      </div>
      <div class="table-wrapper">
        <el-table :data="tableData" @sort-change="handleSortChange">
          <el-table-column prop="ID" label="ID"/>
          <el-table-column prop="path" label="路径" sortable="custom"/>
          <el-table-column prop="api_group" label="分组" sortable="custom"/>
          <el-table-column prop="method" label="请求方法" sortable="custom"/>
          <el-table-column prop="description" label="描述"/>
          <el-table-column label="操作">
            <template #default="scope">
              <el-button type="primary" text icon="Edit" size="small" @click="editDialog(scope.row)">编辑</el-button>
              <el-button type="danger" text icon="Delete" size="small" @click="handleDeleteApi(scope.row)"
              >删除
              </el-button
              >
            </template>
          </el-table-column>
        </el-table>
      </div>
      <div class="pager-wrapper">
        <el-pagination
          background
          :layout="paginationData.layout"
          :page-sizes="paginationData.pageSizes"
          :total="paginationData.total"
          :page-size="paginationData.pageSize"
          :currentPage="paginationData.currentPage"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>
    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      :before-close="handleClose"
      :close-on-click-modal="false"
      width="38%"
    >
      <warning-bar title="新增接口，需要在角色管理内配置权限才可使用"/>
      <el-form ref="formRef" :model="opFormData" :rules="addFormRules" label-width="80px">
        <el-form-item label="API路径" prop="path">
          <el-input v-model="opFormData.path"/>
        </el-form-item>
        <el-form-item prop="method" label="请求方法">
          <el-select
            v-model="opFormData.method"
            :disabled="opFormData.isDisabled"
            filterable
            multiple
            collapse-tags
            collapse-tags-tooltip
            :max-collapse-tags="4"
            placeholder="请选择方法"
            style="width: 100%"
            :clearable="true"
          >
            <el-option v-for="item in methodOptions" :key="item.value" :label="item.label" :value="item.value"/>
          </el-select>
        </el-form-item>
        <el-form-item label="API分组" prop="api_group">
          <el-input v-model="opFormData.api_group"/>
        </el-form-item>
        <el-form-item label="API描述" prop="description">
          <el-input v-model="opFormData.description"/>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取消</el-button>
          <el-button type="primary" @click="operateAction(formRef)">确认</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import {reactive, ref} from "vue"
import {type FormInstance, type FormRules, ElMessage, ElMessageBox} from "element-plus"
import {usePagination} from "@/hooks/usePagination"
import {type ApiData, getApisApi, addApiApi, deleteApiApi, editApiApi, getApiGroups} from "@/api/system/api"
import WarningBar from "@/components/WarningBar/warningBar.vue"

defineOptions({
  name: "Api"
})

const {paginationData, changeCurrentPage, changePageSize} = usePagination()

const loading = ref(false)
const searchFormData = reactive({
  path: "",
  api_group: [],
  method: [],
  description: "",
  orderKey: "",
  // 默认升序
  desc: false
})

const methodOptions = [
  {value: "GET", label: "GET"},
  {value: "POST", label: "POST"},
  {value: "PUT", label: "PUT"},
  {value: "DELETE", label: "DELETE"}
]

const handleSearch = () => {
  paginationData.currentPage = 1
  paginationData.pageSize = 10
  getTableData()
}


const tableData = ref<ApiData[]>([])
const ApiGroupOptions = ref<any[]>([])

const getTableData = async () => {
  loading.value = true
  try {
    const res = await getApisApi({
      path: searchFormData.path || undefined,
      api_group: searchFormData.api_group.join(",") || undefined,
      method: searchFormData.method.join(",") || undefined,
      description: searchFormData.description || undefined,
      orderKey: searchFormData.orderKey || undefined,
      desc: searchFormData.desc || undefined,
      page: paginationData.currentPage,
      pageSize: paginationData.pageSize
    })

    tableData.value = res.data.list
    paginationData.total = res.data.total

    const res2 = await getApiGroups()
    ApiGroupOptions.value = res2.data.list
  } catch (error) {
    console.log(error)
  }
  loading.value = false
}
getTableData()

// 分页
const handleSizeChange = (value: number) => {
  changePageSize(value)
  getTableData()
}

const handleCurrentChange = (value: number) => {
  changeCurrentPage(value)
  getTableData()
}

// 排序
const handleSortChange = (column: any) => {
  searchFormData.orderKey = column.prop
  if (column.order === "descending") {
    searchFormData.desc = true
  } else {
    searchFormData.desc = false
  }
  getTableData()
}

// 对话框
const formRef = ref<FormInstance>()
const opFormData = reactive({
  path: "",
  api_group: "",
  method: [],
  description: "",
  isDisabled: false
})

enum operationKind {
  Add = "Add",
  Edit = "Edit"
}

let oKind: operationKind
const addFormRules: FormRules = reactive({
  path: [{required: true, trigger: "blur", message: "路径不能为空"}],
  api_group: [{required: true, trigger: "blur", message: "分组不能为空"}],
  method: [{required: true, trigger: "change", message: "方法不能为空"}],
  description: [{required: true, trigger: "blur", message: "描述不能为空"}]
})

const initForm = () => {
  formRef.value?.resetFields()
  opFormData.path = ""
  opFormData.api_group = ""
  opFormData.method = []
  opFormData.description = ""
  opFormData.isDisabled = false
}

const dialogVisible = ref(false)
const dialogTitle = ref("")
const handleClose = (done: Function) => {
  initForm()
  done()
}

const addDialog = () => {
  dialogTitle.value = "新增接口"
  oKind = operationKind.Add
  dialogVisible.value = true
}

const closeDialog = () => {
  dialogVisible.value = false
  initForm()
}

const operateAction = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate(async (valid) => {
    if (valid) {
      const methodString = opFormData.method.join(",") // 将 string 数组转换为以逗号分隔的字符串
      if (oKind === "Add") {
        const res = await addApiApi({...opFormData, method: methodString})
        if (res.code === 0) {
          ElMessage({type: "success", message: res.msg})
        }
      } else if (oKind === "Edit") {
        const res = await editApiApi({id: activeRow.ID, ...opFormData, method: methodString})
        if (res.code === 0) {
          ElMessage({type: "success", message: res.msg})
        }
      }
      // 由于把method提成了一个数组，所以直接重新请求接口，而不是手动修改数据
      getTableData()
      closeDialog()
    }
  })
}

// 删除api
const handleDeleteApi = (row: ApiData) => {
  ElMessageBox.confirm("此操作将永久删除所有角色下该api, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  })
    .then(() => {
      deleteApiApi({id: row.ID}).then((res) => {
        if (res.code === 0) {
          ElMessage({type: "success", message: res.msg})
          const index = tableData.value.indexOf(row)
          tableData.value.splice(index, 1)
        }
      })
    })
    .catch(() => {
    })
}

// 编辑dialog
let activeRow: ApiData
const editDialog = (row: ApiData) => {
  opFormData.isDisabled = true
  dialogTitle.value = "编辑接口"
  oKind = operationKind.Edit
  opFormData.api_group = row.api_group
  opFormData.description = row.description
  opFormData.method.push(row.method)
  opFormData.path = row.path
  activeRow = row
  dialogVisible.value = true
}
</script>

<style lang="scss" scoped>
.search-wrapper {
  margin-bottom: 20px;

  :deep(.el-card__body) {
    padding-bottom: 2px;
  }
}
</style>
