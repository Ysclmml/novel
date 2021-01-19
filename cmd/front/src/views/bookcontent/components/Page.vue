<template>
  <div>
    <div>
      <div class="readBody cf">
        <div class="readMain cf">
          <!-- 左侧工具栏 -->
          <div class="read_menu">
            <div id="readMenu" class="menu_left">
              <ul>
                <!-- todo: 跳转到目录页 -->
                <li>
                  <router-link :to="{name: 'BookIndexList', params: {bookId}}" class="ico_catalog" title="目录">
                    <svg-icon icon-class="list" class="menu-icon" />
                    <b>目录</b>
                  </router-link>
                </li>
                <li>
                  <a class="ico_page" title="返回书页" @click="goBookDetail(bookId)">
                    <svg-icon icon-class="book" class="menu-icon" />
                    <b>书页</b>
                  </a>
                </li>
                <li class="li_shelf" title="加入书架">
                  <a class="ico_shelf">
                    <svg-icon icon-class="marker" class="menu-icon" />
                    <b>加书架</b>
                  </a>
                </li>
                <li class="li_shelfed" style="display: none;">
                  <svg-icon icon-class="marker" class="menu-icon" />
                  <a class="ico_shelfed" href="#" title="已收藏"><b>已收藏</b></a>
                </li>
                <li>
                  <router-link :to="{name: 'BookComment', params: { bookId}}" class="ico_comment" title="评论">
                    <svg-icon icon-class="comment" class="menu-icon" />
                    <b>评论</b>
                  </router-link>
                </li>
                <li>
                  <a class="ico_setup" title="设置" @click="readSettings.isShow=!readSettings.isShow">
                    <svg-icon icon-class="settings" class="menu-icon" />
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
              <a href="/">首页 </a>&gt; <router-link :to="{name: 'BookClass', query: {cat_id: book.cat_id}}"> {{ book.cat_name }}</router-link>&gt;
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
                  <div
                    class="readBox"
                    :style="{'font-size': readSettings.readFont + 'px', 'font-family': readSettings.readFontFamily}"
                    v-html="bookContent.content">
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
              <el-button
                :disabled="preIndexId === '0'"
                :title="preIndexId === '0' ? '当前已经是第一章': '上一章'"
                @click="goIndexPage(preIndexId)"
              >
                上一章
              </el-button>
              <el-button @click="goIndexListPage()">目录</el-button>
              <el-button
                :disabled="nextIndexId === '0'"
                :title="nextIndexId === '0' ? '当前已经最后一章了': '下一章'"
                @click="goIndexPage(nextIndexId)"
              >
                下一章
              </el-button>
            </div>
          </div>
        </div>
      </div>
      <!-- 设置的弹窗 -->
      <div v-show="readSettings.isShow" class="readPopup setupBox">
        <div @click="readSettings.isShow=false">
          <svg-icon icon-class="close" class="closePopup"></svg-icon>
        </div>
        <div class="popupTit">
          <h3>设置</h3>
        </div>
        <div class="setupList">
          <ul>
            <li class="readTheme">
              <em class="tit">阅读主题：</em>
              <!-- 下面可以把属性全写在一个对象里, 然后做遍历, 优化, 现在就不做优化了. -->
              <a class="white" :class="{current: readSettings.readBgColor === BgColor.White}" title="白色" @click="SetBackUpColor(BgColor.White)" />
              <a class="green" :class="{current: readSettings.readBgColor === BgColor.Green}" title="绿色" @click="SetBackUpColor(BgColor.Green)" />
              <a class="pink" :class="{current: readSettings.readBgColor === BgColor.Pink}" title="粉色" @click="SetBackUpColor(BgColor.Pink)" />
              <a class="yellow" :class="{current: readSettings.readBgColor === BgColor.Yellow}" title="黄色" @click="SetBackUpColor(BgColor.Yellow)" />
              <a class="gray" :class="{current: readSettings.readBgColor === BgColor.Gray}" title="灰色" @click="SetBackUpColor(BgColor.Gray)" />
              <a class="night" :class="{current: readSettings.readBgColor === BgColor.Night}" title="夜间" @click="SetBackUpColor(BgColor.Night)" />
            </li>
            <li class="setFont setBtn">
              <em class="tit">正文字体：</em>
              <a :class="{setYahei: true, current: readSettings.readFontFamily === FontType.MicrosoftYaHei}" @click="SetReadFontFamily(FontType.MicrosoftYaHei)">雅黑</a>
              <a :class="{setSimsun: true, current: readSettings.readFontFamily === FontType.Simsun}" @click="SetReadFontFamily(FontType.Simsun)">宋体</a>
              <a :class="{setKs: true, current: readSettings.readFontFamily === FontType.KaiTi}" @click="SetReadFontFamily(FontType.KaiTi)">楷书</a>
            </li>
            <li class="fontSize setBtn">
              <em class="tit">字体大小：</em>
              <a class="small" @click="SetReadFont(readSettings.readFont - 1)">A-</a><span class="current_font">{{ readSettings.readFont }}</span>
              <a class="big" @click="SetReadFont(readSettings.readFont + 1)">A+</a>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

import {getBookContent} from "@/api/book";
import myLocal from "@/utils/mylocal";

// 背景颜色映射
const BgColor = {
  White: 1,
  Green: 2,
  Pink: 3,
  Yellow: 4,
  Gray: 5,
  Night: 6,
}

// 背景颜色映射
const FontType = {
  MicrosoftYaHei: 'microsoft yahei', // 微软雅黑
  KaiTi: 'kaiti',  // 楷体
  Simsun: 'Simsun',  // 宋体
}

export default {
  name: "Page",
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
      nextIndexId: '',
      readSettings: {
        isShow: false,
        readBgColor: BgColor.White, // 默认白色
        readFont: 16,  // 16px
        readFontFamily: 0, // 白色主题
      },
      BgColor, // 背景颜色映射
      FontType,
    }
  },
  created() {
    this.initial()
    this.getContentDetail()
  },
  mounted() {
    window.addEventListener('scroll', this.handleScroll)  // 监听（绑定）滚轮滚动事件
  },
  destroyed() {
    window.removeEventListener('scroll', this.handleScroll)   //  离开页面清除（移除）滚轮滚动事件

    // 清楚body设置的样式属性
    document.body.className = ''
  },
  methods: {
    initial() {
      this.bookId = this.$route.params.bookId
      this.indexId = this.$route.params.indexId
      // 读取本地阅读设置
      let colorNum = myLocal.getLocal('colorNum')
      if (colorNum) {
        this.readSettings.readBgColor = colorNum
      }
      this.SetBackUpColor(this.readSettings.readBgColor)
    },
    goBookDetail(bookId) {
      this.$router.push({
        name: 'BookDetail',
        params: { bookId }
      })
    },
    // 上一章或下一章
    goIndexPage(indexId) {
      this.$router.push({
        name: 'BookContent',
        params: { bookId: this.bookId, indexId: indexId}
      })
    },
    goIndexListPage() {
      this.$router.push({
        name: 'BookIndexList',
        params: {bookId: this.bookId}
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
    },
    SetReadFontFamily: function (fontType) {
      // todo: 全局管理localStorage的值.
      myLocal.setLocal('fontType', fontType)
      this.readSettings.readFontFamily = fontType
    },
    // 设置背景颜色
    SetBackUpColor: function (colorNum) {
      myLocal.setLocal('colorNum', colorNum)
      document.body.className = 'read_style_' + colorNum

    },
    // 设置字体大小
    SetReadFont: function (fontNum) {
      if (fontNum < 8) {
        fontNum = 8
      }
      if (fontNum > 48) {
        fontNum = 48
      }
      myLocal.setLocal('fontNum', fontNum)
      this.readSettings.readFont = fontNum
    },
  },
}

</script>

<style scoped src="@/style/read.css">
</style>

<style lang="scss" scoped>

  .menu-icon {
    font-size: 22px;
    margin: 10px calc(50% - 11px) 0;
  }
  /* 翻页 */
  .nextPageBox {
    margin: 30px auto;
    text-align: center;
    width: 98%;

    .page-button {
      width: 26%;
      height: 58px;
      line-height: 58px;
      font-size: 18px;
      //display: inline-block;
      border-radius: 3px;
      text-align: center;
      opacity:.95;
      border: 1px solid rgba(0,0,0,.1);
    }

  }
</style>

<style lang="scss">
  /*颜色：浅黄白、护眼绿、粉色、浅黄、浅灰、夜间黑*/
  // 下面的样式如果写在scope里面, 给body设置属性就会失败. 
  .read_style_1 { background-color: #ebe5d8 }
  .read_style_2 { background-color: #cbdec9 }
  .read_style_3 { background-color: #edd4d4 }
  .read_style_4 { background-color: #e0cfa3 }
  .read_style_5 { background-color: #d3d3d3 }
  .read_style_6 { background-color: #0e0f0f }

  .nextPageBox {
    .el-button:focus, .el-button:hover {
      color: white;
      border-color: #c6e2ff;
      background-color: #ff8800;
    }
  }
</style>
