<template>
  <div>
    <div class="main box_center cf mb50">
      <div class="nav_sub">
        <a href="/" />小说网站名字 &gt;<a href="">{{ bookDetail.cat_name }}</a>&gt;<a href="">{{ bookDetail.book_name }}</a>
      </div>
      <div class="channelWrap channelBookInfo cf">
        <div class="bookCover cf">
          <a :href="bookDetail.pic_url" class="book_cover">
            <img class="cover" :src="bookDetail.pic_url" :alt="bookDetail.book_name">
          </a>
          <div class="book_info">
            <div class="tit">
              <h1>{{ bookDetail.book_name }}</h1>
              <a class="author">{{ bookDetail.author_name }} 著</a>
            </div>
            <ul class="list">
              <li>
                <span class="item">类别：<em>{{ bookDetail.cat_name }}</em></span>
                <span class="item">状态：<em>{{ bookDetail.book_status ? '已完结' : '连载中' }}</em></span>
                <span class="item">总点击：<em>{{ bookDetail.visit_count }}</em></span>
                <span class="item">总字数：<em />{{ bookDetail.word_count }}</span>
              </li>
            </ul>
            <div class="intro_txt">
              <p v-html="bookDetail.book_desc" />
              <a class="icon_hide" href="#" onclick=""><i />收起</a>
              <a class="icon_show" href="#" onclick=""><i />展开</a>
            </div>
            <div id="optBtn" class="btns">
              <a href="" class="btn_ora">点击阅读</a>
              <span id="cFavs"><a href="" class="btn_ora_white btn_addsj">加入书架</a>
              </span>
            </div>
          </div>
        </div>
      </div>

      <div class="channelBookContent cf">
        <!-- left start -->
        <div class="wrap_left fl">
          <div class="wrap_bg">
            <!-- 章节目录 start -->
            <div class="pad20_nobt">
              <div class="bookChapter">
                <div class="book_tit">
                  <div class="fl">
                    <h3>最新章节</h3><span>({{ bookAbout.book_index_count }}章)</span>
                  </div>
                  <a class="fr" href="">全部目录</a>
                </div>
                <ul class="list cf">
                  <li>
                    <span class="fl font16"> <a href="">{{ bookDetail.last_index_name }}</a></span>
                    <span class="black9 fr">更新时间：{{ bookDetail.last_index_update_time | timeFormat }}</span>
                  </li>
                  <li class="zj_yl" v-html="bookAbout.book_content" />
                <!-- 此处是该章节预览，截取最前面的42个字 -->
                </ul>
              </div>
            </div>
            <!-- 章节目录 end -->
            <!-- 评论区 -->
            <Comment :book-id="bookId" :is-show-page="false"></Comment>
          </div>
        </div>
        <!-- left end -->
        <!-- right start -->
        <div class="wrap_right fr">
          <AuthorPart :book-detail="bookDetail" :book-id="bookId" />
          <ClassifyRec :book-id="bookId" :cat-id="bookDetail.cat_id" />
        </div>
      <!-- right end -->
      </div>
    </div>
  </div>
</template>

<script>
import {dateFormat} from "@/utils";
import {addVisitCount, queryBookDetail, queryBookIndexAbout} from "@/api/book"
import Comment from "@/views/comment/components/Comment"
import AuthorPart from "@/views/bookdetail/components/AuthorPart";
import ClassifyRec from "@/views/bookdetail/components/ClassifyRec";

export default {
  name: "Page",
  components: {ClassifyRec, AuthorPart, Comment},
  filters: {
    timeFormat(timeString) {
      const d = new Date(timeString)
      console.log(timeString, '...timeFormat')
      return dateFormat(d.getTime(), 'yy/MM/dd hh:mm:ss')
    }
  },
  data() {
    return {
      bookId: 0,
      bookDetail: {},
      bookAbout: {
        book_index_count: 0,
        book_content: ''
      }
    }
  },
  async created() {
    this.bookId = this.$route.params.bookId
    await this.getBookDetail()
    await this.getLatestUpdate()
    this.addVisit()

  },
  methods: {
    async getBookDetail() {
      const { data } = await queryBookDetail(this.bookId)
      this.bookDetail = data
    },
    async getLatestUpdate() {
      const { data } = await queryBookIndexAbout(this.bookId, this.bookDetail.last_index_id)
      this.bookAbout = data
    },
    // 添加访问数
    addVisit() {
      addVisitCount(this.bookId).then().catch(err => {console.log('err: addVisit ----', err)})
    }
  }
}
</script>

<style scoped>

</style>
