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
import Page from "@/views/bookdetail/components/Page";

export default {
  name: "Index",
  components: {Page, Top, Footer}
}
</script>

<style scoped>

</style>
