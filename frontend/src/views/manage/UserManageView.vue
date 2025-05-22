<template>
  <div class="w-full p-4">
    <div class="mb-4 flex gap-4">
      <el-button type="primary" @click="addUserDrawer = true">添加用户</el-button>
    </div>
    <el-card>
      <div class="mb-4 flex gap-4">
        <el-input v-model="filter.user_id" placeholder="ID" clearable @clear="search">
          <template #prefix>
            <el-icon><Search /></el-icon>
          </template>
        </el-input>
        <el-input v-model="filter.user_name" placeholder="用户名" clearable @clear="search">
          <template #prefix>
            <el-icon><Search /></el-icon>
          </template>
        </el-input>
        <el-input v-model="filter.user_email" placeholder="邮箱" clearable @clear="search">
          <template #prefix>
            <el-icon><Search /></el-icon>
          </template>
        </el-input>
        <el-button type="primary" @click="search">搜索</el-button>
      </div>

      <el-table :data="tableData" stripe border class="w-full" v-loading="loading">
        <el-table-column prop="user_id" label="ID" />
        <el-table-column prop="user_name" label="用户名" />
        <el-table-column prop="user_email" label="邮箱" />
        <el-table-column prop="user_role" label="角色">
          <template #default="scope">
            <div>{{ roleToText(scope.row.user_role) }}</div>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template #default="scope">
            <div class="flex justify-start gap-3">
              <el-link type="primary" @click="handleUserDetail(scope.row)" plain :underline="false"
                >详情</el-link
              >
              <el-link type="danger" @click="handleDeleteUser(scope.row)" plain :underline="false"
                >删除</el-link
              >
            </div>
          </template>
        </el-table-column>
      </el-table>

      <el-pagination
        v-model:current-page="page"
        v-model:page-size="pageSize"
        :total="total"
        class="mt-4"
        @size-change="search"
        @current-change="search"
      />
    </el-card>

    <el-drawer v-model="addUserDrawer" title="添加用户" size="50%">
      <el-form label-width="100px">
        <el-form-item label="用户名">
          <el-input v-model="newUserInfo.user_name" placeholder="请输入用户名" />
        </el-form-item>
        <el-form-item label="邮箱">
          <el-input v-model="newUserInfo.user_email" placeholder="请输入邮箱" />
        </el-form-item>
        <el-form-item label="密码">
          <el-input
            v-model="newUserInfo.user_password"
            type="password"
            clearable
            placeholder="请输入密码"
            readonly
            @focus="(e: any) => e.target.removeAttribute('readonly')"
          />
        </el-form-item>
        <el-form-item label="角色">
          <el-select v-model="newUserInfo.user_role" placeholder="请选择角色">
            <el-option label="普通用户" value="user" />
            <el-option label="管理员" value="admin" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleAddUser">添加</el-button>
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer v-model="userDetailDrawer" title="用户详情" size="50%">
      <el-form label-width="100px" autocomplete="off">
        <div class="mb-4 flex items-center justify-center">
          <el-avatar
            :size="100"
            :src="currentUser.user_avatar_url"
            class="cursor-pointer border-4 border-white shadow-lg"
            @click="handleClickAvatar"
            title="点击更换头像"
          />
        </div>
        <el-form-item label="用户ID">
          <span>{{ currentUser.user_id }}</span>
        </el-form-item>
        <el-form-item label="用户名">
          <el-input v-model="currentUser.user_name" />
        </el-form-item>
        <el-form-item label="邮箱">
          <el-input v-model="currentUser.user_email" />
        </el-form-item>
        <el-form-item label="密码">
          <el-input
            v-model="currentUser.user_password"
            type="password"
            show-password
            placeholder="不修改密码则不填"
            readonly
            @focus="(e: any) => e.target.removeAttribute('readonly')"
          />
        </el-form-item>
        <el-form-item label="角色">
          <el-select v-model="currentUser.user_role" placeholder="请选择角色">
            <el-option label="管理员" value="admin" />
            <el-option label="普通用户" value="user" />
          </el-select>
        </el-form-item>
        <el-form-item label="个人简介">
          <el-input type="textarea" v-model="currentUser.user_profile" />
        </el-form-item>
        <el-form-item label="注册时间">
          <span>{{ currentUser.create_at }}</span>
        </el-form-item>
        <el-form-item label="更新时间">
          <span>{{ currentUser.update_at }}</span>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleUpdateUser">更新用户信息</el-button>
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup lang="ts">
import api from '@/api'
import { ref, onMounted } from 'vue'
import { Search } from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus'

interface User {
  user_id: number
  user_name: string
  user_email: string
  user_password: string
  user_role: string
  user_avatar_url: string
  avatar_base64: string
  user_profile: string
  create_at: string
  update_at: string
}

const currentUser = ref<User>({
  user_id: NaN,
  user_name: '',
  user_email: '',
  user_password: '',
  user_role: '',
  user_avatar_url: '',
  avatar_base64: '',
  user_profile: '',
  create_at: '',
  update_at: '',
})

interface NewUserInfo {
  user_name: string
  user_email: string
  user_password: string
  user_role: string
}

const newUserInfo = ref<NewUserInfo>({
  user_name: '',
  user_email: '',
  user_password: '',
  user_role: '',
})

const filter = ref({
  user_id: '',
  user_name: '',
  user_email: '',
})

const tableData = ref<User[]>([])
const total = ref(0)
const page = ref(1)
const pageSize = ref(10)
const addUserDrawer = ref(false)
const userDetailDrawer = ref(false)
const loading = ref(false)

onMounted(() => {
  document.title = 'OnlineJudge - 用户管理'
  search()
})

const search = () => {
  loading.value = true
  api.user
    .getUserList(page.value, pageSize.value, filter.value)
    .then((res) => {
      tableData.value = res.data.users
      total.value = res.data.total
    })
    .catch(() => {
      ElMessage.error('获取用户列表失败')
    })
    .finally(() => {
      loading.value = false
    })
}

const handleAddUser = async () => {
  try {
    await api.user.addUser(newUserInfo.value)
    newUserInfo.value = {
      user_name: '',
      user_email: '',
      user_password: '',
      user_role: '',
    }
    addUserDrawer.value = false
    ElMessage.success('添加成功')
    search()
  } catch {
    ElMessage.error('添加失败')
  }
}

const handleUserDetail = (row: User) => {
  currentUser.value = JSON.parse(
    JSON.stringify({
      user_id: row.user_id,
      user_name: row.user_name,
      user_email: row.user_email,
      user_password: '',
      user_role: row.user_role,
      user_avatar_url: row.user_avatar_url,
      avatar_base64: '',
      user_profile: row.user_profile,
      create_at: row.create_at,
      update_at: row.update_at,
    }),
  )
  userDetailDrawer.value = true
}

const handleDeleteUser = (row: User) => {
  ElMessageBox.confirm('确定删除该用户吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  }).then(() => {
    api.user
      .deleteUser(row.user_id)
      .then(() => {
        ElMessage.success('删除成功')
        search()
      })
      .catch(() => {
        ElMessage.error('删除失败')
      })
  })
}

const roleToText = (role: string) => {
  return role === 'admin' ? '管理员' : '普通用户'
}

const handleUpdateUser = async () => {
  try {
    await api.user.updateUser({
      user_id: currentUser.value.user_id,
      user_name: currentUser.value.user_name,
      user_email: currentUser.value.user_email,
      avatar_base64: currentUser.value.avatar_base64,
      user_profile: currentUser.value.user_profile,
      user_role: currentUser.value.user_role,
      user_password: currentUser.value.user_password,
    })
    ElMessage.success('更新成功')
    search()
  } catch {
    ElMessage.error('更新失败')
  }
}

const handleClickAvatar = () => {
  // 选择文件，并转为base64
  const file = document.createElement('input')
  file.type = 'file'
  file.accept = 'image/jpeg, image/png'
  file.onchange = () => {
    currentUser.value.user_avatar_url = URL.createObjectURL(file.files![0])
    const reader = new FileReader()
    reader.readAsDataURL(file.files![0])
    reader.onload = (e) => {
      currentUser.value.avatar_base64 = e.target?.result as string
      currentUser.value.avatar_base64 = currentUser.value.avatar_base64.split(',')[1]
    }
  }
  file.click()
}
</script>
