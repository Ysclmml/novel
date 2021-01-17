<template>
  <div id="app">
    <!-- 暂时根据路由的全路径来解决相同路由但是不刷新的问题 -->
    <router-view :key="$route.fullPath" />
  </div>
</template>

<script>
</script>

<style lang="scss">
@import url(~@/style/base.css);
@import url(~@/style/main.css);
@import url(~@/style/book.css);
</style>
