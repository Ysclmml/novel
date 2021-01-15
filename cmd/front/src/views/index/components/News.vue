<template>
  <div>
    <dl id="indexNews" class="hot_notice">
      <dd v-for="(item, index) in newsList" :key="index" style="text-align:left;">
        <span>[{{ item.cat_name }}]</span><a href="" v-html="item.title"></a>
      </dd>
    </dl>
  </div>
</template>

<script>
import { listIndexNews } from "@/api/news";

export default {
  name: "News",
  data() {
    return {
      newsList: []
    }
  },
  mounted() {
    this.getListIndexNews()
  },
  methods: {
    async getListIndexNews() {
      const { data } = await listIndexNews()
      this.newsList = data
    }
  }
}
</script>

<style scoped>

</style>
