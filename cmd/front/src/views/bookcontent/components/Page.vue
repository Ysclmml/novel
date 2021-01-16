<template>
  <div>
    <div>
      <div class="readBody cf">
        <div class="readMain cf">
          <div class="read_menu">
            <div id="readMenu" class="menu_left" style="">
              <ul>
                <!-- todo: 跳转到目录页 -->
                <li>
                  <a class="ico_catalog" title="目录">
                    <svg-icon icon-class="list" class="menu-icon"></svg-icon>
                    <b>目录</b></a>
                </li>
                <li>
                  <a class="ico_page" title="返回书页" @click="goBookDetail(bookId)">
                    <svg-icon icon-class="book" class="menu-icon"></svg-icon>
                    <b>书页</b>
                  </a>
                </li>
                <li class="li_shelf" title="加入书架">
                  <a class="ico_shelf">
                    <svg-icon icon-class="marker" class="menu-icon"></svg-icon>
                    <b>加书架</b>
                  </a></li>
                <li class="li_shelfed" style="display: none;">
                  <svg-icon icon-class="marker" class="menu-icon"></svg-icon>
                  <a class="ico_shelfed" href="#" title="已收藏"><b>已收藏</b></a>
                </li>
                <li>
                  <a class="ico_comment" title="评论">
                    <svg-icon icon-class="comment" class="menu-icon"></svg-icon>
                    <b>评论</b>
                  </a>
                </li>
                <li>
                  <a class="ico_setup" title="设置">
                    <svg-icon icon-class="settings" class="menu-icon"></svg-icon>
                    <b>设置</b>
                  </a>
                </li>
              </ul>
            </div>
            <div class="menu_right" style="position: fixed; bottom: 0">
              <ul>
                <li><a class="ico_pagePrev" title="上一章"><i>上一章</i></a></li>
                <li><a class="ico_pageNext" title="下一章"><i>下一章</i></a></li>
              </ul>
            </div>
          </div>
          <div class="readWrap">
            <div class="bookNav">
              <a href="/">首页 </a>&gt; <a href=""> {{ book.cat_name }}</a>&gt;
              <a @click="goBookDetail(bookId)"> {{ book.book_name }}</a>
            </div>
            <div>
              <div class="textbox cf">
                <div class="book_title">
                  <h1> {{ bookIndex.index_name }}</h1>
                  <div class="textinfo">
                    <!-- todo: 书籍类别信息 -->
                    类别：<a>{{ book.cat_name }}</a>
                    <!-- todo: 书籍按照作者搜索功能 -->
                    作者：<a>{{ book.author_name }}</a>
                    <span>字数：{{ bookIndex.wordCount }}</span>
                    <span>更新时间：{{ bookIndex.update_time | timeFormat('yy/MM/dd hh:mm:ss') }}</span>
                  </div>
                </div>
                <div v-if="book.is_vip" class="txtwrap">
                  <div class="readBox" style="font-size: 16px; font-family: microsoft yahei">
                    <p />
                    <div class="pc_bar" style="display: none;">
                      <a class="icon_pc">
                        <span><i class="icon_yb" /><em>捧场</em></span>
                      </a>
                    </div>
                  </div>
                  <!-- todo: 添加查看章节的是否收费, 然后根据用户是否购买决定是否显示章节 -->
                  <div class="orderBox">
                    <h3>此章为VIP章节，需要订阅后才能继续阅读</h3>
                    <form method="post" action="">
                      <ul class="order_list">
                        <li>价格：<span class="red">{{ bookIndex.book_price }}</span></li>
                        <li class="btns"><a class="btn_red">购买</a></li>
                      </ul>
                    </form>
                  </div>
                </div>
                <div class="txtwrap">
                  <div class="readBox" style="font-size: 16px; font-family: microsoft yahei" v-html="bookContent.content">
                    <div class="pc_bar" style="display: none;">
                      <a class="icon_pc">
                        <span><i class="icon_yb" /><em>捧场</em></span>
                      </a>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="nextPageBox">
              <!-- todo: 目录页, 上下页 -->
              <a class="prev">上一章</a>
              <a class="dir">目录</a>
              <a class="next">下一章</a>
            </div>
          </div>
        </div>
      </div>
      <div class="readPopup qrBox" style="display: none">
        <a class="closePopup" />
        <div class="popupTit">
          <h3>手机阅读</h3>
        </div>
        <div class="qrList">
          <ul />
        </div>
      </div>
      <div class="readPopup setupBox" style="display: none;">
        <a class="closePopup" />
        <div class="popupTit">
          <h3>设置</h3>
        </div>
        <div class="setupList">
          <ul>
            <li class="readTheme">
              <em class="tit">阅读主题：</em>
              <a class="white current" title="白色" />
              <a class="green" title="绿色" />
              <a class="pink" title="粉色" />
              <a class="yellow" title="黄色" />
              <a class="gray" title="灰色"/>
              <a class="night" title="夜间"/>
            </li>
            <li class="setFont setBtn">
              <em class="tit">正文字体：</em> <a class="setYahei current">雅黑</a>
              <a class="setSimsun">宋体</a>
              <a class="setKs">楷书</a>
            </li>
            <li class="fontSize setBtn">
              <em class="tit">字体大小：</em>
              <a class="small">A-</a><span class="current_font">16</span>
              <a class="big">A+</a>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

import {getBookContent} from "@/api/book";
import SvgIcon from "@/components/svg/SvgIcon";

export default {
  name: "Page",
  components: {SvgIcon},
  data() {
    return {
      bookId: '',  // 书籍id
      indexId: '',   // 书籍id, string
      book: {},    // 书籍信息
      bookIndex: {},  // 书籍章节信息
      bookContent: {  // 书籍内容
        content: ''
      },
      preIndexId: '',
      nextIndexId: ''
    }
  },
  created() {
    this.initData()
    this.getContentDetail()
  },
  mounted() {
    window.addEventListener('scroll', this.handleScroll)  // 监听（绑定）滚轮滚动事件
  },
  destroyed() {
    window.removeEventListener('scroll', this.handleScroll)   //  离开页面清除（移除）滚轮滚动事件
  },
  methods: {
    initData() {
      this.bookId = this.$route.params.bookId
      this.indexId = this.$route.params.indexId
    },
    goBookDetail(bookId) {
      this.$router.push({
        name: 'BookDetail',
        params: { bookId }
      })
    },
    async getContentDetail() {
      const { data } = await getBookContent(this.bookId, this.indexId)
      this.book = data.book_detail
      this.bookIndex = data.index_detail
      this.bookContent = data.book_content
      this.preIndexId = data.pre_book_index_id
      this.nextIndexId = data.next_book_index_id
      console.log(data, 'data....')
    },
    handleScroll() {
      let clientHeight = document.documentElement.clientHeight || document.body.clientHeight
      // 比例: 160 / 720, 基础比例, 当滚轮到整体比例的160/720时候, 固定
      let base = 160 / 720
      // 设备/屏幕高度
      let scrollY = window.pageYOffset ||
        document.documentElement.scrollTop ||
        document.body.scrollTop
      let readMenu = document.getElementById('readMenu') // 滚动区域
      if (scrollY / clientHeight >= base) {
        readMenu.style.position = 'fixed'
        readMenu.style.top = '2px'
      } else {
        readMenu.style.position = ''
        readMenu.style.top = ''
      }
    }
  },
}

</script>

<style scoped lang="scss">
  .menu-icon {
    font-size: 22px;
    margin: 10px calc(50% - 11px) 0;
  }
</style>
