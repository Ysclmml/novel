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
            <td class="style"><a href="">[{{ item.cat_name }}]</a></td>
            <td class="name"><a href="">{{ item.book_name }}</a></td>
            <td class="chapter">
              <a href="">{{ item.last_index_name }}</a>
              <i class="" />
            </td>
            <td class="author"><a href="">{{ item.author_name }}</a></td>
            <td class="time">{{ item.last_index_update_time | timeFormat }}</td>
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
  filters: {
    timeFormat(time) {
      const d = new Date(time)
      return dateFormat(d.getTime(), 'MM/dd hh:mm')
    }
  },
  props: {
    booksList: {
      type: Array,
      default: () => { return [] }
    }
  }
}
</script>

<style scoped>

</style>
