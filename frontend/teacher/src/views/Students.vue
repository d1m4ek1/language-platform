<script setup>
import { useClient } from '../stores/useClient';

const store = useClient()
</script>

<template>
<div class='component'>
  <section class='component-content'>
    <h2>Ваши ученики</h2>
    <ol class='list-students'>
      <li v-for='(item, idx) in store.dataClient.subStudents' :key='`student_item_${idx}`' class='student-item'>
        <a :href='`/profile/${item.id}`'></a>
        <figure>
          <img :src='item.avatar' alt='avatar'>
          <figcaption>
            <p>{{item.firstName}}</p>
            <p>{{item.lastName}}</p>
          </figcaption>
        </figure>
      </li>
    </ol>
  </section>
</div>
</template>

<style scoped>
.component-content > h2 {
  margin-bottom: 30px;
}
figure > img {
  border-radius: 50%;
  width: 60px;
  height: 60px;
}
.list-students {
  display: flex;
  flex-wrap: wrap;
}
.student-item {
  display: flex;
  justify-content: space-between;
  position: relative;
  padding: 10px;
  border-radius: 10px;
  box-shadow: 0 0 10px -2px #0003;
  background-color: white;
  min-width: 300px;
  transition: .3s;
  margin: 5px;
  flex: 1 0 25%;
}
.student-item:hover {
  transform: scale(1.02);
}
.student-item > figure {
  display: flex;
  align-items: center;
}
.student-item > figure > figcaption {
  margin-left: 20px;
}
.student-item > a {
  position: absolute;
  top: 0;
  width: 100%;
  height: 100%;
}
</style>