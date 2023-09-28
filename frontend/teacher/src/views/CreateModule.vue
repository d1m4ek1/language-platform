<script>
import AddLesson from '../components/AddLesson/AddLesson.vue';
import { useModule } from '../stores/useModule';
import FileLoader from '../components/FileLoader.vue';
import APIFileLoader from '../models/fileLoader';
import { useClient } from '../stores/useClient';

const store = useModule();

export default {
  props: {
    isEditModule: Boolean
  },
  data() {
    return {
      dataModule: {
        name: '',
        language: 'english',
        description: '',
        price: 0,
        free: false,
        hidden: false,
        mainVideo: '',
        mainImg: '',
        lessonItems: [],
        subStudents: []
      },
      files: new FormData(),
      isUploadLessonFiles: false,
      moduleId: 0,
      timeoutWatchDataModule: null,
      errorCreateModule: null,
      isButtonCreateModuleDisabled: true,
      isEditDataLoaded: false,
      storeClient: useClient(),
      subStudentsObjects: []
    };
  },
  watch: {
    isEditModule: {
      immediate: true,
      async handler(n) {
        if (n) {
          await store.sendGetModuleByID(this.moduleId).then((response) => {
            const data = response.moduleData;

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
                name: i.name,
                deadlineDate: new Date(i.deadlineDate)
              })).sort((a, b) => a.lessonNumber > b.lessonNumber) : [];

            this.dataModule.mainImg = data.mainImage;
            this.dataModule.mainVideo = data.mainVideo;
            store.subStudents = data.subStudents;

            this.isEditDataLoaded = true;
          });
          Object.keys(this.dataModule).forEach((key) => {
            if (store[key] === undefined) return;

            this.dataModule[key] = store[key];
          });
        }
      }
    },
    dataModule: {
      handler(n) {
        clearTimeout(this.timeoutWatchDataModule);
        this.timeoutWatchDataModule = setTimeout(() => {
          if (this.dataModule.name !== '' && this.dataModule.lessonItems.length !== 0
            && this.dataModule.subStudents.length !== 0) {
            this.isButtonCreateModuleDisabled = false;
          }
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
      this.dataModule.price = +value;
      event.target.value = value;
    },
    setLessonComponent(idx) {
      this.dataModule.lessonItems.push({
        lessonNumber: this.dataModule.lessonItems.length + 1,
        name: 'Название урока',
        exerciseItems: [],
        taskDescription: ''
      });
    },
    addAndRemoveSubStudents(id, e) {
      e.preventDefault();
      if (this.dataModule.subStudents.includes(id)) {
        this.dataModule.subStudents = this.dataModule.subStudents.filter(i => i !== id);
      } else {
        this.dataModule.subStudents.push(id);
      }
    },
    async fileUpload() {
      const responseFileLoader = await APIFileLoader.fileLoad(
        this.files,
        'main_files',
        `module_id=${this.moduleId}`
      );
      const jsonFileLoader = await responseFileLoader.json();
      if (!jsonFileLoader.success)
        return (this.errorCreateModule = jsonFileLoader);
    },
    async saveModule() {
      if (this.isButtonCreateModuleDisabled) return;
      this.isButtonCreateModuleDisabled = true;

      const deleteFiles = {};
      if (this.files.get('mainFile') !== null) {
        if (this.files.get('mainFile').type.includes('image/')) deleteFiles['mainFileImg'] = this.dataModule.mainImg;
        if (this.files.get('mainFile').type.includes('video/')) deleteFiles['mainFileVideo'] = this.dataModule.mainVideo;
      }

      const responseSaveModule = await store.sendSaveModule(deleteFiles);
      const jsonSaveModule = await responseSaveModule.json();
      if (jsonSaveModule.status !== 200)
        return (this.errorCreateModule = jsonSaveModule);
      this.moduleId = jsonSaveModule.moduleId;

      this.isUploadLessonFiles = true;

      setTimeout(async () => {
        const responseSaveLessons = await store.sendSaveLessonForSavedModule(
          this.moduleId
        );
        const jsonSaveLessons = await responseSaveLessons.json();
        if (!jsonSaveLessons.success)
          return (this.errorCreateModule = jsonSaveLessons);

        // if (jsonFileLoader.success) {
        //   window.location.href = "/teacher/list-modules"
        // }
      }, 3000);

      await this.fileUpload();
    },
    async createModule() {
      if (this.isButtonCreateModuleDisabled) return;
      this.isButtonCreateModuleDisabled = true;

      const responseCreateModule = await store.sendCreateModule();
      const jsonCreateModule = await responseCreateModule.json();
      if (!jsonCreateModule.success)
        return (this.errorCreateModule = jsonCreateModule);
      this.moduleId = jsonCreateModule.moduleId;

      this.isUploadLessonFiles = true;

      setTimeout(async () => {
        const responseAddLessons = await store.sendAddLessonsToCreatedModule(
          this.moduleId
        );
        const jsonAddLessons = await responseAddLessons.json();
        if (!jsonAddLessons.success)
          return (this.errorCreateModule = jsonAddLessons);

        // if (jsonFileLoader.success) {
        //   window.location.href = "/teacher/list-modules"
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

<template>
  <div class='component'>
    <section class='title'>
      <h1 v-if='!isEditModule'>Создать курс</h1>
      <h1 v-else>Редактировать курс</h1>
    </section>
    <section class='component-block-input'>
      <h2>Главная информация для курса</h2>
      <div class='data-input'>
        <p>Название курса</p>
        <input
          v-model='dataModule.name'
          class='input'
          maxlength='200'
          name='name_module'
          type='text'
        />
      </div>

      <div class='data-input'>
        <p>Выберите язык курса</p>
        <select v-model='dataModule.language' name='language_module'>
          <option value='english'>Английский</option>
        </select>
      </div>

      <div class='data-input'>
        <p>Описание курса</p>
        <textarea
          v-model='dataModule.description'
          maxlength='5000'
          name='module_description'
        ></textarea>
      </div>

      <div class='data-input'>
        <p>Цена курса</p>
        <div class='block_checkbox'>
          <input
            id='free_module'
            v-model='dataModule.free'
            class='input_checkbox'
            name='free_module'
            type='checkbox'
          />
          <label for='free_module'>Бесплатный курс?</label>
        </div>
        <input
          id='price'
          :class="{ 'input-disabled': this.dataModule.free }"
          :disabled='this.dataModule.free'
          class='input number_input'
          maxlength='6'
          name='price_module'
          type='text'
          @input='setPrice($event)'
        />
      </div>

      <div class='data-input'>
        <div class='block_checkbox'>
          <input
            id='hidden_module'
            v-model='dataModule.hidden'
            class='input_checkbox'
            name='hidden_module'
            type='checkbox'
          />
          <label for='hidden_module'>Скрыть курс?</label>
        </div>
      </div>

      <div class='file-loader-section'>
        <FileLoader
          :edit-preview='dataModule.mainImg'
          :id-for="'file_loader_image'"
          :name-form-data="'mainFile'"
          :type-file="'image'"
          @on-upload-file='onUploadFile'
        ></FileLoader>
        <FileLoader
          :edit-preview='dataModule.mainVideo'
          :id-for="'file_loader_video'"
          :name-form-data="'mainFile'"
          :type-file="'video'"
          @on-upload-file='onUploadFile'
        ></FileLoader>
      </div>
    </section>
    <section class='create-lessons'>
      <div
        v-for='(item, idx) in this.dataModule.lessonItems'
        :key='`add_lesson_${item.lessonNumber}`'
      >
        <AddLesson
          :is-edit-module='isEditDataLoaded'
          :is-upload-lesson-files='isUploadLessonFiles'
          :lesson-id='item.id'
          :lesson-number='item.lessonNumber'
          :module-id='moduleId'
        ></AddLesson>
      </div>
      <button class='btn-blue' @click='setLessonComponent()'>
        Добавить урок
      </button>
    </section>
    <section class='component-block-input'>
      <p>Добавьте участников этого курса</p>
      <ol v-if='isEditModule' class='list-students'>
        <li v-for='(item, idx) in storeClient.dataClient.studentItems' :key='`student_item_${idx}`'
            :class='{"sub-student": dataModule.subStudents.includes(item.id)}' class='student-item'>
          <a href='' @click='addAndRemoveSubStudents(item.id, $event)'></a>
          <figure>
            <img :src='item.avatar' alt='avatar'>
            <figcaption>
              <p>{{ item.firstName }}</p>
              <p>{{ item.lastName }}</p>
            </figcaption>
          </figure>
        </li>
      </ol>
      <ol v-else class='list-students'>
        <li v-for='(item, idx) in storeClient.dataClient.studentItems' :key='`student_edit_item_${idx}`'
            :class='{"sub-student": dataModule.subStudents.includes(item.id)}' class='student-item'>
          <a href='' @click='addAndRemoveSubStudents(item.id, $event)'></a>
          <figure>
            <img :src='item.avatar' alt='avatar'>
            <figcaption>
              <p>{{ item.firstName }}</p>
              <p>{{ item.lastName }}</p>
            </figcaption>
          </figure>
        </li>
      </ol>
    </section>
    <section class='finally'>
      <p v-if='errorCreateModule !== null' class='error-message'>
        {{ errorCreateModule.error }}
      </p>
      <button
        v-if='!isEditModule'
        :class="{ 'btn-disabled': isButtonCreateModuleDisabled }"
        :disabled='isButtonCreateModuleDisabled'
        class='btn-blue'
        @click='createModule'
      >
        Создать
      </button>
      <button v-else :class='{"btn-disabled" : isButtonCreateModuleDisabled}' :disabled='isButtonCreateModuleDisabled'
              class='btn-blue' @click='saveModule'>
        Сохранить
      </button>
    </section>
  </div>
</template>

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

.list-students {
  display: flex;
  flex-wrap: wrap;
  width: 100%;
  max-height: 500px;
  overflow-y: auto;
  overflow-x: hidden;
  padding: 10px;
  border-radius: 10px;
  border: 1px solid #cecece;
  margin-top: 10px;
}

.list-students figure > img {
  border-radius: 50%;
  width: 60px;
  height: 60px;
}

.sub-student {
  background-color: #caeeff !important;
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
  left: 0;
  width: 100%;
  height: 100%;
}
</style>
