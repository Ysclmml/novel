<template>
  <div>
    <Top></Top>
    <Page></Page>
    <Footer></Footer>
  </div>
</template>

<script>
import Page from './components/Page'
import Top from '@/views/common/Top'
import Footer from '@/views/common/Footer'

export default {
  name: 'Index',
  components: {Top, Footer, Page}
}
</script>
