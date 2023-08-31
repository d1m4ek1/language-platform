<script lang='ts' setup>

</script>

<template>
  <li class='lesson-item'>
    <a href='#'></a>
    <h4>Lesson theme</h4>
    <p>date of created</p>
    <p>deadline</p>
    <p>status</p>
  </li>
</template>

<style scoped>
.lesson-item {
  position: relative;
  display: flex;
  justify-content: space-between;
  padding: 10px;
  box-shadow: 0 0 10px -2px #0003;
  border-radius: 10px;
  background-color: white;
  transition: .3s;
  margin-bottom: 10px;
}
.lesson-item > a {
  position: absolute;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
}
.lesson-item:hover {
  transform: scale(1.01);
}
</style>