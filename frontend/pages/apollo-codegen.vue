<script lang="ts" setup>
import {
  Todo,
 } from '~/src/generated/graphql'

const datas = ref<Todo[]>([])

/**
 * codegenを使う場合
 * ・フロントの.graphqlファイルにクエリドキュメント(ex.getTodos)を記述
 * ・pnpm graphql:generate で生成
 * ・importで自動生成されたcomposable useGetXXXQuery() 取得
 * ・composable実行
 * ・ Promiseレベルで受け取りたい&ロード状態を管理したい時
 * ・ const { result, loading, error } = useGetTodosQuery();
 * ・ レスポンスハンドラを実装したい時
 * ・ const { onResult, onError } = useGetTodosQuery();
 */

const query = gql`
  query testQuery {
    test
  }
`
// uselazyasyncquery
// @see https://apollo.nuxtjs.org/getting-started/composables/#uselazyasyncquery
// 公式より: unlike useAsyncQuery which is best used for initially fetching data in SSR applications
const fetchLazyTodo = async () => {
  const options = {}
  // Promise状態のresultを処理するとレスポンスが入ってないためハンドラを使う
  // loadingは自身で管理する必要がある
  const { data } = useLazyAsyncQuery(query, options)
  console.log('lazy fetched!!')
  console.log(data)
}

// useQuery
// @see https://apollo.nuxtjs.org/getting-started/composables/#usequery
// 公式より: unlike useAsyncQuery which is best used for initially fetching data in SSR applications
const fetchTodo = () => {
  const options = {}
  // Promise状態のresultを処理するとレスポンスが入ってないためハンドラを使う
  const { onResult } = useQuery(query, options)
  onResult((result) => {
    console.log('fetched!!')
    console.log(result)
    // datas.value = result.data?.todos as unknown as Todo[]
  })
}

onMounted(() => {
  console.log('mounted')
  // 初期表示で取得したい場合
  fetchTodo()
})
</script>

<template>
  <p>here is a apllo with code-gen client test</p>
  <NuxtLink to="/">Index</NuxtLink>
  <NuxtLink to="/apollo">Apollo</NuxtLink>

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

  <div>
    <p>useQuery</p>
    <p>初期表示時に実行される。</p>
    <p>Queryは初期表示時に実行されているため、ボタン押下してもfetchTodoのクエリは実行されない</p>
    <button @click="fetchTodo">fetch</button>
  </div>
  <div>
    <p>useLazyQuery</p>
    <p>初期表示時に実行されない</p>
    <p>ボタン押下の度に実行される</p>
    <button @click="fetchLazyTodo">Lazy fetch</button>
  </div>
</template>
