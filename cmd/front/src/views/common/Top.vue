<template>
  <div class="header">
    <div class="topMain">
      <div class="box_center cf">
        <a href="/?m=2" class="logo fl"><img src="/favicon.ico"></a>
        <div class="searchBar fl">
          <div class="search cf">
            <input id="searchKey" type="text" placeholder="书名、作者、关键字" class="s_int" name="searchKey">
            <label id="btnSearch" class="search_btn" onclick="searchByK()"><i class="icon" /></label>
          </div>
        </div>
        <div id="headerUserInfo" class="bookShelf fr">
          <a class="sj_link" href="/user/favorites.html">
            <svg-icon icon-class="icon_sj" class="icon-shelf" />
            我的书架
          </a>
          <span class="user_link"><i class="line mr20">|</i><a href="/user/login.html" class="mr15">登录</a><a
            href="/user/register.html"
          >注册</a></span>
        </div>
      </div>
    </div>
    <div id="mainNav" class="mainNav">
      <div class="box_center cf">
        <ul id="navModule" class="nav">
          <li><a href="/">首页</a></li>
          <li><router-link :to="{name: 'BookClass'}">全部作品</router-link></li>
          <li><router-link :to="{name: 'BookRanking'}">排行榜</router-link></li>
          <li class=""><a href="/pay/index.html">充值</a></li>
          <li><a href="/author/index.html">作家专区</a></li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
import SvgIcon from "@/components/svg/SvgIcon";
export default {
  name: 'Top',
  components: {SvgIcon}
}
</script>

<style scoped lang="scss">
.sj_link {
  .icon-shelf {
    font-size: 23px;
    margin-left: 5px;
    margin-top: 5px;
  }
}
</style>
