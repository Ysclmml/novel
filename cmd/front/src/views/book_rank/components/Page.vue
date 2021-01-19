<template>
  <div class="main box_center cf mb50">
    <div class="channelRankingContent cf">
      <div class="wrap_left fl">
        <div class="wrap_bg">
          <!-- 榜单详情 start -->
          <div class="pad20">
            <div class="book_tit">
              <div class="fl">
                <h3 id="rankName" class="font26 mt5 mb5">{{ rankName[currentType] }}</h3>
              </div>
              <a class="fr" />
            </div>
            <div class="updateTable rankTable">
              <table cellpadding="0" cellspacing="0">
                <thead>
                  <tr>
                    <th class="rank">排名</th>
                    <th class="style">类别</th>
                    <th class="name">书名</th>
                    <th class="chapter">最新章节</th>
                    <th class="author">作者</th>
                    <th class="word">字数</th>
                  </tr>
                </thead>
                <tbody id="bookRankList">
                  <tr v-for="(book, index) in rankList" :key="index">
                    <td class="rank"><i :class="'num' + (index + 1)">{{ index + 1 }}</i></td>
                    <td class="style"><router-link :to="{name: 'BookClass', query: {cat_id: book.cat_id}}">[{{ book.cat_name }}]</router-link></td>
                    <td class="name"><router-link :to="{name: 'BookDetail', params: {bookId: book.id}}">{{ book.book_name }}</router-link></td>
                    <td class="chapter"><router-link :to="{name: 'BookContent', params: {bookId: book.id, indexId: book.last_index_id}}">{{ book.last_index_name }}</router-link>
                    </td>
                    <td class="author"><a>{{ book.author_name }}</a></td>
                    <td class="word">{{ book.word_count }}</td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
          <!-- 榜单详情 end -->
        </div>
      </div>

      <div class="wrap_right fr">
        <div class="wrap_inner wrap_right_cont mb20">
          <div class="title cf noborder">
            <h3 class="on">排行榜</h3>
          </div>
          <div class="rightList2">
            <ul id="rankType">
              <li><a :class="{on: currentType === rankType.ClickBook}" @click="listRank(rankType.ClickBook)">点击榜</a></li>
              <li><a :class="{on: currentType === rankType.NewBook}" @click="listRank(rankType.NewBook)">新书榜</a></li>
              <li><a :class="{on: currentType === rankType.UpdateBook}" @click="listRank(rankType.UpdateBook)">更新榜</a></li>
              <li><a :class="{on: currentType === rankType.CommentBook}" @click="listRank(rankType.CommentBook)">评论榜</a></li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {listBooksRank} from "@/api/book";

const rankType = {
  ClickBook: 0,
  NewBook: 1,
  UpdateBook: 2,
  CommentBook: 3
}

const rankName = {
  [rankType.ClickBook]: '点击榜',
  [rankType.NewBook]: '新书榜',
  [rankType.UpdateBook]: '更新榜',
  [rankType.CommentBook]: '评论榜'
}

export default {
  name: "Page",
  data() {
    return {
      rankName,
      rankType,
      currentType: 0,  // 当前的排行类型
      rankList: []   // 书籍列表
    }
  },
  created() {
    let urlRankType = this.$route.query.rankType
    if (urlRankType) {
      this.currentType = Number(urlRankType)
    }
    this.listRank(this.currentType)
  },
  methods: {
    async listRank(type) {
      this.currentType = type
      const { data } = await listBooksRank(type)
      this.rankList = data
      console.log(this.rankList)
    }
  }
}
</script>

<style scoped>

</style>
