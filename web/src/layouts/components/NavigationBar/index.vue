<template>
  <div class="navigation-bar">
    <Hamburger :is-active="sidebar.opened" class="hamburger" @toggle-click="toggleSidebar" />
    <Breadcrumb class="breadcrumb" />
    <div class="right-menu">
      <Screenfull v-if="showScreenfull" class="right-menu-item" />
      <ThemeSwitch v-if="showThemeSwitch" class="right-menu-item" />
      <!--      系统设置-->
      <el-dropdown class="right-menu-item">
        <div class="right-menu-avatar">
          <el-icon :size="20">
            <Setting />
          </el-icon>
        </div>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item v-for="r in Routers" @click="toRoute(r.component)" style="width: 120px">
              {{ r.label }}
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
      <!--     用户中心 -->
      <el-dropdown class="right-menu-item">
        <div class="right-menu-avatar">
          <el-avatar :icon="UserFilled" :size="20" />
          <span>{{ userStore.username }}</span>
        </div>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item @click="toRoute('Profile')">个人中心</el-dropdown-item>
            <el-dropdown-item divided @click="logout">
              <span style="display: block">退出登录</span>
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useRouter } from "vue-router"
import { storeToRefs } from "pinia"
import { useAppStore } from "@/store/modules/app"
import { useUserStore } from "@/store/modules/user"
import Breadcrumb from "../Breadcrumb/index.vue"
import Hamburger from "../Hamburger/index.vue"
import ThemeSwitch from "@/components/ThemeSwitch/index.vue"
import Screenfull from "@/components/Screenfull/index.vue"
import { joinInBlacklistApi } from "@/api/system/jwt"
import { useSettingsStore } from "@/store/modules/settings"
import { UserFilled } from "@element-plus/icons-vue"

const router = useRouter()
const appStore = useAppStore()
const settingsStore = useSettingsStore()
const userStore = useUserStore()

const { sidebar } = storeToRefs(appStore)
const { showThemeSwitch, showScreenfull } = storeToRefs(settingsStore)

/* 跳转路由 */
const Routers = [
  { label: "用户管理", component: "User" },
  { label: "角色管理", component: "Role" },
  { label: "菜单管理", component: "Menu" },
  { label: "API 管理", component: "Api" }
]

/** 切换侧边栏 */
const toggleSidebar = () => {
  appStore.toggleSidebar(false)
}

const logout = () => {
  // userStore.logout()
  // router.push("/login")
  joinInBlacklistApi()
    .then(() => {
      userStore.logout()
      router.push("/login")
    })
    .catch(() => {})
}

const toRoute = (componentName: string) => {
  router.push({ name: componentName })
}
</script>

<style lang="scss" scoped>
.navigation-bar {
  height: var(--base-navigationbar-height);
  overflow: hidden;
  background: #fff;

  .hamburger {
    display: flex;
    align-items: center;
    height: 100%;
    float: left;
    padding: 0 15px;
    cursor: pointer;
  }

  .breadcrumb {
    float: left;
    // 参考 Bootstrap 的响应式设计将宽度设置为 576
    @media screen and (max-width: 576px) {
      display: none;
    }
  }

  .right-menu {
    float: right;
    margin-right: 10px;
    height: 100%;
    display: flex;
    align-items: center;
    color: #606266;

    .right-menu-item {
      padding: 0 10px;
      cursor: pointer;

      .right-menu-avatar {
        display: flex;
        align-items: center;

        .el-avatar {
          margin-right: 10px;
        }

        span {
          font-size: 16px;
        }
      }
    }
  }
}

.el-button.is-plain {
  --el-button-hover-border-color: #ffffff;
}
</style>
