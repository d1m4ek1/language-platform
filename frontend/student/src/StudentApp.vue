<script lang="ts" setup>
import Header from './components/Header.vue';
import { useClient } from './stores/useClient.ts';

const store = useClient()
store.sendGetDataClient()
</script>

<template>
  <Header/>
  <main class='component-router-loader'>
    <RouterView/>
  </main>
</template>
