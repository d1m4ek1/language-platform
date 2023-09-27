<script setup>
import { RouterView } from 'vue-router';
import Header from './components/Header.vue';
import { useClient } from './stores/useClient';

const store = useClient()
store.sendGetDataClient()
</script>

<template>
  <Header></Header>
  <main class="component-router-loader">
    <RouterView />
  </main>
</template>
