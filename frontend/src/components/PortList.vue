<template>
    <div class="main-container">
        <div class="mt-4 port-number-container" style="height: 32px;border: none;">
            <el-input v-model="portNumber" placeholder="请输入端口号查询相关进程" class="input-with-select"
                @keyup.enter.native="onSearch">
                <template #append>
                    <el-button :icon="Search" @click="onSearch" />
                </template>
            </el-input>
        </div>
        <div class="main-table">
            <el-table :data="tableData" border style="width: 100%;height: 100%;" empty-text="暂无数据">
            <el-table-column prop="Pid" label="PID" width="100" />
            <el-table-column prop="Command" label="程序名称" width="180">
                <template #default="scope">
                    <span v-if="scope.row.Command.indexOf('x81') > 0">未知程序</span>
                    <span v-else>{{ scope.row.Command }}</span>
                </template>
            </el-table-column>
            <el-table-column prop="User" label="用户" />
            <el-table-column prop="IpProtocol" label="IP协议" />
            <!-- 超出隐藏 -->
            <el-table-column prop="Connection" :show-overflow-tooltip="true" label="连接" width="300"/>
            <el-table-column label="操作" width="120">
                <template #default="scope">
                    <el-button @click="onKillProcess(scope.row)">
                        终止进程
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
        </div>
    </div>
</template>


<script setup>
import { ElMessage, ElMessageBox } from 'element-plus'
import {
    Search,
    Menu as IconMenu,
    Location,
    Setting
} from '@element-plus/icons-vue'
import { ref, onMounted, onActivated, reactive, watch } from 'vue';
import { GetProcessListByPortNumber, KillProcessByPortNumber } from '../../wailsjs/go/main/App'

const portNumber = ref()
const tableData = ref()

onActivated(() => {
    console.log("PortList onActivated !!!!!!!")
})

watch(portNumber, (newP, oldP) => {
    if (!newP) {
        tableData.value = []
    }
})
const onSearch = () => {
    getProcessListByPortNumber()
}

const getProcessListByPortNumber = () => {
    GetProcessListByPortNumber(portNumber.value).then(result => {

        tableData.value = result

    })
}


const onKillProcess = (row) => {
    ElMessageBox.confirm('确认终止进程 ' + row.Command + ' 吗？', '提示', {
        // if you want to disable its autofocus
        // autofocus: false,
        confirmButtonText: '确认',
        cancelButtonText: "取消"
    }).then(() => {
        KillProcessByPortNumber(row.Pid).then(success => {
            if (success) {
                ElMessage({
                    type: "success",
                    message: "已终止"
                })

                onSearch()
            } else {
                ElMessage({
                    type: "warn",
                    message: "终止失败"
                })
            }
        })

    }).catch(() => {
        ElMessage({
            type: "info",
            message: "已取消"
        })
    })
}
</script>


<style>
.main-container {
    /* width: 100vw; */
    height: 100vh;
}

.main-table {
    height: calc(100% - 32px);
}
/* 去除端口号搜索框、放大镜搜索按钮的圆角边框 */
.port-number-container .el-input__wrapper,
.port-number-container .el-input-group__append {
    border-radius: 0 !important;
}
</style>