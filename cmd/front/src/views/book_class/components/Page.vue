<template>
  <div class="main box_center cf">
    <div class="channelWrap classTable cf">
      <div class="so_tag">
        <ul class="list">
          <li id="workDirection" class="so_pd">
            <span class="tit">作品频道：</span>
            <a :class="{on: !queryParams.work_direction}" @click="queryParams.work_direction = false;queryParams.cat_id = 0">男频</a>
            <a :class="{on: queryParams.work_direction}" @click="queryParams.work_direction = true;queryParams.cat_id = 0">女频</a>
          </li>
          <li class="so_class">
            <span class="tit">作品分类：</span>
            <span v-show="queryParams.work_direction" id="girlCategoryList" class="so_girl">
              <a :class="{on: queryParams.cat_id === '0'}" @click="queryParams.cat_id = '0'">不限</a>
              <a v-for="(cate, index) in girlCateList"
                 :key="index" :class="{on: queryParams.cat_id === cate.id}"
                 @click="queryParams.cat_id = cate.id"
              >
                {{ cate.name }}
              </a>
            </span>
            <span v-show="!queryParams.work_direction" class="so_boy">
              <a :class="{on: queryParams.cat_id == '0'}" @click="queryParams.cat_id = '0'">不限</a>
              <a v-for="(cate, index) in boyCateList"
                 :key="index"
                 :class="{on: queryParams.cat_id === cate.id}"
                 @click="queryParams.cat_id = cate.id"
              >
                {{ cate.name }}
              </a>
            </span>
          </li>
          <!--
            <li class="so_free">
            <span class="tit">是否免费：</span>
            <a href="?b=0&s=0&wb=0&wd=0&up=0&fr=0&so=0&ms=2" class="on">不限</a>
            <a href="?b=0&s=0&wb=0&wd=0&up=0&fr=1&so=0&ms=2" class="">免费作品</a>
            <a href="?b=0&s=0&wb=0&wd=0&up=0&fr=2&so=0&ms=2" class="">收费作品</a>
            </li>
          -->
          <li class="so_progress">
            <span class="tit">是否完结：</span>
            <a :class="{on: queryParams.book_status == null}" @click="queryParams.book_status = null">不限</a>
            <a :class="{on: queryParams.book_status === false}" @click="queryParams.book_status = false">连载中</a>
            <a :class="{on: queryParams.book_status === true}" @click="queryParams.book_status = true">已完结</a>
          </li>
          <li class="so_words">
            <span class="tit">作品字数：</span>
            <a :class="{on: queryParams.word_count_max === 0}" @click="queryParams.word_count_max = 0;queryParams.word_count_min = 0">不限</a>
            <a :class="{on: queryParams.word_count_max === 300000}" @click="queryParams.word_count_max = 300000;queryParams.word_count_min = 0">30万字以下</a>
            <a :class="{on: queryParams.word_count_max === 500000}" @click="queryParams.word_count_max = 500000;queryParams.word_count_min = 300000">30-50万字</a>
            <a :class="{on: queryParams.word_count_max === 1000000}" @click="queryParams.word_count_max = 1000000;queryParams.word_count_min = 500000">50-100万字</a>
            <a :class="{on: queryParams.word_count_min === 1000000}" @click="queryParams.word_count_max = 0;queryParams.word_count_min = 1000000">100万字以上</a>
          </li>
          <li class="so_update">
            <span class="tit">更新时间：</span>
            <a :class="{on: queryParams.update_period === 0}" @click="queryParams.update_period = 0">不限</a>
            <a :class="{on: queryParams.update_period === 3}" @click="queryParams.update_period = 3">三日内</a>
            <a :class="{on: queryParams.update_period === 7}" @click="queryParams.update_period = 7">七日内</a>
            <a :class="{on: queryParams.update_period === 14}" @click="queryParams.update_period = 14">半月内</a>
            <a :class="{on: queryParams.update_period === 30}" @click="queryParams.update_period = 30">一月内</a>
          </li>
          <li class="so_sort">
            <span class="tit">排序方式：</span>
            <a :class="{on: queryParams.sort_type === 0}" @click="queryParams.sort_type = 0">不限</a>
            <a :class="{on: queryParams.sort_type === 1}" @click="queryParams.sort_type = 1">更新时间</a>
            <a :class="{on: queryParams.sort_type === 2}" @click="queryParams.sort_type = 2">总字数</a>
            <a :class="{on: queryParams.sort_type === 3}" @click="queryParams.sort_type = 3">点击量</a>
          </li>
        </ul>
      </div>
    </div>

    <div class="channelWrap channelClassContent cf">
      <div class="updateTable rankTable">
        <table cellpadding="0" cellspacing="0">
          <thead>
            <tr>
              <th class="rank">序号</th>
              <th class="style">类别</th>
              <th class="name">书名</th>
              <th class="chapter">最新章节</th>
              <th class="author">作者</th>
              <th class="word">字数</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(book, index) in searchList" :key="index">
              <td class="rank"><i>{{ index + 1 }}</i></td>
              <td class="style"><a>[{{ book.cat_name }}]</a></td>
              <td class="name"><a>{{ book.book_name }}</a></td>
              <td class="chapter"><a>{{ book.last_index_name }}</a>
              </td>
              <td class="author"><a href="javascript:void(0)">{{ book.author_name }}</a></td>
              <td class="word">{{ book.word_count }}</td>
            </tr>
          </tbody>
        </table>

        <el-pagination
          class="pageBox cf"
          background
          layout="prev, pager, next"
          :total="searchCount"
          :page-size="queryParams.page_size"
          :current-page.sync="queryParams.page"
          @current-change="handleCurrentChange"
        >
        </el-pagination>
      </div>
    </div>
  </div>
</template>

<script>

import {listBookCategory, searchBookPage} from "@/api/book"

const queryParams = {
  work_direction: false,
  cat_id: '0',
  is_vip: null,
  book_status: null,
  word_count_min: 0,
  word_count_max: 0,
  update_period: 0,
  sort_type: 0,  // 0 不限, 1 更新时间, 2 字数  3 点击
  page: 1,
  page_size: 20
}

export default {
  name: "Page",
  data() {
    return {
      boyCateList:[],
      girlCateList:[],
      searchList: [],
      searchCount: 0,
      searchParams: new Map(),
      queryParams
    }
  },
  watch: {
    queryParams: {
      handler() {
        this.getSearchList()
      },
      deep: true
    }
  },
  async created() {
    listBookCategory(queryParams).then(({ data }) => {
      for (let cat of data) {
        if (cat.work_direction) {
          this.girlCateList.push(cat)
        } else {
          this.boyCateList.push(cat)
        }
      }
      // 初始化初始请求参数
      let cat_id = this.$route.query.cat_id
      let work_direction = false
      if (cat_id) {
        cat_id = cat_id.toString()
        for (let cat of this.girlCateList) {
          if (cat.id === cat_id) {
            work_direction = true
            break
          }
        }
        this.queryParams = Object.assign(this.queryParams, {cat_id, work_direction})
      } else {
        this.getSearchList(queryParams)
      }
    })
  },
  methods: {
    async getSearchList() {
      const { data } = await searchBookPage(queryParams)
      this.searchList = data.result
      this.searchCount = data.total
    },
    handleCurrentChange(currentPage) {
      this.queryParams.page = currentPage
    }
  }
}
</script>

<style scoped>

</style>
