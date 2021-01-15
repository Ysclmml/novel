<!-- 评论页面 -->
<template>
  <div>
    <Top></Top>
    <Page></Page>
    <Footer></Footer>
  </div>
</template>

<script>
import Top from '@/views/common/Top'
import Footer from '@/views/common/Footer'
import Page from './components/Page'

export default {
  name: 'Index',
  components: {Top, Footer, Page}
}
</script>

<style scoped>

</style>
