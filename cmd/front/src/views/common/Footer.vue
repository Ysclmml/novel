<template>
  <div>
    <div class="friend_link">
      <div class="box_center cf" id="friendLink">
        <span>友情链接：<a href="">网站a</a><a href="">网站2</a></span>
      </div>
    </div>
    <div class="footer">
      <div class="box_center cf">
        <div class="copyright">
          <ul>
            <li class="menu"><a href="/?to=mobile">手机站</a>
              <i class="line">|</i>
              <a href="/">网站首页</a>
              <i class="line">|</i>
              <a href="/about/default.html">关于我们</a>
              <i class="line">|</i>
              <a href="/about/contact.html">联系我们</a>
              <i class="line">|</i>
              <a href="/user/feedback.html">反馈留言</a><i class="line">|</i>
              <a href="/author/index.html">作家专区</a><i class="line">|</i><a href="/mobile/fiction_house.apk">客户端</a>
            </li>
            <li>Copyright (C) {{ website.domain }} All rights reserved&nbsp;&nbsp;{{ website.name }} 版权所有</li>
          </ul>
        </div>
      </div>
    </div>
  </div>

</template>

<script>
export default {
  // eslint-disable-next-line vue/match-component-file-name
  name: "Footer",
  data() {
    return {
      website: {
        domain: "novel.com",
        name: 'novel'
      }
    }
  }
}
</script>

<style scoped>
.friend_link {
  padding: 12px 0 0;
  line-height: 2.4;
  text-align: center
}

.friend_link a {
  margin: 0 10px;
  display: inline-block
}
</style>
