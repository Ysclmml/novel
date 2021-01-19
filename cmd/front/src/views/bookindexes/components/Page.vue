<template>
  <div class="main box_center cf">
    <div class="nav_sub">
      <a href="/">网站名字</a>&gt;<a>{{ book.cat_name }}</a>&gt;
      <a>{{ book.book_name }}</a>&gt;<a>作品目录</a>
    </div>
    <div class="channelWrap channelChapterlist cf mb50">
      <div class="bookMain">
        <div class="bookCover cf">
          <div class="book_info1">
            <div class="tit">
              <h1>{{ book.book_name }}</h1><!-- <i class="vip_b">VIP</i> -->
            </div>
            <ul class="list">
              <li>
                <span>作者：<a>{{ book.author_name }}</a></span>
                <span>类别：<a>{{ book.cat_name }}</a></span>
                <span>状态：<em class="black3"> {{ book.book_status ? '已完结' : '连载中' }}</em></span>
                <span>总点击：<em class="black3">{{ book.visit_count }}</em></span>
                <span>总字数：<em class="black3">{{ book.word_count }}</em></span>
              </li>
            </ul>
          </div>
        </div>
        <div class="dirWrap cf">
          <h3>正文({{ bookIndexesCount }})</h3>
          <div class="dirList">
            <ul>
              <li v-for="(bookIndex, index) in bookIndexList" :key="index">
                <router-link :to="{name: 'BookContent', params: {bookId, indexId: bookIndex.id}}">
                  <span />{{ bookIndex.index_name }}<i v-if="!bookIndex.is_vip" class="red"> [免费]</i>
                </router-link>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {getIndexList, queryBookDetail} from "@/api/book";

export default {
  name: "Page",
  data() {
    return {
      book: {},  // 书籍信息
      bookId: '',
      bookIndexList: [], // 书籍目录列表
      bookIndexesCount: 0, // 目录数量
    }
  },
  created() {
    this.bookId = this.$route.params.bookId
    this.getBookDetail()
    this.getList()
  },
  methods: {
    async getList() {
      const { data } = await getIndexList(this.bookId)
      console.log('getList', data)
      this.bookIndexList = data.result
      this.bookIndexesCount = data.total
    },
    async getBookDetail() {
      const { data } = await queryBookDetail(this.bookId)
      console.log('queryBookDetail', data)
      this.book = data
    }
  }
}
</script>

<style scoped>

</style>
