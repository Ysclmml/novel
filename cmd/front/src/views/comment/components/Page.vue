<template>
  <div>
    <div class="main box_center cf mb50">
      <div class="channelBookContent cf">
        <div class="wrap_left fl">
          <div class="wrap_bg">
            <Comment :book-id="bookId" :is-show-page="true"></Comment>
          </div>
        </div>
        <div class="wrap_right fr">
          <AuthorPart :book-id="bookId"></AuthorPart>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Comment from "./Comment";
import AuthorPart from "@/views/bookdetail/components/AuthorPart";
export default {
  name: "Page",
  components: {AuthorPart, Comment},
  data() {
    return {
      bookId: 0,
    }
  },
  async created() {
    this.bookId = this.$route.params.bookId
  },
}
</script>

<style lang="scss">

</style>
