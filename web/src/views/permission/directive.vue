<script lang="ts" setup>
</script>

<template>
  <div class="app-container">
    <!-- v-permission 示例 -->
      <div>
        <el-tag v-permission="['admin']" type="success" size="large" effect="plain">
          这里采用了 v-permission="['admin']" 所以只有 admin 可以看见这句话
        </el-tag>
      </div>
      <div>
        <el-tag v-permission="['普通用户']" type="success" size="large" effect="plain">
          这里采用了 v-permission="['普通用户']" 所以只有 普通用户 可以看见这句话
        </el-tag>
      </div>
      <div class="margin-top-15">
        <el-tag v-permission="['admin', '普通用户']" type="success" size="large" effect="plain">
          这里采用了 v-permission="['admin', '普通用户']" 所以 admin 和 普通用户 都可以看见这句话
        </el-tag>
      </div>
    </div>
</template>

<style lang="scss" scoped>
.margin-top-15 {
  margin-top: 15px;
}
.margin-top-30 {
  margin-top: 30px;
}
</style>
