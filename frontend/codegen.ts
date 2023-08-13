import type { CodegenConfig } from '@graphql-codegen/cli'

// @see https://the-guild.dev/graphql/codegen/docs/config-reference/codegen-config
const config: CodegenConfig = {
  // grahpql型定義の取得先
  // http://<container>:<port>/query
  schema: 'http://go-docker-go-1:8080/query',
  // query,mutationの取得先
  documents: ['pages/**/*.vue'],
  // for better experience with the watcher
  ignoreNoDocuments: true,
  generates: {
    // defined type of used by vue-apollo
    'src/generated/graphql.ts': {
      // https://the-guild.dev/graphql/codegen/plugins
      plugins: ['typescript', 'typescript-operations', 'typescript-vue-apollo'],
      config: {
        withCompositionFunctions: true,
        vueCompositionApiImportFrom: 'vue',
        scalars: {
          ULID: 'string',
          Uint: 'number',
          Uint32: 'number',
          Uint64: 'number',
          Float64: 'number',
        },
      },
    },
    // defined types for introspection
    './graphql.schema.json': {
      plugins: ['introspection'],
    },
  },
}

export default config
