// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ssr: true,
  modules: ['@nuxtjs/apollo'],

  apollo: {
    clients: {
      default: {
        // Graphql endpoint
        httpEndpoint: 'http://localhost:8080/query',
        // Rest endpoint
        // httpEndpoint: 'https://localhost:8080',
      }
    },
  },
})
