import { type VxeColumnPropTypes } from "vxe-table/types/column"

const solts: VxeColumnPropTypes.Slots = {
  default: ({ row, column }) => {
    const cellValue = row[column.field]
    let type = "success"
    let value = cellValue
    switch (cellValue) {
      case "PUT":
        type = "warning"
        value = "修改"
        break
      case "DELETE":
        type = "danger"
        value = "删除"
        break
      case "POST":
        type = "success"
        value = "新增"
        break
    }
    return [<span class={`el-tag el-tag--${type} el-tag--plain`}>{value}</span>]
  }
}

export default solts
