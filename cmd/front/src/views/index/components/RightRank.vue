<!-- 右侧排行榜信息 -->
<template>
  <div>
    <div class="rightBox ">
      <div class="title cf">
        <h3 class="on">{{ rankName[rankType] }}</h3>
      </div>
      <div class="rightList">
        <ul>
          <li v-for="(item, index) in booksList" :key="index" :class="`num${index+1}` + (index === 0 ?' on' : '')">
            <div class="book_name">
              <i>{{ index + 1 }}</i><a class="name" href="">{{ item.book_name }}</a>
            </div>
            <div class="book_intro">
              <div class="cover">
                <a href=""><img :src="item.pic_url" :alt="item.book_name"></a>
              </div>
              <a class="txt" href="" v-html="item.book_desc" />
            </div>
          </li>
        </ul>
        <div class="more"><a href="">查看更多&gt;</a></div>
      </div>
    </div>
  </div>
</template>

<script>
const rankName = {
  click: '点击榜单',
  new: '新书榜单',
  update: '更新榜单',
}

export default {
  name: "RightRank",
  props: {
    rankType: {
      type: String,
      default: 'click'  // new | click | update 对应 新书榜 | 点击榜 | 更新榜
    },
    booksList: {
      type: Array,
      default: () => { return [] }
    }
  },
  data() {
    return {
      rankName: rankName
    }
  }
}
</script>

<style scoped>

</style>
