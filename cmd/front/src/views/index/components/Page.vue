<template>
  <div class="main box_center cf">
    <div class="channelWrap channelBanner cf">
      <div class="leftBox">
        <div class="sliderContent">
          <dl id="carouseBig" class="scBigImg">
            <dd
              v-for="(item, idx) in bannerPic"
              :key="item.book_id"
              :class="bannerSelectIdx === idx ? 'on' : ''"
            >
              <a href="" class="">
                <img :src="item.pic_url" alt="" @click="goBookDetail(item.book_id)">
              </a>
            </dd>
          </dl>
          <div id="carouseSmall" class="scSmallImg">
            <ul>
              <li v-for="(item, idx) in bannerPic" :key="item.book_id" :class="bannerSelectIdx === idx ? 'on' : ''">
                <img :src="item.pic_url" alt="" @mouseover="bannerSelectIdx=idx" @click="goBookDetail(item.book_id)">
              </li>
            </ul>
          </div>
        </div>
        <div class="hot_articles">
          <dl id="topBooks1" class="hot_recommend">
            <dt>
              <router-link v-show="bannerWord[0]" :to="{name: 'BookDetail', params: { bookId: bannerWord[0] ? bannerWord[0].book_id : '0'}}">
                {{ bannerWord[0] ? bannerWord[0].book_name : '' }}
              </router-link>
            </dt>
            <dd>
              <router-link
                v-for="(item, idx) in bannerWord.slice(1, 3)"
                :key="idx"
                :to="{name: 'BookDetail', params: { bookId: item.book_id }}"
              >
                {{ item.book_name }}
              </router-link>
            </dd>
            <dd>
              <router-link
                v-for="(item, idx) in bannerWord.slice(3, 5)"
                :key="idx"
                :to="{name: 'BookDetail', params: { bookId: item.book_id }}"
              >
                {{ item.book_name }}
              </router-link>
            </dd>
          </dl>
          <dl id="topBooks2" class="hot_recommend">
            <dt>
              <router-link :to="{name: 'BookDetail', params: { bookId: bannerWord[5] ? bannerWord[5].book_id : '0' }}" />{{ bannerWord[5] ? bannerWord[5].book_name : '' }}
            </dt>
            <dd>
              <router-link
                v-for="(item, idx) in bannerWord.slice(6, 8)"
                :key="idx"
                :to="{name: 'BookDetail', params: { bookId: item.book_id }}"
              >
                {{ item.book_name }}
              </router-link>
            </dd>
            <dd>
              <router-link
                v-for="(item, idx) in bannerWord.slice(8, 10)"
                :key="idx"
                :to="{name: 'BookDetail', params: { bookId: item.book_id }}"
              >
                {{ item.book_name }}
              </router-link>
            </dd>
          </dl>
          <News />
        </div>
      </div>
      <div class="rightBox">
        <div
          id="weekcommend"
          class="title cf"
        >
          <h3>本周强推</h3>
        </div>
        <div class="rightList">
          <ul id="currentWeek">
            <li v-for="(item, index) in weekRecBooks" :key="index"
                :class="`num${index+1}` + (index === 0 ?' on' : '') "
            >
              <div class="book_name"><i>{{ index + 1 }}</i><a class="name" @click="goBookDetail(item.book_id)">{{ item.book_name }}</a></div>
              <div class="book_intro">
                <div class="cover">
                  <a><img :src="item.pic_url" alt="" @click="goBookDetail(item.book_id)"></a>
                </div>
                <a class="txt" @click="goBookDetail(item.book_id)" v-html="item.book_desc" />
              </div>
            </li>
          </ul>
        </div>
      </div>
    </div>

    <div class="channelWrap channelPic cf">
      <div class="leftBox">
        <div class="title">
          <h2 class="on">热门推荐</h2>
        </div>
        <div id="hotRecBooks" class="picRecommend cf">
          <div v-for="(item, index) in hotRecBooks" :key="index" class="itemsList">
            <a class="items_img">
              <img :src="item.pic_url" :alt="item.book_name" @click="goBookDetail(item.book_id)">
            </a>
            <div class="items_txt">
              <h4><a @click="goBookDetail(item.book_id)">{{ item.book_name }}</a></h4>
              <p class="author"><a class="">作者：{{ item.author_name }}</a></p>
              <p class="intro">
                <a @click="goBookDetail(item.book_id)" v-html="item.book_desc" />
              </p>
            </div>
          </div>
        </div>
      </div>
      <RightRank rank-type="click" :books-list="rightBooks.clickBooks" />
    </div>

    <div class="channelWrap channelPic cf">
      <div class="leftBox">
        <div class="title">
          <h2>精品推荐</h2>
        </div>
        <div id="classicBooks" class="picRecommend cf">
          <div v-for="(item, index) in classicRecBooks" :key="index" class="itemsList">
            <a class="items_img">
              <img :src="item.pic_url" :alt="item.book_name" @click="goBookDetail(item.book_id)">
            </a>
            <div class="items_txt">
              <h4><a @click="goBookDetail(item.book_id)">{{ item.book_name }}</a></h4>
              <p class="author"><a @click="goBookDetail(item.book_id)">作者：{{ item.author_name }}</a></p>
              <p class="intro">
                <a @click="goBookDetail(item.book_id)" v-html="item.book_desc" />
              </p>
            </div>
          </div>
        </div>
      </div>
      <RightRank rank-type="new" :books-list="rightBooks.newBooks" />
    </div>

    <div class="channelWrap channelTable cf">
      <UpdatePart :books-list="rightBooks.updateBooks" />
      <RightRank rank-type="update" :books-list="rightBooks.updateBooks.slice(0, 10)" />
    </div>
  </div>
</template>

<script>
import {listBookSetting, listClickRank, listNewRank, listUpdateRank} from "@/api/book";
import News from "@/views/index/components/News";
import RightRank from "@/views/index/components/RightRank";
import UpdatePart from "@/views/index/components/UpdatePart";

export default {
  name: 'Page',
  components: {UpdatePart, RightRank, News},
  data() {
    return {
      bannerPic: [],  // 轮播图
      bannerWord: [],  // 轮播图右侧的文字版书籍推荐
      weekRecBooks: [],  // 本周强推
      hotRecBooks: [],  // 热门推荐
      classicRecBooks: [],  // 精品推荐
      bannerSelectIdx: 0, // 当前轮播图选中的索引
      rightBooks: {
        clickBooks: [],
        newBooks: [],
        updateBooks: []
      }
    }
  },
  created() {
    this.getBookSetting()
    this.getRightList()
    this.bannerStart()
  },
  methods: {
    async getBookSetting() {
      // 首页设置分为5部分, 分别是左侧的轮播图第一部分, 第二部分轮播图右侧的文字版书籍推荐,第三部分本周强推部分, 第四部分  热门推荐部分, 第五部分 精品推荐
      listBookSetting().then(res => {
        this.bannerPic = res.data[0]
        this.bannerWord = res.data[1]
        this.weekRecBooks = res.data[2]
        this.hotRecBooks = res.data[3]
        this.classicRecBooks = res.data[4]
      })
    },
    async getRightList() {
      listClickRank().then(({ data }) => {
          this.rightBooks.clickBooks = data
      })

      listNewRank().then(({ data }) => {
        this.rightBooks.newBooks = data
      })

      listUpdateRank().then(({ data }) => {
        this.rightBooks.updateBooks = data
      })
    },
    // 轮播图自动轮播
    bannerStart() {
      let t = setInterval(() => {
        this.bannerSelectIdx = (this.bannerSelectIdx + 1) % 4
      }, 3000)
    },
    goBookDetail(bookId) {
      this.$router.push({
        name: 'BookDetail',
        params: { bookId }
      })
    }
  }
}
</script>

<style scoped>

</style>
