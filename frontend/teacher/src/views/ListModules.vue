<script setup>
import { onMounted, ref } from 'vue';
import { useListModule } from '../stores/useListModule';
import ModuleCard from '../components/ListModule/ModuleCard.vue';
import { useModule } from '../stores/useModule';
import CreateModule from './CreateModule.vue';

const store = useListModule();
const storeModule = useModule();

let list = ref([]);
let moduleId = ref(0);

onMounted(() => {
  if (
    store.getListOfModules === undefined ||
    store.getListOfModules.length === 0
  ) {
    store.sendGetListOfModules().then((response) => {
      if (response) list.value = store.listOfModules;
    });
  } else {
    list.value = store.getListOfModules;
  }
});

const isFlagFilter = ref('all');

function setFlagFilter(param) {
  isFlagFilter.value = param;
  if (
    store.getListOfModules !== undefined &&
    store.getListOfModules.length !== 0
  )
    filterListOfModules();
}

function filterListOfModules() {
  if (isFlagFilter.value === 'all') list.value = store.getListOfModules;
  if (isFlagFilter.value === 'public')
    list.value = store.getListOfModules.filter((i) => !i.hidden);
  if (isFlagFilter.value === 'hidden')
    list.value = store.getListOfModules.filter((i) => i.hidden);
}

function onEditModule(data) {
  storeModule.moduleId = data.moduleId;
  isFlagFilter.value = 'edit';
}
</script>

<template>
  <div class="component">
    <section class="title">
      <h1>Список курсов</h1>
    </section>
    <section class="component-content">
      <h2 class="folders">
        <span
          :class="{ underline: isFlagFilter === 'all' }"
          class="link"
          @click="setFlagFilter('all')"
          >Все курсы</span
        >/
        <span
          :class="{ underline: isFlagFilter === 'public' }"
          class="link"
          @click="setFlagFilter('public')"
          >Публичные курсы</span
        >/
        <span
          :class="{ underline: isFlagFilter === 'hidden' }"
          class="link"
          @click="setFlagFilter('hidden')"
          >Скрытые курсы</span
        >/
        <span :class="{ 'underline edit-active': isFlagFilter === 'edit' }"
          >Редактирование курса</span
        >
      </h2>

      <div
        v-if="
          store.getListOfModules !== undefined &&
          store.getListOfModules.length !== 0 &&
          isFlagFilter !== 'edit'
        "
        class="module-list"
      >
        <p v-if="list.length === 0 && isFlagFilter === 'public'">
          У вас нет публичных курсов
        </p>
        <p v-if="list.length === 0 && isFlagFilter === 'hidden'">
          У вас нет скрытых курсов
        </p>
        <div v-for="(item, idx) in list" :key="`module_card_${idx}`">
          <ModuleCard
            :module-data="item"
            @on-edit-module="onEditModule"
          ></ModuleCard>
        </div>
      </div>
      <div v-else-if="isFlagFilter === 'edit'" class="module-edit">
        <CreateModule :is-edit-module="true"></CreateModule>
      </div>
      <div v-else class="module-list">
        <p>
          У вас пока нет курсов.
          <RouterLink to="/create-module">Создать курс?</RouterLink>
        </p>
      </div>
    </section>
  </div>
</template>

<style scoped>
.folders {
  padding-bottom: 50px;
}

.folders > span:first-child {
  padding-right: 10px;
}

.folders > span:not(:first-child) {
  padding: 0 10px;
}

.module-list {
  display: flex;
  flex-wrap: wrap;
}

.edit-active {
  color: #acca8d;
}
</style>
