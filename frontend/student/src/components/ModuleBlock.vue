<script lang='ts' setup>
import LessonItem from './LessonItem.vue';

const def = "/default/avatar/1.png"
</script>

<template>
  <div class='module-block'>
    <section class='module-block-header'>
      <h2>Name module</h2>
      <div class='teacher-info'>
        <div class='img'>
          <img :src='def' alt='avatar'>
        </div>
        <p>Teacher Name</p>
      </div>
    </section>
    <section class='module-block-content'>
      <p>Список уроков</p>
      <ul class='lesson-items'>
        <LessonItem v-for='_ in 10'/>
      </ul>
    </section>
  </div>
</template>

<style scoped>
.module-block {
  width: 500px;
  height: 500px;
  display: flex;
  flex-direction: column;
  border-radius: 10px;
  border: 1px solid #b5b5b5;
  overflow: hidden;
}
.module-block-header {
  background-color: #dbe8ff;
  width: 100%;
  height: 80px;
  padding: 10px;
}
.module-block-header > h1 {
  margin-bottom: 10px;
}
.teacher-info {
  display: flex;
  align-items: center;
}
.teacher-info > p {
  margin-left: 10px;
}
.img {
  width: 30px;
  height: 30px;
  border-radius: 50%;
  overflow: hidden;
}
.img > img {
  width: 100%;
  height: 100%;
}

.module-block-content {
  height: 84%;
  padding: 10px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}
.module-block-content > p {
  margin-bottom: 10px;
}
.lesson-items {
  display: flex;
  flex-direction: column;
  overflow-y: auto;
  border-radius: 10px;
  border: 1px solid #b5b5b5;
  padding: 10px;
  height: 95%;
}
</style>