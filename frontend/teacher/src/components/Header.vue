<script setup>
import { RouterLink } from 'vue-router'
import { useClient } from '../stores/useClient';

const store = useClient()

async function signout() {
  const response = await fetch("/auth/signout", {
    method: "GET"
  })

  const json = await response.json()

  if (json.success) {
    window.location.href = "/"
  }
}
</script>

<template>
  <header class="header-vertical">
    <nav class="header-vertical-navigation">
      <RouterLink to="/">{{ store.getFullName }}</RouterLink>
      <ul class="header-vertical-navigation-list">
        <li><RouterLink to="/create-course">Создать курс</RouterLink></li>
        <li><RouterLink to='/list-courses'>Список курсов</RouterLink></li>
        <li>Проверка уроков</li>
        <li>Ученики</li>
        <li><RouterLink to='/profile'>Профиль</RouterLink></li>
        <li><RouterLink to='/profile/settings'>Настройки</RouterLink></li>
        <li><a class='btn-signout' href='#' @click='signout()'>Выйти</a></li>
      </ul>
    </nav>
  </header>
</template>