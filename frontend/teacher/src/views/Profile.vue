<script setup>
import { useClient } from '../stores/useClient';

const store = useClient()
</script>

<template>
  <div class='component'>
    <section class='title'>
      <figure class='block-avatar'>
        <img :src='store.getAvatar' alt='Аватар'>
        <figcaption>
          <h1>{{ store.getFullName }}</h1>
          <p>{{store.getTextIsTeacher}}</p>
          <p>Дата регистрации: {{ store.getDate }}</p>
        </figcaption>
      </figure>
    </section>
    <section class='component-content'>
      <div class='about-me'>
        <h3>О себе</h3>
        <pre>{{store.getAboutMe}}</pre>
      </div>
      <div class='language-items'>
        <h3>Преподавание языков</h3>
        <ul v-if='store.getLanguageItems.length !== 0'>
          <li v-for='(item, idx) in store.getLanguageItems' :key='`lang_item_${idx}`'>
            {{item}}
          </li>
        </ul>
        <p v-else>Вы не указали преподаваемые языки</p>
      </div>
    </section>
  </div>
</template>

<style scoped>
figure {
  display: flex;
  align-items: center;
}
figure > img {
  width: 90px;
  height: 90px;
  border-radius: 50%;
}
figcaption {
  margin-left: 20px;
}
figcaption > h1 {
  margin-bottom: 10px;
}
figcaption > p:last-child {
  font-size: 14px;
  color: #c7c7c7;
}

.component-content > div:not(:last-child) {
  margin-bottom: 20px;
}

.component-content > div > h3 {
  margin-bottom: 10px;
}

ul > li {
  list-style: disc;
  margin-left: 20px;
}
</style>