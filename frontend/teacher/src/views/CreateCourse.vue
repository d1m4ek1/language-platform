<template>
  <div class='component'>
    <section class='title'>
      <h1 v-if='!isEditCourse'>Создать курс</h1>
      <h1 v-else>Редактировать курс</h1>
    </section>
    <section class='component-block-input'>
      <h2>Главная информация для курса</h2>
      <div class='data-input'>
        <p>Название курса</p>
        <input
          v-model='dataCourse.name'
          class='input'
          maxlength='200'
          name='name_course'
          type='text'
        />
      </div>

      <div class='data-input'>
        <p>Выберите язык курса</p>
        <select v-model='dataCourse.language' name='language_course'>
          <option value='english'>Английский</option>
        </select>
      </div>

      <div class='data-input'>
        <p>Описание курса</p>
        <textarea
          v-model='dataCourse.description'
          maxlength='5000'
          name='course_description'
        ></textarea>
      </div>

      <div class='data-input'>
        <p>Цена курса</p>
        <div class='block_checkbox'>
          <input
            id='free_course'
            v-model='dataCourse.free'
            class='input_checkbox'
            name='free_course'
            type='checkbox'
          />
          <label for='free_course'>Бесплатный курс?</label>
        </div>
        <input
          id='price'
          :class="{ 'input-disabled': this.dataCourse.free }"
          :disabled='this.dataCourse.free'
          class='input number_input'
          maxlength='6'
          name='price_course'
          type='text'
          @input='setPrice($event)'
        />
      </div>

      <div class='data-input'>
        <div class='block_checkbox'>
          <input
            id='hidden_course'
            v-model='dataCourse.hidden'
            class='input_checkbox'
            name='hidden_course'
            type='checkbox'
          />
          <label for='hidden_course'>Скрыть курс?</label>
        </div>
      </div>

      <div class='file-loader-section'>
        <FileLoader
          :edit-preview='dataCourse.mainImg'
          :id-for="'file_loader_image'"
          :name-form-data="'mainFile'"
          :type-file="'image'"
          @on-upload-file='onUploadFile'
        ></FileLoader>
        <FileLoader
          :edit-preview='dataCourse.mainVideo'
          :id-for="'file_loader_video'"
          :name-form-data="'mainFile'"
          :type-file="'video'"
          @on-upload-file='onUploadFile'
        ></FileLoader>
      </div>
    </section>
    <section class='create-lessons'>
      <div
        v-for='(item, idx) in this.dataCourse.lessonItems'
        :key='`add_lesson_${item.lessonNumber}`'
      >
        <AddLesson
          :course-id='courseId'
          :is-edit-course='isEditDataLoaded'
          :is-upload-lesson-files='isUploadLessonFiles'
          :lesson-id='item.id'
          :lesson-number='item.lessonNumber'
        ></AddLesson>
      </div>
      <button class='btn-blue' @click='setLessonComponent()'>
        Добавить урок
      </button>
    </section>
    <section class='finally'>
      <p v-if='errorCreateCourse !== null' class='error-message'>
        {{ errorCreateCourse.error }}
      </p>
      <button
        v-if='!isEditCourse'
        :class="{ 'btn-disabled': isButtonCreateCourseDisabled }"
        :disabled='isButtonCreateCourseDisabled'
        class='btn-blue'
        @click='createCourse'
      >
        Создать
      </button>
      <button v-else :class='{"btn-disabled" : isButtonCreateCourseDisabled}' :disabled='isButtonCreateCourseDisabled'
              class='btn-blue' @click='saveCourse'>
        Сохранить
      </button>
    </section>
  </div>
</template>

<script>
import AddLesson from '../components/AddLesson/AddLesson.vue';
import { useCourse } from '../stores/useCourse';
import FileLoader from '../components/FileLoader.vue';
import APIFileLoader from '../models/fileLoader';

const store = useCourse();

export default {
  props: {
    isEditCourse: Boolean
  },
  data() {
    return {
      dataCourse: {
        name: '',
        language: 'english',
        description: '',
        price: 0,
        free: false,
        hidden: false,
        mainVideo: '',
        mainImg: '',
        lessonItems: []
      },
      files: new FormData(),
      isUploadLessonFiles: false,
      courseId: 0,
      timeoutWatchDataCourse: null,
      errorCreateCourse: null,
      isButtonCreateCourseDisabled: false,
      isEditDataLoaded: false
    };
  },
  watch: {
    isEditCourse: {
      immediate: true,
      async handler(n) {
        if (n) {
          await store.sendGetCourseByID(this.courseId).then((response) => {
            const data = response.courseData;

            store.name = data.name;
            store.language = data.language;
            store.description = data.description || '';

            store.price = data.price || 0;
            document.getElementById('price').value = store.price || '';

            store.hidden = data.hidden || false;
            store.free = !data.price;
            store.lessonItems = response.lessonDataItems !== undefined ?
              response.lessonDataItems.map((i) => ({
                exerciseItems: JSON.parse(i.exerciseJson),
                id: i.id,
                lessonNumber: i.lessonNumber,
                name: i.name
              })).sort((a, b) => a.lessonNumber > b.lessonNumber) : [];

            this.dataCourse.mainImg = data.mainImg;
            this.dataCourse.mainVideo = data.mainVideo;
            this.isEditDataLoaded = true;
          });
          Object.keys(this.dataCourse).forEach((key) => {
            if (store[key] === undefined) return;

            this.dataCourse[key] = store[key];
          });
        }
      }
    },
    dataCourse: {
      handler(n) {
        clearTimeout(this.timeoutWatchDataCourse);
        this.timeoutWatchDataCourse = setTimeout(() => {
          store.$patch({
            ...n
          });
        }, 1000);
      },
      immediate: true,
      deep: true
    }
  },
  methods: {
    onUploadFile(data = {}) {
      if (data.remove) {
        this.files.delete(data.nameFormData);
      }

      if (!data.success) {
        return;
      }

      this.files.append(data.nameFormData, data.file);
    },
    setPrice(event) {
      const value = event.target.value
        .replace(/[a-zа-яё\s]/gi, '')
        .replace(/[^a-zа-яё0-9\s]/gi, '');
      this.dataCourse.price = +value;
      event.target.value = value;
    },
    setLessonComponent(idx) {
      this.dataCourse.lessonItems.push({
        lessonNumber: this.dataCourse.lessonItems.length + 1,
        name: 'Название урока',
        exerciseItems: [],
        taskDescription: ''
      });
    },
    async fileUpload() {
      const responseFileLoader = await APIFileLoader.fileLoad(
        this.files,
        'main_files',
        `course_id=${this.courseId}`
      );
      const jsonFileLoader = await responseFileLoader.json();
      if (!jsonFileLoader.success)
        return (this.errorCreateCourse = jsonFileLoader);
    },
    async saveCourse() {
      if (this.isButtonCreateCourseDisabled) return;
      this.isButtonCreateCourseDisabled = true;

      const deleteFiles = {};
      if (this.files.get('mainFile') !== null) {
        if (this.files.get('mainFile').type.includes('image/')) deleteFiles['mainFileImg'] = this.dataCourse.mainImg;
        if (this.files.get('mainFile').type.includes('video/')) deleteFiles['mainFileVideo'] = this.dataCourse.mainVideo;
      }

      const responseSaveCourse = await store.sendSaveCourse(deleteFiles);
      const jsonSaveCourse = await responseSaveCourse.json();
      if (jsonSaveCourse.status !== 200)
        return (this.errorCreateCourse = jsonSaveCourse);
      this.courseId = jsonSaveCourse.courseId;

      this.isUploadLessonFiles = true;

      setTimeout(async () => {
        const responseSaveLessons = await store.sendSaveLessonForSavedCourse(
          this.courseId
        );
        const jsonSaveLessons = await responseSaveLessons.json();
        if (!jsonSaveLessons.success)
          return (this.errorCreateCourse = jsonSaveLessons);

        // if (jsonFileLoader.success) {
        //   window.location.href = "/teacher/list-courses"
        // }
      }, 3000);

      await this.fileUpload();
    },
    async createCourse() {
      if (this.isButtonCreateCourseDisabled) return;
      this.isButtonCreateCourseDisabled = true;

      const responseCreateCourse = await store.sendCreateCourse();
      const jsonCreateCourse = await responseCreateCourse.json();
      if (!jsonCreateCourse.success)
        return (this.errorCreateCourse = jsonCreateCourse);
      this.courseId = jsonCreateCourse.courseId;

      this.isUploadLessonFiles = true;

      setTimeout(async () => {
        const responseAddLessons = await store.sendAddLessonsToCreatedCourse(
          this.courseId
        );
        const jsonAddLessons = await responseAddLessons.json();
        if (!jsonAddLessons.success)
          return (this.errorCreateCourse = jsonAddLessons);

        // if (jsonFileLoader.success) {
        //   window.location.href = "/teacher/list-courses"
        // }
      }, 3000);

      await this.fileUpload();
    }
  },
  components: {
    FileLoader,
    AddLesson
  }
};
</script>

<style scoped>
.finally > p {
  margin-bottom: 10px;
}

.file-loader-section {
  display: flex;
  flex-direction: column;
}

.file-loader-section .file-loader-block:not(:last-child) {
  margin-bottom: 10px;
}
</style>
