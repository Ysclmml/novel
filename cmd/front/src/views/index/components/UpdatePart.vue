<!-- 最新更新部分 -->
<template>
  <div class="leftBox">
    <div class="title">
      <h2>最新更新</h2>
    </div>
    <div class="updateTable" id="newRankBooks2">
      <table
        cellpadding="0"
        cellspacing="0"
      >
        <thead>
          <tr>
            <th class="style">类别</th>
            <th class="name">书名</th>
            <th class="chapter">最新章节</th>
            <th class="author">作者</th>
            <th class="time">更新时间</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in booksList" :key="index">
            <td class="style"><a @click="goBookDetail(item.id)">[{{ item.cat_name }}]</a></td>
            <td class="name"><a @click="goBookDetail(item.id)">{{ item.book_name }}</a></td>
            <td class="chapter">
              <a>{{ item.last_index_name }}</a>
              <i class="" />
            </td>
            <td class="author"><a>{{ item.author_name }}</a></td>
            <td class="time">{{ item.last_index_update_time | timeFormat('MM/dd hh:mm') }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import {dateFormat} from "@/utils";

export default {
  name: "UpdatePart",
  props: {
    booksList: {
      type: Array,
      default: () => { return [] }
    }
  },
  methods: {
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
