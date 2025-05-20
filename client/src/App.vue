<template>
  <div id="app">
    <h3>Добавить нового пользователя:</h3>
    <form @submit.prevent="addUser">
      <input id="newUserName" type="text" placeholder="Имя пользователя" v-model="newUserName"/>
      <button type="submit">Добавить</button>
    </form>
    <h3>Пользователи:</h3>
    <ul>
      <li v-for="user in users" :key="user.name">{{ user.name }}</li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

type User = {
  name:string
}

const users = ref<User[]>([])
const newUserName = ref('')

onMounted(fetchUsers)

async function fetchUsers () {
  try {
    const res = await fetch('http://localhost:8080/users')
    const data = await res.json()
    users.value = data
  } catch (err) {
    console.error('Ошибка загрузки пользователей:', err)
  }
}

async function addUser () {
  if (!newUserName.value.trim()) return

  const newUser = { name: newUserName.value.trim() }

  try {
    await fetch('http://localhost:8080/users', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(newUser)
    })

    fetchUsers()
    newUserName.value = ''
  } catch (err) {
    console.error('Ошибка добавления пользователя:', err)
  }
}
</script>

<style scoped>
#app {
  font-family: Arial, sans-serif;
}

#newUserName {
  margin-right: 10px;
  width: 500px;
  height: 33px;
  font-size: 1.2rem;
  border-radius: 5px;
}

form {
  display: flex;
  justify-content: center;
  width: 700px;
  background-color: #d6e2ef;
  padding:10px 20px;
  border-radius: 10px;
  margin-bottom: 40px;
}

ul {
  height: 60vh;
  overflow: auto;
}

</style>
