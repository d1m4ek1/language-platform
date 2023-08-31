<script lang='ts' setup>

import CourseBlock from '../components/CourseBlock.vue';
</script>

<template>
  <div class="component">
    <section class="title">
      <h1>Мое обучение</h1>
    </section>
    <section class="component-content">
      <h2>Список курсов</h2>
      <div class='course-block-list'>
        <CourseBlock/>
      </div>
    </section>
  </div>
</template>

<style scoped>
.course-block-list {
  margin-top: 20px;
}
</style>