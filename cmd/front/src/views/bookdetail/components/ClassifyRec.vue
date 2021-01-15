<!-- 同类推荐 -->
<template>
  <div>
    <div id="RelateBookOther" class="wrap_inner wrap_right_cont mb20">
      <div class="title cf">
        <h3 class="on">同类推荐</h3>
      </div>
      <div class="tj_bar">
        <ul>
          <li v-for="(item, index) in recBooks" :key="index">
            <div class="book_intro">
              <div class="cover">
                <a @click="goBookDetail(item.id)"><img :src="item.pic_url" :alt="item.book_name"></a>
              </div>
              <div class="dec">
                <a class="book_name" @click="goBookDetail(item.id)">{{ item.book_name }}</a>
                <a class="txt" @click="goBookDetail(item.id)" v-html="item.book_desc"></a>
              </div>
            </div>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
import {listRecBookByCatId} from "@/api/book";

export default {
  name: "ClassifyRec",
  props: {
    bookId: {
      type: String,
      default: '0'
    },
    catId: {
      type: Number,
      default: 0
    }
  },
  data() {
    return {
      recBooks: []
    }
  },
  watch: {
    catId(value) {
      if (value > 0) {
        this.getClassifyRecBooks()
      }
    }
  },
  methods: {
    async getClassifyRecBooks() {
      const { data } = await listRecBookByCatId(this.bookId, this.catId)
      this.recBooks = data
    },
    // 跳转到对应的路由中
    goBookDetail(bookId) {
      this.$router.push({
        name: 'BookDetail',
        params: {
          bookId
        }
      })
    }
  }
}
</script>

<style scoped>

</style>
