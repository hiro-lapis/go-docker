<script lang="ts" setup>
type Todo = {
  id: number
  text: string
}
type Todos = {
  todos: Todo[]
}
const query = gql`
  query getTodos {
    todos {
      id
      text
    }
  }
`
const datas = ref<Todo[]>([])
// useAsyncQuery pattern
// @see https://apollo.nuxtjs.org/getting-started/composables/#useasyncquery
// 公式より: useAsyncQuery is primarily used for querying data when a page or component is initially loaded
const asyncFetchTodo = async () => {
  // for grqphql, useAsyncQuery, (REST API =useFetch)
  // passing generic type to useAsyncQuery, it will be inferred
  const { data } = await useAsyncQuery<Todos>(query, {})
  // APIで取得したdataもvalueに格納されている
  datas.value = data.value?.todos as unknown as Todo[]
  console.log('fetched!!')
  console.log(datas.value)
  console.log('data.value')
  console.log(data.value)
}

// useQuery
// @see https://apollo.nuxtjs.org/getting-started/composables/#usequery
// 公式より: unlike useAsyncQuery which is best used for initially fetching data in SSR applications
const fetchTodo = async () => {
  // resultを素直に受け取ると、画面初期表示時にはデータ取得しないで代入処理が走ってしまうので注意
  const { onResult } = await useQuery<Todos>(query, {})
  // on result is a function that will be called when the query is finished
  onResult((result) => {
    datas.value = result.data?.todos as unknown as Todo[]
  })
}
onMounted(async () => {
  console.log('mounted')
  console.log(datas.value)
  await fetchTodo()
  console.log('done')
  console.log(datas.value)
})
// fetchTodo()
</script>

<template>
  <p>here is a apllo client test</p>
  <NuxtLink to="/">Index</NuxtLink>
  <div>
    <template v-if="datas && datas.length > 0">
      <p>{{ datas }}</p>
      <ul>
        <li v-for="todo in datas" :key="todo.id">
          ID:{{ todo.id }} : {{ todo.text }}
        </li>
      </ul>
    </template>
  </div>

  <button @click="fetchTodo">fetch</button>
</template>
