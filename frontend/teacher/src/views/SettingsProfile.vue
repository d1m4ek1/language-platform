<template>
  <div class='component'>
    <section class='title'>
      <figure class='block-avatar'>
        <div class='control-avatar'>
          <FileLoader :edit-preview='dataSettings.avatar' :id-for='"file_loader_avatar"' :name-form-data='"avatar"'
                      :type-file='"image"'
                      @on-upload-file='onUploadFile'></FileLoader>
        </div>
        <figcaption>
          <div class='data-input'>
            <input v-model='dataSettings.firstName' class='input' name='first_name' placeholder='Имя' type='text'>
            <input v-model='dataSettings.lastName' class='input' name='last_name' placeholder='Фамилия' type='text'>
          </div>
          <p>{{ store.getTextIsTeacher }}</p>
          <p>Дата регистрации: {{ store.getDate }}</p>
        </figcaption>
      </figure>
    </section>
    <section class='component-content'>
      <div class='about-me'>
        <h3>О себе</h3>
        <div class='contenteditable' contenteditable='true' @input='setAboutMe'
             v-html='`<pre>${store.getAboutMe}</pre>`'></div>
      </div>
      <div class='language-items'>
        <h3>Преподавание языков</h3>
        <ul v-if='dataSettings.languageItems !== null'>
          <li v-for='(item, idx) in dataSettings.languageItems' :key='`lang_item_${idx}`'>
            <div class='list-select'>
              <select v-model='item.lang'>
                <option v-for='(lang, idx) in selectOptionsLang' :key='`option_${lang}_${idx}`'
                        :disabled='lang.disabled' :value='lang.lang' @click='setDisabledSelectLanguageOption'>
                  {{ lang.value }}
                </option>
              </select>
            </div>
            <div class='list-select'>
              <select v-model='item.lvl'>
                <option v-for='(level, idx) in levelLanguageObject[item.lang]' :disabled='item.lvl === level'
                        :value='level'>{{ level }}
                </option>
              </select>
            </div>
            <button v-if='dataSettings.languageItems.length !== 1' class='btn-red' @click='deleteLanguageItem(idx)'>
              Удалить
            </button>
          </li>
        </ul>
        <button v-if='dataSettings.languageItems.length !== selectOptionsLang.length' class='btn-blue'
                @click='addLanguageItem'>Добавить язык
        </button>
      </div>
      <div class='interface-lang'>
        <h3>Язык интерфейса</h3>
        <select v-model='dataSettings.interfaceLang'>
          <option value='ru'>Русский</option>
          <option disabled value='en'>Английский</option>
        </select>
      </div>
      <div class='save-data-settings'>
        <h3>Сохранить изменения</h3>
        <button class='btn-blue' @click='saveDataSettings'>Сохранить</button>
      </div>
    </section>
  </div>
</template>

<script setup>
import { useClient } from '../stores/useClient';
import { ref, watch } from 'vue';
import FileLoader from '../components/FileLoader.vue';
import APIFileLoader from '../models/fileLoader';
import { storeToRefs } from 'pinia';

const store = useClient();
const { isGetDataClient } = storeToRefs(store);

const avatarFile = ref(new FormData());
const dataSettings = ref({
  firstName: '',
  lastName: '',
  aboutMe: '',
  languageItems: [],
  interfaceLang: '',
  avatar: '',
  deleteAvatar: '',
  isTeacher: null
});
const selectOptionsLang = ref([
  {
    lang: 'en',
    value: 'Английский',
    disabled: false
  }
  // {
  //   lang: 'arab',
  //   value: 'Арабский',
  //   disabled: false
  // },
  // {
  //   lang: 'french',
  //   value: 'Французкий',
  //   disabled: false
  // }
]);
const levelLanguageObject = {
  en: ['A0', 'A1', 'B1', 'B2', 'C1', 'C2']
};

watch(isGetDataClient, (n) => {
  if (n) {
    dataSettings.value = {
      firstName: store.dataClient.firstName,
      lastName: store.dataClient.lastName,
      aboutMe: store.dataClient.aboutMe,
      languageItems: store.dataClient.languageItems,
      interfaceLang: store.dataClient.interfaceLang,
      avatar: store.dataClient.avatar,
      deleteAvatar: '',
      isTeacher: store.dataClient.isTeacher
    };
  }
}, { immediate: true });

function onUploadFile(data) {
  if (data.remove) {
    return avatarFile.value.delete(data.nameFormData);
  }

  if (data.success) {
    avatarFile.value.append(data.nameFormData, data.file);
  }
}

function setAboutMe(e) {
  dataSettings.value.aboutMe = e.target.innerText;
}

function addLanguageItem() {
  if (dataSettings.value.languageItems.length === selectOptionsLang.value.length) return;

  for (let i = 0; i < selectOptionsLang.value.length; i++) {
    const option = selectOptionsLang.value[i];

    if (!option.disabled) {
      return dataSettings.value.languageItems.push({
        lang: '',
        lvl: ''
      });
    }
  }
}

function deleteLanguageItem(idx) {
  if (dataSettings.value.languageItems.length === 1) return;

  dataSettings.value.languageItems.splice(idx, 1);
  setDisabledSelectLanguageOption();
}

function setDisabledSelectLanguageOption() {
  for (let j = 0; j < selectOptionsLang.value.length; j++) {
    const option = selectOptionsLang.value[j];

    option.disabled = false;
  }

  for (let i = 0; i < dataSettings.value.languageItems.length; i++) {
    const item = dataSettings.value.languageItems[i];

    selectOptionsLang.value = selectOptionsLang.value.map(option => {
      if (item.lang === option.lang) option.disabled = true;
      return option;
    });
  }
}

setDisabledSelectLanguageOption();

const saveDataSettings = async () => {
  if (avatarFile.value.get('avatar') !== null) {
    try {
      await APIFileLoader.fileLoad(avatarFile.value, 'avatar')
        .then(response => response.json())
        .then(response => {
          dataSettings.value.deleteAvatar = dataSettings.value.avatar;
          dataSettings.value.avatar = response.avatar;
        });
    } catch (err) {
      console.error(err);
    }
  }

  Object.keys(dataSettings.value).forEach(key => {
    if (dataSettings.value[key] !== '' && store.dataClient[key] !== undefined) {
      store.dataClient[key] = dataSettings.value[key];
    }

  });

  await store.sandSaveDataClient(dataSettings.value);
};
</script>

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

figcaption > .data-input {
  margin-bottom: 10px;
}

figcaption > .data-input > .input:first-child {
  margin-right: 10px;
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

.language-items > ul > li {
  display: flex;
  align-items: center;
  margin-bottom: 5px;
}

.language-items > ul > li > .list-select {
  margin-right: 10px;
}

.language-items > ul {
  margin-bottom: 10px;
}

.control-avatar {
  width: 200px;
  height: 90px;
  position: relative;
}

.control-avatar > img {
  border-radius: 50%;
  width: 100%;
}

.control-avatar > button {
  position: absolute;
  top: 0;
  right: 0;
}

.file-loader-block {
  width: 200px;
  height: 90px;
  min-height: 90px;
}

.list-select > select {
  width: 200px;
}
</style>