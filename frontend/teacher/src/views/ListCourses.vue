<script setup>
import { onMounted, ref } from 'vue';
import { useListCourses } from '../stores/useListCourses';
import CourseCard from '../components/ListCourses/CourseCard.vue';
import { useCourse } from '../stores/useCourse';
import CreateCourse from './CreateCourse.vue';

const store = useListCourses();
const storeCourse = useCourse();

let list = ref([]);
let courseId = ref(0);

onMounted(() => {
  if (
    store.getListOfCourses === undefined ||
    store.getListOfCourses.length === 0
  ) {
    store.sendGetListOfCourses().then((response) => {
      if (response) list.value = store.listOfCourses;
    });
  } else {
    list.value = store.getListOfCourses;
  }
});

const isFlagFilter = ref('all');

function setFlagFilter(param) {
  isFlagFilter.value = param;
  if (
    store.getListOfCourses !== undefined &&
    store.getListOfCourses.length !== 0
  )
    filterListOfCourses();
}

function filterListOfCourses() {
  if (isFlagFilter.value === 'all') list.value = store.getListOfCourses;
  if (isFlagFilter.value === 'public')
    list.value = store.getListOfCourses.filter((i) => !i.hidden);
  if (isFlagFilter.value === 'hidden')
    list.value = store.getListOfCourses.filter((i) => i.hidden);
}

function onEditCourse(data) {
  storeCourse.courseId = data.courseId;
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
          store.getListOfCourses !== undefined &&
          store.getListOfCourses.length !== 0 &&
          isFlagFilter !== 'edit'
        "
        class="course-list"
      >
        <p v-if="list.length === 0 && isFlagFilter === 'public'">
          У вас нет публичных курсов
        </p>
        <p v-if="list.length === 0 && isFlagFilter === 'hidden'">
          У вас нет скрытых курсов
        </p>
        <div v-for="(item, idx) in list" :key="`course_card_${idx}`">
          <CourseCard
            :course-data="item"
            @on-edit-course="onEditCourse"
          ></CourseCard>
        </div>
      </div>
      <div v-else-if="isFlagFilter === 'edit'" class="course-edit">
        <CreateCourse :is-edit-course="true"></CreateCourse>
      </div>
      <div v-else class="course-list">
        <p>
          У вас пока нет курсов.
          <RouterLink to="/teacher/create-course">Создать курс?</RouterLink>
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

.course-list {
  display: flex;
  flex-wrap: wrap;
}

.edit-active {
  color: #acca8d;
}
</style>
