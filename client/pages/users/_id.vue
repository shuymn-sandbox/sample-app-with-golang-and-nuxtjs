<template>
  <div>
    <p>{{ user.id }}</p>
  </div>
</template>

<script>
export default {
  async asyncData({ params, app }) {
    const { id } = params
    const user = await app.$axios.$get(`http://localhost:1323/users/${id}`)
    return { user }
  }
}
</script>
