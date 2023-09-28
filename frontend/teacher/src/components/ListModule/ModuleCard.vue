<script setup>
import { ref } from 'vue';

const props = defineProps({
  courseData: {
    courseID: Number(),
    mainImage: String(),
    mainVideo: String(),
    name: String(),
    description: String(),
    price: Number(),
    createdDate: String(),
    editingDate: String(),
    hidden: Boolean(),
    language: String()
  }
});

const emit = defineEmits();

const createdDate = ref(props.courseData.createdDate.split('T')[0]);
const editingDate = ref(props.courseData.editingDate.split('T')[0]);

const price = ref(new Intl.NumberFormat('ru-RU', {
  style: 'currency',
  currency: 'RUB'
}).format(props.courseData.price));

function onEditCourse() {
  emit('onEditCourse', {
    courseId: props.courseData.courseID
  });
}
</script>

<template>
  <div class='course-card'>
    <div class='img-preview'>
      <div class='course-header'>
        <h1>{{ courseData.name }}</h1>
        <div class='info-date'>
          <p>Создан: {{ createdDate }}</p>
          <p>Редактирован: {{ editingDate }}</p>
        </div>
      </div>
      <img :src='courseData.mainImage'>
      <div class='course-card-control'>
        <button class='btn-blue' @click='onEditCourse'>Редактировать</button>
      </div>
    </div>
    <div class='course-card-content'>
      <div class='course-info'>
        <p v-if='courseData.hidden' class='course-hidden'>Курс скрыт</p>
        <p v-else class='course-public'>Курс публичный</p>
        <p>{{ courseData.description }}</p>
      </div>
      <div class='course-price'>
        <p v-if='courseData.price !== undefined'>Цена: {{ price }}</p>
        <p v-else>Бесплатно</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.course-card {
  max-width: 400px;
  min-width: 200px;
  width: 100%;
  height: 400px;
  border: 1px solid #cecece;
  border-radius: 10px;
  overflow: hidden;
  margin: 5px;
}

.img-preview {
  overflow: hidden;
  position: relative;
  width: 100%;
  height: 200px;
}

.img-preview > img {
  width: 100%;
}

.img-preview > .course-header {
  width: 100%;
  color: white;
  position: absolute;
  bottom: 0;
  padding: 10px 10px;
  background-color: rgba(0, 0, 0, 0.20);
}

.info-date {
  display: flex;
  justify-content: space-between;
}

.info-date > p {
  font-size: 12px;
}

.info-date > p:first-child {
  padding-right: 10px;
}

.img-preview > .course-card-control {
  position: absolute;
  top: 10px;
  right: 10px;
  z-index: 10;
}

.course-card-content {
  padding: 10px;
  display: flex;
  justify-content: space-between;
  flex-direction: column;
  height: 50%;
}

.course-hidden, .course-public {
  padding: 3px;
  border-radius: 5px;
  color: white;
  text-align: center;
}

.course-public {
  background-color: #acca8d;
}

.course-hidden {
  background-color: #cecece;
}
</style>