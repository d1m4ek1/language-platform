<script lang='ts' setup>
import { RouterLink } from 'vue-router';
import { useClient } from '../stores/useClient.ts';
import { storeToRefs } from 'pinia';

const store = useClient()
const {getFullName} = storeToRefs(store)

async function signout() {
  const response = await fetch("/auth/signout", {
    method: "GET"
  })

  const json: IStatusResponse = await response.json()

  if (json.success) {
    window.location.href = "/"
  }
}
</script>

<template>
  <header class="header-vertical">
    <nav class="header-vertical-navigation">
      <RouterLink to="/">{{ getFullName }}</RouterLink>
      <ul class="header-vertical-navigation-list">
        <li><RouterLink to="/my-education">Мое обучение</RouterLink></li>
        <li>Профиль</li>
        <li>Настройки профиля</li>
        <li><a class='btn-signout' href='#' @click='signout'>Выйти</a></li>
      </ul>
    </nav>
  </header>
</template>