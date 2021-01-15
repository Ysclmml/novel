<template>
  <div>
    <!-- 作品评论区 start -->
    <div class="pad20">
      <div class="bookComment">
        <div class="book_tit">
          <div class="fl">
            <h3>作品评论区</h3><span id="bookCommentTotal">({{ commentCount }}条)</span>
          </div>
          <a class="fr" href="#txtComment">发表评论</a>
        </div>
        <div v-show="commentList.length === 0" class="no_comment">
          <img src="/images/no_comment.png" alt="">
          <span class="block">暂无评论</span>
        </div>
        <div class="commentBar">
          <div v-for="(item, index) in commentList" :key="index" class="comment_list cf">
            <div class="user_heads fl">
              <img :src="item.create_user_photo || '/images/man.png'" class="user_head" alt="">
              <span class="user_level1" style="display: none;">见习</span>
            </div>
            <ul class="pl_bar fr">
              <li class="name">{{ item.create_user_name }}</li>
              <li class="dec">{{ item.comment_content }}🌹</li>
              <li class="other cf">
                <span class="time fl">{{ item.create_time }}</span>
                <span class="fr">
                  <a href="" class="zan" style="display: none;">赞<i class="num">(0)</i></a>
                </span>
              </li>
            </ul>
          </div>
        </div>

        <el-pagination
          v-show="isShowPage"
          class="pageBox cf mt15 mr10"
          background
          layout="prev, pager, next"
          :total="commentCount"
          :page-size="pageSize"
          :current-page.sync="currentPage"
          @current-change="handleCurrentChange"
        >
        </el-pagination>

        <!-- 无评论时此处隐藏 -->
        <div v-show="commentList.length > 0 && !isShowPage" class="more_bar">
          <!-- 跳转到评论页面 -->
          <router-link :to="{name: 'BookComment', params: {bookId}}">查看全部评论&gt;</router-link>
        </div>

        <div id="reply_bar" class="reply_bar">
          <div class="tit">
            <span class="fl font16">发表评论</span>
            <!-- 未登录状态下不可发表评论，显示以下链接 -->
            <span class="fr black9" style="display:none; ">
              请先 <a class="orange" href="">登录</a><em class="ml10 mr10">|</em>
              <a class="orange" href="">注册</a></span>
          </div>
          <textarea id="txtComment" name="txtComment" rows="2" cols="20" class="replay_text"
                    placeholder="我来说两句..."
          />
          <div class="reply_btn">
            <span class="fl black9"><em id="emCommentNum" class="ml5">0/1000</em> 字</span>
            <span class="fr"><a class="btn_ora" href="">发表</a></span>
          </div>
        </div>
      </div>
    </div>
    <!-- 作品评论区 end -->
  </div>
</template>

<script>
import {listCommentByPage} from "@/api/book";

export default {
  name: "Comment",
  props: {
    bookId: {
      type: String,
      default: '0'
    },
    // 用来控制是否显示分页组件, 如果不显示就显示(查看更多按钮)
    isShowPage: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      commentList: [],
      commentCount: 0,
      pageSize: 10,  // 每页的评论数量
      currentPage: 1 // 当前页码
    }
  },
  mounted() {
    const pageSize = this.isShowPage ? this.pageSize : 3
    this.getCommentList(1, pageSize)
  },
  methods: {
    async getCommentList(page, pageSize) {
      // 默认只取3条数据
      const { data } = await listCommentByPage(this.bookId, page, pageSize)
      this.commentList = data.result
      this.commentCount = data.total
    },
    // 页码改变请求评论信息
    handleCurrentChange() {
      this.getCommentList(this.currentPage, this.pageSize)
    }
  }
}
</script>

<style lang="scss">
.bookComment .el-pagination.is-background .el-pager li:not(.disabled).active{
  background-color: #f80;
}
</style>
