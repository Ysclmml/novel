<template>
  <div>
    <Top></Top>
    <Body></Body>
    <Footer></Footer>
  </div>
</template>

<script>
import Body from './components/Page'
import Top from '@/views/common/Top'
import Footer from '@/views/common/Footer'

export default {
  name: 'Index',
  components: {Top, Footer, Body}
}
</script>
