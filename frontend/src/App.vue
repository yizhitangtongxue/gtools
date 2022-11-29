<template>
  <div class="common-layout">
    <el-container>
      <!-- <el-aside width="200px"> -->
      <el-aside width="auto">
        <div class="main-menu">
          <!-- <el-radio-group v-model="isCollapse" style="margin-bottom: 20px">
            <el-radio-button :label="false">expand</el-radio-button>
            <el-radio-button :label="true">collapse</el-radio-button>
          </el-radio-group> -->
          <el-menu class="el-menu-vertical-demo" router :collapse="isCollapse" @open="handleOpen" @close="handleClose">
            <el-menu-item index="/">
              <el-icon>
                <location />
              </el-icon>
            </el-menu-item>
            <el-menu-item index="/hello">
              <el-icon>
                <setting />
              </el-icon>
            </el-menu-item>
          </el-menu>
        </div>
      </el-aside>
      <el-main>
        <router-view v-slot="{ Component }">
          <!-- <transition name="fade"> -->
          <keep-alive>
            <component :is="Component" />
          </keep-alive>
          <!-- </transition> -->
        </router-view>
      </el-main>
    </el-container>
  </div>


</template>

<script setup>
import {
  Document,
  Menu as IconMenu,
  Location,
  Setting
} from '@element-plus/icons-vue'
import { ref, onMounted, reactive, watch } from 'vue';
const isCollapse = ref(true)


const handleOpen = (key, keyPath) => {
  console.log(key, keyPath)
}
const handleClose = (key, keyPath) => {
  console.log(key, keyPath)
}

</script>


<style scoped>
/* .main-menu {
  height: 100vh;
  min-height: 100vh;
} */

.main-menu>ul {
  background-color: white;
  border-right: none;
  height: 100vh;
}

.el-main {
  --el-main-padding: none;
  min-height: 100vh;
}

/* .fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
} */
</style>