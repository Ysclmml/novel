<template>
  <div>
    <!-- 作者专栏s -->
    <div class="wrap_inner author_info mb20">
      <div class="author_head cf">
        <a href="" class="head"><img id="authorLogoImg" src="/images/author_head.png" alt="作者头像"></a>
        <div class="msg">
          <span class="icon_qyzz">签约作家</span>
          <h4>
            <!-- todo 添加作者的搜索页面跳转... -->
            <a href="">{{ bookInfo.author_name }}</a>
          </h4>
        </div>
      </div>
      <div class="author_intro cf">
        <h4>作者有话说</h4>
        <div id="authorNote" class="intro_txt">
          亲亲们，你们的支持是我最大的动力！求点击、求推荐、求书评哦！
        </div>
      </div>
      <!-- 如果作者没有其他作品就下方代码整个不显示 -->
    </div>
    <!-- 作者专栏e -->
  </div>
</template>

<script>
import {queryBookDetail} from "@/api/book";

export default {
  name: "AuthorPart",
  props: {
    bookDetail: {
      type: Object,
      default: () => { return { id: '0'} }
    },
    bookId: {
      type: String,
      default: '0'
    }
  },
  data() {
    return {
      bookInfo: {}
    }
  },
  created() {
    // 如果bookDetail有值, 就不去请求, 否则需要请求接口数据
    if (this.bookDetail.id === '0') {
      this.getBookDetail()
    } else {
      this.bookInfo = this.bookDetail
    }
  },
  methods: {
    async getBookDetail() {
      const { data } = await queryBookDetail(this.bookId)
      this.bookInfo = data
    }
  }
}
</script>

<style scoped>

</style>
