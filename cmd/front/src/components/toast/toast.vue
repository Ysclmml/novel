<template>
  <div class="toast">
    <div class="app-toast"
         v-if="isShow"
         :class="{'info': type=== 'info','success': type=== 'success','wraning': type=== 'wraning','danger': type=== 'danger'}">
      {{text}}
    </div>
  </div>
</template>

<script>
export default {

}
</script>

<style scoped>
.app-toast {
  position: fixed;
  left: 50%;
  top: 10px;
  background: #ccc;
  padding: 10px;
  border-radius: 5px;
  transform: translate(-50%, 0);
  color: #fff;
}
.info {
  background: #00aaee;
}
.success {
  background: #00ee6b;
}
.wraning {
  background: #eea300;
}
.danger {
  background: #ee000c;
}
</style>