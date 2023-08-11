<script setup lang="ts">

const message = ref('')
const placeholder = ref()

onMounted( async () => {
  console.log('index mounted')
  await $fetch('https://jsonplaceholder.typicode.com/todos/1')
    .then((res) => {
        placeholder.value = res
    }).then((json) => console.log(json))

  await $fetch('http://localhost:8080')
    .then((res) => {
        message.value = res.message
        console.log(res)
    })
    .then((json) => console.log(json))
})
// onMounted(() => {
//   console.log('index mounted')
//   $fetch('https://jsonplaceholder.typicode.com/todos/1')
//     .then((res) => res.json())
//     .then((json) => console.log(json))
// })
</script>
<template>
    <div>
      <h1>Welcome to the homepage</h1>
      <div>
        <div>
          <NuxtLink to="/about">About</NuxtLink>
        </div>
        <div>
          <NuxtLink to="/apollo">Apollo client test</NuxtLink>
        </div>
      </div>
      <AppAlert>
        This is an auto-imported component
      </AppAlert>
    <p>if you can see like below, interact with backend is succeeded</p>
    <img src="img/example.png" alt="" :style="{ width: '300px'}">

    <template v-if="message">
        <p>message:{{ message }}</p>
    </template>
    <template v-if="placeholder">
        <p>jsonplaceholder response:{{ placeholder }}</p>
    </template>
    </div>
</template>
