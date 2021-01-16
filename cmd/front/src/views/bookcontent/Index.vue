<template>
  <div>
    <Top />
    <Page />
    <Footer />
  </div>
</template>

<script>
import Page from './components/Page'
import Top from '@/views/common/Top'
import Footer from '@/views/common/Footer'

export default {
name: "Index",
  components: {Top, Footer, Page}
}
</script>

<style scoped>

</style>
